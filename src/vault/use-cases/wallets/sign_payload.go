package wallets

import (
	"context"
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcec/v2/ecdsa"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/lugondev/signer-hashicorp-vault-plugin/src/pkg/errors"
	usecases "github.com/lugondev/signer-hashicorp-vault-plugin/src/vault/use-cases"

	"github.com/hashicorp/vault/sdk/logical"
	"github.com/lugondev/signer-hashicorp-vault-plugin/src/pkg/log"

	"github.com/ethereum/go-ethereum/crypto"
)

// signPayloadUseCase is a use case to sign an arbitrary payload usign an existing wallet
type signPayloadUseCase struct {
	getWalletUC usecases.GetWalletUseCase
}

// NewSignPayloadUseCase creates a new SignUseCase
func NewSignPayloadUseCase(getWalletUC usecases.GetWalletUseCase) usecases.SignPayloadUseCase {
	return &signPayloadUseCase{
		getWalletUC: getWalletUC,
	}
}

func (uc *signPayloadUseCase) WithStorage(storage logical.Storage) usecases.SignPayloadUseCase {
	uc.getWalletUC = uc.getWalletUC.WithStorage(storage)
	return uc
}

// Execute signs an arbitrary payload using an existing wallet
func (uc *signPayloadUseCase) Execute(ctx context.Context, pubkey, namespace, data string) (string, error) {
	logger := log.FromContext(ctx).With("namespace", namespace).With("pubkey", pubkey)
	logger.Debug("signing message", "type", "wallets/sign_payload")

	dataBytes, err := hexutil.Decode(data)
	if err != nil {
		errMessage := "data must be a hex string"
		logger.With("error", err).Error(errMessage)
		return "", errors.InvalidParameterError(errMessage)
	}

	wallet, err := uc.getWalletUC.Execute(ctx, pubkey, namespace)
	if err != nil {
		return "", err
	}

	ecdsaPrivKey, err := crypto.HexToECDSA(wallet.PrivateKey)
	if err != nil {
		errMessage := "failed to parse private key"
		logger.With("error", err).Error(errMessage)
		return "", errors.CryptoOperationError(errMessage)
	}

	key, _ := btcec.PrivKeyFromBytes(ecdsaPrivKey.D.Bytes())
	signature := ecdsa.Sign(key, dataBytes)
	if err != nil {
		errMessage := "failed to sign payload"
		logger.With("error", err).Error(errMessage)
		return "", errors.CryptoOperationError(errMessage)
	}

	logger.Info("payload signed successfully")
	return hexutil.Encode(signature.Serialize()), nil
}

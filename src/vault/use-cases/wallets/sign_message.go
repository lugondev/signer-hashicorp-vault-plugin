package wallets

import (
	"context"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/lugondev/signer-hashicorp-vault-plugin/src/pkg/errors"
	usecases "github.com/lugondev/signer-hashicorp-vault-plugin/src/vault/use-cases"

	"github.com/hashicorp/vault/sdk/logical"
	"github.com/lugondev/signer-hashicorp-vault-plugin/src/pkg/log"

	"github.com/ethereum/go-ethereum/crypto"
)

// signMessageUseCase is a use case to sign an arbitrary payload usign an existing Ethereum account
type signMessageUseCase struct {
	getWalletUC usecases.GetWalletUseCase
}

// NewSignMessageUseCase creates a new SignUseCase
func NewSignMessageUseCase(getWalletUC usecases.GetWalletUseCase) usecases.SignMessageUseCase {
	return &signMessageUseCase{
		getWalletUC: getWalletUC,
	}
}

func (uc *signMessageUseCase) WithStorage(storage logical.Storage) usecases.SignMessageUseCase {
	uc.getWalletUC = uc.getWalletUC.WithStorage(storage)
	return uc
}

// Execute signs an arbitrary payload using an existing Ethereum account
func (uc *signMessageUseCase) Execute(ctx context.Context, pubkey, namespace, data string) (string, error) {
	logger := log.FromContext(ctx).With("namespace", namespace).With("pubkey", pubkey)
	logger.Debug("signing message", "type", "wallets/sign_message")

	dataBytes, err := hexutil.Decode(data)
	if err != nil {
		errMessage := "data must be a hex string"
		logger.With("error", err).Error(errMessage)
		return "", errors.InvalidParameterError(errMessage)
	}

	account, err := uc.getWalletUC.Execute(ctx, pubkey, namespace)
	if err != nil {
		return "", err
	}

	ecdsaPrivKey, err := crypto.HexToECDSA(account.PrivateKey)
	if err != nil {
		errMessage := "failed to parse private key"
		logger.With("error", err).Error(errMessage)
		return "", errors.CryptoOperationError(errMessage)
	}

	signature, err := crypto.Sign(crypto.Keccak256(dataBytes), ecdsaPrivKey)
	if err != nil {
		errMessage := "failed to sign payload"
		logger.With("error", err).Error(errMessage)
		return "", errors.CryptoOperationError(errMessage)
	}

	logger.Info("payload signed successfully")
	return hexutil.Encode(signature), nil
}

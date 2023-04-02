package ethereum

import (
	"context"
	"github.com/ethereum/go-ethereum/core/types"
	signing "github.com/lugondev/signer-hashicorp-vault-plugin/src/pkg/crypto/ethereum"
	"github.com/lugondev/signer-hashicorp-vault-plugin/src/pkg/errors"
	usecases "github.com/lugondev/signer-hashicorp-vault-plugin/src/vault/use-cases"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/hashicorp/vault/sdk/logical"
	"github.com/lugondev/signer-hashicorp-vault-plugin/src/pkg/log"
)

// signTxUseCase is a use case to sign an ethereum transaction using an existing account
type signTxUseCase struct {
	getAccountUC usecases.GetAccountUseCase
}

// NewSignTransactionUseCase creates a new SignTransactionUseCase
func NewSignTransactionUseCase(getAccountUC usecases.GetAccountUseCase) usecases.SignTransactionUseCase {
	return &signTxUseCase{
		getAccountUC: getAccountUC,
	}
}

func (uc *signTxUseCase) WithStorage(storage logical.Storage) usecases.SignTransactionUseCase {
	uc.getAccountUC = uc.getAccountUC.WithStorage(storage)
	return uc
}

// Execute signs an ethereum transaction
func (uc *signTxUseCase) Execute(ctx context.Context, address, namespace, chainID string, tx *types.Transaction) (string, error) {
	logger := log.FromContext(ctx).With("namespace", namespace).With("address", address)
	logger.Debug("signing ethereum transaction")

	account, err := uc.getAccountUC.Execute(ctx, address, namespace)
	if err != nil {
		return "", err
	}

	ecdsaPrivKey, err := crypto.HexToECDSA(account.PrivateKey)
	if err != nil {
		errMessage := "failed to parse private key"
		logger.With("error", err).Error(errMessage)
		return "", errors.CryptoOperationError(errMessage)
	}

	signature, err := signing.SignTransaction(tx, ecdsaPrivKey, signing.GetEIP155Signer(chainID))
	if err != nil {
		errMessage := "failed to sign ethereum transaction"
		logger.With("error", err).Error(errMessage)
		return "", errors.CryptoOperationError(errMessage)
	}

	logger.Info("ethereum transaction signed successfully")
	return hexutil.Encode(signature), nil
}

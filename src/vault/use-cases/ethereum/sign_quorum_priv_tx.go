package ethereum

import (
	"context"
	"github.com/lugondev/signer-hashicorp-vault-plugin/src/pkg/errors"
	usecases "github.com/lugondev/signer-hashicorp-vault-plugin/src/vault/use-cases"

	"github.com/consensys/quorum/common/hexutil"
	"github.com/consensys/quorum/core/types"
	"github.com/consensys/quorum/crypto"
	"github.com/hashicorp/vault/sdk/logical"
	signing "github.com/lugondev/signer-hashicorp-vault-plugin/src/pkg/crypto/ethereum"
	"github.com/lugondev/signer-hashicorp-vault-plugin/src/pkg/log"
)

// signQuorumPrivateTxUseCase is a use case to sign a Quorum private transaction using an existing account
type signQuorumPrivateTxUseCase struct {
	getAccountUC usecases.GetAccountUseCase
}

// NewSignQuorumPrivateTransactionUseCase creates a new signQuorumPrivateTxUseCase
func NewSignQuorumPrivateTransactionUseCase(getAccountUC usecases.GetAccountUseCase) usecases.SignQuorumPrivateTransactionUseCase {
	return &signQuorumPrivateTxUseCase{
		getAccountUC: getAccountUC,
	}
}

func (uc *signQuorumPrivateTxUseCase) WithStorage(storage logical.Storage) usecases.SignQuorumPrivateTransactionUseCase {
	uc.getAccountUC = uc.getAccountUC.WithStorage(storage)
	return uc
}

// Execute signs a Quorum private transaction
func (uc *signQuorumPrivateTxUseCase) Execute(ctx context.Context, address, namespace string, tx *types.Transaction) (string, error) {
	logger := log.FromContext(ctx).With("namespace", namespace).With("address", address)
	logger.Debug("signing quorum private transaction")

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

	signature, err := signing.SignQuorumPrivateTransaction(tx, ecdsaPrivKey, signing.GetQuorumPrivateTxSigner())
	if err != nil {
		errMessage := "failed to sign quorum private transaction"
		logger.With("error", err).Error(errMessage)
		return "", errors.CryptoOperationError(errMessage)
	}

	logger.Info("quorum private transaction signed successfully")
	return hexutil.Encode(signature), nil
}

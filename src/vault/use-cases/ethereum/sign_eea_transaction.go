package ethereum

import (
	"context"
	apputils "github.com/ConsenSys/orchestrate-hashicorp-vault-plugin/src/utils"
	usecases "github.com/ConsenSys/orchestrate-hashicorp-vault-plugin/src/vault/use-cases"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/hashicorp/vault/sdk/logical"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/pkg/crypto/ethereum/signing"
	"gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/pkg/types/entities"
)

// signEEATxUseCase is a use case to sign a Quorum private transaction using an existing account
type signEEATxUseCase struct {
	getAccountUC usecases.GetAccountUseCase
}

// NewSignEEATransactionUseCase creates a new signEEATxUseCase
func NewSignEEATransactionUseCase(getAccountUC usecases.GetAccountUseCase) usecases.SignEEATransactionUseCase {
	return &signEEATxUseCase{
		getAccountUC: getAccountUC,
	}
}

func (uc signEEATxUseCase) WithStorage(storage logical.Storage) usecases.SignEEATransactionUseCase {
	uc.getAccountUC = uc.getAccountUC.WithStorage(storage)
	return &uc
}

// Execute signs an EEA transaction
func (uc *signEEATxUseCase) Execute(
	ctx context.Context,
	address, namespace, chainID string,
	tx *ethtypes.Transaction,
	privateArgs *entities.PrivateETHTransactionParams,
) (string, error) {
	logger := apputils.Logger(ctx).With("namespace", namespace).With("address", address)
	logger.Debug("signing eea private transaction")

	account, err := uc.getAccountUC.Execute(ctx, address, namespace)
	if err != nil {
		return "", err
	}

	ecdsaPrivKey, err := crypto.HexToECDSA(account.PrivateKey)
	if err != nil {
		errMessage := "failed to parse private key"
		logger.With("error", err).Error(errMessage)
		return "", err
	}

	signature, err := signing.SignEEATransaction(tx, privateArgs, chainID, ecdsaPrivKey)
	if err != nil {
		return "", err
	}

	logger.Info("eea private transaction signed successfully")
	return hexutil.Encode(signature), nil
}
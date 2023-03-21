package ethereum

import (
	"context"
	"github.com/lugondev/signer-hashicorp-vault-plugin/src/pkg/errors"
	usecases "github.com/lugondev/signer-hashicorp-vault-plugin/src/vault/use-cases"

	"github.com/hashicorp/vault/sdk/logical"
	"github.com/lugondev/signer-hashicorp-vault-plugin/src/pkg/log"
	"github.com/lugondev/signer-hashicorp-vault-plugin/src/vault/storage"
)

// listAccountsUseCase is a use case to get a list of Ethereum accounts
type listAccountsUseCase struct {
	storage logical.Storage
}

// NewListAccountUseCase creates a new ListAccountsUseCase
func NewListAccountsUseCase() usecases.ListAccountsUseCase {
	return &listAccountsUseCase{}
}

func (uc listAccountsUseCase) WithStorage(storage logical.Storage) usecases.ListAccountsUseCase {
	uc.storage = storage
	return &uc
}

// Execute gets a list of Ethereum accounts
func (uc *listAccountsUseCase) Execute(ctx context.Context, namespace string) ([]string, error) {
	logger := log.FromContext(ctx).With("namespace", namespace)
	logger.Debug("listing Ethereum accounts")

	keys, err := uc.storage.List(ctx, storage.ComputeEthereumStorageKey("", namespace))
	if err != nil {
		errMessage := "failed to list keys"
		logger.With("error", err).Error(errMessage)
		return nil, errors.StorageError(errMessage)
	}

	return keys, nil
}

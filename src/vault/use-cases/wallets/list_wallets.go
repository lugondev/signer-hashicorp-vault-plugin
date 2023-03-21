package wallets

import (
	"context"
	"github.com/lugondev/signer-hashicorp-vault-plugin/src/pkg/errors"
	usecases "github.com/lugondev/signer-hashicorp-vault-plugin/src/vault/use-cases"

	"github.com/hashicorp/vault/sdk/logical"
	"github.com/lugondev/signer-hashicorp-vault-plugin/src/pkg/log"
	"github.com/lugondev/signer-hashicorp-vault-plugin/src/vault/storage"
)

// listWalletsUseCase is a use case to get a list of wallets
type listWalletsUseCase struct {
	storage logical.Storage
}

// NewListWalletsUseCase creates a new ListWalletsUseCase
func NewListWalletsUseCase() usecases.ListWalletsUseCase {
	return &listWalletsUseCase{}
}

func (uc *listWalletsUseCase) WithStorage(storage logical.Storage) usecases.ListWalletsUseCase {
	uc.storage = storage
	return uc
}

// Execute gets a list of wallets
func (uc *listWalletsUseCase) Execute(ctx context.Context, namespace string) ([]string, error) {
	logger := log.FromContext(ctx).With("namespace", namespace)
	logger.Debug("listing Wallets")

	keys, err := uc.storage.List(ctx, storage.ComputeWalletsStorageKey("", namespace))
	if err != nil {
		errMessage := "failed to list keys"
		logger.With("error", err).Error(errMessage)
		return nil, errors.StorageError(errMessage)
	}

	return keys, nil
}

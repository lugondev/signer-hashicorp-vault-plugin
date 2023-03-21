package keys

import (
	"context"
	"github.com/lugondev/signer-hashicorp-vault-plugin/src/pkg/errors"

	"github.com/hashicorp/vault/sdk/logical"
	"github.com/lugondev/signer-hashicorp-vault-plugin/src/pkg/log"
	"github.com/lugondev/signer-hashicorp-vault-plugin/src/vault/storage"
	usecases "github.com/lugondev/signer-hashicorp-vault-plugin/src/vault/use-cases"
)

type listKeysUseCase struct {
	storage logical.Storage
}

func NewListKeysUseCase() usecases.ListKeysUseCase {
	return &listKeysUseCase{}
}

func (uc listKeysUseCase) WithStorage(storage logical.Storage) usecases.ListKeysUseCase {
	uc.storage = storage
	return &uc
}

func (uc *listKeysUseCase) Execute(ctx context.Context, namespace string) ([]string, error) {
	logger := log.FromContext(ctx).With("namespace", namespace)
	logger.Debug("listing key pairs")

	keys, err := uc.storage.List(ctx, storage.ComputeKeysStorageKey("", namespace))
	if err != nil {
		errMessage := "failed to list keys"
		logger.With("error", err).Error(errMessage)
		return nil, errors.StorageError(errMessage)
	}

	return keys, nil
}

package keys

import (
	"context"
	"github.com/hashicorp/vault/sdk/logical"
	"github.com/lugondev/signer-hashicorp-vault-plugin/src/pkg/log"
	"github.com/lugondev/signer-hashicorp-vault-plugin/src/vault/entities"
	"github.com/lugondev/signer-hashicorp-vault-plugin/src/vault/storage"
	usecases "github.com/lugondev/signer-hashicorp-vault-plugin/src/vault/use-cases"
)

type updateKeyUseCase struct {
	storage  logical.Storage
	getKeyUC usecases.GetKeyUseCase
}

func NewUpdateKeyUseCase(getKeyUC usecases.GetKeyUseCase) usecases.UpdateKeyUseCase {
	return &updateKeyUseCase{
		getKeyUC: getKeyUC,
	}
}

func (uc updateKeyUseCase) WithStorage(storage logical.Storage) usecases.UpdateKeyUseCase {
	uc.storage = storage
	uc.getKeyUC = uc.getKeyUC.WithStorage(storage)
	return &uc
}

func (uc *updateKeyUseCase) Execute(ctx context.Context, namespace, id string, tags map[string]string) (*entities.Key, error) {
	logger := log.FromContext(ctx).With("namespace", namespace).With("id", id)
	logger.Debug("updating key")

	key, err := uc.getKeyUC.Execute(ctx, id, namespace)
	if err != nil {
		return nil, err
	}

	key.Tags = tags
	err = storage.StoreJSON(ctx, uc.storage, storage.ComputeKeysStorageKey(id, namespace), key)
	if err != nil {
		return nil, err
	}

	logger.Info("key pair updated successfully")
	return key, nil
}

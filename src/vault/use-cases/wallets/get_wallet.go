package wallets

import (
	"context"
	usecases "github.com/lugondev/signer-hashicorp-vault-plugin/src/vault/use-cases"

	"github.com/hashicorp/vault/sdk/logical"
	"github.com/lugondev/signer-hashicorp-vault-plugin/src/pkg/log"
	"github.com/lugondev/signer-hashicorp-vault-plugin/src/vault/entities"
	"github.com/lugondev/signer-hashicorp-vault-plugin/src/vault/storage"
)

// getWalletUseCase is a use case to get wallets1
type getWalletUseCase struct {
	storage logical.Storage
}

// NewGetWalletUseCase creates a new GetWalletUseCase
func NewGetWalletUseCase() usecases.GetWalletUseCase {
	return &getWalletUseCase{}
}

func (uc *getWalletUseCase) WithStorage(storage logical.Storage) usecases.GetWalletUseCase {
	uc.storage = storage
	return uc
}

// Execute creates a wallet and stores it in the Vault
func (uc *getWalletUseCase) Execute(ctx context.Context, compressedPublicKey, namespace string) (*entities.Wallet, error) {
	logger := log.FromContext(ctx).With("namespace", namespace).With("compressedPublicKey", compressedPublicKey)
	logger.Debug("getting wallets")

	account := &entities.Wallet{}
	err := storage.GetJSON(ctx, uc.storage, storage.ComputeWalletsStorageKey(compressedPublicKey, namespace), account)
	if err != nil {
		return nil, err
	}

	logger.Debug("wallet found successfully")
	return account, nil
}

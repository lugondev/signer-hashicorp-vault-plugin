package usecases

import (
	"context"
	"github.com/hashicorp/vault/sdk/logical"

	"github.com/lugondev/signer-hashicorp-vault-plugin/src/vault/entities"
)

// //go:generate mockgen -source=ethereum.go -destination=mocks/ethereum.go -package=mocks

type WalletsUseCases interface {
	CreateWallet() CreateWalletUseCase
	GetWallet() GetWalletUseCase
	ListWallets() ListWalletsUseCase
	ListWalletsNamespaces() ListWalletsNamespacesUseCase
	SignPayload() SignPayloadUseCase
}

type CreateWalletUseCase interface {
	Execute(ctx context.Context, namespace, importedPrivKey string) (*entities.Wallet, error)
	WithStorage(storage logical.Storage) CreateWalletUseCase
}

type GetWalletUseCase interface {
	Execute(ctx context.Context, address, namespace string) (*entities.Wallet, error)
	WithStorage(storage logical.Storage) GetWalletUseCase
}

type ListWalletsUseCase interface {
	Execute(ctx context.Context, namespace string) ([]string, error)
	WithStorage(storage logical.Storage) ListWalletsUseCase
}

type SignPayloadUseCase interface {
	Execute(ctx context.Context, address, namespace, data string) (string, error)
	WithStorage(storage logical.Storage) SignPayloadUseCase
}

type ListWalletsNamespacesUseCase interface {
	Execute(ctx context.Context) ([]string, error)
	WithStorage(storage logical.Storage) ListWalletsNamespacesUseCase
}

package usecases

import (
	"context"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/hashicorp/vault/sdk/logical"

	"github.com/lugondev/signer-hashicorp-vault-plugin/src/vault/entities"
)

//go:generate mockgen -source=ethereum.go -destination=mocks/ethereum.go -package=mocks

type ETHUseCases interface {
	CreateAccount() CreateAccountUseCase
	GetAccount() GetAccountUseCase
	ListAccounts() ListAccountsUseCase
	ListNamespaces() ListNamespacesUseCase
	SignPayload() SignUseCase
	SignTransaction() SignTransactionUseCase
}

type CreateAccountUseCase interface {
	Execute(ctx context.Context, namespace, importedPrivKey string) (*entities.ETHAccount, error)
	WithStorage(storage logical.Storage) CreateAccountUseCase
}

type GetAccountUseCase interface {
	Execute(ctx context.Context, address, namespace string) (*entities.ETHAccount, error)
	WithStorage(storage logical.Storage) GetAccountUseCase
}

type ListAccountsUseCase interface {
	Execute(ctx context.Context, namespace string) ([]string, error)
	WithStorage(storage logical.Storage) ListAccountsUseCase
}

type SignUseCase interface {
	Execute(ctx context.Context, address, namespace, data string) (string, error)
	WithStorage(storage logical.Storage) SignUseCase
}

type SignTransactionUseCase interface {
	Execute(ctx context.Context, address, namespace, chainID string, tx *types.Transaction) (string, error)
	WithStorage(storage logical.Storage) SignTransactionUseCase
}

type ListNamespacesUseCase interface {
	Execute(ctx context.Context) ([]string, error)
	WithStorage(storage logical.Storage) ListNamespacesUseCase
}

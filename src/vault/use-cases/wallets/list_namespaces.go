package wallets

import (
	"context"
	usecases "github.com/lugondev/signer-hashicorp-vault-plugin/src/vault/use-cases"
	"strings"

	"github.com/hashicorp/vault/sdk/logical"
	"github.com/lugondev/signer-hashicorp-vault-plugin/src/pkg/log"
	"github.com/lugondev/signer-hashicorp-vault-plugin/src/vault/storage"
)

// listWalletsNamespacesUseCase is a use case to get a list of wallets
type listWalletsNamespacesUseCase struct {
	storage logical.Storage
}

// NewListWalletNamespacesUseCase creates a new ListAccountsUseCase
func NewListWalletNamespacesUseCase() usecases.ListWalletsNamespacesUseCase {
	return &listWalletsNamespacesUseCase{}
}

func (uc *listWalletsNamespacesUseCase) WithStorage(storage logical.Storage) usecases.ListWalletsNamespacesUseCase {
	uc.storage = storage
	return uc
}

// Execute get a list of all available namespaces
func (uc *listWalletsNamespacesUseCase) Execute(ctx context.Context) ([]string, error) {
	logger := log.FromContext(ctx)
	logger.Debug("listing wallets namespaces")

	namespaceSet := make(map[string]bool)
	err := storage.GetWalletsNamespaces(ctx, uc.storage, "", namespaceSet)
	if err != nil {
		return nil, err
	}

	namespaces := make([]string, 0, len(namespaceSet))
	for namespace := range namespaceSet {
		if namespace != "" {
			namespace = strings.TrimSuffix(namespace, "/")
		}
		namespaces = append(namespaces, namespace)
	}

	logger.Debug("wallets namespaces found successfully")
	return namespaces, nil

}

package ethereum

import (
	"context"
	usecases "github.com/lugondev/signer-hashicorp-vault-plugin/src/vault/use-cases"
	"strings"

	"github.com/hashicorp/vault/sdk/logical"
	"github.com/lugondev/signer-hashicorp-vault-plugin/src/pkg/log"
	"github.com/lugondev/signer-hashicorp-vault-plugin/src/vault/storage"
)

// listNamespacesUseCase is a use case to get a list of Ethereum accounts
type listNamespacesUseCase struct {
	storage logical.Storage
}

// NewListNamespacesUseCase creates a new ListAccountsUseCase
func NewListNamespacesUseCase() usecases.ListNamespacesUseCase {
	return &listNamespacesUseCase{}
}

func (uc *listNamespacesUseCase) WithStorage(storage logical.Storage) usecases.ListNamespacesUseCase {
	uc.storage = storage
	return uc
}

// Execute get a list of all available namespaces
func (uc *listNamespacesUseCase) Execute(ctx context.Context) ([]string, error) {
	logger := log.FromContext(ctx)
	logger.Debug("listing ethereum namespaces")

	namespaceSet := make(map[string]bool)
	err := storage.GetEthereumNamespaces(ctx, uc.storage, "", namespaceSet)
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

	logger.Debug("ethereum namespaces found successfully")
	return namespaces, nil

}

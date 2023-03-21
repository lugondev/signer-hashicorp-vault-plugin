package wallets

import (
	"fmt"

	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
	"github.com/lugondev/signer-hashicorp-vault-plugin/src/pkg/log"
	"github.com/lugondev/signer-hashicorp-vault-plugin/src/service/formatters"
	usecases "github.com/lugondev/signer-hashicorp-vault-plugin/src/vault/use-cases"
)

type controller struct {
	useCases usecases.WalletsUseCases
	logger   log.Logger
}

func NewController(useCases usecases.WalletsUseCases, logger log.Logger) *controller {
	if logger == nil {
		logger = log.Default()
	}

	return &controller{
		useCases: useCases,
		logger:   logger.Named("wallets"),
	}
}

// Paths returns the list of paths
func (c *controller) Paths() []*framework.Path {
	return framework.PathAppend(
		[]*framework.Path{
			c.pathImportWallet(),
			c.pathWallets(),
			c.pathWallet(),
			c.pathSignPayload(),
			c.pathNamespaces(),
		},
	)
}

func (c *controller) pathWallets() *framework.Path {
	return &framework.Path{
		Pattern:      "wallets/?",
		HelpSynopsis: "Creates a new wallets or list them",
		Operations: map[logical.Operation]framework.OperationHandler{
			logical.CreateOperation: c.NewCreateOperation(),
			logical.UpdateOperation: c.NewCreateOperation(),
			logical.ListOperation:   c.NewListOperation(),
			logical.ReadOperation:   c.NewListOperation(),
		},
	}
}

func (c *controller) pathNamespaces() *framework.Path {
	return &framework.Path{
		Pattern:      "namespaces/wallets/?",
		HelpSynopsis: "Lists all wallets namespaces",
		Operations: map[logical.Operation]framework.OperationHandler{
			logical.ListOperation: c.NewListNamespacesOperation(),
			logical.ReadOperation: c.NewListNamespacesOperation(),
		},
	}
}

func (c *controller) pathWallet() *framework.Path {
	return &framework.Path{
		Pattern:      fmt.Sprintf("wallets/%s", framework.GenericNameRegex(formatters.IDLabel)),
		HelpSynopsis: "Get, update or delete a wallets",
		Fields: map[string]*framework.FieldSchema{
			formatters.IDLabel: formatters.AddressFieldSchema,
		},
		Operations: map[logical.Operation]framework.OperationHandler{
			logical.ReadOperation: c.NewGetOperation(),
		},
	}
}

func (c *controller) pathImportWallet() *framework.Path {
	return &framework.Path{
		Pattern: "wallets/import",
		Fields: map[string]*framework.FieldSchema{
			formatters.PrivateKeyLabel: {
				Type:        framework.TypeString,
				Description: "Private key in hexadecimal format",
				Required:    true,
			},
		},
		Operations: map[logical.Operation]framework.OperationHandler{
			logical.CreateOperation: c.NewImportOperation(),
			logical.UpdateOperation: c.NewImportOperation(),
		},
		HelpSynopsis: "Imports a wallets",
	}
}

func (c *controller) pathSignPayload() *framework.Path {
	return &framework.Path{
		Pattern: fmt.Sprintf("wallets/%s/sign", framework.GenericNameRegex(formatters.IDLabel)),
		Fields: map[string]*framework.FieldSchema{
			formatters.IDLabel: formatters.AddressFieldSchema,
			formatters.DataLabel: {
				Type:        framework.TypeString,
				Description: "data to sign",
				Required:    true,
			},
		},
		Operations: map[logical.Operation]framework.OperationHandler{
			logical.CreateOperation: c.NewSignPayloadOperation(),
			logical.UpdateOperation: c.NewSignPayloadOperation(),
		},
		HelpSynopsis: "Signs an arbitrary message using an existing wallet",
	}
}

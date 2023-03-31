package ethereum

import (
	"fmt"

	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
	"github.com/lugondev/signer-hashicorp-vault-plugin/src/pkg/log"
	"github.com/lugondev/signer-hashicorp-vault-plugin/src/service/formatters"
	usecases "github.com/lugondev/signer-hashicorp-vault-plugin/src/vault/use-cases"
)

type controller struct {
	useCases usecases.ETHUseCases
	logger   log.Logger
}

func NewController(useCases usecases.ETHUseCases, logger log.Logger) *controller {
	if logger == nil {
		logger = log.Default()
	}

	return &controller{
		useCases: useCases,
		logger:   logger.Named("ethereum"),
	}
}

// Paths returns the list of paths
func (c *controller) Paths() []*framework.Path {
	return framework.PathAppend(
		[]*framework.Path{
			c.pathAccounts(),
			c.pathImportAccount(),
			c.pathAccount(),
			c.pathSignPayload(),
			c.pathSignTransaction(),
			c.pathNamespaces(),
		},
	)
}

func (c *controller) pathAccounts() *framework.Path {
	return &framework.Path{
		Pattern:      "ethereum/?",
		HelpSynopsis: "Creates a new Ethereum account or list them",
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
		Pattern:      "namespaces/ethereum/?",
		HelpSynopsis: "Lists all ethereum namespaces",
		Operations: map[logical.Operation]framework.OperationHandler{
			logical.ListOperation: c.NewListNamespacesOperation(),
			logical.ReadOperation: c.NewListNamespacesOperation(),
		},
	}
}

func (c *controller) pathAccount() *framework.Path {
	return &framework.Path{
		Pattern:      fmt.Sprintf("ethereum/%s", framework.GenericNameRegex(formatters.IDLabel)),
		HelpSynopsis: "Get, update or delete an Ethereum account",
		Fields: map[string]*framework.FieldSchema{
			formatters.IDLabel: formatters.AddressFieldSchema,
		},
		Operations: map[logical.Operation]framework.OperationHandler{
			logical.ReadOperation: c.NewGetOperation(),
		},
	}
}

func (c *controller) pathImportAccount() *framework.Path {
	return &framework.Path{
		Pattern: "ethereum/import",
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
		HelpSynopsis: "Imports an Ethereum account",
	}
}

func (c *controller) pathSignPayload() *framework.Path {
	return &framework.Path{
		Pattern: fmt.Sprintf("ethereum/%s/sign", framework.GenericNameRegex(formatters.IDLabel)),
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
		HelpSynopsis: "Signs an arbitrary message using an existing Ethereum account",
	}
}

func (c *controller) pathSignTransaction() *framework.Path {
	return &framework.Path{
		Pattern: fmt.Sprintf("ethereum/%s/sign-transaction", framework.GenericNameRegex(formatters.IDLabel)),
		Fields: map[string]*framework.FieldSchema{
			formatters.IDLabel:       formatters.AddressFieldSchema,
			formatters.NonceLabel:    formatters.NonceFieldSchema,
			formatters.ToLabel:       formatters.ToFieldSchema,
			formatters.AmountLabel:   formatters.AmountFieldSchema,
			formatters.GasPriceLabel: formatters.GasPriceFieldSchema,
			formatters.GasLimitLabel: formatters.GasLimitFieldSchema,
			formatters.ChainIDLabel:  formatters.ChainIDFieldSchema,
			formatters.DataLabel:     formatters.DataFieldSchema,
		},
		Operations: map[logical.Operation]framework.OperationHandler{
			logical.CreateOperation: c.NewSignTransactionOperation(),
			logical.UpdateOperation: c.NewSignTransactionOperation(),
		},
		HelpSynopsis: "Signs an Ethereum transaction using an existing account",
	}
}

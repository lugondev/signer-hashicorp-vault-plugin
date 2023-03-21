package keys

import (
	"fmt"

	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
	"github.com/lugondev/signer-hashicorp-vault-plugin/src/pkg/log"
	"github.com/lugondev/signer-hashicorp-vault-plugin/src/service/formatters"
	usecases "github.com/lugondev/signer-hashicorp-vault-plugin/src/vault/use-cases"
)

type controller struct {
	useCases usecases.KeysUseCases
	logger   log.Logger
}

func NewController(useCases usecases.KeysUseCases, logger log.Logger) *controller {
	if logger == nil {
		logger = log.Default()
	}

	return &controller{
		useCases: useCases,
		logger:   logger.Named("keys"),
	}
}

// Paths returns the list of paths
func (c *controller) Paths() []*framework.Path {
	return framework.PathAppend(
		[]*framework.Path{
			c.pathKeys(),
			c.pathImportKey(),
			c.pathKey(),
			c.pathSignPayload(),
			c.pathNamespaces(),
			c.pathDestroy(),
		},
	)
}

func (c *controller) pathKeys() *framework.Path {
	return &framework.Path{
		Pattern:      "keys/?",
		HelpSynopsis: "Creates a new key pair or list them",
		Operations: map[logical.Operation]framework.OperationHandler{
			logical.CreateOperation: c.NewCreateOperation(),
			logical.UpdateOperation: c.NewCreateOperation(),
			logical.ListOperation:   c.NewListOperation(),
			logical.ReadOperation:   c.NewListOperation(),
		},
		Fields: map[string]*framework.FieldSchema{
			formatters.IDLabel: formatters.IDFieldSchema,
			formatters.CurveLabel: {
				Type:        framework.TypeString,
				Description: "Elliptic curve",
				Required:    true,
			},
			formatters.AlgorithmLabel: {
				Type:        framework.TypeString,
				Description: "Signing algorithm",
				Required:    true,
			},
			formatters.TagsLabel: formatters.TagsFieldSchema,
		},
	}
}

func (c *controller) pathNamespaces() *framework.Path {
	return &framework.Path{
		Pattern:      "namespaces/keys/?",
		HelpSynopsis: "Lists all keys namespaces",
		Operations: map[logical.Operation]framework.OperationHandler{
			logical.ListOperation: c.NewListNamespacesOperation(),
			logical.ReadOperation: c.NewListNamespacesOperation(),
		},
	}
}

func (c *controller) pathKey() *framework.Path {
	return &framework.Path{
		Pattern:      fmt.Sprintf("keys/%s", framework.GenericNameRegex(formatters.IDLabel)),
		HelpSynopsis: "Get, update or delete a key pair",
		Fields: map[string]*framework.FieldSchema{
			formatters.IDLabel:   formatters.IDFieldSchema,
			formatters.TagsLabel: formatters.TagsFieldSchema,
		},
		Operations: map[logical.Operation]framework.OperationHandler{
			logical.ReadOperation:   c.NewGetOperation(),
			logical.UpdateOperation: c.NewUpdateOperation(),
		},
	}
}

func (c *controller) pathImportKey() *framework.Path {
	return &framework.Path{
		Pattern: "keys/import",
		Fields: map[string]*framework.FieldSchema{
			formatters.PrivateKeyLabel: {
				Type:        framework.TypeString,
				Description: "Private key in base64 format",
				Required:    true,
			},
			formatters.IDLabel: formatters.IDFieldSchema,
			formatters.CurveLabel: {
				Type:        framework.TypeString,
				Description: "Elliptic curve",
				Required:    true,
			},
			formatters.AlgorithmLabel: {
				Type:        framework.TypeString,
				Description: "Signing algorithm",
				Required:    true,
			},
			formatters.TagsLabel: formatters.TagsFieldSchema,
		},
		Operations: map[logical.Operation]framework.OperationHandler{
			logical.CreateOperation: c.NewImportOperation(),
			logical.UpdateOperation: c.NewImportOperation(),
		},
		HelpSynopsis: "Imports a key pair",
	}
}

func (c *controller) pathSignPayload() *framework.Path {
	return &framework.Path{
		Pattern: fmt.Sprintf("keys/%s/sign", framework.GenericNameRegex(formatters.IDLabel)),
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
		HelpSynopsis: "Signs an arbitrary message using an existing key pair",
	}
}

func (c *controller) pathDestroy() *framework.Path {
	return &framework.Path{
		Pattern: fmt.Sprintf("keys/%s/destroy", framework.GenericNameRegex(formatters.IDLabel)),
		Fields: map[string]*framework.FieldSchema{
			formatters.IDLabel: formatters.AddressFieldSchema,
		},
		Operations: map[logical.Operation]framework.OperationHandler{
			logical.DeleteOperation: c.NewDestroyOperation(),
		},
		HelpSynopsis: "Destroys an existing key pair",
	}
}

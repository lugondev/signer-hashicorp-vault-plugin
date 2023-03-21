package wallets

import (
	"context"
	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
	"github.com/lugondev/signer-hashicorp-vault-plugin/src/pkg/log"
	"github.com/lugondev/signer-hashicorp-vault-plugin/src/service/errors"
	"github.com/lugondev/signer-hashicorp-vault-plugin/src/utils"
)

func (c *controller) NewListNamespacesOperation() *framework.PathOperation {
	return &framework.PathOperation{
		Callback:    c.listNamespacesHandler(),
		Summary:     "Gets a list of all wallets namespaces",
		Description: "Gets a list of all wallets namespaces",
		Examples: []framework.RequestExample{
			{
				Description: "Gets all wallets namespaces",
				Response: &framework.Response{
					Description: "Success",
					Example:     logical.ListResponse([]string{"ns1", "ns2"}),
				},
			},
		},
		Responses: map[int][]framework.Response{
			200: {framework.Response{
				Description: "Success",
				Example:     logical.ListResponse([]string{"ns1", "ns2"}),
			}},
			500: {utils.Example500Response()},
		},
	}
}

func (c *controller) listNamespacesHandler() framework.OperationFunc {
	return func(ctx context.Context, req *logical.Request, data *framework.FieldData) (*logical.Response, error) {
		ctx = log.Context(ctx, c.logger)
		namespaces, err := c.useCases.ListWalletsNamespaces().WithStorage(req.Storage).Execute(ctx)
		if err != nil {
			return errors.ParseHTTPError(err)
		}

		return logical.ListResponse(namespaces), nil
	}
}

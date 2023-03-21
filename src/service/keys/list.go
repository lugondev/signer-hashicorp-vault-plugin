package keys

import (
	"context"
	"github.com/lugondev/signer-hashicorp-vault-plugin/src/service/errors"

	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
	"github.com/lugondev/signer-hashicorp-vault-plugin/src/pkg/log"
	"github.com/lugondev/signer-hashicorp-vault-plugin/src/service/formatters"
	"github.com/lugondev/signer-hashicorp-vault-plugin/src/utils"
)

func (c *controller) NewListOperation() *framework.PathOperation {
	return &framework.PathOperation{
		Callback:    c.listHandler(),
		Summary:     "Gets a list of all key pair ids",
		Description: "Gets a list of all key pair ids optionally filtered by namespace",
		Examples: []framework.RequestExample{
			{
				Description: "Gets all key pair ids",
				Response: &framework.Response{
					Description: "Success",
					Example:     logical.ListResponse([]string{utils.ExampleKey().ID}),
				},
			},
		},
		Responses: map[int][]framework.Response{
			200: {framework.Response{
				Description: "Success",
				Example:     logical.ListResponse([]string{utils.ExampleKey().ID}),
			}},
			500: {utils.Example500Response()},
		},
	}
}

func (c *controller) listHandler() framework.OperationFunc {
	return func(ctx context.Context, req *logical.Request, data *framework.FieldData) (*logical.Response, error) {
		namespace := formatters.GetRequestNamespace(req)

		ctx = log.Context(ctx, c.logger)
		keys, err := c.useCases.ListKeys().WithStorage(req.Storage).Execute(ctx, namespace)
		if err != nil {
			return errors.ParseHTTPError(err)
		}

		return logical.ListResponse(keys), nil
	}
}

package wallets

import (
	"context"
	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
	"github.com/lugondev/signer-hashicorp-vault-plugin/src/pkg/log"
	"github.com/lugondev/signer-hashicorp-vault-plugin/src/service/errors"
	"github.com/lugondev/signer-hashicorp-vault-plugin/src/service/formatters"
	"github.com/lugondev/signer-hashicorp-vault-plugin/src/utils"
)

func (c *controller) NewGetOperation() *framework.PathOperation {
	exampleAccount := utils.ExampleETHAccount()
	successExample := utils.Example200Response()

	return &framework.PathOperation{
		Callback:    c.getHandler(),
		Summary:     "Gets a wallet",
		Description: "Gets a wallet stored in the vault at the given address and namespace",
		Examples: []framework.RequestExample{
			{
				Description: "Gets an account on the tenant0 namespace",
				Data: map[string]interface{}{
					formatters.IDLabel: exampleAccount.Address,
				},
				Response: successExample,
			},
		},
		Responses: map[int][]framework.Response{
			200: {*successExample},
			400: {utils.Example400Response()},
			404: {utils.Example404Response()},
			500: {utils.Example500Response()},
		},
	}
}

func (c *controller) getHandler() framework.OperationFunc {
	return func(ctx context.Context, req *logical.Request, data *framework.FieldData) (*logical.Response, error) {
		address := data.Get(formatters.IDLabel).(string)
		namespace := formatters.GetRequestNamespace(req)

		ctx = log.Context(ctx, c.logger)
		account, err := c.useCases.GetWallet().WithStorage(req.Storage).Execute(ctx, address, namespace)
		if err != nil {
			return errors.ParseHTTPError(err)
		}

		return formatters.FormatWalletResponse(account), nil
	}
}

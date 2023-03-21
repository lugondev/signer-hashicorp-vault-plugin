package ethereum

import (
	"context"
	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
	"github.com/lugondev/signer-hashicorp-vault-plugin/src/pkg/log"
	"github.com/lugondev/signer-hashicorp-vault-plugin/src/service/errors"
	"github.com/lugondev/signer-hashicorp-vault-plugin/src/service/formatters"
	"github.com/lugondev/signer-hashicorp-vault-plugin/src/utils"
)

func (c *controller) NewImportOperation() *framework.PathOperation {
	exampleAccount := utils.ExampleETHAccount()
	successExample := utils.Example200Response()

	return &framework.PathOperation{
		Callback:    c.importHandler(),
		Summary:     "Imports an Ethereum account",
		Description: "Imports an Ethereum account given a private key, storing it in the Vault and computing its public key and address",
		Examples: []framework.RequestExample{
			{
				Description: "Imports an account on the tenant0 namespace",
				Data: map[string]interface{}{
					formatters.PrivateKeyLabel: exampleAccount.PrivateKey,
				},
				Response: successExample,
			},
		},
		Responses: map[int][]framework.Response{
			200: {*successExample},
			400: {utils.Example400Response()},
			500: {utils.Example500Response()},
		},
	}
}

func (c *controller) importHandler() framework.OperationFunc {
	return func(ctx context.Context, req *logical.Request, data *framework.FieldData) (*logical.Response, error) {
		privateKeyString := data.Get(formatters.PrivateKeyLabel).(string)
		namespace := formatters.GetRequestNamespace(req)

		if privateKeyString == "" {
			return logical.ErrorResponse("private_key must be provided"), nil
		}

		ctx = log.Context(ctx, c.logger)
		account, err := c.useCases.CreateAccount().WithStorage(req.Storage).Execute(ctx, namespace, privateKeyString)
		if err != nil {
			return errors.ParseHTTPError(err)
		}

		return formatters.FormatAccountResponse(account), nil
	}
}

package wallets

import (
	"context"
	"github.com/lugondev/signer-hashicorp-vault-plugin/src/service/errors"

	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
	errors2 "github.com/lugondev/signer-hashicorp-vault-plugin/src/pkg/errors"
	"github.com/lugondev/signer-hashicorp-vault-plugin/src/pkg/log"
	"github.com/lugondev/signer-hashicorp-vault-plugin/src/service/formatters"
	"github.com/lugondev/signer-hashicorp-vault-plugin/src/utils"
)

func (c *controller) NewSignPayloadOperation() *framework.PathOperation {
	exampleAccount := utils.ExampleETHAccount()

	return &framework.PathOperation{
		Callback:    c.signPayloadHandler(),
		Summary:     "Signs an arbitrary message using an existing wallet",
		Description: "Signs an arbitrary message using ECDSA and the private key of an existing wallets",
		Examples: []framework.RequestExample{
			{
				Description: "Signs by pubkey",
				Data: map[string]interface{}{
					formatters.IDLabel:       exampleAccount.Address,
					formatters.DataLabel:     "my data to sign",
					formatters.TypeSignLabel: "type sign: eth_sign, ecdsa, taproot",
				},
				Response: utils.Example200ResponseSignature(),
			},
		},
		Responses: map[int][]framework.Response{
			200: {*utils.Example200ResponseSignature()},
			400: {utils.Example400Response()},
			404: {utils.Example404Response()},
			500: {utils.Example500Response()},
		},
	}
}

func (c *controller) signPayloadHandler() framework.OperationFunc {
	return func(ctx context.Context, req *logical.Request, data *framework.FieldData) (*logical.Response, error) {
		pubkey := data.Get(formatters.IDLabel).(string)
		payload := data.Get(formatters.DataLabel).(string)
		typeSign := data.Get(formatters.TypeSignLabel).(string)
		namespace := formatters.GetRequestNamespace(req)

		if payload == "" {
			return errors.ParseHTTPError(errors2.InvalidFormatError("payload must be provided"))
		}
		if typeSign == "" || (typeSign != "ecdsa" && typeSign != "taproot" && typeSign != "eth_sign") {
			return errors.ParseHTTPError(errors2.InvalidFormatError("type sign must be provided"))
		}

		c.logger.Info("signPayloadHandler", "type_sign", typeSign, "payload", payload)
		ctx = log.Context(ctx, c.logger)

		var signature string
		var err error
		if typeSign == "ecdsa" {
			signature, err = c.useCases.SignPayload().WithStorage(req.Storage).Execute(ctx, pubkey, namespace, payload)
		}
		if typeSign == "taproot" {
			signature, err = c.useCases.SignTaproot().WithStorage(req.Storage).Execute(ctx, pubkey, namespace, payload)
		}
		if typeSign == "eth_sign" {
			signature, err = c.useCases.SignMessage().WithStorage(req.Storage).Execute(ctx, pubkey, namespace, payload)
		}
		if err != nil {
			return errors.ParseHTTPError(err)
		}

		return formatters.FormatSignatureResponse(signature), nil
	}
}

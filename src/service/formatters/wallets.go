package formatters

import (
	"github.com/hashicorp/vault/sdk/logical"
	"github.com/lugondev/signer-hashicorp-vault-plugin/src/vault/entities"
)

func FormatWalletResponse(account *entities.Wallet) *logical.Response {
	return &logical.Response{
		Data: map[string]interface{}{
			PublicKeyLabel:           account.PublicKey,
			CompressedPublicKeyLabel: account.CompressedPublicKey,
			NamespaceLabel:           account.Namespace,
		},
	}
}

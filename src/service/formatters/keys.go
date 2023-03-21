package formatters

import (
	"github.com/hashicorp/vault/sdk/logical"
	"github.com/lugondev/signer-hashicorp-vault-plugin/src/vault/entities"
)

func FormatKeyResponse(key *entities.Key) *logical.Response {
	return &logical.Response{
		Data: map[string]interface{}{
			IDLabel:        key.ID,
			CurveLabel:     key.Curve,
			AlgorithmLabel: key.Algorithm,
			PublicKeyLabel: key.PublicKey,
			NamespaceLabel: key.Namespace,
			TagsLabel:      key.Tags,
			VersionLabel:   key.Version,
			CreatedAtLabel: key.CreatedAt,
			UpdatedAtLabel: key.UpdatedAt,
		},
	}
}

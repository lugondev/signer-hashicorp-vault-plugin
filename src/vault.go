package src

import (
	"context"
	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
	"github.com/lugondev/signer-hashicorp-vault-plugin/src/service/ethereum"
	"github.com/lugondev/signer-hashicorp-vault-plugin/src/service/keys"
	"github.com/lugondev/signer-hashicorp-vault-plugin/src/service/wallets"
	"github.com/lugondev/signer-hashicorp-vault-plugin/src/vault/builder"
)

// NewVaultBackend returns the Hashicorp Vault backend
func NewVaultBackend(ctx context.Context, conf *logical.BackendConfig) (logical.Backend, error) {
	vaultPlugin := &framework.Backend{
		Help: "Signer Hashicorp Vault Plugin",
		PathsSpecial: &logical.Paths{
			SealWrapStorage: []string{
				"ethereum/",
				"keys/",
				"wallets/",
			},
		},
		Secrets:     []*framework.Secret{},
		BackendType: logical.TypeLogical,
	}

	ethUseCases := builder.NewEthereumUseCases()
	keysUseCases := builder.NewKeysUseCases()
	walletsUseCases := builder.NewWalletsUseCases()

	ethereumController := ethereum.NewController(ethUseCases, conf.Logger)
	keysController := keys.NewController(keysUseCases, conf.Logger)
	walletsController := wallets.NewController(walletsUseCases, conf.Logger)

	vaultPlugin.Paths = framework.PathAppend(
		ethereumController.Paths(),
		keysController.Paths(),
		walletsController.Paths(),
	)

	if err := vaultPlugin.Setup(ctx, conf); err != nil {
		return nil, err
	}

	return vaultPlugin, nil
}

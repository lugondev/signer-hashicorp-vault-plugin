package builder

import (
	"github.com/lugondev/signer-hashicorp-vault-plugin/src/vault/use-cases"
	"github.com/lugondev/signer-hashicorp-vault-plugin/src/vault/use-cases/wallets"
)

type walletsUseCases struct {
	createWallet   usecases.CreateWalletUseCase
	getWallet      usecases.GetWalletUseCase
	listWallets    usecases.ListWalletsUseCase
	listNamespaces usecases.ListWalletsNamespacesUseCase
	sign           usecases.SignPayloadUseCase
}

func NewWalletsUseCases() usecases.WalletsUseCases {
	getWallet := wallets.NewGetWalletUseCase()
	return &walletsUseCases{
		createWallet:   wallets.NewCreateWalletUseCase(),
		getWallet:      getWallet,
		listWallets:    wallets.NewListWalletsUseCase(),
		listNamespaces: wallets.NewListWalletNamespacesUseCase(),
		sign:           wallets.NewSignPayloadUseCase(getWallet),
	}
}

func (ucs *walletsUseCases) CreateWallet() usecases.CreateWalletUseCase {
	return ucs.createWallet
}

func (ucs *walletsUseCases) GetWallet() usecases.GetWalletUseCase {
	return ucs.getWallet
}

func (ucs *walletsUseCases) ListWallets() usecases.ListWalletsUseCase {
	return ucs.listWallets
}

func (ucs *walletsUseCases) ListWalletsNamespaces() usecases.ListWalletsNamespacesUseCase {
	return ucs.listNamespaces
}

func (ucs *walletsUseCases) SignPayload() usecases.SignPayloadUseCase {
	return ucs.sign
}

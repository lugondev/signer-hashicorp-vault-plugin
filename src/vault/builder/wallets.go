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
	signPayload    usecases.SignPayloadUseCase
	signMessage    usecases.SignMessageUseCase
	signTaproot    usecases.SignTaprootUseCase
}

func NewWalletsUseCases() usecases.WalletsUseCases {
	getWallet := wallets.NewGetWalletUseCase()
	return &walletsUseCases{
		createWallet:   wallets.NewCreateWalletUseCase(),
		getWallet:      getWallet,
		listWallets:    wallets.NewListWalletsUseCase(),
		listNamespaces: wallets.NewListWalletNamespacesUseCase(),
		signPayload:    wallets.NewSignPayloadUseCase(getWallet),
		signMessage:    wallets.NewSignMessageUseCase(getWallet),
		signTaproot:    wallets.NewSignTaprootUseCase(getWallet),
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
	return ucs.signPayload
}

func (ucs *walletsUseCases) SignMessage() usecases.SignMessageUseCase {
	return ucs.signMessage
}

func (ucs *walletsUseCases) SignTaproot() usecases.SignTaprootUseCase {
	return ucs.signTaproot
}

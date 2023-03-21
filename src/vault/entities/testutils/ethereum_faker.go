package testutils

import (
	"github.com/consensys/quorum/common"
	"github.com/lugondev/signer-hashicorp-vault-plugin/src/vault/entities"
)

func FakePrivateETHTransactionParams() *entities.PrivateETHTransactionParams {
	return &entities.PrivateETHTransactionParams{
		PrivateFrom:   "A1aVtMxLCUHmBVHXoZzzBgPbW/wj5axDpW9X8l91SGo=",
		PrivateFor:    []string{"A1aVtMxLCUHmBVHXoZzzBgPbW/wj5axDpW9X8l91SGo="},
		PrivateTxType: entities.PrivateTxTypeRestricted,
	}
}

func FakeETHAccount() *entities.ETHAccount {
	return &entities.ETHAccount{
		Namespace:           "_",
		Address:             common.HexToAddress(RandHexString(12)).String(),
		PublicKey:           common.HexToHash(RandHexString(12)).String(),
		CompressedPublicKey: common.HexToHash(RandHexString(12)).String(),
	}
}

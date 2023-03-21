package utils

import (
	"math/rand"
	"time"

	"github.com/consensys/quorum/common"
	"github.com/lugondev/signer-hashicorp-vault-plugin/src/vault/entities"
)

func FakeETHAccount() *entities.ETHAccount {
	return &entities.ETHAccount{
		Address:             common.HexToAddress(randHexString(12)).String(),
		PublicKey:           common.HexToHash(randHexString(12)).String(),
		CompressedPublicKey: common.HexToHash(randHexString(12)).String(),
		Namespace:           "_",
	}
}

func FakeKey() *entities.Key {
	return &entities.Key{
		Algorithm:  entities.EDDSA,
		Curve:      entities.Babyjubjub,
		PublicKey:  "X9Yz_5-O42-eOodHCUBhA4VMD2ZQy5CMAQ6lXqvDUZE=",
		PrivateKey: "X9Yz_5-O42-eOodHCUBhA4VMD2ZQy5CMAQ6lXqvDUZGGbioek5qYuzJzTNZpTHrVjjFk7iFe3FYwfpxZyNPxtIaFB5gb9VP9IcHZewwNZly821re7RkmB8pGdjywygPH",
		Namespace:  "_",
		ID:         "my-key",
		Tags: map[string]string{
			"tag1": "tagValue1",
			"tag2": "tagValue2",
		},
		Version:   1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func randHexString(n int) string {
	var letterRunes = []rune("abcdef0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

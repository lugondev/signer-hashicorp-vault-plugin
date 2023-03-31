package ethereum

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

func SignTransaction(tx *types.Transaction, privKey *ecdsa.PrivateKey, signer types.Signer) ([]byte, error) {
	h := signer.Hash(tx)
	decodedSignature, err := crypto.Sign(h[:], privKey)
	if err != nil {
		return nil, err
	}

	return decodedSignature, nil
}

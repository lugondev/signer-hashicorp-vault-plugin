package wallets

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	crypto2 "github.com/ethereum/go-ethereum/crypto"
	"github.com/lugondev/signer-hashicorp-vault-plugin/src/pkg/errors"

	"github.com/lugondev/signer-hashicorp-vault-plugin/src/pkg/log"
	"github.com/lugondev/signer-hashicorp-vault-plugin/src/vault/entities"
	"github.com/lugondev/signer-hashicorp-vault-plugin/src/vault/storage"
	usecases "github.com/lugondev/signer-hashicorp-vault-plugin/src/vault/use-cases"

	"github.com/hashicorp/vault/sdk/logical"

	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/ethereum/go-ethereum/crypto"
)

// createWalletUseCase is a use case to create a new wallet
type createWalletUseCase struct {
	storage logical.Storage
}

// NewCreateWalletUseCase creates a new CreateWalletUseCase
func NewCreateWalletUseCase() usecases.CreateWalletUseCase {
	return &createWalletUseCase{}
}

func (uc *createWalletUseCase) WithStorage(storage logical.Storage) usecases.CreateWalletUseCase {
	uc.storage = storage
	return uc
}

// Execute creates a new wallet and stores it in the Vault
func (uc *createWalletUseCase) Execute(ctx context.Context, namespace, importedPrivKey string) (*entities.Wallet, error) {
	logger := log.FromContext(ctx).With("namespace", namespace)
	logger.Debug("creating new wallet")

	var privKey = new(ecdsa.PrivateKey)
	var err error
	if importedPrivKey == "" {
		privKey, err = crypto.GenerateKey()
		if err != nil {
			errMessage := "failed to generate private key"
			logger.With("error", err).Error(errMessage)
			return nil, errors.CryptoOperationError(errMessage)
		}
	} else {
		privKey, err = crypto.HexToECDSA(importedPrivKey)
		if err != nil {
			errMessage := "failed to import private key, please verify that the provided private key is valid"
			logger.With("error", err).Error(errMessage)
			return nil, errors.InvalidParameterError(errMessage)
		}
	}
	typeWallet := "imported"
	if importedPrivKey == "" {
		typeWallet = "generated"
	}

	wallet := &entities.Wallet{
		PrivateKey:          hex.EncodeToString(crypto.FromECDSA(privKey)),
		PublicKey:           hexutil.Encode(crypto.FromECDSAPub(&privKey.PublicKey)),
		CompressedPublicKey: hexutil.Encode(crypto2.CompressPubkey(&privKey.PublicKey)),
		Namespace:           namespace,
		Type:                typeWallet,
	}

	err = storage.StoreJSON(ctx, uc.storage, storage.ComputeWalletsStorageKey(wallet.CompressedPublicKey, wallet.Namespace), wallet)
	if err != nil {
		return nil, err
	}

	logger.With("compressedPublicKey", wallet.CompressedPublicKey).Info("wallet created successfully")
	return wallet, nil
}

package wallets

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"github.com/lugondev/signer-hashicorp-vault-plugin/src/pkg/errors"

	"github.com/lugondev/signer-hashicorp-vault-plugin/src/pkg/log"
	"github.com/lugondev/signer-hashicorp-vault-plugin/src/vault/entities"
	"github.com/lugondev/signer-hashicorp-vault-plugin/src/vault/storage"
	usecases "github.com/lugondev/signer-hashicorp-vault-plugin/src/vault/use-cases"

	"github.com/hashicorp/vault/sdk/logical"

	"github.com/consensys/quorum/common/hexutil"

	"github.com/consensys/quorum/crypto"
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

	account := &entities.Wallet{
		PrivateKey:          hex.EncodeToString(crypto.FromECDSA(privKey)),
		PublicKey:           hexutil.Encode(crypto.FromECDSAPub(&privKey.PublicKey)),
		CompressedPublicKey: hexutil.Encode(crypto.CompressPubkey(&privKey.PublicKey)),
		Namespace:           namespace,
	}

	err = storage.StoreJSON(ctx, uc.storage, storage.ComputeWalletsStorageKey(account.CompressedPublicKey, account.Namespace), account)
	if err != nil {
		return nil, err
	}

	logger.With("compressedPublicKey", account.CompressedPublicKey).Info("Wallet created successfully")
	return account, nil
}
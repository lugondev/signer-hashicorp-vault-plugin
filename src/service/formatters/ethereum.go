package formatters

import (
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/lugondev/signer-hashicorp-vault-plugin/src/pkg/errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
	"github.com/lugondev/signer-hashicorp-vault-plugin/src/vault/entities"
)

func FormatAccountResponse(account *entities.ETHAccount) *logical.Response {
	return &logical.Response{
		Data: map[string]interface{}{
			AddressLabel:             account.Address,
			PublicKeyLabel:           account.PublicKey,
			CompressedPublicKeyLabel: account.CompressedPublicKey,
			NamespaceLabel:           account.Namespace,
		},
	}
}

func FormatSignETHTransactionRequest(requestData *framework.FieldData) (*types.Transaction, error) {
	amount, ok := new(big.Int).SetString(requestData.Get(AmountLabel).(string), 10)
	if !ok {
		return nil, errors.InvalidFormatError("invalid amount")
	}

	gasPrice, ok := new(big.Int).SetString(requestData.Get(GasPriceLabel).(string), 10)
	if !ok {
		return nil, errors.InvalidFormatError("invalid gas price")
	}

	data, err := hexutil.Decode(requestData.Get(DataLabel).(string))
	if err != nil {
		return nil, errors.InvalidFormatError("invalid data")
	}

	nonce := requestData.Get(NonceLabel).(int)
	gasLimit := requestData.Get(GasLimitLabel).(int)
	to := requestData.Get(ToLabel).(string)
	if to == "" {
		return types.NewTx(&types.LegacyTx{
			Nonce:    uint64(nonce),
			GasPrice: gasPrice,
			Gas:      uint64(gasLimit),
			Value:    amount,
			Data:     data,
		}), nil
	}

	toAddress := common.HexToAddress(to)
	return types.NewTx(&types.LegacyTx{
		Nonce:    uint64(nonce),
		GasPrice: gasPrice,
		Gas:      uint64(gasLimit),
		To:       &toAddress,
		Value:    amount,
		Data:     data,
	}), nil
}

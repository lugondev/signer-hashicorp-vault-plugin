package errors

import (
	"errors"
	"github.com/hashicorp/vault/sdk/logical"
	pkgerrors "github.com/lugondev/signer-hashicorp-vault-plugin/src/pkg/errors"
)

func ParseHTTPError(err error) (*logical.Response, error) {
	switch {
	case pkgerrors.IsNotFoundError(err):
		return logical.ErrorResponse(err.Error()), logical.ErrUnsupportedPath
	case
		pkgerrors.IsInvalidFormatError(err),
		pkgerrors.IsInvalidParameterError(err),
		pkgerrors.IsEncodingError(err),
		pkgerrors.IsAlreadyExistsError(err):
		return logical.ErrorResponse(err.Error()), logical.ErrInvalidRequest
	default:
		return nil, errors.New("internal server error. Please retry or contact an administrator")
	}
}

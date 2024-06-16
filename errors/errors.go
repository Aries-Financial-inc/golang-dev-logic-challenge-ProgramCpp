package errors

import "errors"

var (
	ErrInvalidOptionsType    = errors.New("invalid option type")
	ErrInvalidStrikePrice    = errors.New("invalid strike price")
	ErrInvalidAskPrice       = errors.New("invalid ask price")
	ErrInvalidBidPrice       = errors.New("invalid bid price")
	ErrAskBidMismatch        = errors.New("ask price must be greater than bid price")
	ErrInvalidExpirationDate = errors.New("invalid expiration date")
	ErrInvalidLongShort      = errors.New("invalid longShort")
)

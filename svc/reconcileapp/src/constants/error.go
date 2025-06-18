package constants

import (
	"errors"

	"github.com/hasemeneh/PoC-OnlineStore/helper/response"
)

var (
	KeyNotFoundError       = errors.New("Key Not Found")
	UserNotFoundError      = response.NewResponseError("User Not Found", 404)
	InsufficientSaldoError = response.NewResponseError("Insufficient Saldo", 400)
)

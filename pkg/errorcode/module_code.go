package errorcode

import (
	"net/http"

	hwikerror "github.com/hwikpass/hwik-go/core/error"
)

var (
	AlternativesMissingRequiredParam = hwikerror.NewCustomError(http.StatusOK, "FMSADM0001", "필수값 누락")
	ErrorRentalServiceCall           = hwikerror.NewCustomError(http.StatusOK, "FMSADM0002", "error occurred when call Rental service")
)

package app

import (
	"strings"

	"github.com/beego/beego/v2/core/validation"
	"github.com/hwikpass/encryptor/logger"
	errHttp "github.com/hwikpass/hwik-go/core/error"
)

// ValidError ValidError struct.
type ValidError struct {
	Message string
}

// ValidErrors error array.
type ValidErrors []*ValidError

// Error : get error message.
func (v *ValidError) Error() string {
	return v.Message
}

func (v ValidErrors) Error() string {
	return strings.Join(v.Errors(), ",")
}

// Errors : get error array.
func (v ValidErrors) Errors() []string {
	var errs []string
	for _, err := range v {
		errs = append(errs, err.Error())
	}

	return errs
}

// Validate : validate paramter.
func Validate(v interface{}) error {
	valid := validation.Validation{}
	b, err := valid.Valid(v)
	if err != nil {
		logger.Warn(err)
		return errHttp.ErrorValidation.WithDetails(err.Error())
	}

	if !b {
		for _, err := range valid.Errors {
			logger.Warn(err)
			return errHttp.ErrorValidation.WithDetails(err.Error())
		}
	}

	return nil
}

package app

import (
	"github.com/hwikpass/encryptor/logger"
	"github.com/labstack/echo/v4"
)

// Bind : Bind request body to struct.
func Bind(c echo.Context, v interface{}) error {
	if err := c.Bind(v); err != nil {
		logger.Warn(err)
		return err
	}

	return nil
}

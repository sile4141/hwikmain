package dao

import (
	"net/http/httptest"
	"os"

	app "github.com/hwikpass/hwik-go"
	"github.com/hwikpass/hwik-sample/model"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
)

type TestSuiteDao struct {
	suite.Suite
	e   *echo.Echo
	c   echo.Context
	dao *Dao
}

func (suite *TestSuiteDao) SetupSuite() {
	_ = os.Chdir("../")
	os.Setenv("ENV", "dev")
	os.Setenv("LOG_LEVEL", "DEBUG")

	suite.e, _ = app.Init(false, true)
	req := httptest.NewRequest(echo.POST, "/", nil)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	suite.c = suite.e.NewContext(req, rec)
	var err error
	suite.dao, err = NewTx(suite.c.Request().Context())
	if err != nil {
		panic(err)
	}

	model.InitConfig()
}

func (suite *TestSuiteDao) TearDownSuite() {
	defer func() {
		_ = suite.dao.Tx.Rollback()
	}()
}

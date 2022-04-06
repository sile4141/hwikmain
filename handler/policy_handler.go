// Package handler : 정책을 관리하는 handler.
package handler

import (
	"net/http"
	"strconv"

	"github.com/hwikpass/hwik-sample/model"

	hwikerror "github.com/hwikpass/hwik-go/core/error"
	res "github.com/hwikpass/hwik-go/core/response"
	"github.com/hwikpass/hwik-sample/pkg/app"
	"github.com/hwikpass/hwik-sample/pkg/util"
	"github.com/hwikpass/hwik-sample/service"
	"github.com/labstack/echo/v4"
)

// ListPolicy : 정책관리목록조회.
func ListPolicy() echo.HandlerFunc {
	return func(c echo.Context) error {
		param := &model.ReqPolicy{
			ServiceID:       c.QueryParam("serviceID"),
			EquipmentKindCD: c.QueryParam("equipmentKindCD"),
		}
		serviceRegionID, err := strconv.Atoi(c.QueryParam("serviceRegionID"))
		if err != nil {
			serviceRegionID = 0
		}
		param.ServiceRegionID = serviceRegionID
		if err = app.Validate(param); err != nil {
			return hwikerror.NewHTTPCustomError(err, &c)
		}

		svc, err := service.New(c.Request().Context())
		if err != nil {
			return hwikerror.NewHTTPCustomError(hwikerror.ErrorDatabase, &c)
		}

		result, err := svc.ListPolicy(param)
		if err != nil {
			return hwikerror.NewHTTPCustomError(err, &c)
		}

		return c.JSON(http.StatusOK, res.GetResponse(SC0001, result))
	}
}

// GetPolicy : 정책기준조회.
func GetPolicy() echo.HandlerFunc {
	return func(c echo.Context) error {
		param := &model.ReqPolicy{}
		param.ServiceID = c.QueryParam("serviceID")
		param.EquipmentKindCD = c.QueryParam("equipmentKindCD")
		serviceRegionID, err := strconv.Atoi(c.QueryParam("serviceRegionID"))
		if err != nil {
			serviceRegionID = 0
		}
		param.ServiceRegionID = serviceRegionID
		if err = app.Validate(param); err != nil {
			return hwikerror.NewHTTPCustomError(err, &c)
		}

		svc, err := service.New(c.Request().Context())
		if err != nil {
			return hwikerror.NewHTTPCustomError(hwikerror.ErrorDatabase, &c)
		}

		result, err := svc.GetPolicy(param)
		if err != nil {
			return hwikerror.NewHTTPCustomError(err, &c)
		}

		return c.JSON(http.StatusOK, res.GetResponse(SC0001, result))
	}
}

// CreatePolicy : 정책기준입력.
func CreatePolicy() echo.HandlerFunc {
	return func(c echo.Context) error {
		memberNumber := strconv.Itoa(util.GetMemberInfo(c).MemberNumber)
		param := &model.PolicyStandard{}
		if err := app.Bind(c, &param); err != nil {
			return hwikerror.NewHTTPCustomError(hwikerror.ErrorValidation.WithDetails(err.Error()), &c)
		}
		param.Creater = memberNumber
		param.Updater = memberNumber
		if err := app.Validate(param); err != nil {
			return hwikerror.NewHTTPCustomError(err, &c)
		}

		svc, err := service.NewTx(c.Request().Context())
		if err != nil {
			return hwikerror.NewHTTPCustomError(hwikerror.ErrorDatabase, &c)
		}
		defer svc.Rollback()
		svc.IMemberService = &service.MemberServiceImpl{}

		if err = svc.CreatePolicy(param); err != nil {
			return hwikerror.NewHTTPCustomError(err, &c)
		}
		if err = svc.Commit(); err != nil {
			return hwikerror.NewHTTPCustomError(err, &c)
		}

		return c.JSON(http.StatusOK, res.GetResponse(SC0001))
	}
}

// ModifyPolicy : 정책기준수정.
func ModifyPolicy() echo.HandlerFunc {
	return func(c echo.Context) error {
		memberNumber := strconv.Itoa(util.GetMemberInfo(c).MemberNumber)
		param := &model.PolicyStandard{}
		err := app.Bind(c, &param)
		if err != nil {
			return hwikerror.NewHTTPCustomError(hwikerror.ErrorValidation.WithDetails(err.Error()), &c)
		}
		param.Updater = memberNumber
		if err = app.Validate(param); err != nil {
			return hwikerror.NewHTTPCustomError(err, &c)
		}

		svc, err := service.NewTx(c.Request().Context())
		if err != nil {
			return hwikerror.NewHTTPCustomError(hwikerror.ErrorDatabase, &c)
		}
		defer svc.Rollback()

		if err = svc.ModifyPolicy(param); err != nil {
			return hwikerror.NewHTTPCustomError(err, &c)
		}
		if err = svc.Commit(); err != nil {
			return hwikerror.NewHTTPCustomError(err, &c)
		}

		return c.JSON(http.StatusOK, res.GetResponse(SC0001))
	}
}

// DeletePolicy : 정책기준삭제.
func DeletePolicy() echo.HandlerFunc {
	return func(c echo.Context) error {
		param := &model.PolicyStandard{
			ServiceID:       c.QueryParam("serviceID"),
			EquipmentKindCD: c.QueryParam("equipmentKindCD"),
		}
		serviceRegionID, err := strconv.Atoi(c.QueryParam("serviceRegionID"))
		if err != nil {
			serviceRegionID = 0
		}
		param.ServiceRegionID = serviceRegionID
		if err = app.Validate(param); err != nil {
			return hwikerror.NewHTTPCustomError(err, &c)
		}

		svc, err := service.NewTx(c.Request().Context())
		if err != nil {
			return hwikerror.NewHTTPCustomError(hwikerror.ErrorDatabase, &c)
		}
		defer svc.Rollback()

		if err = svc.DeletePolicy(param); err != nil {
			return hwikerror.NewHTTPCustomError(err, &c)
		}
		if err = svc.Commit(); err != nil {
			return hwikerror.NewHTTPCustomError(err, &c)
		}

		return c.JSON(http.StatusOK, res.GetResponse(SC0001))
	}
}

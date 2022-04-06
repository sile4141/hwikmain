package router

import (
	interceptor "github.com/hwikpass/hwik-go/core/interceptor"
	"github.com/hwikpass/hwik-sample/handler"
	"github.com/labstack/echo/v4"
)

// SetPolicyRouter : 정책관리 라우터.
func SetPolicyRouter(e *echo.Echo) {
	// 정책관리목록조회.
	e.GET("/v2/sample/e/policy/standards", handler.ListPolicy(), interceptor.GetMemberWithToken)
	// 정책기준조회.
	e.GET("/v2/sample/e/policy/standard", handler.GetPolicy(), interceptor.GetMemberWithToken)
	// 정책기준입력.
	e.POST("/v2/sample/e/policy/standard", handler.CreatePolicy(), interceptor.GetMemberWithToken)
	// 정책기준수정.
	e.PUT("/v2/sample/e/policy/standard", handler.ModifyPolicy(), interceptor.GetMemberWithToken)
	// 정책기준삭제.
	e.DELETE("/v2/sample/e/policy/standard", handler.DeletePolicy(), interceptor.GetMemberWithToken)
}

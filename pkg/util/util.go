package util

import (
	"github.com/labstack/echo/v4"
)

// Member : member info. GetMemberWithToken 미들웨어 실행 후 사용자 정보가 저장된다.
type Member struct {
	MemberNumber int    `json:"memberNumber"`
	MemberNm     string `json:"memberNm"`
	ServiceID    string `json:"serviceId"`
	Result       string `json:"result"`
}

func GetMemberInfo(c echo.Context) Member {
	return c.Get("memberInfo").(Member)
}

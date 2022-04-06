package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	apmmanager "github.com/hwikpass/hwik-go/core/apm"
	"github.com/hwikpass/hwik-go/core/logger"
	"github.com/hwikpass/hwik-sample/model"
	"github.com/hwikpass/hwik-sample/pkg/errorcode"

	hwikerror "github.com/hwikpass/hwik-go/core/error"
)

// MemberServiceInterface : 회원 서비스 Interface.
type MemberServiceInterface interface {
	// CallRentalSuccess : 대여서비스에 대여성공 메시지 호출.
	CallRentalSuccess(param *model.IotPolicy) error
}

// MemberServiceImpl : 회원 서비스 구현체.
type MemberServiceImpl struct {
	Svc *Service
}

// CallRentalSuccess : 대여서비스에 대여성공 메시지 호출.
func (m *MemberServiceImpl) CallRentalSuccess(param *model.IotPolicy) error {
	rentalURL := "https://i-gn-dev.hwikx.com"

	url := fmt.Sprintf("%s/v2/rental/i/rentals", rentalURL)

	body, _ := json.Marshal(param)

	logger.Debug(fmt.Sprintf("CallRentalSuccess param: %s", string(body)))
	req, err := http.NewRequestWithContext(m.Svc.Ctx, http.MethodPut, url, bytes.NewBuffer(body))
	if err != nil {
		return hwikerror.InternalErrorWrap(err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := apmmanager.HTTPCallWithAPM()

	res, err := client.Do(req)
	if err != nil {
		return hwikerror.InternalErrorWrap(err)
	}
	defer func() { _ = res.Body.Close() }()

	var resultData map[string]interface{}
	respBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return hwikerror.InternalErrorWrap(err)
	}

	if err = json.Unmarshal(respBody, &resultData); err != nil {
		return hwikerror.InternalErrorWrap(err)
	}

	logger.Debug(fmt.Sprintf("CallRentalSuccess result: %+v", resultData))

	if res.StatusCode != 200 {
		return errorcode.ErrorRentalServiceCall.WithDetails(fmt.Sprintf("%v", resultData["detailMessage"]))
	}

	return nil
}

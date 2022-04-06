package code

// CheckReason : 점검사유 코드.
var CheckReason = newCheckReason()

func newCheckReason() *checkReason {
	return &checkReason{
		ByCommError:    "1",  // 통신 실패.
		ByAdmin:        "2",  // 관리자 처리.
		ByErrorCode:    "3",  // 기기오류코드에 따른 점검 처리.
		ByRelocation:   "20", // 이동배치.
		ByRepair:       "21", // 입고.
		ByUnrepairable: "22", // 수리불가.
	}
}

type checkReason struct {
	ByCommError    string
	ByAdmin        string
	ByErrorCode    string
	ByRelocation   string
	ByRepair       string
	ByUnrepairable string
}

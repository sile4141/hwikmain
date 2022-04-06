package service

import (
	"github.com/golang/mock/gomock"
	mock_dao "github.com/hwikpass/hwik-sample/mocks/dao"
	mock_service "github.com/hwikpass/hwik-sample/mocks/service"
	"github.com/hwikpass/hwik-sample/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/volatiletech/null/v8"
)

type PolicyTestSuite struct {
	suite.Suite
	ctrl              *gomock.Controller
	svc               *Service
	mockDao           *mock_dao.MockDaoInterface
	mockMemberService *mock_service.MockMemberServiceInterface
}

func (suite *PolicyTestSuite) SetupTest() {
	suite.ctrl = gomock.NewController(suite.T())
	suite.svc = &Service{}
}

func (suite *PolicyTestSuite) BeforeTest(suiteName, testName string) {
	suite.mockDao = mock_dao.NewMockDaoInterface(suite.ctrl)
	suite.svc.IDao = suite.mockDao
	suite.mockMemberService = mock_service.NewMockMemberServiceInterface(suite.ctrl)
	suite.svc.IMemberService = suite.mockMemberService
}

func (suite *PolicyTestSuite) TestListPolicy_성공_정책관리목록조회() {
	suite.mockDao.EXPECT().ListPolicy(gomock.Any()).Return([]*model.ResPolicies{
		{
			ServiceID:            "hwikgo",
			ServiceRegionID:      1,
			ServiceRegionName:    "서울",
			EquipmentKindCD:      "15",
			EquipmentKindCDName:  "킥보드",
			PolicyUpdateDatetime: "2022-01-01 00:00:00",
			LastUpdateDatetime:   null.NewString("2022-01-01 00:00:00", true),
		},
	}, nil).AnyTimes()
	suite.mockDao = getDaoMock(suite.mockDao)

	_, err := suite.svc.ListPolicy(&model.ReqPolicy{
		ServiceID:       "hwikgo",
		ServiceRegionID: 1,
		EquipmentKindCD: "15",
	})
	assert.Nil(suite.T(), err)
}

func (suite *PolicyTestSuite) TestGetPolicy_성공_정책기준조회() {
	suite.mockDao.EXPECT().GetPolicy(gomock.Any(), gomock.Any(), gomock.Any()).Return(&model.PolicyStandard{
		ServiceID:                    "hwikgo",
		ServiceRegionID:              1,
		EquipmentKindCD:              "13",
		TripStoreSecond:              3,
		RentalSendSecond:             60,
		StandbySendSecond:            900,
		DayMaxSpeed:                  20,
		NightMaxSpeed:                20,
		LowbatterySendYn:             "Y",
		FalldownSendYn:               "Y",
		ServiceLowbattery:            30,
		DrivingLowbatteryAlert:       "20",
		DrivingLowbattery:            10,
		ValiditytimeSecond:           30,
		OpenTime:                     "070000",
		CloseTime:                    "230000",
		OutMaxSpeed:                  15,
		OutNotificationSecond:        10,
		OutSendSecond:                10,
		LogLevelCD:                   "D",
		LightOnTime:                  "180000",
		LightOffTime:                 "070000",
		IndicatorOnTime:              "180000",
		IndicatorOffTime:             "070000",
		NightStartTime:               "180001",
		NightEndTime:                 "065959",
		UvaSecond:                    300,
		UvcSecond:                    600,
		TheftAccel:                   5,
		TheftSendSecond:              60,
		TheftPlaySecond:              20,
		TheftBeepPreset:              "030001",
		TheftHeadlampPreset:          "030001",
		TheftIndicatorPreset:         "030001",
		TheftIndicatorColor:          "800",
		FalldownAngle:                70,
		FalldownCheckSecond:          3,
		LowbatteryMaxSpeed:           5,
		RentalSuccessIndicatorPreset: "030001",
		RentalSuccessColor:           "800",
		RentalSuccessBeep:            "010001",
		RentalFailIndicatorPreset:    "030001",
		RentalFailIndicatorColor:     "810",
		RentalFailBeepPreset:         "020001",
		ReturnSuccessAudio:           "returnSucc",
		ReturnSuccessIndicatorPreset: "030001",
		ReturnSuccessIndicatorColor:  "800",
		ReturnFailAudio:              "returnFail",
		ReturnFailIndicatorPreset:    "030001",
		ReturnFailIndicatorColor:     "800",
		StandbyIndicatorPreset:       "03020C",
		StandbyIndicatorColor:        "600",
		NonmovementDetectSecond:      60,
	}, nil).AnyTimes()

	suite.mockDao = getDaoMock(suite.mockDao)
	_, err := suite.svc.GetPolicy(&model.ReqPolicy{
		ServiceID:       "hwikgo",
		ServiceRegionID: 1,
		EquipmentKindCD: "15",
	})
	assert.Nil(suite.T(), err)
}

func (suite *PolicyTestSuite) TestCreatePolicy_성공_정책기준입력() {
	suite.mockDao.EXPECT().CreatePolicy(gomock.Any()).Return(nil).AnyTimes()
	suite.mockDao.EXPECT().ModifyPolicy(gomock.Any()).Return(nil).AnyTimes()
	suite.mockMemberService.EXPECT().CallRentalSuccess(gomock.Any()).Return(nil).AnyTimes()
	suite.mockDao.EXPECT().GetPolicy("hwikgo", 1, "15").Return(&model.PolicyStandard{
		ServiceID:       "hwikgo",
		ServiceRegionID: 1,
		EquipmentKindCD: "15",
	}, nil)

	suite.mockDao = getDaoMock(suite.mockDao)
	err := suite.svc.CreatePolicy(&model.PolicyStandard{
		ServiceID:       "hwikgo",
		ServiceRegionID: 1,
		EquipmentKindCD: "15",
	})
	assert.Nil(suite.T(), err)
}

func (suite *PolicyTestSuite) TestCreatePolicy_성공_정책기준수정() {
	suite.mockDao.EXPECT().CreatePolicy(gomock.Any()).Return(nil).AnyTimes()
	suite.mockDao.EXPECT().ModifyPolicy(gomock.Any()).Return(nil).AnyTimes()
	suite.mockMemberService.EXPECT().CallRentalSuccess(gomock.Any()).Return(nil).AnyTimes()

	suite.mockDao.EXPECT().GetPolicy("hwikgo", 1, "15").Return(nil, nil)
	suite.mockDao = getDaoMock(suite.mockDao)

	err := suite.svc.CreatePolicy(&model.PolicyStandard{
		ServiceID:       "hwikgo",
		ServiceRegionID: 1,
		EquipmentKindCD: "15",
	})
	assert.Nil(suite.T(), err)

}

func (suite *PolicyTestSuite) TestModifyPolicy_성공_정책기준수정() {
	suite.mockDao = getDaoMock(suite.mockDao)
	err := suite.svc.ModifyPolicy(&model.PolicyStandard{})
	assert.Nil(suite.T(), err)
}

func (suite *PolicyTestSuite) TestDeletePolicy_성공_정책기준삭제() {
	suite.mockDao = getDaoMock(suite.mockDao)
	err := suite.svc.DeletePolicy(&model.PolicyStandard{})
	assert.Nil(suite.T(), err)
}

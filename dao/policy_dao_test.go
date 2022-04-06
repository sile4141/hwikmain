package dao

import (
	"github.com/hwikpass/hwik-sample/model"
	"github.com/stretchr/testify/assert"
)

func (suite *TestSuiteDao) TestListPolicy() {
	_, err := suite.dao.ListPolicy(&model.ReqPolicy{
		ServiceID:       "hwigo",
		ServiceRegionID: 1,
		EquipmentKindCD: "15",
	})
	assert.ErrorIs(suite.T(), nil, err)
}

func (suite *TestSuiteDao) TestGetPolicy() {
	_, err := suite.dao.GetPolicy("hwigo", 1, "15")
	assert.ErrorIs(suite.T(), nil, err)
}

func (suite *TestSuiteDao) TestCreatePolicy() {
	err := suite.dao.CreatePolicy(&model.PolicyStandard{
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
	})
	assert.ErrorIs(suite.T(), nil, err)
}

func (suite *TestSuiteDao) TestModifyPolicy() {
	err := suite.dao.ModifyPolicy(&model.PolicyStandard{
		ServiceID:                    "hwikgo",
		ServiceRegionID:              1,
		EquipmentKindCD:              "15",
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
	})
	assert.ErrorIs(suite.T(), nil, err)
}

func (suite *TestSuiteDao) TestDeletePolicy() {
	err := suite.dao.DeletePolicy(&model.PolicyStandard{
		ServiceID:       "hwikgo",
		ServiceRegionID: 1,
		EquipmentKindCD: "15"})
	assert.ErrorIs(suite.T(), nil, err)
}

package model

import (
	"github.com/volatiletech/null/v8"
)

type PolicyStandard struct {
	ServiceID                    string      `json:"serviceID" db:"service_id" valid:"MaxSize(30);Required"`
	ServiceRegionID              int         `json:"serviceRegionID" db:"service_region_id" valid:"Required"`
	EquipmentKindCD              string      `json:"equipmentKindCD" db:"equipment_kind_cd" valid:"MaxSize(20);Required"`
	TripStoreSecond              int         `json:"tripStoreSecond" db:"trip_store_second"`
	RentalSendSecond             int         `json:"rentalSendSecond" db:"rental_send_second"`
	StandbySendSecond            int         `json:"standbySendSecond" db:"standby_send_second"`
	DayMaxSpeed                  int         `json:"dayMaxSpeed" db:"day_max_speed"`
	NightMaxSpeed                int         `json:"nightMaxSpeed" db:"night_max_speed"`
	LowbatterySendYn             string      `json:"lowbatterySendYn" db:"lowbattery_send_yn" valid:"MaxSize(1);"`
	FalldownSendYn               string      `json:"falldownSendYn" db:"falldown_send_yn" valid:"MaxSize(1);"`
	ServiceLowbattery            int         `json:"serviceLowbattery" db:"service_lowbattery"`
	DrivingLowbatteryAlert       string      `json:"drivingLowbatteryAlert" db:"driving_lowbattery_alert" valid:"MaxSize(50);"`
	DrivingLowbattery            int         `json:"drivingLowbattery" db:"driving_lowbattery"`
	ValiditytimeSecond           int         `json:"validitytimeSecond" db:"validitytime_second"`
	OpenTime                     string      `json:"openTime"  db:"open_time" valid:"MaxSize(6);"`
	CloseTime                    string      `json:"closeTime"  db:"close_time" valid:"MaxSize(6);"`
	OutMaxSpeed                  int         `json:"outMaxSpeed"  db:"out_max_speed"`
	OutNotificationSecond        int         `json:"outNotificationSecond"  db:"out_notification_second"`
	OutSendSecond                int         `json:"outSendSecond" db:"out_send_second"`
	LogLevelCD                   string      `json:"logLevelCD" db:"log_level_cd" valid:"MaxSize(20);"`
	LightOnTime                  string      `json:"lightOnTime" db:"light_on_time" valid:"MaxSize(6);"`
	LightOffTime                 string      `json:"lightOffTime" db:"light_off_time" valid:"MaxSize(6);"`
	IndicatorOnTime              string      `json:"indicatorOnTime" db:"indicator_on_time" valid:"MaxSize(6);"`
	IndicatorOffTime             string      `json:"indicatorOffTime" db:"indicator_off_time" valid:"MaxSize(6);"`
	NightStartTime               string      `json:"nightStartTime" db:"night_start_time" valid:"MaxSize(6);"`
	NightEndTime                 string      `json:"nightEndTime" db:"night_end_time" valid:"MaxSize(6);"`
	UvaSecond                    int         `json:"uvaSecond" db:"uva_second"`
	UvcSecond                    int         `json:"uvcSecond" db:"uvc_second"`
	LowbatteryMaxSpeed           int         `json:"lowbatteryMaxSpeed" db:"lowbattery_max_speed"`
	TheftAccel                   int         `json:"theftAccel" db:"theft_accel"`
	TheftSendSecond              int         `json:"theftSendSecond" db:"theft_send_second"`
	TheftPlaySecond              int         `json:"theftPlaySecond" db:"theft_play_second"`
	TheftBeepPreset              string      `json:"theftBeepPreset" db:"theft_beep_preset" valid:"MaxSize(50);"`
	TheftHeadlampPreset          string      `json:"theftHeadlampPreset" db:"theft_headlamp_preset" valid:"MaxSize(50);"`
	TheftIndicatorPreset         string      `json:"theftIndicatorPreset" db:"theft_indicator_preset" valid:"MaxSize(50);"`
	TheftIndicatorColor          string      `json:"theftIndicatorColor" db:"theft_indicator_color" valid:"MaxSize(50);"`
	FalldownAngle                int         `json:"falldownAngle" db:"falldown_angle"`
	FalldownCheckSecond          int         `json:"falldownCheckSecond" db:"falldown_check_second"`
	RentalSuccessIndicatorPreset string      `json:"rentalSuccessIndicatorPreset" db:"rental_success_indicator_preset" valid:"MaxSize(50);"`
	RentalSuccessColor           string      `json:"rentalSuccessColor" db:"rental_success_color" valid:"MaxSize(50);"`
	RentalSuccessBeep            string      `json:"rentalSuccessBeep" db:"rental_success_beep" valid:"MaxSize(50);"`
	RentalFailIndicatorPreset    string      `json:"rentalFailIndicatorPreset" db:"rental_fail_indicator_preset" valid:"MaxSize(50);"`
	RentalFailIndicatorColor     string      `json:"rentalFailIndicatorColor" db:"rental_fail_indicator_color" valid:"MaxSize(50);"`
	RentalFailBeepPreset         string      `json:"rentalFailBeepPreset" db:"rental_fail_beep_preset" valid:"MaxSize(50);"`
	ReturnSuccessAudio           string      `json:"returnSuccessAudio" db:"return_success_audio" valid:"MaxSize(50);"`
	ReturnSuccessIndicatorPreset string      `json:"returnSuccessIndicatorPreset" db:"return_success_indicator_preset" valid:"MaxSize(50);"`
	ReturnSuccessIndicatorColor  string      `json:"returnSuccessIndicatorColor" db:"return_success_indicator_color" valid:"MaxSize(50);"`
	ReturnFailAudio              string      `json:"returnFailAudio" db:"return_fail_audio" valid:"MaxSize(50);"`
	ReturnFailIndicatorPreset    string      `json:"returnFailIndicatorPreset" db:"return_fail_indicator_preset" valid:"MaxSize(50);"`
	ReturnFailIndicatorColor     string      `json:"returnFailIndicatorColor" db:"return_fail_indicator_color" valid:"MaxSize(50);"`
	StandbyIndicatorPreset       string      `json:"standbyIndicatorPreset" db:"standby_indicator_preset" valid:"MaxSize(50);"`
	StandbyIndicatorColor        string      `json:"standbyIndicatorColor" db:"standby_indicator_color" valid:"MaxSize(50);"`
	NonmovementDetectSecond      int         `json:"nonmovementDetectSecond" db:"nonmovement_detect_second"`
	Creater                      string      `json:"creater" db:"creater" valid:"MaxSize(20);"`
	CreateDatetime               string      `json:"createDatetime" db:"create_datetime"`
	Updater                      string      `json:"updater" db:"updater" valid:"MaxSize(20);"`
	UpdateDatetime               string      `json:"updateDatetime" db:"update_datetime"`
	LastUpdateDatetime           null.String `json:"lastUpdateDatetime" db:"last_update_datetime"`
}

const InsertPolicyStandard = `
insert into policy_standard( service_id, service_region_id, equipment_kind_cd, trip_store_second, rental_send_second
                           , standby_send_second, day_max_speed, night_max_speed, lowbattery_send_yn, falldown_send_yn
                           , service_lowbattery, driving_lowbattery_alert, driving_lowbattery, validitytime_second
                           , open_time, close_time, out_max_speed, out_notification_second, out_send_second
                           , log_level_cd, light_on_time, light_off_time, indicator_on_time, indicator_off_time
                           , night_start_time, night_end_time, uva_second, uvc_second
                           , lowbattery_max_speed, theft_accel, theft_send_second, theft_play_second, theft_beep_preset
                           , theft_headlamp_preset, theft_indicator_preset, theft_indicator_color, falldown_angle
                           , falldown_check_second, rental_success_indicator_preset, rental_success_color
                           , rental_success_beep, rental_fail_indicator_preset, rental_fail_indicator_color
                           , rental_fail_beep_preset, return_success_audio, return_success_indicator_preset
                           , return_success_indicator_color, return_fail_audio, return_fail_indicator_preset
                           , return_fail_indicator_color, standby_indicator_color, standby_indicator_preset, nonmovement_detect_second, creater
                           , create_datetime, updater, update_datetime )
values ( :service_id, :service_region_id, :equipment_kind_cd, :trip_store_second, :rental_send_second, :standby_send_second
       , :day_max_speed, :night_max_speed, :lowbattery_send_yn, :falldown_send_yn, :service_lowbattery, :driving_lowbattery_alert
       , :driving_lowbattery, :validitytime_second, :open_time, :close_time, :out_max_speed, :out_notification_second, :out_send_second
       , :log_level_cd, :light_on_time, :light_off_time, :indicator_on_time, :indicator_off_time
       , :night_start_time, :night_end_time, :uva_second, :uvc_second, :lowbattery_max_speed, :theft_accel, :theft_send_second
       , :theft_play_second, :theft_beep_preset, :theft_headlamp_preset, :theft_indicator_preset, :theft_indicator_color
       , :falldown_angle, :falldown_check_second, :rental_success_indicator_preset, :rental_success_color, :rental_success_beep
       , :rental_fail_indicator_preset, :rental_fail_indicator_color, :rental_fail_beep_preset, :return_success_audio, :return_success_indicator_preset
       , :return_success_indicator_color, :return_fail_audio, :return_fail_indicator_preset, :return_fail_indicator_color
       , :standby_indicator_color, :standby_indicator_preset, :nonmovement_detect_second, :creater, now(), :updater, now() )
`

const UpdatePolicyStandard = `
update policy_standard
set trip_store_second               = :trip_store_second
  , rental_send_second              = :rental_send_second
  , standby_send_second             = :standby_send_second
  , day_max_speed                   = :day_max_speed
  , night_max_speed                 = :night_max_speed
  , lowbattery_send_yn              = :lowbattery_send_yn
  , falldown_send_yn                = :falldown_send_yn
  , service_lowbattery              = :service_lowbattery
  , driving_lowbattery_alert        = :driving_lowbattery_alert
  , driving_lowbattery              = :driving_lowbattery
  , validitytime_second             = :validitytime_second
  , open_time                       = :open_time
  , close_time                      = :close_time
  , out_max_speed                   = :out_max_speed
  , out_notification_second         = :out_notification_second
  , out_send_second                 = :out_send_second
  , log_level_cd                    = :log_level_cd
  , light_on_time                   = :light_on_time
  , light_off_time                  = :light_off_time
  , indicator_on_time               = :indicator_on_time
  , indicator_off_time              = :indicator_off_time
  , night_start_time                = :night_start_time
  , night_end_time                  = :night_end_time
  , uva_second                      = :uva_second
  , uvc_second                      = :uvc_second
  , lowbattery_max_speed            = :lowbattery_max_speed
  , theft_accel                     = :theft_accel
  , theft_send_second               = :theft_send_second
  , theft_play_second               = :theft_play_second
  , theft_beep_preset               = :theft_beep_preset
  , theft_headlamp_preset           = :theft_headlamp_preset
  , theft_indicator_preset          = :theft_indicator_preset
  , theft_indicator_color           = :theft_indicator_color
  , falldown_angle                  = :falldown_angle
  , falldown_check_second           = :falldown_check_second
  , rental_success_indicator_preset = :rental_success_indicator_preset
  , rental_success_color            = :rental_success_color
  , rental_success_beep             = :rental_success_beep
  , rental_fail_indicator_preset    = :rental_fail_indicator_preset
  , rental_fail_indicator_color     = :rental_fail_indicator_color
  , rental_fail_beep_preset         = :rental_fail_beep_preset
  , return_success_audio            = :return_success_audio
  , return_success_indicator_preset = :return_success_indicator_preset
  , return_success_indicator_color  = :return_success_indicator_color
  , return_fail_audio               = :return_fail_audio
  , return_fail_indicator_preset    = :return_fail_indicator_preset
  , return_fail_indicator_color     = :return_fail_indicator_color
  , standby_indicator_color         = :standby_indicator_color
  , standby_indicator_preset        = :standby_indicator_preset
  , nonmovement_detect_second       = :nonmovement_detect_second
  , updater                         = :updater
  , update_datetime                 = now()
where service_id = :service_id
  and service_region_id = :service_region_id
  and equipment_kind_cd = :equipment_kind_cd
`

const DeletePolicyStandard = `
delete from policy_standard
 where service_id = :service_id 
   and service_region_id = :service_region_id 
   and equipment_kind_cd = :equipment_kind_cd
`

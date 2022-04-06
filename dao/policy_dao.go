// Package dao: 정책을 관리하는 dao.
package dao

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/hwikpass/hwik-go/core/apm"
	hwikquerybuilder "github.com/hwikpass/hwik-query-builder"

	hwikerror "github.com/hwikpass/hwik-go/core/error"
	"github.com/hwikpass/hwik-go/core/logger"
	"github.com/hwikpass/hwik-sample/model"
)

// ListPolicy : 정책관리목록조회.
func (dao *Dao) ListPolicy(param *model.ReqPolicy) ([]*model.ResPolicies, error) {
	headQuery := `
select a.service_id
     , a.service_region_id
     , b.service_region_name
     , a.equipment_kind_cd
     , c.detail_cd_name as equipment_kind_cd_name
     , date_format(a.update_datetime, '%Y-%m-%d %H:%i:%s') as policy_update_datetime
     , date_format((select v1.create_datetime from policy_request_history v1 where v1.service_id = a.service_id and v1.service_region_id = a.service_region_id and v1.equipment_kind_cd = a.equipment_kind_cd order by v1.control_request_id desc limit 1 ), '%Y-%m-%d %H:%i:%s')  as last_update_datetime
  from policy_standard a
     , service_area b
     , detail_cd c`

	var hqb hwikquerybuilder.HwikQueryBuilder
	query := hqb.And(
		`
       a.service_id = b.service_id
   and a.service_region_id = b.service_region_id
   and c.group_cd = 'equipment_kind_cd'
   and c.detail_cd = a.equipment_kind_cd
`,
		hqb.NewExp("a.service_id", "=", param.ServiceID),
		hqb.NewExp("a.service_region_id", "=", param.ServiceRegionID),
		hqb.NewExp("a.equipment_kind_cd", "=", param.EquipmentKindCD),
	).OrderBy(
		hqb.NewOrder("a.service_id", "ASC"),
		hqb.NewOrder("a.service_region_id", "ASC"),
		hqb.NewOrder("a.equipment_kind_cd", "ASC"),
	).BindSQL(headQuery)

	logger.LoggingQuery(query)
	apm.RegistQueryWithContextToAPM(dao.Ctx, query, fmt.Sprintf("%+v", param))

	result := []*model.ResPolicies{}
	err := dao.GetDB().Select(&result, query)
	if err != nil {
		return nil, hwikerror.InternalErrorWrap(err)
	}

	return result, nil
}

// GetPolicy : 정책기준조회.
func (dao *Dao) GetPolicy(serviceID string, serviceRegionID int, equipmentKindCD string) (*model.PolicyStandard, error) {
	query := `
select a.service_id
     , a.service_region_id
     , a.equipment_kind_cd
     , a.trip_store_second
     , a.rental_send_second
     , a.standby_send_second
     , a.day_max_speed
     , a.night_max_speed
     , a.lowbattery_send_yn
     , a.falldown_send_yn
     , a.service_lowbattery
     , a.driving_lowbattery_alert
     , a.driving_lowbattery
     , a.validitytime_second
     , a.open_time
     , a.close_time
     , a.out_max_speed
     , a.out_notification_second
     , a.out_send_second
     , a.log_level_cd
     , a.light_on_time
     , a.light_off_time
     , a.indicator_on_time
     , a.indicator_off_time
     , a.night_start_time
     , a.night_end_time
     , a.uva_second
     , a.uvc_second
     , a.lowbattery_max_speed            
     , a.theft_accel                     
     , a.theft_send_second               
     , a.theft_play_second               
     , a.theft_beep_preset               
     , a.theft_headlamp_preset           
     , a.theft_indicator_preset          
     , a.theft_indicator_color           
     , a.falldown_angle                  
     , a.falldown_check_second           
     , a.rental_success_indicator_preset 
     , a.rental_success_color            
     , a.rental_success_beep             
     , a.rental_fail_indicator_preset    
     , a.rental_fail_indicator_color     
     , a.rental_fail_beep_preset         
     , a.return_success_audio            
     , a.return_success_indicator_preset 
     , a.return_success_indicator_color  
     , a.return_fail_audio               
     , a.return_fail_indicator_preset    
     , a.return_fail_indicator_color    
     , a.standby_indicator_color         
     , a.standby_indicator_preset   
     , a.nonmovement_detect_second   
     , date_format((select v1.create_datetime from policy_request_history v1 where v1.service_id = a.service_id and v1.service_region_id = a.service_region_id and v1.equipment_kind_cd = a.equipment_kind_cd order by v1.control_request_id desc limit 1 ), '%Y-%m-%d %H:%i:%s')  as last_update_datetime
  from policy_standard a
 where a.service_id = ?
   and a.service_region_id = ?
   and a.equipment_kind_cd = ?
`

	logger.LoggingQuery(query, serviceID, serviceRegionID, equipmentKindCD)
	apm.RegistQueryWithContextToAPM(dao.Ctx, query, fmt.Sprintf("%s %d %s", serviceID, serviceRegionID, equipmentKindCD))

	result := &model.PolicyStandard{}
	err := dao.GetDB().Get(result, query, serviceID, serviceRegionID, equipmentKindCD)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	} else if err != nil {
		return nil, hwikerror.InternalErrorWrap(err)
	}

	return result, nil
}

// CreatePolicy : 정책기준입력.
func (dao *Dao) CreatePolicy(param *model.PolicyStandard) error {
	query := model.InsertPolicyStandard
	logger.LoggingQuery(query, param)
	apm.RegistQueryWithContextToAPM(dao.Ctx, query, fmt.Sprintf("%+v", param))

	result, err := dao.GetDB().NamedExec(query, param)
	if err != nil {
		return hwikerror.InternalErrorWrap(err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return hwikerror.InternalErrorWrap(err)
	} else if rowsAffected == 0 {
		return hwikerror.AlternativesRowsAffectedZero
	}
	return nil
}

// ModifyPolicy : 정책기준수정.
func (dao *Dao) ModifyPolicy(param *model.PolicyStandard) error {
	query := model.InsertPolicyStandard
	logger.LoggingQuery(query, param)
	apm.RegistQueryWithContextToAPM(dao.Ctx, query, fmt.Sprintf("%+v", param))

	result, err := dao.GetDB().NamedExec(query, param)
	if err != nil {
		return hwikerror.InternalErrorWrap(err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return hwikerror.InternalErrorWrap(err)
	} else if rowsAffected == 0 {
		return hwikerror.AlternativesRowsAffectedZero
	}
	return nil
}

// DeletePolicy : 정책기준삭제.
func (dao *Dao) DeletePolicy(param *model.PolicyStandard) error {
	query := model.DeletePolicyStandard
	logger.LoggingQuery(query, param)
	apm.RegistQueryWithContextToAPM(dao.Ctx, query, fmt.Sprintf("%+v", param))

	result, err := dao.GetDB().NamedExec(query, param)
	if err != nil {
		return hwikerror.InternalErrorWrap(err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return hwikerror.InternalErrorWrap(err)
	} else if rowsAffected == 0 {
		return hwikerror.AlternativesRowsAffectedZero
	}
	return nil
}

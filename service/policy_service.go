// Package service : 정책을 관리하는 service.
package service

import (
	"github.com/hwikpass/hwik-sample/model"
)

// ListPolicy : 정책관리목록조회.
func (svc *Service) ListPolicy(param *model.ReqPolicy) ([]*model.ResPolicies, error) {
	return svc.IDao.ListPolicy(param)
}

// GetPolicy : 정책기준조회.
func (svc *Service) GetPolicy(param *model.ReqPolicy) (*model.PolicyStandard, error) {
	return svc.IDao.GetPolicy(param.ServiceID, param.ServiceRegionID, param.EquipmentKindCD)
}

// CreatePolicy : 정책기준입력.
func (svc *Service) CreatePolicy(param *model.PolicyStandard) error {
	policy, err := svc.IDao.GetPolicy(param.ServiceID, param.ServiceRegionID, param.EquipmentKindCD)
	if err != nil {
		return err
	}

	if policy == nil {
		if err := svc.IDao.CreatePolicy(param); err != nil {
			return err
		}
	} else {
		if err := svc.IDao.ModifyPolicy(param); err != nil {
			return err
		}
	}

	// 멤버서비스에 대여 호출.
	if err = svc.IMemberService.CallRentalSuccess(&model.IotPolicy{
		EquipmentID: 1,
	}); err != nil {
		return err
	}

	return nil
}

// ModifyPolicy : 정책기준수정.
func (svc *Service) ModifyPolicy(param *model.PolicyStandard) error {
	return svc.IDao.ModifyPolicy(param)
}

// DeletePolicy : 정책기준삭제.
func (svc *Service) DeletePolicy(param *model.PolicyStandard) error {
	return svc.IDao.DeletePolicy(param)
}

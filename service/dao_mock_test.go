package service

import (
	mock_dao "github.com/hwikpass/hwik-sample/mocks/dao"
	"github.com/hwikpass/hwik-sample/model"
)

func getDaoMock(mockDao *mock_dao.MockDaoInterface) *mock_dao.MockDaoInterface {
	mockDao.EXPECT().GetPolicy("", 0, "").Return(&model.PolicyStandard{}, nil).AnyTimes()
	mockDao.EXPECT().CreatePolicy(&model.PolicyStandard{}).Return(nil).AnyTimes()
	mockDao.EXPECT().ModifyPolicy(&model.PolicyStandard{}).Return(nil).AnyTimes()
	mockDao.EXPECT().DeletePolicy(&model.PolicyStandard{}).Return(nil).AnyTimes()
	return mockDao
}

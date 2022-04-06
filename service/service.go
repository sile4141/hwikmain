package service

import (
	"context"

	"github.com/hwikpass/hwik-go/core/logger"
	"github.com/hwikpass/hwik-sample/dao"
	"github.com/hwikpass/hwik-sample/pkg/aws"
	"go.mongodb.org/mongo-driver/mongo"
)

// Service : Service.
type Service struct {
	Ctx            context.Context
	Dao            *dao.Dao
	IDao           dao.InterfaceDao
	IAWSSDK        aws.SDKInterface
	IMemberService MemberServiceInterface
	Client         *mongo.Client
}

// New : 서비스생성.
func New(c context.Context) (*Service, error) {
	svc := &Service{Ctx: c}
	svc.IAWSSDK = &aws.SDKImpl{}
	var err error
	svc.Dao, err = dao.New(c)
	svc.IDao = svc.Dao
	if err != nil {
		return nil, err
	}

	return svc, nil
}

// NewTx : 서비스생성(트랜잭션 처리).
func NewTx(c context.Context) (*Service, error) {
	svc := &Service{Ctx: c}
	svc.IAWSSDK = &aws.SDKImpl{}
	var err error
	svc.Dao, err = dao.NewTx(c)
	svc.IDao = svc.Dao
	if err != nil {
		return nil, err
	}

	return svc, nil
}

// Rollback : Rollback.
func (svc *Service) Rollback() {
	if svc.Dao.Tx == nil {
		return
	}
	if err := svc.Dao.Tx.Rollback(); err != nil {
		if err.Error() == "sql: transaction has already been committed or rolled back" {
			// 아래 Info logging은 시스템 담당자의 필요에 의해 변경.
			logger.Info(err.Error())
			return
		}
		logger.Error(err)
	}
}

// Commit : Commit.
func (svc *Service) Commit() error {
	if svc.Dao.Tx == nil {
		return nil
	}
	if err := svc.Dao.Tx.Commit(); err != nil {
		if err.Error() == "sql: transaction has already been committed or rolled back" {
			// 아래 Info logging은 시스템 담당자의 필요에 의해 변경.
			logger.Info(err.Error())
			return nil
		}
		logger.Error(err)
		return err
	}
	return nil
}

// GetDB : Get Database.
func (svc *Service) GetDB() dao.DBTX {
	return svc.Dao.Queries.Db
}

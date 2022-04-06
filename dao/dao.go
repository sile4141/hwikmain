package dao

import (
	"context"
	"database/sql"

	"github.com/hwikpass/hwik-go/core/dbcp"
	"github.com/hwikpass/hwik-sample/model"
	"github.com/jmoiron/sqlx"
)

// Dao :.
type Dao struct {
	Ctx     context.Context
	Db      *sqlx.DB
	Tx      *sqlx.Tx
	Queries *Queries
}

// New :.
func New(c context.Context) (*Dao, error) {
	db, err := dbcp.GetConnection()
	if err != nil {
		return nil, err
	}
	queries := &Queries{
		Db: db,
	}

	return &Dao{Ctx: c, Queries: queries}, nil
}

// NewTx :.
func NewTx(c context.Context) (*Dao, error) {
	db, err := dbcp.GetConnection()
	if err != nil {
		return nil, err
	}

	tx, err := db.BeginTxx(c, nil)
	if err != nil {
		return nil, err
	}
	queries := &Queries{
		Db: tx,
	}

	return &Dao{Ctx: c, Queries: queries, Tx: tx}, nil
}

// GetDB :.
func (dao *Dao) GetDB() DBTX {
	return dao.Queries.Db
}

// DBTX :.
type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryxContext(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
	Get(dest interface{}, query string, args ...interface{}) error
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	NamedExec(query string, arg interface{}) (sql.Result, error)
	NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
	PrepareNamedContext(ctx context.Context, query string) (*sqlx.NamedStmt, error)
	NamedQuery(query string, arg interface{}) (*sqlx.Rows, error)
	BindNamed(query string, arg interface{}) (string, []interface{}, error)
	Select(dest interface{}, query string, args ...interface{}) error
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	Rebind(query string) string
}

// Queries :.
type Queries struct {
	Db DBTX
}

type InterfaceDao interface {
	// ListPolicy : 정책관리목록조회.
	ListPolicy(param *model.ReqPolicy) ([]*model.ResPolicies, error)
	// GetPolicy : 정책기준조회.
	GetPolicy(serviceID string, serviceRegionID int, equipmentKindCD string) (*model.PolicyStandard, error)
	// CreatePolicy : 정책기준입력.
	CreatePolicy(param *model.PolicyStandard) error
	// ModifyPolicy : 정책기준수정.
	ModifyPolicy(param *model.PolicyStandard) error
	// DeletePolicy : 정책기준삭제.
	DeletePolicy(param *model.PolicyStandard) error
}

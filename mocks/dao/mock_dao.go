// Code generated by MockGen. DO NOT EDIT.
// Source: ../dao/dao.go

// Package mock_dao is a generated GoMock package.
package mock_dao

import (
	context "context"
	sql "database/sql"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/hwikpass/hwik-sample/model"
	sqlx "github.com/jmoiron/sqlx"
)

// MockDBTX is a mock of DBTX interface.
type MockDBTX struct {
	ctrl     *gomock.Controller
	recorder *MockDBTXMockRecorder
}

// MockDBTXMockRecorder is the mock recorder for MockDBTX.
type MockDBTXMockRecorder struct {
	mock *MockDBTX
}

// NewMockDBTX creates a new mock instance.
func NewMockDBTX(ctrl *gomock.Controller) *MockDBTX {
	mock := &MockDBTX{ctrl: ctrl}
	mock.recorder = &MockDBTXMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDBTX) EXPECT() *MockDBTXMockRecorder {
	return m.recorder
}

// BindNamed mocks base method.
func (m *MockDBTX) BindNamed(query string, arg interface{}) (string, []interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BindNamed", query, arg)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].([]interface{})
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// BindNamed indicates an expected call of BindNamed.
func (mr *MockDBTXMockRecorder) BindNamed(query, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BindNamed", reflect.TypeOf((*MockDBTX)(nil).BindNamed), query, arg)
}

// Exec mocks base method.
func (m *MockDBTX) Exec(query string, args ...interface{}) (sql.Result, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Exec", varargs...)
	ret0, _ := ret[0].(sql.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exec indicates an expected call of Exec.
func (mr *MockDBTXMockRecorder) Exec(query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exec", reflect.TypeOf((*MockDBTX)(nil).Exec), varargs...)
}

// ExecContext mocks base method.
func (m *MockDBTX) ExecContext(arg0 context.Context, arg1 string, arg2 ...interface{}) (sql.Result, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ExecContext", varargs...)
	ret0, _ := ret[0].(sql.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExecContext indicates an expected call of ExecContext.
func (mr *MockDBTXMockRecorder) ExecContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecContext", reflect.TypeOf((*MockDBTX)(nil).ExecContext), varargs...)
}

// Get mocks base method.
func (m *MockDBTX) Get(dest interface{}, query string, args ...interface{}) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{dest, query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Get", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Get indicates an expected call of Get.
func (mr *MockDBTXMockRecorder) Get(dest, query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{dest, query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockDBTX)(nil).Get), varargs...)
}

// GetContext mocks base method.
func (m *MockDBTX) GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, dest, query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetContext indicates an expected call of GetContext.
func (mr *MockDBTXMockRecorder) GetContext(ctx, dest, query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, dest, query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContext", reflect.TypeOf((*MockDBTX)(nil).GetContext), varargs...)
}

// NamedExec mocks base method.
func (m *MockDBTX) NamedExec(query string, arg interface{}) (sql.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NamedExec", query, arg)
	ret0, _ := ret[0].(sql.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NamedExec indicates an expected call of NamedExec.
func (mr *MockDBTXMockRecorder) NamedExec(query, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NamedExec", reflect.TypeOf((*MockDBTX)(nil).NamedExec), query, arg)
}

// NamedExecContext mocks base method.
func (m *MockDBTX) NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NamedExecContext", ctx, query, arg)
	ret0, _ := ret[0].(sql.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NamedExecContext indicates an expected call of NamedExecContext.
func (mr *MockDBTXMockRecorder) NamedExecContext(ctx, query, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NamedExecContext", reflect.TypeOf((*MockDBTX)(nil).NamedExecContext), ctx, query, arg)
}

// NamedQuery mocks base method.
func (m *MockDBTX) NamedQuery(query string, arg interface{}) (*sqlx.Rows, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NamedQuery", query, arg)
	ret0, _ := ret[0].(*sqlx.Rows)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NamedQuery indicates an expected call of NamedQuery.
func (mr *MockDBTXMockRecorder) NamedQuery(query, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NamedQuery", reflect.TypeOf((*MockDBTX)(nil).NamedQuery), query, arg)
}

// PrepareContext mocks base method.
func (m *MockDBTX) PrepareContext(arg0 context.Context, arg1 string) (*sql.Stmt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrepareContext", arg0, arg1)
	ret0, _ := ret[0].(*sql.Stmt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PrepareContext indicates an expected call of PrepareContext.
func (mr *MockDBTXMockRecorder) PrepareContext(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrepareContext", reflect.TypeOf((*MockDBTX)(nil).PrepareContext), arg0, arg1)
}

// PrepareNamedContext mocks base method.
func (m *MockDBTX) PrepareNamedContext(ctx context.Context, query string) (*sqlx.NamedStmt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrepareNamedContext", ctx, query)
	ret0, _ := ret[0].(*sqlx.NamedStmt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PrepareNamedContext indicates an expected call of PrepareNamedContext.
func (mr *MockDBTXMockRecorder) PrepareNamedContext(ctx, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrepareNamedContext", reflect.TypeOf((*MockDBTX)(nil).PrepareNamedContext), ctx, query)
}

// QueryContext mocks base method.
func (m *MockDBTX) QueryContext(arg0 context.Context, arg1 string, arg2 ...interface{}) (*sql.Rows, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "QueryContext", varargs...)
	ret0, _ := ret[0].(*sql.Rows)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryContext indicates an expected call of QueryContext.
func (mr *MockDBTXMockRecorder) QueryContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryContext", reflect.TypeOf((*MockDBTX)(nil).QueryContext), varargs...)
}

// QueryRowContext mocks base method.
func (m *MockDBTX) QueryRowContext(arg0 context.Context, arg1 string, arg2 ...interface{}) *sql.Row {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "QueryRowContext", varargs...)
	ret0, _ := ret[0].(*sql.Row)
	return ret0
}

// QueryRowContext indicates an expected call of QueryRowContext.
func (mr *MockDBTXMockRecorder) QueryRowContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryRowContext", reflect.TypeOf((*MockDBTX)(nil).QueryRowContext), varargs...)
}

// QueryxContext mocks base method.
func (m *MockDBTX) QueryxContext(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "QueryxContext", varargs...)
	ret0, _ := ret[0].(*sqlx.Rows)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryxContext indicates an expected call of QueryxContext.
func (mr *MockDBTXMockRecorder) QueryxContext(ctx, query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryxContext", reflect.TypeOf((*MockDBTX)(nil).QueryxContext), varargs...)
}

// Rebind mocks base method.
func (m *MockDBTX) Rebind(query string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Rebind", query)
	ret0, _ := ret[0].(string)
	return ret0
}

// Rebind indicates an expected call of Rebind.
func (mr *MockDBTXMockRecorder) Rebind(query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rebind", reflect.TypeOf((*MockDBTX)(nil).Rebind), query)
}

// Select mocks base method.
func (m *MockDBTX) Select(dest interface{}, query string, args ...interface{}) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{dest, query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Select", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Select indicates an expected call of Select.
func (mr *MockDBTXMockRecorder) Select(dest, query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{dest, query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Select", reflect.TypeOf((*MockDBTX)(nil).Select), varargs...)
}

// SelectContext mocks base method.
func (m *MockDBTX) SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, dest, query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SelectContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// SelectContext indicates an expected call of SelectContext.
func (mr *MockDBTXMockRecorder) SelectContext(ctx, dest, query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, dest, query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectContext", reflect.TypeOf((*MockDBTX)(nil).SelectContext), varargs...)
}

// MockDaoInterface is a mock of DaoInterface interface.
type MockDaoInterface struct {
	ctrl     *gomock.Controller
	recorder *MockDaoInterfaceMockRecorder
}

// MockDaoInterfaceMockRecorder is the mock recorder for MockDaoInterface.
type MockDaoInterfaceMockRecorder struct {
	mock *MockDaoInterface
}

// NewMockDaoInterface creates a new mock instance.
func NewMockDaoInterface(ctrl *gomock.Controller) *MockDaoInterface {
	mock := &MockDaoInterface{ctrl: ctrl}
	mock.recorder = &MockDaoInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDaoInterface) EXPECT() *MockDaoInterfaceMockRecorder {
	return m.recorder
}

// CreatePolicy mocks base method.
func (m *MockDaoInterface) CreatePolicy(param *model.PolicyStandard) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePolicy", param)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreatePolicy indicates an expected call of CreatePolicy.
func (mr *MockDaoInterfaceMockRecorder) CreatePolicy(param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePolicy", reflect.TypeOf((*MockDaoInterface)(nil).CreatePolicy), param)
}

// DeletePolicy mocks base method.
func (m *MockDaoInterface) DeletePolicy(param *model.PolicyStandard) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePolicy", param)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePolicy indicates an expected call of DeletePolicy.
func (mr *MockDaoInterfaceMockRecorder) DeletePolicy(param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePolicy", reflect.TypeOf((*MockDaoInterface)(nil).DeletePolicy), param)
}

// GetPolicy mocks base method.
func (m *MockDaoInterface) GetPolicy(serviceID string, serviceRegionID int, equipmentKindCD string) (*model.PolicyStandard, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPolicy", serviceID, serviceRegionID, equipmentKindCD)
	ret0, _ := ret[0].(*model.PolicyStandard)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPolicy indicates an expected call of GetPolicy.
func (mr *MockDaoInterfaceMockRecorder) GetPolicy(serviceID, serviceRegionID, equipmentKindCD interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPolicy", reflect.TypeOf((*MockDaoInterface)(nil).GetPolicy), serviceID, serviceRegionID, equipmentKindCD)
}

// ListPolicy mocks base method.
func (m *MockDaoInterface) ListPolicy(param *model.ReqPolicy) ([]*model.ResPolicies, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPolicy", param)
	ret0, _ := ret[0].([]*model.ResPolicies)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPolicy indicates an expected call of ListPolicy.
func (mr *MockDaoInterfaceMockRecorder) ListPolicy(param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPolicy", reflect.TypeOf((*MockDaoInterface)(nil).ListPolicy), param)
}

// ModifyPolicy mocks base method.
func (m *MockDaoInterface) ModifyPolicy(param *model.PolicyStandard) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModifyPolicy", param)
	ret0, _ := ret[0].(error)
	return ret0
}

// ModifyPolicy indicates an expected call of ModifyPolicy.
func (mr *MockDaoInterfaceMockRecorder) ModifyPolicy(param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModifyPolicy", reflect.TypeOf((*MockDaoInterface)(nil).ModifyPolicy), param)
}

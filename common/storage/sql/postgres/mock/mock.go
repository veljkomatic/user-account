// Code generated by MockGen. DO NOT EDIT.
// Source: /Users/veljkomatic/user-account/common/storage/sql/postgres/interface.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	sql "database/sql"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	bun "github.com/uptrace/bun"
	schema "github.com/uptrace/bun/schema"
)

// MockIDB is a mock of IDB interface.
type MockIDB struct {
	ctrl     *gomock.Controller
	recorder *MockIDBMockRecorder
}

// MockIDBMockRecorder is the mock recorder for MockIDB.
type MockIDBMockRecorder struct {
	mock *MockIDB
}

// NewMockIDB creates a new mock instance.
func NewMockIDB(ctrl *gomock.Controller) *MockIDB {
	mock := &MockIDB{ctrl: ctrl}
	mock.recorder = &MockIDBMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIDB) EXPECT() *MockIDBMockRecorder {
	return m.recorder
}

// BeginTx mocks base method.
func (m *MockIDB) BeginTx(ctx context.Context, opts *sql.TxOptions) (bun.Tx, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BeginTx", ctx, opts)
	ret0, _ := ret[0].(bun.Tx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BeginTx indicates an expected call of BeginTx.
func (mr *MockIDBMockRecorder) BeginTx(ctx, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeginTx", reflect.TypeOf((*MockIDB)(nil).BeginTx), ctx, opts)
}

// Dialect mocks base method.
func (m *MockIDB) Dialect() schema.Dialect {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Dialect")
	ret0, _ := ret[0].(schema.Dialect)
	return ret0
}

// Dialect indicates an expected call of Dialect.
func (mr *MockIDBMockRecorder) Dialect() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Dialect", reflect.TypeOf((*MockIDB)(nil).Dialect))
}

// ExecContext mocks base method.
func (m *MockIDB) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ExecContext", varargs...)
	ret0, _ := ret[0].(sql.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExecContext indicates an expected call of ExecContext.
func (mr *MockIDBMockRecorder) ExecContext(ctx, query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecContext", reflect.TypeOf((*MockIDB)(nil).ExecContext), varargs...)
}

// NewAddColumn mocks base method.
func (m *MockIDB) NewAddColumn() *bun.AddColumnQuery {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewAddColumn")
	ret0, _ := ret[0].(*bun.AddColumnQuery)
	return ret0
}

// NewAddColumn indicates an expected call of NewAddColumn.
func (mr *MockIDBMockRecorder) NewAddColumn() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewAddColumn", reflect.TypeOf((*MockIDB)(nil).NewAddColumn))
}

// NewCreateIndex mocks base method.
func (m *MockIDB) NewCreateIndex() *bun.CreateIndexQuery {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewCreateIndex")
	ret0, _ := ret[0].(*bun.CreateIndexQuery)
	return ret0
}

// NewCreateIndex indicates an expected call of NewCreateIndex.
func (mr *MockIDBMockRecorder) NewCreateIndex() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewCreateIndex", reflect.TypeOf((*MockIDB)(nil).NewCreateIndex))
}

// NewCreateTable mocks base method.
func (m *MockIDB) NewCreateTable() *bun.CreateTableQuery {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewCreateTable")
	ret0, _ := ret[0].(*bun.CreateTableQuery)
	return ret0
}

// NewCreateTable indicates an expected call of NewCreateTable.
func (mr *MockIDBMockRecorder) NewCreateTable() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewCreateTable", reflect.TypeOf((*MockIDB)(nil).NewCreateTable))
}

// NewDelete mocks base method.
func (m *MockIDB) NewDelete() *bun.DeleteQuery {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewDelete")
	ret0, _ := ret[0].(*bun.DeleteQuery)
	return ret0
}

// NewDelete indicates an expected call of NewDelete.
func (mr *MockIDBMockRecorder) NewDelete() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewDelete", reflect.TypeOf((*MockIDB)(nil).NewDelete))
}

// NewDropColumn mocks base method.
func (m *MockIDB) NewDropColumn() *bun.DropColumnQuery {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewDropColumn")
	ret0, _ := ret[0].(*bun.DropColumnQuery)
	return ret0
}

// NewDropColumn indicates an expected call of NewDropColumn.
func (mr *MockIDBMockRecorder) NewDropColumn() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewDropColumn", reflect.TypeOf((*MockIDB)(nil).NewDropColumn))
}

// NewDropIndex mocks base method.
func (m *MockIDB) NewDropIndex() *bun.DropIndexQuery {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewDropIndex")
	ret0, _ := ret[0].(*bun.DropIndexQuery)
	return ret0
}

// NewDropIndex indicates an expected call of NewDropIndex.
func (mr *MockIDBMockRecorder) NewDropIndex() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewDropIndex", reflect.TypeOf((*MockIDB)(nil).NewDropIndex))
}

// NewDropTable mocks base method.
func (m *MockIDB) NewDropTable() *bun.DropTableQuery {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewDropTable")
	ret0, _ := ret[0].(*bun.DropTableQuery)
	return ret0
}

// NewDropTable indicates an expected call of NewDropTable.
func (mr *MockIDBMockRecorder) NewDropTable() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewDropTable", reflect.TypeOf((*MockIDB)(nil).NewDropTable))
}

// NewInsert mocks base method.
func (m *MockIDB) NewInsert() *bun.InsertQuery {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewInsert")
	ret0, _ := ret[0].(*bun.InsertQuery)
	return ret0
}

// NewInsert indicates an expected call of NewInsert.
func (mr *MockIDBMockRecorder) NewInsert() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewInsert", reflect.TypeOf((*MockIDB)(nil).NewInsert))
}

// NewRaw mocks base method.
func (m *MockIDB) NewRaw(query string, args ...interface{}) *bun.RawQuery {
	m.ctrl.T.Helper()
	varargs := []interface{}{query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "NewRaw", varargs...)
	ret0, _ := ret[0].(*bun.RawQuery)
	return ret0
}

// NewRaw indicates an expected call of NewRaw.
func (mr *MockIDBMockRecorder) NewRaw(query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewRaw", reflect.TypeOf((*MockIDB)(nil).NewRaw), varargs...)
}

// NewSelect mocks base method.
func (m *MockIDB) NewSelect() *bun.SelectQuery {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewSelect")
	ret0, _ := ret[0].(*bun.SelectQuery)
	return ret0
}

// NewSelect indicates an expected call of NewSelect.
func (mr *MockIDBMockRecorder) NewSelect() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewSelect", reflect.TypeOf((*MockIDB)(nil).NewSelect))
}

// NewTruncateTable mocks base method.
func (m *MockIDB) NewTruncateTable() *bun.TruncateTableQuery {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewTruncateTable")
	ret0, _ := ret[0].(*bun.TruncateTableQuery)
	return ret0
}

// NewTruncateTable indicates an expected call of NewTruncateTable.
func (mr *MockIDBMockRecorder) NewTruncateTable() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewTruncateTable", reflect.TypeOf((*MockIDB)(nil).NewTruncateTable))
}

// NewUpdate mocks base method.
func (m *MockIDB) NewUpdate() *bun.UpdateQuery {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewUpdate")
	ret0, _ := ret[0].(*bun.UpdateQuery)
	return ret0
}

// NewUpdate indicates an expected call of NewUpdate.
func (mr *MockIDBMockRecorder) NewUpdate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewUpdate", reflect.TypeOf((*MockIDB)(nil).NewUpdate))
}

// NewValues mocks base method.
func (m *MockIDB) NewValues(model interface{}) *bun.ValuesQuery {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewValues", model)
	ret0, _ := ret[0].(*bun.ValuesQuery)
	return ret0
}

// NewValues indicates an expected call of NewValues.
func (mr *MockIDBMockRecorder) NewValues(model interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewValues", reflect.TypeOf((*MockIDB)(nil).NewValues), model)
}

// QueryContext mocks base method.
func (m *MockIDB) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "QueryContext", varargs...)
	ret0, _ := ret[0].(*sql.Rows)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryContext indicates an expected call of QueryContext.
func (mr *MockIDBMockRecorder) QueryContext(ctx, query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryContext", reflect.TypeOf((*MockIDB)(nil).QueryContext), varargs...)
}

// QueryRowContext mocks base method.
func (m *MockIDB) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "QueryRowContext", varargs...)
	ret0, _ := ret[0].(*sql.Row)
	return ret0
}

// QueryRowContext indicates an expected call of QueryRowContext.
func (mr *MockIDBMockRecorder) QueryRowContext(ctx, query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryRowContext", reflect.TypeOf((*MockIDB)(nil).QueryRowContext), varargs...)
}

// RunInTx mocks base method.
func (m *MockIDB) RunInTx(ctx context.Context, opts *sql.TxOptions, f func(context.Context, bun.Tx) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunInTx", ctx, opts, f)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunInTx indicates an expected call of RunInTx.
func (mr *MockIDBMockRecorder) RunInTx(ctx, opts, f interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunInTx", reflect.TypeOf((*MockIDB)(nil).RunInTx), ctx, opts, f)
}


// MockQuery is a mock of Query interface.
type MockQuery struct {
	ctrl     *gomock.Controller
	recorder *MockQueryMockRecorder
}

// MockQueryMockRecorder is the mock recorder for MockQuery.
type MockQueryMockRecorder struct {
	mock *MockQuery
}

// NewMockQuery creates a new mock instance.
func NewMockQuery(ctrl *gomock.Controller) *MockQuery {
	mock := &MockQuery{ctrl: ctrl}
	mock.recorder = &MockQueryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQuery) EXPECT() *MockQueryMockRecorder {
	return m.recorder
}

func (m *MockQuery) Model(model interface{}) *MockQueryMockRecorder {
	m.ctrl.T.Helper()
	m.ctrl.RecordCallWithMethodType(m, "Model", reflect.TypeOf((*MockQuery)(nil).Model), model)
	return m.recorder
}

func (mr *MockQueryMockRecorder) Model(model interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Model", reflect.TypeOf((*MockQuery)(nil).Model), model)
}

func (mr *MockQueryMockRecorder) Where(query string, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := []interface{}{query}
	varargs = append(varargs, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Where", reflect.TypeOf((*MockQuery)(nil).Where), varargs...)
}

func (mr *MockQueryMockRecorder) Scan(ctx context.Context, dest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Scan", reflect.TypeOf((*MockQuery)(nil).Scan), ctx, dest)
}

func (m *MockQuery) Where(query string, args ...interface{}) *MockQueryMockRecorder {
	m.ctrl.T.Helper()
	m.ctrl.RecordCallWithMethodType(m, "Where", reflect.TypeOf((*MockQuery)(nil).Where), query, args)
	return m.recorder
}

func (m *MockQuery) Scan(ctx context.Context, dest interface{}) *MockQueryMockRecorder {
	m.ctrl.T.Helper()
	m.ctrl.RecordCallWithMethodType(m, "Scan", reflect.TypeOf((*MockQuery)(nil).Scan), ctx, dest)
	return m.recorder
}

// AppendQuery mocks base method.
func (m *MockQuery) AppendQuery(fmter schema.Formatter, b []byte) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppendQuery", fmter, b)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AppendQuery indicates an expected call of AppendQuery.
func (mr *MockQueryMockRecorder) AppendQuery(fmter, b interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppendQuery", reflect.TypeOf((*MockQuery)(nil).AppendQuery), fmter, b)
}

// GetModel mocks base method.
func (m *MockQuery) GetModel() schema.Model {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetModel")
	ret0, _ := ret[0].(schema.Model)
	return ret0
}

// GetModel indicates an expected call of GetModel.
func (mr *MockQueryMockRecorder) GetModel() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetModel", reflect.TypeOf((*MockQuery)(nil).GetModel))
}

// GetTableName mocks base method.
func (m *MockQuery) GetTableName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTableName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetTableName indicates an expected call of GetTableName.
func (mr *MockQueryMockRecorder) GetTableName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTableName", reflect.TypeOf((*MockQuery)(nil).GetTableName))
}

// Operation mocks base method.
func (m *MockQuery) Operation() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Operation")
	ret0, _ := ret[0].(string)
	return ret0
}

// Operation indicates an expected call of Operation.
func (mr *MockQueryMockRecorder) Operation() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Operation", reflect.TypeOf((*MockQuery)(nil).Operation))
}

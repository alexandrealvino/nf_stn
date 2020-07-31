// Code generated by MockGen. DO NOT EDIT.
// Source: db_queries.go

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	entities "nf_stn/entities"
	reflect "reflect"
)

// MockDataBase is a mock of DataBase interface
type MockDataBase struct {
	ctrl     *gomock.Controller
	recorder *MockDataBaseMockRecorder
}

// MockDataBaseMockRecorder is the mock recorder for MockDataBase
type MockDataBaseMockRecorder struct {
	mock *MockDataBase
}

// NewMockDataBase creates a new mock instance
func NewMockDataBase(ctrl *gomock.Controller) *MockDataBase {
	mock := &MockDataBase{ctrl: ctrl}
	mock.recorder = &MockDataBaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDataBase) EXPECT() *MockDataBaseMockRecorder {
	return m.recorder
}

// Init mocks base method
func (m *MockDataBase) Init() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Init")
}

// Init indicates an expected call of Init
func (mr *MockDataBaseMockRecorder) Init() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockDataBase)(nil).Init))
}

// GetAll mocks base method
func (m *MockDataBase) GetAll() ([]entities.Invoice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]entities.Invoice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll
func (mr *MockDataBaseMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockDataBase)(nil).GetAll))
}

// GetInvoiceByDocument mocks base method
func (m *MockDataBase) GetInvoiceByDocument(document string) (entities.Invoice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInvoiceByDocument", document)
	ret0, _ := ret[0].(entities.Invoice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInvoiceByDocument indicates an expected call of GetInvoiceByDocument
func (mr *MockDataBaseMockRecorder) GetInvoiceByDocument(document interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInvoiceByDocument", reflect.TypeOf((*MockDataBase)(nil).GetInvoiceByDocument), document)
}

// GetUser mocks base method
func (m *MockDataBase) GetUser(username, password string) (int, string, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", username, password)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[0].(string)
	ret2, _ := ret[1].(string)
	ret3, _ := ret[2].(error)
	return ret0, ret1, ret2, ret3
}

// GetUser indicates an expected call of GetUser
func (mr *MockDataBaseMockRecorder) GetUser(username, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockDataBase)(nil).GetUser), username, password)
}

// GetInvoiceByID mocks base method
func (m *MockDataBase) GetInvoiceByID(id int) (entities.Invoice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInvoiceByID", id)
	ret0, _ := ret[0].(entities.Invoice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInvoiceByID indicates an expected call of GetInvoiceByID
func (mr *MockDataBaseMockRecorder) GetInvoiceByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInvoiceByID", reflect.TypeOf((*MockDataBase)(nil).GetInvoiceByID), id)
}

// InsertInvoice mocks base method
func (m *MockDataBase) InsertInvoice(invoice entities.Invoice) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertInvoice", invoice)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertInvoice indicates an expected call of InsertInvoice
func (mr *MockDataBaseMockRecorder) InsertInvoice(invoice interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertInvoice", reflect.TypeOf((*MockDataBase)(nil).InsertInvoice), invoice)
}

// DeleteInvoice mocks base method
func (m *MockDataBase) DeleteInvoice(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteInvoice", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteInvoice indicates an expected call of DeleteInvoice
func (mr *MockDataBaseMockRecorder) DeleteInvoice(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteInvoice", reflect.TypeOf((*MockDataBase)(nil).DeleteInvoice), id)
}

// UpdateInvoice mocks base method
func (m *MockDataBase) UpdateInvoice(invoice entities.Invoice) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateInvoice", invoice)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateInvoice indicates an expected call of UpdateInvoice
func (mr *MockDataBaseMockRecorder) UpdateInvoice(invoice interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateInvoice", reflect.TypeOf((*MockDataBase)(nil).UpdateInvoice), invoice)
}

// PatchInvoice mocks base method
func (m *MockDataBase) PatchInvoice(invoice entities.Invoice) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PatchInvoice", invoice)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchInvoice indicates an expected call of PatchInvoice
func (mr *MockDataBaseMockRecorder) PatchInvoice(invoice interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchInvoice", reflect.TypeOf((*MockDataBase)(nil).PatchInvoice), invoice)
}

// InvoiceExists mocks base method
func (m *MockDataBase) InvoiceExists(document string) (entities.Invoice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InvoiceExists", document)
	ret0, _ := ret[0].(entities.Invoice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InvoiceExists indicates an expected call of InvoiceExists
func (mr *MockDataBaseMockRecorder) InvoiceExists(document interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InvoiceExists", reflect.TypeOf((*MockDataBase)(nil).InvoiceExists), document)
}

// ClearTable mocks base method
func (m *MockDataBase) ClearTable() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ClearTable")
}

// ClearTable indicates an expected call of ClearTable
func (mr *MockDataBaseMockRecorder) ClearTable() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClearTable", reflect.TypeOf((*MockDataBase)(nil).ClearTable))
}

// Pagination mocks base method
func (m *MockDataBase) Pagination(offset int) ([]entities.Invoice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Pagination", offset)
	ret0, _ := ret[0].([]entities.Invoice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Pagination indicates an expected call of Pagination
func (mr *MockDataBaseMockRecorder) Pagination(offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pagination", reflect.TypeOf((*MockDataBase)(nil).Pagination), offset)
}

// PaginationOrderByMonth mocks base method
func (m *MockDataBase) PaginationOrderByMonth(offset int) ([]entities.Invoice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PaginationOrderByMonth", offset)
	ret0, _ := ret[0].([]entities.Invoice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PaginationOrderByMonth indicates an expected call of PaginationOrderByMonth
func (mr *MockDataBaseMockRecorder) PaginationOrderByMonth(offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PaginationOrderByMonth", reflect.TypeOf((*MockDataBase)(nil).PaginationOrderByMonth), offset)
}

// PaginationOrderByYear mocks base method
func (m *MockDataBase) PaginationOrderByYear(offset int) ([]entities.Invoice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PaginationOrderByYear", offset)
	ret0, _ := ret[0].([]entities.Invoice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PaginationOrderByYear indicates an expected call of PaginationOrderByYear
func (mr *MockDataBaseMockRecorder) PaginationOrderByYear(offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PaginationOrderByYear", reflect.TypeOf((*MockDataBase)(nil).PaginationOrderByYear), offset)
}

// PaginationOrderByDocument mocks base method
func (m *MockDataBase) PaginationOrderByDocument(offset int) ([]entities.Invoice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PaginationOrderByDocument", offset)
	ret0, _ := ret[0].([]entities.Invoice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PaginationOrderByDocument indicates an expected call of PaginationOrderByDocument
func (mr *MockDataBaseMockRecorder) PaginationOrderByDocument(offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PaginationOrderByDocument", reflect.TypeOf((*MockDataBase)(nil).PaginationOrderByDocument), offset)
}

// PaginationByMonth mocks base method
func (m *MockDataBase) PaginationByMonth(offset, referenceMonth int) ([]entities.Invoice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PaginationByMonth", offset, referenceMonth)
	ret0, _ := ret[0].([]entities.Invoice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PaginationByMonth indicates an expected call of PaginationByMonth
func (mr *MockDataBaseMockRecorder) PaginationByMonth(offset, referenceMonth interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PaginationByMonth", reflect.TypeOf((*MockDataBase)(nil).PaginationByMonth), offset, referenceMonth)
}

// PaginationByYear mocks base method
func (m *MockDataBase) PaginationByYear(offset, referenceYear int) ([]entities.Invoice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PaginationByYear", offset, referenceYear)
	ret0, _ := ret[0].([]entities.Invoice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PaginationByYear indicates an expected call of PaginationByYear
func (mr *MockDataBaseMockRecorder) PaginationByYear(offset, referenceYear interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PaginationByYear", reflect.TypeOf((*MockDataBase)(nil).PaginationByYear), offset, referenceYear)
}

// PaginationByDocument mocks base method
func (m *MockDataBase) PaginationByDocument(offset int, document string) ([]entities.Invoice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PaginationByDocument", offset, document)
	ret0, _ := ret[0].([]entities.Invoice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PaginationByDocument indicates an expected call of PaginationByDocument
func (mr *MockDataBaseMockRecorder) PaginationByDocument(offset, document interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PaginationByDocument", reflect.TypeOf((*MockDataBase)(nil).PaginationByDocument), offset, document)
}

// PaginationOrderByMonthYear mocks base method
func (m *MockDataBase) PaginationOrderByMonthYear(offset int) ([]entities.Invoice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PaginationOrderByMonthYear", offset)
	ret0, _ := ret[0].([]entities.Invoice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PaginationOrderByMonthYear indicates an expected call of PaginationOrderByMonthYear
func (mr *MockDataBaseMockRecorder) PaginationOrderByMonthYear(offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PaginationOrderByMonthYear", reflect.TypeOf((*MockDataBase)(nil).PaginationOrderByMonthYear), offset)
}

// PaginationOrderByMonthDocument mocks base method
func (m *MockDataBase) PaginationOrderByMonthDocument(offset int) ([]entities.Invoice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PaginationOrderByMonthDocument", offset)
	ret0, _ := ret[0].([]entities.Invoice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PaginationOrderByMonthDocument indicates an expected call of PaginationOrderByMonthDocument
func (mr *MockDataBaseMockRecorder) PaginationOrderByMonthDocument(offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PaginationOrderByMonthDocument", reflect.TypeOf((*MockDataBase)(nil).PaginationOrderByMonthDocument), offset)
}

// PaginationOrderByYearDocument mocks base method
func (m *MockDataBase) PaginationOrderByYearDocument(offset int) ([]entities.Invoice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PaginationOrderByYearDocument", offset)
	ret0, _ := ret[0].([]entities.Invoice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PaginationOrderByYearDocument indicates an expected call of PaginationOrderByYearDocument
func (mr *MockDataBaseMockRecorder) PaginationOrderByYearDocument(offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PaginationOrderByYearDocument", reflect.TypeOf((*MockDataBase)(nil).PaginationOrderByYearDocument), offset)
}

// NewMockDbClient mocks
func NewMockDbClient(ctrl *gomock.Controller) *MockDataBase {
	mock := &MockDataBase{ctrl: ctrl}
	mock.recorder = &MockDataBaseMockRecorder{mock}
	return mock
}
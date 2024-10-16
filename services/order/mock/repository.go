// Code generated by MockGen. DO NOT EDIT.
// Source: ./services/order/repository.go
//
// Generated by this command:
//
//	mockgen -source=./services/order/repository.go -destination=./services/order/mock/repository.go
//

// Package mock_order is a generated GoMock package.
package mock_order

import (
	reflect "reflect"

	entity "github.com/sendydwi/online-book-store/services/order/entity"
	gomock "go.uber.org/mock/gomock"
)

// MockOrderRepositoryInterface is a mock of OrderRepositoryInterface interface.
type MockOrderRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockOrderRepositoryInterfaceMockRecorder
}

// MockOrderRepositoryInterfaceMockRecorder is the mock recorder for MockOrderRepositoryInterface.
type MockOrderRepositoryInterfaceMockRecorder struct {
	mock *MockOrderRepositoryInterface
}

// NewMockOrderRepositoryInterface creates a new mock instance.
func NewMockOrderRepositoryInterface(ctrl *gomock.Controller) *MockOrderRepositoryInterface {
	mock := &MockOrderRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockOrderRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrderRepositoryInterface) EXPECT() *MockOrderRepositoryInterfaceMockRecorder {
	return m.recorder
}

// CreateOrder mocks base method.
func (m *MockOrderRepositoryInterface) CreateOrder(order entity.Order, orderItem []*entity.OrderItem) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrder", order, orderItem)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateOrder indicates an expected call of CreateOrder.
func (mr *MockOrderRepositoryInterfaceMockRecorder) CreateOrder(order, orderItem any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrder", reflect.TypeOf((*MockOrderRepositoryInterface)(nil).CreateOrder), order, orderItem)
}

// DeleteOrder mocks base method.
func (m *MockOrderRepositoryInterface) DeleteOrder(order entity.Order) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOrder", order)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteOrder indicates an expected call of DeleteOrder.
func (mr *MockOrderRepositoryInterfaceMockRecorder) DeleteOrder(order any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOrder", reflect.TypeOf((*MockOrderRepositoryInterface)(nil).DeleteOrder), order)
}

// GetOrderById mocks base method.
func (m *MockOrderRepositoryInterface) GetOrderById(orderId string) (*entity.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrderById", orderId)
	ret0, _ := ret[0].(*entity.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrderById indicates an expected call of GetOrderById.
func (mr *MockOrderRepositoryInterfaceMockRecorder) GetOrderById(orderId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrderById", reflect.TypeOf((*MockOrderRepositoryInterface)(nil).GetOrderById), orderId)
}

// GetOrderByUserId mocks base method.
func (m *MockOrderRepositoryInterface) GetOrderByUserId(userId string) ([]entity.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrderByUserId", userId)
	ret0, _ := ret[0].([]entity.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrderByUserId indicates an expected call of GetOrderByUserId.
func (mr *MockOrderRepositoryInterfaceMockRecorder) GetOrderByUserId(userId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrderByUserId", reflect.TypeOf((*MockOrderRepositoryInterface)(nil).GetOrderByUserId), userId)
}

// GetOrderItemByOrderId mocks base method.
func (m *MockOrderRepositoryInterface) GetOrderItemByOrderId(orderId string) ([]entity.OrderItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrderItemByOrderId", orderId)
	ret0, _ := ret[0].([]entity.OrderItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrderItemByOrderId indicates an expected call of GetOrderItemByOrderId.
func (mr *MockOrderRepositoryInterfaceMockRecorder) GetOrderItemByOrderId(orderId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrderItemByOrderId", reflect.TypeOf((*MockOrderRepositoryInterface)(nil).GetOrderItemByOrderId), orderId)
}

// Code generated by MockGen. DO NOT EDIT.
// Source: ./services/cart/service.go
//
// Generated by this command:
//
//	mockgen -source=./services/cart/service.go -destination=./services/cart/mock/service.go
//

// Package mock_cart is a generated GoMock package.
package mock_cart

import (
	reflect "reflect"

	apicart "github.com/sendydwi/online-book-store/api/cart"
	entity "github.com/sendydwi/online-book-store/services/cart/entity"
	gomock "go.uber.org/mock/gomock"
)

// MockCartServiceInterface is a mock of CartServiceInterface interface.
type MockCartServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockCartServiceInterfaceMockRecorder
}

// MockCartServiceInterfaceMockRecorder is the mock recorder for MockCartServiceInterface.
type MockCartServiceInterfaceMockRecorder struct {
	mock *MockCartServiceInterface
}

// NewMockCartServiceInterface creates a new mock instance.
func NewMockCartServiceInterface(ctrl *gomock.Controller) *MockCartServiceInterface {
	mock := &MockCartServiceInterface{ctrl: ctrl}
	mock.recorder = &MockCartServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCartServiceInterface) EXPECT() *MockCartServiceInterfaceMockRecorder {
	return m.recorder
}

// GetCartItem mocks base method.
func (m *MockCartServiceInterface) GetCartItem(userId string) (*apicart.GetCartResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCartItem", userId)
	ret0, _ := ret[0].(*apicart.GetCartResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCartItem indicates an expected call of GetCartItem.
func (mr *MockCartServiceInterfaceMockRecorder) GetCartItem(userId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCartItem", reflect.TypeOf((*MockCartServiceInterface)(nil).GetCartItem), userId)
}

// UpdateCartItem mocks base method.
func (m *MockCartServiceInterface) UpdateCartItem(updateRequest apicart.CartUpdateRequest, userId string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCartItem", updateRequest, userId)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCartItem indicates an expected call of UpdateCartItem.
func (mr *MockCartServiceInterfaceMockRecorder) UpdateCartItem(updateRequest, userId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCartItem", reflect.TypeOf((*MockCartServiceInterface)(nil).UpdateCartItem), updateRequest, userId)
}

// UpdateCartStatusToOrdered mocks base method.
func (m *MockCartServiceInterface) UpdateCartStatusToOrdered(userId string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCartStatusToOrdered", userId)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCartStatusToOrdered indicates an expected call of UpdateCartStatusToOrdered.
func (mr *MockCartServiceInterfaceMockRecorder) UpdateCartStatusToOrdered(userId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCartStatusToOrdered", reflect.TypeOf((*MockCartServiceInterface)(nil).UpdateCartStatusToOrdered), userId)
}

// GetCurrentCart mocks base method.
func (m *MockCartServiceInterface) GetCurrentCart(userId string) (*entity.Cart, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrentCart", userId)
	ret0, _ := ret[0].(*entity.Cart)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCurrentCart indicates an expected call of getCurrentCart.
func (mr *MockCartServiceInterfaceMockRecorder) GetCurrentCart(userId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentCart", reflect.TypeOf((*MockCartServiceInterface)(nil).GetCurrentCart), userId)
}

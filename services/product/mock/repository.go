// Code generated by MockGen. DO NOT EDIT.
// Source: ./services/product/repository.go
//
// Generated by this command:
//
//	mockgen -source=./services/product/repository.go -destination=./services/product/mock/repository.go
//

// Package mock_product is a generated GoMock package.
package mock_product

import (
	reflect "reflect"

	entity "github.com/sendydwi/online-book-store/services/product/entity"
	gomock "go.uber.org/mock/gomock"
)

// MockProductRepositoryInterface is a mock of ProductRepositoryInterface interface.
type MockProductRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockProductRepositoryInterfaceMockRecorder
}

// MockProductRepositoryInterfaceMockRecorder is the mock recorder for MockProductRepositoryInterface.
type MockProductRepositoryInterfaceMockRecorder struct {
	mock *MockProductRepositoryInterface
}

// NewMockProductRepositoryInterface creates a new mock instance.
func NewMockProductRepositoryInterface(ctrl *gomock.Controller) *MockProductRepositoryInterface {
	mock := &MockProductRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockProductRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductRepositoryInterface) EXPECT() *MockProductRepositoryInterfaceMockRecorder {
	return m.recorder
}

// GetProductById mocks base method.
func (m *MockProductRepositoryInterface) GetProductById(productId int) (*entity.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProductById", productId)
	ret0, _ := ret[0].(*entity.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProductById indicates an expected call of GetProductById.
func (mr *MockProductRepositoryInterfaceMockRecorder) GetProductById(productId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductById", reflect.TypeOf((*MockProductRepositoryInterface)(nil).GetProductById), productId)
}

// GetProductList mocks base method.
func (m *MockProductRepositoryInterface) GetProductList(page, size int) (*[]entity.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProductList", page, size)
	ret0, _ := ret[0].(*[]entity.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProductList indicates an expected call of GetProductList.
func (mr *MockProductRepositoryInterfaceMockRecorder) GetProductList(page, size any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductList", reflect.TypeOf((*MockProductRepositoryInterface)(nil).GetProductList), page, size)
}

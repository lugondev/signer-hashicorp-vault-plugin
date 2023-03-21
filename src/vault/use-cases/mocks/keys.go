// Code generated by MockGen. DO NOT EDIT.
// Source: keys.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	entities "github.com/lugondev/signer-hashicorp-vault-plugin/src/vault/entities"
	usecases "github.com/lugondev/signer-hashicorp-vault-plugin/src/vault/use-cases"
	gomock "github.com/golang/mock/gomock"
	logical "github.com/hashicorp/vault/sdk/logical"
	reflect "reflect"
)

// MockKeysUseCases is a mock of KeysUseCases interface
type MockKeysUseCases struct {
	ctrl     *gomock.Controller
	recorder *MockKeysUseCasesMockRecorder
}

// MockKeysUseCasesMockRecorder is the mock recorder for MockKeysUseCases
type MockKeysUseCasesMockRecorder struct {
	mock *MockKeysUseCases
}

// NewMockKeysUseCases creates a new mock instance
func NewMockKeysUseCases(ctrl *gomock.Controller) *MockKeysUseCases {
	mock := &MockKeysUseCases{ctrl: ctrl}
	mock.recorder = &MockKeysUseCasesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockKeysUseCases) EXPECT() *MockKeysUseCasesMockRecorder {
	return m.recorder
}

// CreateKey mocks base method
func (m *MockKeysUseCases) CreateKey() usecases.CreateKeyUseCase {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateKey")
	ret0, _ := ret[0].(usecases.CreateKeyUseCase)
	return ret0
}

// CreateKey indicates an expected call of CreateKey
func (mr *MockKeysUseCasesMockRecorder) CreateKey() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateKey", reflect.TypeOf((*MockKeysUseCases)(nil).CreateKey))
}

// GetKey mocks base method
func (m *MockKeysUseCases) GetKey() usecases.GetKeyUseCase {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetKey")
	ret0, _ := ret[0].(usecases.GetKeyUseCase)
	return ret0
}

// GetKey indicates an expected call of GetKey
func (mr *MockKeysUseCasesMockRecorder) GetKey() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetKey", reflect.TypeOf((*MockKeysUseCases)(nil).GetKey))
}

// ListKeys mocks base method
func (m *MockKeysUseCases) ListKeys() usecases.ListKeysUseCase {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListKeys")
	ret0, _ := ret[0].(usecases.ListKeysUseCase)
	return ret0
}

// ListKeys indicates an expected call of ListKeys
func (mr *MockKeysUseCasesMockRecorder) ListKeys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListKeys", reflect.TypeOf((*MockKeysUseCases)(nil).ListKeys))
}

// ListNamespaces mocks base method
func (m *MockKeysUseCases) ListNamespaces() usecases.ListKeysNamespacesUseCase {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListNamespaces")
	ret0, _ := ret[0].(usecases.ListKeysNamespacesUseCase)
	return ret0
}

// ListNamespaces indicates an expected call of ListNamespaces
func (mr *MockKeysUseCasesMockRecorder) ListNamespaces() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListNamespaces", reflect.TypeOf((*MockKeysUseCases)(nil).ListNamespaces))
}

// SignPayload mocks base method
func (m *MockKeysUseCases) SignPayload() usecases.KeysSignUseCase {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignPayload")
	ret0, _ := ret[0].(usecases.KeysSignUseCase)
	return ret0
}

// SignPayload indicates an expected call of SignPayload
func (mr *MockKeysUseCasesMockRecorder) SignPayload() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignPayload", reflect.TypeOf((*MockKeysUseCases)(nil).SignPayload))
}

// UpdateKey mocks base method
func (m *MockKeysUseCases) UpdateKey() usecases.UpdateKeyUseCase {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateKey")
	ret0, _ := ret[0].(usecases.UpdateKeyUseCase)
	return ret0
}

// UpdateKey indicates an expected call of UpdateKey
func (mr *MockKeysUseCasesMockRecorder) UpdateKey() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateKey", reflect.TypeOf((*MockKeysUseCases)(nil).UpdateKey))
}

// DestroyKey mocks base method
func (m *MockKeysUseCases) DestroyKey() usecases.DestroyKeyUseCase {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DestroyKey")
	ret0, _ := ret[0].(usecases.DestroyKeyUseCase)
	return ret0
}

// DestroyKey indicates an expected call of DestroyKey
func (mr *MockKeysUseCasesMockRecorder) DestroyKey() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DestroyKey", reflect.TypeOf((*MockKeysUseCases)(nil).DestroyKey))
}

// MockCreateKeyUseCase is a mock of CreateKeyUseCase interface
type MockCreateKeyUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockCreateKeyUseCaseMockRecorder
}

// MockCreateKeyUseCaseMockRecorder is the mock recorder for MockCreateKeyUseCase
type MockCreateKeyUseCaseMockRecorder struct {
	mock *MockCreateKeyUseCase
}

// NewMockCreateKeyUseCase creates a new mock instance
func NewMockCreateKeyUseCase(ctrl *gomock.Controller) *MockCreateKeyUseCase {
	mock := &MockCreateKeyUseCase{ctrl: ctrl}
	mock.recorder = &MockCreateKeyUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCreateKeyUseCase) EXPECT() *MockCreateKeyUseCaseMockRecorder {
	return m.recorder
}

// Execute mocks base method
func (m *MockCreateKeyUseCase) Execute(ctx context.Context, namespace, id, algo, curve, importedPrivKey string, tags map[string]string) (*entities.Key, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", ctx, namespace, id, algo, curve, importedPrivKey, tags)
	ret0, _ := ret[0].(*entities.Key)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute
func (mr *MockCreateKeyUseCaseMockRecorder) Execute(ctx, namespace, id, algo, curve, importedPrivKey, tags interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockCreateKeyUseCase)(nil).Execute), ctx, namespace, id, algo, curve, importedPrivKey, tags)
}

// WithStorage mocks base method
func (m *MockCreateKeyUseCase) WithStorage(storage logical.Storage) usecases.CreateKeyUseCase {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithStorage", storage)
	ret0, _ := ret[0].(usecases.CreateKeyUseCase)
	return ret0
}

// WithStorage indicates an expected call of WithStorage
func (mr *MockCreateKeyUseCaseMockRecorder) WithStorage(storage interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithStorage", reflect.TypeOf((*MockCreateKeyUseCase)(nil).WithStorage), storage)
}

// MockUpdateKeyUseCase is a mock of UpdateKeyUseCase interface
type MockUpdateKeyUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockUpdateKeyUseCaseMockRecorder
}

// MockUpdateKeyUseCaseMockRecorder is the mock recorder for MockUpdateKeyUseCase
type MockUpdateKeyUseCaseMockRecorder struct {
	mock *MockUpdateKeyUseCase
}

// NewMockUpdateKeyUseCase creates a new mock instance
func NewMockUpdateKeyUseCase(ctrl *gomock.Controller) *MockUpdateKeyUseCase {
	mock := &MockUpdateKeyUseCase{ctrl: ctrl}
	mock.recorder = &MockUpdateKeyUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUpdateKeyUseCase) EXPECT() *MockUpdateKeyUseCaseMockRecorder {
	return m.recorder
}

// Execute mocks base method
func (m *MockUpdateKeyUseCase) Execute(ctx context.Context, namespace, id string, tags map[string]string) (*entities.Key, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", ctx, namespace, id, tags)
	ret0, _ := ret[0].(*entities.Key)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute
func (mr *MockUpdateKeyUseCaseMockRecorder) Execute(ctx, namespace, id, tags interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockUpdateKeyUseCase)(nil).Execute), ctx, namespace, id, tags)
}

// WithStorage mocks base method
func (m *MockUpdateKeyUseCase) WithStorage(storage logical.Storage) usecases.UpdateKeyUseCase {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithStorage", storage)
	ret0, _ := ret[0].(usecases.UpdateKeyUseCase)
	return ret0
}

// WithStorage indicates an expected call of WithStorage
func (mr *MockUpdateKeyUseCaseMockRecorder) WithStorage(storage interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithStorage", reflect.TypeOf((*MockUpdateKeyUseCase)(nil).WithStorage), storage)
}

// MockDestroyKeyUseCase is a mock of DestroyKeyUseCase interface
type MockDestroyKeyUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockDestroyKeyUseCaseMockRecorder
}

// MockDestroyKeyUseCaseMockRecorder is the mock recorder for MockDestroyKeyUseCase
type MockDestroyKeyUseCaseMockRecorder struct {
	mock *MockDestroyKeyUseCase
}

// NewMockDestroyKeyUseCase creates a new mock instance
func NewMockDestroyKeyUseCase(ctrl *gomock.Controller) *MockDestroyKeyUseCase {
	mock := &MockDestroyKeyUseCase{ctrl: ctrl}
	mock.recorder = &MockDestroyKeyUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDestroyKeyUseCase) EXPECT() *MockDestroyKeyUseCaseMockRecorder {
	return m.recorder
}

// Execute mocks base method
func (m *MockDestroyKeyUseCase) Execute(ctx context.Context, namespace, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", ctx, namespace, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Execute indicates an expected call of Execute
func (mr *MockDestroyKeyUseCaseMockRecorder) Execute(ctx, namespace, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockDestroyKeyUseCase)(nil).Execute), ctx, namespace, id)
}

// WithStorage mocks base method
func (m *MockDestroyKeyUseCase) WithStorage(storage logical.Storage) usecases.DestroyKeyUseCase {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithStorage", storage)
	ret0, _ := ret[0].(usecases.DestroyKeyUseCase)
	return ret0
}

// WithStorage indicates an expected call of WithStorage
func (mr *MockDestroyKeyUseCaseMockRecorder) WithStorage(storage interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithStorage", reflect.TypeOf((*MockDestroyKeyUseCase)(nil).WithStorage), storage)
}

// MockGetKeyUseCase is a mock of GetKeyUseCase interface
type MockGetKeyUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockGetKeyUseCaseMockRecorder
}

// MockGetKeyUseCaseMockRecorder is the mock recorder for MockGetKeyUseCase
type MockGetKeyUseCaseMockRecorder struct {
	mock *MockGetKeyUseCase
}

// NewMockGetKeyUseCase creates a new mock instance
func NewMockGetKeyUseCase(ctrl *gomock.Controller) *MockGetKeyUseCase {
	mock := &MockGetKeyUseCase{ctrl: ctrl}
	mock.recorder = &MockGetKeyUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGetKeyUseCase) EXPECT() *MockGetKeyUseCaseMockRecorder {
	return m.recorder
}

// Execute mocks base method
func (m *MockGetKeyUseCase) Execute(ctx context.Context, id, namespace string) (*entities.Key, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", ctx, id, namespace)
	ret0, _ := ret[0].(*entities.Key)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute
func (mr *MockGetKeyUseCaseMockRecorder) Execute(ctx, id, namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockGetKeyUseCase)(nil).Execute), ctx, id, namespace)
}

// WithStorage mocks base method
func (m *MockGetKeyUseCase) WithStorage(storage logical.Storage) usecases.GetKeyUseCase {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithStorage", storage)
	ret0, _ := ret[0].(usecases.GetKeyUseCase)
	return ret0
}

// WithStorage indicates an expected call of WithStorage
func (mr *MockGetKeyUseCaseMockRecorder) WithStorage(storage interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithStorage", reflect.TypeOf((*MockGetKeyUseCase)(nil).WithStorage), storage)
}

// MockListKeysUseCase is a mock of ListKeysUseCase interface
type MockListKeysUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockListKeysUseCaseMockRecorder
}

// MockListKeysUseCaseMockRecorder is the mock recorder for MockListKeysUseCase
type MockListKeysUseCaseMockRecorder struct {
	mock *MockListKeysUseCase
}

// NewMockListKeysUseCase creates a new mock instance
func NewMockListKeysUseCase(ctrl *gomock.Controller) *MockListKeysUseCase {
	mock := &MockListKeysUseCase{ctrl: ctrl}
	mock.recorder = &MockListKeysUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockListKeysUseCase) EXPECT() *MockListKeysUseCaseMockRecorder {
	return m.recorder
}

// Execute mocks base method
func (m *MockListKeysUseCase) Execute(ctx context.Context, namespace string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", ctx, namespace)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute
func (mr *MockListKeysUseCaseMockRecorder) Execute(ctx, namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockListKeysUseCase)(nil).Execute), ctx, namespace)
}

// WithStorage mocks base method
func (m *MockListKeysUseCase) WithStorage(storage logical.Storage) usecases.ListKeysUseCase {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithStorage", storage)
	ret0, _ := ret[0].(usecases.ListKeysUseCase)
	return ret0
}

// WithStorage indicates an expected call of WithStorage
func (mr *MockListKeysUseCaseMockRecorder) WithStorage(storage interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithStorage", reflect.TypeOf((*MockListKeysUseCase)(nil).WithStorage), storage)
}

// MockKeysSignUseCase is a mock of KeysSignUseCase interface
type MockKeysSignUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockKeysSignUseCaseMockRecorder
}

// MockKeysSignUseCaseMockRecorder is the mock recorder for MockKeysSignUseCase
type MockKeysSignUseCaseMockRecorder struct {
	mock *MockKeysSignUseCase
}

// NewMockKeysSignUseCase creates a new mock instance
func NewMockKeysSignUseCase(ctrl *gomock.Controller) *MockKeysSignUseCase {
	mock := &MockKeysSignUseCase{ctrl: ctrl}
	mock.recorder = &MockKeysSignUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockKeysSignUseCase) EXPECT() *MockKeysSignUseCaseMockRecorder {
	return m.recorder
}

// Execute mocks base method
func (m *MockKeysSignUseCase) Execute(ctx context.Context, id, namespace, data string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", ctx, id, namespace, data)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute
func (mr *MockKeysSignUseCaseMockRecorder) Execute(ctx, id, namespace, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockKeysSignUseCase)(nil).Execute), ctx, id, namespace, data)
}

// WithStorage mocks base method
func (m *MockKeysSignUseCase) WithStorage(storage logical.Storage) usecases.KeysSignUseCase {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithStorage", storage)
	ret0, _ := ret[0].(usecases.KeysSignUseCase)
	return ret0
}

// WithStorage indicates an expected call of WithStorage
func (mr *MockKeysSignUseCaseMockRecorder) WithStorage(storage interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithStorage", reflect.TypeOf((*MockKeysSignUseCase)(nil).WithStorage), storage)
}

// MockListKeysNamespacesUseCase is a mock of ListKeysNamespacesUseCase interface
type MockListKeysNamespacesUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockListKeysNamespacesUseCaseMockRecorder
}

// MockListKeysNamespacesUseCaseMockRecorder is the mock recorder for MockListKeysNamespacesUseCase
type MockListKeysNamespacesUseCaseMockRecorder struct {
	mock *MockListKeysNamespacesUseCase
}

// NewMockListKeysNamespacesUseCase creates a new mock instance
func NewMockListKeysNamespacesUseCase(ctrl *gomock.Controller) *MockListKeysNamespacesUseCase {
	mock := &MockListKeysNamespacesUseCase{ctrl: ctrl}
	mock.recorder = &MockListKeysNamespacesUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockListKeysNamespacesUseCase) EXPECT() *MockListKeysNamespacesUseCaseMockRecorder {
	return m.recorder
}

// Execute mocks base method
func (m *MockListKeysNamespacesUseCase) Execute(ctx context.Context) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", ctx)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute
func (mr *MockListKeysNamespacesUseCaseMockRecorder) Execute(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockListKeysNamespacesUseCase)(nil).Execute), ctx)
}

// WithStorage mocks base method
func (m *MockListKeysNamespacesUseCase) WithStorage(storage logical.Storage) usecases.ListKeysNamespacesUseCase {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithStorage", storage)
	ret0, _ := ret[0].(usecases.ListKeysNamespacesUseCase)
	return ret0
}

// WithStorage indicates an expected call of WithStorage
func (mr *MockListKeysNamespacesUseCaseMockRecorder) WithStorage(storage interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithStorage", reflect.TypeOf((*MockListKeysNamespacesUseCase)(nil).WithStorage), storage)
}

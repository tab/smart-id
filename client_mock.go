// Code generated by MockGen. DO NOT EDIT.
// Source: client.go
//
// Generated by this command:
//
//	mockgen -source=client.go -destination=client_mock.go -package=smartid
//

// Package smartid is a generated GoMock package.
package smartid

import (
	context "context"
	reflect "reflect"

	models "github.com/tab/smartid/internal/models"
	gomock "go.uber.org/mock/gomock"
)

// MockProvider is a mock of Provider interface.
type MockProvider struct {
	ctrl     *gomock.Controller
	recorder *MockProviderMockRecorder
	isgomock struct{}
}

// MockProviderMockRecorder is the mock recorder for MockProvider.
type MockProviderMockRecorder struct {
	mock *MockProvider
}

// NewMockProvider creates a new mock instance.
func NewMockProvider(ctrl *gomock.Controller) *MockProvider {
	mock := &MockProvider{ctrl: ctrl}
	mock.recorder = &MockProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProvider) EXPECT() *MockProviderMockRecorder {
	return m.recorder
}

// CreateSession mocks base method.
func (m *MockProvider) CreateSession(ctx context.Context, nationalIdentityNumber string) (*models.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSession", ctx, nationalIdentityNumber)
	ret0, _ := ret[0].(*models.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSession indicates an expected call of CreateSession.
func (mr *MockProviderMockRecorder) CreateSession(ctx, nationalIdentityNumber any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSession", reflect.TypeOf((*MockProvider)(nil).CreateSession), ctx, nationalIdentityNumber)
}

// FetchSession mocks base method.
func (m *MockProvider) FetchSession(ctx context.Context, sessionId string) (*Person, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchSession", ctx, sessionId)
	ret0, _ := ret[0].(*Person)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchSession indicates an expected call of FetchSession.
func (mr *MockProviderMockRecorder) FetchSession(ctx, sessionId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchSession", reflect.TypeOf((*MockProvider)(nil).FetchSession), ctx, sessionId)
}

// Validate mocks base method.
func (m *MockProvider) Validate() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validate")
	ret0, _ := ret[0].(error)
	return ret0
}

// Validate indicates an expected call of Validate.
func (mr *MockProviderMockRecorder) Validate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockProvider)(nil).Validate))
}

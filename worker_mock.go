// Code generated by MockGen. DO NOT EDIT.
// Source: worker.go
//
// Generated by this command:
//
//	mockgen -source=worker.go -destination=worker_mock.go -package=smartid
//

// Package smartid is a generated GoMock package.
package smartid

import (
	context "context"
	reflect "reflect"

	config "github.com/tab/smartid/internal/config"
	gomock "go.uber.org/mock/gomock"
)

// MockBackgroundWorker is a mock of BackgroundWorker interface.
type MockBackgroundWorker struct {
	ctrl     *gomock.Controller
	recorder *MockBackgroundWorkerMockRecorder
	isgomock struct{}
}

// MockBackgroundWorkerMockRecorder is the mock recorder for MockBackgroundWorker.
type MockBackgroundWorkerMockRecorder struct {
	mock *MockBackgroundWorker
}

// NewMockBackgroundWorker creates a new mock instance.
func NewMockBackgroundWorker(ctrl *gomock.Controller) *MockBackgroundWorker {
	mock := &MockBackgroundWorker{ctrl: ctrl}
	mock.recorder = &MockBackgroundWorkerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBackgroundWorker) EXPECT() *MockBackgroundWorkerMockRecorder {
	return m.recorder
}

// Process mocks base method.
func (m *MockBackgroundWorker) Process(ctx context.Context, sessionId string) <-chan Result {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Process", ctx, sessionId)
	ret0, _ := ret[0].(<-chan Result)
	return ret0
}

// Process indicates an expected call of Process.
func (mr *MockBackgroundWorkerMockRecorder) Process(ctx, sessionId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Process", reflect.TypeOf((*MockBackgroundWorker)(nil).Process), ctx, sessionId)
}

// Start mocks base method.
func (m *MockBackgroundWorker) Start() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Start")
}

// Start indicates an expected call of Start.
func (mr *MockBackgroundWorkerMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockBackgroundWorker)(nil).Start))
}

// Stop mocks base method.
func (m *MockBackgroundWorker) Stop() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop.
func (mr *MockBackgroundWorkerMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockBackgroundWorker)(nil).Stop))
}

// WithConfig mocks base method.
func (m *MockBackgroundWorker) WithConfig(cfg config.WorkerConfig) *Worker {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithConfig", cfg)
	ret0, _ := ret[0].(*Worker)
	return ret0
}

// WithConfig indicates an expected call of WithConfig.
func (mr *MockBackgroundWorkerMockRecorder) WithConfig(cfg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithConfig", reflect.TypeOf((*MockBackgroundWorker)(nil).WithConfig), cfg)
}

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/kodekulture/wordle-server/handler/token (interfaces: Handler)
//
// Generated by this command:
//
//	mockgen -destination=../../internal/mocks/token_handler.go -package=mocks -mock_names Handler=MockTokenHandler -typed . Handler
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"
	time "time"

	game "github.com/kodekulture/wordle-server/game"
	auth "github.com/lordvidex/x/auth"
	gomock "go.uber.org/mock/gomock"
)

// MockTokenHandler is a mock of Handler interface.
type MockTokenHandler struct {
	ctrl     *gomock.Controller
	recorder *MockTokenHandlerMockRecorder
	isgomock struct{}
}

// MockTokenHandlerMockRecorder is the mock recorder for MockTokenHandler.
type MockTokenHandlerMockRecorder struct {
	mock *MockTokenHandler
}

// NewMockTokenHandler creates a new mock instance.
func NewMockTokenHandler(ctrl *gomock.Controller) *MockTokenHandler {
	mock := &MockTokenHandler{ctrl: ctrl}
	mock.recorder = &MockTokenHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTokenHandler) EXPECT() *MockTokenHandlerMockRecorder {
	return m.recorder
}

// Generate mocks base method.
func (m *MockTokenHandler) Generate(arg0 context.Context, arg1 game.Player, arg2 time.Duration) (auth.Token, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generate", arg0, arg1, arg2)
	ret0, _ := ret[0].(auth.Token)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Generate indicates an expected call of Generate.
func (mr *MockTokenHandlerMockRecorder) Generate(arg0, arg1, arg2 any) *MockTokenHandlerGenerateCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generate", reflect.TypeOf((*MockTokenHandler)(nil).Generate), arg0, arg1, arg2)
	return &MockTokenHandlerGenerateCall{Call: call}
}

// MockTokenHandlerGenerateCall wrap *gomock.Call
type MockTokenHandlerGenerateCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockTokenHandlerGenerateCall) Return(arg0 auth.Token, arg1 error) *MockTokenHandlerGenerateCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockTokenHandlerGenerateCall) Do(f func(context.Context, game.Player, time.Duration) (auth.Token, error)) *MockTokenHandlerGenerateCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockTokenHandlerGenerateCall) DoAndReturn(f func(context.Context, game.Player, time.Duration) (auth.Token, error)) *MockTokenHandlerGenerateCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Validate mocks base method.
func (m *MockTokenHandler) Validate(arg0 context.Context, arg1 auth.Token) (game.Player, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validate", arg0, arg1)
	ret0, _ := ret[0].(game.Player)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Validate indicates an expected call of Validate.
func (mr *MockTokenHandlerMockRecorder) Validate(arg0, arg1 any) *MockTokenHandlerValidateCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockTokenHandler)(nil).Validate), arg0, arg1)
	return &MockTokenHandlerValidateCall{Call: call}
}

// MockTokenHandlerValidateCall wrap *gomock.Call
type MockTokenHandlerValidateCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockTokenHandlerValidateCall) Return(arg0 game.Player, arg1 error) *MockTokenHandlerValidateCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockTokenHandlerValidateCall) Do(f func(context.Context, auth.Token) (game.Player, error)) *MockTokenHandlerValidateCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockTokenHandlerValidateCall) DoAndReturn(f func(context.Context, auth.Token) (game.Player, error)) *MockTokenHandlerValidateCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

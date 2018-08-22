// Code generated by MockGen. DO NOT EDIT.
// Source: go.uber.org/yarpc/api/middleware (interfaces: Router,UnaryInbound,UnaryOutbound,StreamInbound,StreamOutbound)

// Copyright (c) 2018 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Package middlewaretest is a generated GoMock package.
package middlewaretest

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	transport "go.uber.org/yarpc/api/transport"
	reflect "reflect"
)

// MockRouter is a mock of Router interface
type MockRouter struct {
	ctrl     *gomock.Controller
	recorder *MockRouterMockRecorder
}

// MockRouterMockRecorder is the mock recorder for MockRouter
type MockRouterMockRecorder struct {
	mock *MockRouter
}

// NewMockRouter creates a new mock instance
func NewMockRouter(ctrl *gomock.Controller) *MockRouter {
	mock := &MockRouter{ctrl: ctrl}
	mock.recorder = &MockRouterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRouter) EXPECT() *MockRouterMockRecorder {
	return m.recorder
}

// Choose mocks base method
func (m *MockRouter) Choose(arg0 context.Context, arg1 *transport.Request, arg2 transport.Router) (transport.HandlerSpec, error) {
	ret := m.ctrl.Call(m, "Choose", arg0, arg1, arg2)
	ret0, _ := ret[0].(transport.HandlerSpec)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Choose indicates an expected call of Choose
func (mr *MockRouterMockRecorder) Choose(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Choose", reflect.TypeOf((*MockRouter)(nil).Choose), arg0, arg1, arg2)
}

// Procedures mocks base method
func (m *MockRouter) Procedures(arg0 transport.Router) []transport.Procedure {
	ret := m.ctrl.Call(m, "Procedures", arg0)
	ret0, _ := ret[0].([]transport.Procedure)
	return ret0
}

// Procedures indicates an expected call of Procedures
func (mr *MockRouterMockRecorder) Procedures(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Procedures", reflect.TypeOf((*MockRouter)(nil).Procedures), arg0)
}

// MockUnaryInbound is a mock of UnaryInbound interface
type MockUnaryInbound struct {
	ctrl     *gomock.Controller
	recorder *MockUnaryInboundMockRecorder
}

// MockUnaryInboundMockRecorder is the mock recorder for MockUnaryInbound
type MockUnaryInboundMockRecorder struct {
	mock *MockUnaryInbound
}

// NewMockUnaryInbound creates a new mock instance
func NewMockUnaryInbound(ctrl *gomock.Controller) *MockUnaryInbound {
	mock := &MockUnaryInbound{ctrl: ctrl}
	mock.recorder = &MockUnaryInboundMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUnaryInbound) EXPECT() *MockUnaryInboundMockRecorder {
	return m.recorder
}

// Handle mocks base method
func (m *MockUnaryInbound) Handle(arg0 context.Context, arg1 *transport.Request, arg2 transport.ResponseWriter, arg3 transport.UnaryHandler) error {
	ret := m.ctrl.Call(m, "Handle", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// Handle indicates an expected call of Handle
func (mr *MockUnaryInboundMockRecorder) Handle(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Handle", reflect.TypeOf((*MockUnaryInbound)(nil).Handle), arg0, arg1, arg2, arg3)
}

// MockUnaryOutbound is a mock of UnaryOutbound interface
type MockUnaryOutbound struct {
	ctrl     *gomock.Controller
	recorder *MockUnaryOutboundMockRecorder
}

// MockUnaryOutboundMockRecorder is the mock recorder for MockUnaryOutbound
type MockUnaryOutboundMockRecorder struct {
	mock *MockUnaryOutbound
}

// NewMockUnaryOutbound creates a new mock instance
func NewMockUnaryOutbound(ctrl *gomock.Controller) *MockUnaryOutbound {
	mock := &MockUnaryOutbound{ctrl: ctrl}
	mock.recorder = &MockUnaryOutboundMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUnaryOutbound) EXPECT() *MockUnaryOutboundMockRecorder {
	return m.recorder
}

// Call mocks base method
func (m *MockUnaryOutbound) Call(arg0 context.Context, arg1 *transport.Request, arg2 transport.UnaryOutbound) (*transport.Response, error) {
	ret := m.ctrl.Call(m, "Call", arg0, arg1, arg2)
	ret0, _ := ret[0].(*transport.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Call indicates an expected call of Call
func (mr *MockUnaryOutboundMockRecorder) Call(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Call", reflect.TypeOf((*MockUnaryOutbound)(nil).Call), arg0, arg1, arg2)
}

// MockStreamInbound is a mock of StreamInbound interface
type MockStreamInbound struct {
	ctrl     *gomock.Controller
	recorder *MockStreamInboundMockRecorder
}

// MockStreamInboundMockRecorder is the mock recorder for MockStreamInbound
type MockStreamInboundMockRecorder struct {
	mock *MockStreamInbound
}

// NewMockStreamInbound creates a new mock instance
func NewMockStreamInbound(ctrl *gomock.Controller) *MockStreamInbound {
	mock := &MockStreamInbound{ctrl: ctrl}
	mock.recorder = &MockStreamInboundMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStreamInbound) EXPECT() *MockStreamInboundMockRecorder {
	return m.recorder
}

// HandleStream mocks base method
func (m *MockStreamInbound) HandleStream(arg0 *transport.ServerStream, arg1 transport.StreamHandler) error {
	ret := m.ctrl.Call(m, "HandleStream", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// HandleStream indicates an expected call of HandleStream
func (mr *MockStreamInboundMockRecorder) HandleStream(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleStream", reflect.TypeOf((*MockStreamInbound)(nil).HandleStream), arg0, arg1)
}

// MockStreamOutbound is a mock of StreamOutbound interface
type MockStreamOutbound struct {
	ctrl     *gomock.Controller
	recorder *MockStreamOutboundMockRecorder
}

// MockStreamOutboundMockRecorder is the mock recorder for MockStreamOutbound
type MockStreamOutboundMockRecorder struct {
	mock *MockStreamOutbound
}

// NewMockStreamOutbound creates a new mock instance
func NewMockStreamOutbound(ctrl *gomock.Controller) *MockStreamOutbound {
	mock := &MockStreamOutbound{ctrl: ctrl}
	mock.recorder = &MockStreamOutboundMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStreamOutbound) EXPECT() *MockStreamOutboundMockRecorder {
	return m.recorder
}

// CallStream mocks base method
func (m *MockStreamOutbound) CallStream(arg0 context.Context, arg1 *transport.StreamRequest, arg2 transport.StreamOutbound) (*transport.ClientStream, error) {
	ret := m.ctrl.Call(m, "CallStream", arg0, arg1, arg2)
	ret0, _ := ret[0].(*transport.ClientStream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CallStream indicates an expected call of CallStream
func (mr *MockStreamOutboundMockRecorder) CallStream(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CallStream", reflect.TypeOf((*MockStreamOutbound)(nil).CallStream), arg0, arg1, arg2)
}

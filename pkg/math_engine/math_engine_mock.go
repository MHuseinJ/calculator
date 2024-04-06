// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/math_engine/math_engine.go

// Package math_engine is a generated GoMock package.
package math_engine

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockMathEngine is a mock of MathEngine interface.
type MockMathEngine struct {
	ctrl     *gomock.Controller
	recorder *MockMathEngineMockRecorder
}

// MockMathEngineMockRecorder is the mock recorder for MockMathEngine.
type MockMathEngineMockRecorder struct {
	mock *MockMathEngine
}

// NewMockMathEngine creates a new mock instance.
func NewMockMathEngine(ctrl *gomock.Controller) *MockMathEngine {
	mock := &MockMathEngine{ctrl: ctrl}
	mock.recorder = &MockMathEngineMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMathEngine) EXPECT() *MockMathEngineMockRecorder {
	return m.recorder
}

// Abs mocks base method.
func (m *MockMathEngine) Abs(x float64) float64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Abs", x)
	ret0, _ := ret[0].(float64)
	return ret0
}

// Abs indicates an expected call of Abs.
func (mr *MockMathEngineMockRecorder) Abs(x interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Abs", reflect.TypeOf((*MockMathEngine)(nil).Abs), x)
}

// Add mocks base method.
func (m *MockMathEngine) Add(x, y float64) float64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", x, y)
	ret0, _ := ret[0].(float64)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockMathEngineMockRecorder) Add(x, y interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockMathEngine)(nil).Add), x, y)
}

// Cube mocks base method.
func (m *MockMathEngine) Cube(x float64) float64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cube", x)
	ret0, _ := ret[0].(float64)
	return ret0
}

// Cube indicates an expected call of Cube.
func (mr *MockMathEngineMockRecorder) Cube(x interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cube", reflect.TypeOf((*MockMathEngine)(nil).Cube), x)
}

// Cubert mocks base method.
func (m *MockMathEngine) Cubert(x float64) float64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cubert", x)
	ret0, _ := ret[0].(float64)
	return ret0
}

// Cubert indicates an expected call of Cubert.
func (mr *MockMathEngineMockRecorder) Cubert(x interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cubert", reflect.TypeOf((*MockMathEngine)(nil).Cubert), x)
}

// Divide mocks base method.
func (m *MockMathEngine) Divide(x, y float64) float64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Divide", x, y)
	ret0, _ := ret[0].(float64)
	return ret0
}

// Divide indicates an expected call of Divide.
func (mr *MockMathEngineMockRecorder) Divide(x, y interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Divide", reflect.TypeOf((*MockMathEngine)(nil).Divide), x, y)
}

// Multiply mocks base method.
func (m *MockMathEngine) Multiply(x, y float64) float64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Multiply", x, y)
	ret0, _ := ret[0].(float64)
	return ret0
}

// Multiply indicates an expected call of Multiply.
func (mr *MockMathEngineMockRecorder) Multiply(x, y interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Multiply", reflect.TypeOf((*MockMathEngine)(nil).Multiply), x, y)
}

// Neg mocks base method.
func (m *MockMathEngine) Neg(x float64) float64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Neg", x)
	ret0, _ := ret[0].(float64)
	return ret0
}

// Neg indicates an expected call of Neg.
func (mr *MockMathEngineMockRecorder) Neg(x interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Neg", reflect.TypeOf((*MockMathEngine)(nil).Neg), x)
}

// Sqr mocks base method.
func (m *MockMathEngine) Sqr(x float64) float64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sqr", x)
	ret0, _ := ret[0].(float64)
	return ret0
}

// Sqr indicates an expected call of Sqr.
func (mr *MockMathEngineMockRecorder) Sqr(x interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sqr", reflect.TypeOf((*MockMathEngine)(nil).Sqr), x)
}

// Sqrt mocks base method.
func (m *MockMathEngine) Sqrt(x float64) float64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sqrt", x)
	ret0, _ := ret[0].(float64)
	return ret0
}

// Sqrt indicates an expected call of Sqrt.
func (mr *MockMathEngineMockRecorder) Sqrt(x interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sqrt", reflect.TypeOf((*MockMathEngine)(nil).Sqrt), x)
}

// Subtract mocks base method.
func (m *MockMathEngine) Subtract(x, y float64) float64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Subtract", x, y)
	ret0, _ := ret[0].(float64)
	return ret0
}

// Subtract indicates an expected call of Subtract.
func (mr *MockMathEngineMockRecorder) Subtract(x, y interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subtract", reflect.TypeOf((*MockMathEngine)(nil).Subtract), x, y)
}

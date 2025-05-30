// Code generated by mockery; DO NOT EDIT.
// github.com/vektra/mockery
// template: testify

package mocks

import (
	"github.com/gin-gonic/gin"
	mock "github.com/stretchr/testify/mock"
)

// NewMockUserHandler creates a new instance of MockUserHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockUserHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockUserHandler {
	mock := &MockUserHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

// MockUserHandler is an autogenerated mock type for the IUserHandler type
type MockUserHandler struct {
	mock.Mock
}

type MockUserHandler_Expecter struct {
	mock *mock.Mock
}

func (_m *MockUserHandler) EXPECT() *MockUserHandler_Expecter {
	return &MockUserHandler_Expecter{mock: &_m.Mock}
}

// CreateUserHandler provides a mock function for the type MockUserHandler
func (_mock *MockUserHandler) CreateUserHandler(c *gin.Context) {
	_mock.Called(c)
	return
}

// MockUserHandler_CreateUserHandler_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateUserHandler'
type MockUserHandler_CreateUserHandler_Call struct {
	*mock.Call
}

// CreateUserHandler is a helper method to define mock.On call
//   - c
func (_e *MockUserHandler_Expecter) CreateUserHandler(c interface{}) *MockUserHandler_CreateUserHandler_Call {
	return &MockUserHandler_CreateUserHandler_Call{Call: _e.mock.On("CreateUserHandler", c)}
}

func (_c *MockUserHandler_CreateUserHandler_Call) Run(run func(c *gin.Context)) *MockUserHandler_CreateUserHandler_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gin.Context))
	})
	return _c
}

func (_c *MockUserHandler_CreateUserHandler_Call) Return() *MockUserHandler_CreateUserHandler_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockUserHandler_CreateUserHandler_Call) RunAndReturn(run func(c *gin.Context)) *MockUserHandler_CreateUserHandler_Call {
	_c.Run(run)
	return _c
}

// DeleteUserHandler provides a mock function for the type MockUserHandler
func (_mock *MockUserHandler) DeleteUserHandler(c *gin.Context) {
	_mock.Called(c)
	return
}

// MockUserHandler_DeleteUserHandler_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteUserHandler'
type MockUserHandler_DeleteUserHandler_Call struct {
	*mock.Call
}

// DeleteUserHandler is a helper method to define mock.On call
//   - c
func (_e *MockUserHandler_Expecter) DeleteUserHandler(c interface{}) *MockUserHandler_DeleteUserHandler_Call {
	return &MockUserHandler_DeleteUserHandler_Call{Call: _e.mock.On("DeleteUserHandler", c)}
}

func (_c *MockUserHandler_DeleteUserHandler_Call) Run(run func(c *gin.Context)) *MockUserHandler_DeleteUserHandler_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gin.Context))
	})
	return _c
}

func (_c *MockUserHandler_DeleteUserHandler_Call) Return() *MockUserHandler_DeleteUserHandler_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockUserHandler_DeleteUserHandler_Call) RunAndReturn(run func(c *gin.Context)) *MockUserHandler_DeleteUserHandler_Call {
	_c.Run(run)
	return _c
}

// EditUserHandler provides a mock function for the type MockUserHandler
func (_mock *MockUserHandler) EditUserHandler(c *gin.Context) {
	_mock.Called(c)
	return
}

// MockUserHandler_EditUserHandler_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'EditUserHandler'
type MockUserHandler_EditUserHandler_Call struct {
	*mock.Call
}

// EditUserHandler is a helper method to define mock.On call
//   - c
func (_e *MockUserHandler_Expecter) EditUserHandler(c interface{}) *MockUserHandler_EditUserHandler_Call {
	return &MockUserHandler_EditUserHandler_Call{Call: _e.mock.On("EditUserHandler", c)}
}

func (_c *MockUserHandler_EditUserHandler_Call) Run(run func(c *gin.Context)) *MockUserHandler_EditUserHandler_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gin.Context))
	})
	return _c
}

func (_c *MockUserHandler_EditUserHandler_Call) Return() *MockUserHandler_EditUserHandler_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockUserHandler_EditUserHandler_Call) RunAndReturn(run func(c *gin.Context)) *MockUserHandler_EditUserHandler_Call {
	_c.Run(run)
	return _c
}

// ListUserHandler provides a mock function for the type MockUserHandler
func (_mock *MockUserHandler) ListUserHandler(c *gin.Context) {
	_mock.Called(c)
	return
}

// MockUserHandler_ListUserHandler_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListUserHandler'
type MockUserHandler_ListUserHandler_Call struct {
	*mock.Call
}

// ListUserHandler is a helper method to define mock.On call
//   - c
func (_e *MockUserHandler_Expecter) ListUserHandler(c interface{}) *MockUserHandler_ListUserHandler_Call {
	return &MockUserHandler_ListUserHandler_Call{Call: _e.mock.On("ListUserHandler", c)}
}

func (_c *MockUserHandler_ListUserHandler_Call) Run(run func(c *gin.Context)) *MockUserHandler_ListUserHandler_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gin.Context))
	})
	return _c
}

func (_c *MockUserHandler_ListUserHandler_Call) Return() *MockUserHandler_ListUserHandler_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockUserHandler_ListUserHandler_Call) RunAndReturn(run func(c *gin.Context)) *MockUserHandler_ListUserHandler_Call {
	_c.Run(run)
	return _c
}

// LoginUserHandler provides a mock function for the type MockUserHandler
func (_mock *MockUserHandler) LoginUserHandler(c *gin.Context) {
	_mock.Called(c)
	return
}

// MockUserHandler_LoginUserHandler_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LoginUserHandler'
type MockUserHandler_LoginUserHandler_Call struct {
	*mock.Call
}

// LoginUserHandler is a helper method to define mock.On call
//   - c
func (_e *MockUserHandler_Expecter) LoginUserHandler(c interface{}) *MockUserHandler_LoginUserHandler_Call {
	return &MockUserHandler_LoginUserHandler_Call{Call: _e.mock.On("LoginUserHandler", c)}
}

func (_c *MockUserHandler_LoginUserHandler_Call) Run(run func(c *gin.Context)) *MockUserHandler_LoginUserHandler_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gin.Context))
	})
	return _c
}

func (_c *MockUserHandler_LoginUserHandler_Call) Return() *MockUserHandler_LoginUserHandler_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockUserHandler_LoginUserHandler_Call) RunAndReturn(run func(c *gin.Context)) *MockUserHandler_LoginUserHandler_Call {
	_c.Run(run)
	return _c
}

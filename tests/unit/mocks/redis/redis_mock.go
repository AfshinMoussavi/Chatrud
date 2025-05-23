// Code generated by mockery; DO NOT EDIT.
// github.com/vektra/mockery
// template: testify

package mocks

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	mock "github.com/stretchr/testify/mock"
)

// NewMockRedis creates a new instance of MockRedis. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockRedis(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockRedis {
	mock := &MockRedis{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

// MockRedis is an autogenerated mock type for the IRedis type
type MockRedis struct {
	mock.Mock
}

type MockRedis_Expecter struct {
	mock *mock.Mock
}

func (_m *MockRedis) EXPECT() *MockRedis_Expecter {
	return &MockRedis_Expecter{mock: &_m.Mock}
}

// Del provides a mock function for the type MockRedis
func (_mock *MockRedis) Del(ctx context.Context, keys ...string) *redis.IntCmd {
	var tmpRet mock.Arguments
	if len(keys) > 0 {
		tmpRet = _mock.Called(ctx, keys)
	} else {
		tmpRet = _mock.Called(ctx)
	}
	ret := tmpRet

	if len(ret) == 0 {
		panic("no return value specified for Del")
	}

	var r0 *redis.IntCmd
	if returnFunc, ok := ret.Get(0).(func(context.Context, ...string) *redis.IntCmd); ok {
		r0 = returnFunc(ctx, keys...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*redis.IntCmd)
		}
	}
	return r0
}

// MockRedis_Del_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Del'
type MockRedis_Del_Call struct {
	*mock.Call
}

// Del is a helper method to define mock.On call
//   - ctx
//   - keys
func (_e *MockRedis_Expecter) Del(ctx interface{}, keys ...interface{}) *MockRedis_Del_Call {
	return &MockRedis_Del_Call{Call: _e.mock.On("Del",
		append([]interface{}{ctx}, keys...)...)}
}

func (_c *MockRedis_Del_Call) Run(run func(ctx context.Context, keys ...string)) *MockRedis_Del_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := args[1].([]string)
		run(args[0].(context.Context), variadicArgs...)
	})
	return _c
}

func (_c *MockRedis_Del_Call) Return(intCmd *redis.IntCmd) *MockRedis_Del_Call {
	_c.Call.Return(intCmd)
	return _c
}

func (_c *MockRedis_Del_Call) RunAndReturn(run func(ctx context.Context, keys ...string) *redis.IntCmd) *MockRedis_Del_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function for the type MockRedis
func (_mock *MockRedis) Get(ctx context.Context, key string) *redis.StringCmd {
	ret := _mock.Called(ctx, key)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *redis.StringCmd
	if returnFunc, ok := ret.Get(0).(func(context.Context, string) *redis.StringCmd); ok {
		r0 = returnFunc(ctx, key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*redis.StringCmd)
		}
	}
	return r0
}

// MockRedis_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockRedis_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx
//   - key
func (_e *MockRedis_Expecter) Get(ctx interface{}, key interface{}) *MockRedis_Get_Call {
	return &MockRedis_Get_Call{Call: _e.mock.On("Get", ctx, key)}
}

func (_c *MockRedis_Get_Call) Run(run func(ctx context.Context, key string)) *MockRedis_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockRedis_Get_Call) Return(stringCmd *redis.StringCmd) *MockRedis_Get_Call {
	_c.Call.Return(stringCmd)
	return _c
}

func (_c *MockRedis_Get_Call) RunAndReturn(run func(ctx context.Context, key string) *redis.StringCmd) *MockRedis_Get_Call {
	_c.Call.Return(run)
	return _c
}

// Set provides a mock function for the type MockRedis
func (_mock *MockRedis) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	ret := _mock.Called(ctx, key, value, expiration)

	if len(ret) == 0 {
		panic("no return value specified for Set")
	}

	var r0 *redis.StatusCmd
	if returnFunc, ok := ret.Get(0).(func(context.Context, string, interface{}, time.Duration) *redis.StatusCmd); ok {
		r0 = returnFunc(ctx, key, value, expiration)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*redis.StatusCmd)
		}
	}
	return r0
}

// MockRedis_Set_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Set'
type MockRedis_Set_Call struct {
	*mock.Call
}

// Set is a helper method to define mock.On call
//   - ctx
//   - key
//   - value
//   - expiration
func (_e *MockRedis_Expecter) Set(ctx interface{}, key interface{}, value interface{}, expiration interface{}) *MockRedis_Set_Call {
	return &MockRedis_Set_Call{Call: _e.mock.On("Set", ctx, key, value, expiration)}
}

func (_c *MockRedis_Set_Call) Run(run func(ctx context.Context, key string, value interface{}, expiration time.Duration)) *MockRedis_Set_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(interface{}), args[3].(time.Duration))
	})
	return _c
}

func (_c *MockRedis_Set_Call) Return(statusCmd *redis.StatusCmd) *MockRedis_Set_Call {
	_c.Call.Return(statusCmd)
	return _c
}

func (_c *MockRedis_Set_Call) RunAndReturn(run func(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd) *MockRedis_Set_Call {
	_c.Call.Return(run)
	return _c
}

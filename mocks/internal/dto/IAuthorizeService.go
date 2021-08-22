// Code generated by mockery 2.9.0. DO NOT EDIT.

package mocks

import (
	context "context"
	http "net/http"

	dto "github.com/manabie-com/togo/internal/dto"

	mock "github.com/stretchr/testify/mock"

	tools "github.com/manabie-com/togo/internal/tools"
)

// IAuthorizeService is an autogenerated mock type for the IAuthorizeService type
type IAuthorizeService struct {
	mock.Mock
}

// Login provides a mock function with given fields: ctx, req
func (_m *IAuthorizeService) Login(ctx context.Context, req dto.LoginRequest) (*dto.LoginResponse, *tools.TodoError) {
	ret := _m.Called(ctx, req)

	var r0 *dto.LoginResponse
	if rf, ok := ret.Get(0).(func(context.Context, dto.LoginRequest) *dto.LoginResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.LoginResponse)
		}
	}

	var r1 *tools.TodoError
	if rf, ok := ret.Get(1).(func(context.Context, dto.LoginRequest) *tools.TodoError); ok {
		r1 = rf(ctx, req)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*tools.TodoError)
		}
	}

	return r0, r1
}

// Validate provides a mock function with given fields: req
func (_m *IAuthorizeService) Validate(req *http.Request) (context.Context, *tools.TodoError) {
	ret := _m.Called(req)

	var r0 context.Context
	if rf, ok := ret.Get(0).(func(*http.Request) context.Context); ok {
		r0 = rf(req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(context.Context)
		}
	}

	var r1 *tools.TodoError
	if rf, ok := ret.Get(1).(func(*http.Request) *tools.TodoError); ok {
		r1 = rf(req)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*tools.TodoError)
		}
	}

	return r0, r1
}

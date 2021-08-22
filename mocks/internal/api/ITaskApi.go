// Code generated by mockery 2.9.0. DO NOT EDIT.

package mocks

import (
	context "context"
	http "net/http"

	dto "github.com/manabie-com/togo/internal/dto"

	mock "github.com/stretchr/testify/mock"

	tools "github.com/manabie-com/togo/internal/tools"
)

// ITaskApi is an autogenerated mock type for the ITaskApi type
type ITaskApi struct {
	mock.Mock
}

// AddTask provides a mock function with given fields: ctx, req
func (_m *ITaskApi) AddTask(ctx context.Context, req *http.Request) (*dto.AddTaskResponse, *tools.TodoError) {
	ret := _m.Called(ctx, req)

	var r0 *dto.AddTaskResponse
	if rf, ok := ret.Get(0).(func(context.Context, *http.Request) *dto.AddTaskResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.AddTaskResponse)
		}
	}

	var r1 *tools.TodoError
	if rf, ok := ret.Get(1).(func(context.Context, *http.Request) *tools.TodoError); ok {
		r1 = rf(ctx, req)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*tools.TodoError)
		}
	}

	return r0, r1
}

// ListTasksByUserAndDate provides a mock function with given fields: ctx, req
func (_m *ITaskApi) ListTasksByUserAndDate(ctx context.Context, req *http.Request) (*dto.ListTaskResponse, *tools.TodoError) {
	ret := _m.Called(ctx, req)

	var r0 *dto.ListTaskResponse
	if rf, ok := ret.Get(0).(func(context.Context, *http.Request) *dto.ListTaskResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.ListTaskResponse)
		}
	}

	var r1 *tools.TodoError
	if rf, ok := ret.Get(1).(func(context.Context, *http.Request) *tools.TodoError); ok {
		r1 = rf(ctx, req)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*tools.TodoError)
		}
	}

	return r0, r1
}

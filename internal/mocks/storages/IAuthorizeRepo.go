// Code generated by mockery 2.9.0. DO NOT EDIT.

package mocks

import (
	context "context"

	storages "github.com/manabie-com/togo/internal/storages"
	mock "github.com/stretchr/testify/mock"

	tools "github.com/manabie-com/togo/internal/tools"
)

// IAuthorizeRepo is an autogenerated mock type for the IAuthorizeRepo type
type IAuthorizeRepo struct {
	mock.Mock
}

// ValidateUserStore provides a mock function with given fields: ctx, arg
func (_m *IAuthorizeRepo) ValidateUserStore(ctx context.Context, arg storages.ValidateUserParams) (string, *tools.TodoError) {
	ret := _m.Called(ctx, arg)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, storages.ValidateUserParams) string); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 *tools.TodoError
	if rf, ok := ret.Get(1).(func(context.Context, storages.ValidateUserParams) *tools.TodoError); ok {
		r1 = rf(ctx, arg)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*tools.TodoError)
		}
	}

	return r0, r1
}
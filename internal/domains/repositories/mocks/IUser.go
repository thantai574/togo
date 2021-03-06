// Code generated by mockery v2.0.4. DO NOT EDIT.

package mocks

import (
	context "context"

	entities "github.com/manabie-com/togo/internal/domains/entities"
	mock "github.com/stretchr/testify/mock"

	sql "database/sql"
)

// IUser is an autogenerated mock type for the IUser type
type IUser struct {
	mock.Mock
}

// ValidateUser provides a mock function with given fields: ctx, userID, pwd
func (_m *IUser) ValidateUser(ctx context.Context, userID sql.NullString, pwd sql.NullString) (entities.User, error) {
	ret := _m.Called(ctx, userID, pwd)

	var r0 entities.User
	if rf, ok := ret.Get(0).(func(context.Context, sql.NullString, sql.NullString) entities.User); ok {
		r0 = rf(ctx, userID, pwd)
	} else {
		r0 = ret.Get(0).(entities.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, sql.NullString, sql.NullString) error); ok {
		r1 = rf(ctx, userID, pwd)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

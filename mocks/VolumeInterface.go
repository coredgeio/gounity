// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	gounity "github.com/dell/gounity"
	mock "github.com/stretchr/testify/mock"

	types "github.com/dell/gounity/types"
)

// VolumeInterface is an autogenerated mock type for the VolumeInterface type
type VolumeInterface struct {
	mock.Mock
}

// CreateLun provides a mock function with given fields: ctx, name, poolID, description, size, fastVPTieringPolicy, hostIOLimitID, isThinEnabled, isDataReductionEnabled
func (_m *VolumeInterface) CreateLun(ctx context.Context, name string, poolID string, description string, size uint64, fastVPTieringPolicy int, hostIOLimitID string, isThinEnabled bool, isDataReductionEnabled bool) (*types.Volume, error) {
	ret := _m.Called(ctx, name, poolID, description, size, fastVPTieringPolicy, hostIOLimitID, isThinEnabled, isDataReductionEnabled)

	var r0 *types.Volume
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, uint64, int, string, bool, bool) *types.Volume); ok {
		r0 = rf(ctx, name, poolID, description, size, fastVPTieringPolicy, hostIOLimitID, isThinEnabled, isDataReductionEnabled)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Volume)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, uint64, int, string, bool, bool) error); ok {
		r1 = rf(ctx, name, poolID, description, size, fastVPTieringPolicy, hostIOLimitID, isThinEnabled, isDataReductionEnabled)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindHostIOLimitByName provides a mock function with given fields: ctx, hostIoPolicyName
func (_m *VolumeInterface) FindHostIOLimitByName(ctx context.Context, hostIoPolicyName string) (*types.IoLimitPolicy, error) {
	ret := _m.Called(ctx, hostIoPolicyName)

	var r0 *types.IoLimitPolicy
	if rf, ok := ret.Get(0).(func(context.Context, string) *types.IoLimitPolicy); ok {
		r0 = rf(ctx, hostIoPolicyName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.IoLimitPolicy)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, hostIoPolicyName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindVolumeByName provides a mock function with given fields: ctx, volName
func (_m *VolumeInterface) FindVolumeByName(ctx context.Context, volName string) (*types.Volume, error) {
	ret := _m.Called(ctx, volName)

	var r0 *types.Volume
	if rf, ok := ret.Get(0).(func(context.Context, string) *types.Volume); ok {
		r0 = rf(ctx, volName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Volume)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, volName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewVolume provides a mock function with given fields: client
func (_m *VolumeInterface) NewVolume(client *gounity.Client) *gounity.Volume {
	ret := _m.Called(client)

	var r0 *gounity.Volume
	if rf, ok := ret.Get(0).(func(*gounity.Client) *gounity.Volume); ok {
		r0 = rf(client)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gounity.Volume)
		}
	}

	return r0
}

type mockConstructorTestingTNewVolumeInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewVolumeInterface creates a new instance of VolumeInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewVolumeInterface(t mockConstructorTestingTNewVolumeInterface) *VolumeInterface {
	mock := &VolumeInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

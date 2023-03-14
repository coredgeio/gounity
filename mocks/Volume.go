package mocks

import (
	context "context"

	"github.com/dell/gounity/types"
	mock "github.com/stretchr/testify/mock"
)

type Volume struct {
	mock.Mock
}

func (_m *Volume) CreateLun(ctx context.Context, name, poolID, description string, size uint64, fastVPTieringPolicy int,
	hostIOLimitID string, isThinEnabled, isDataReductionEnabled bool) (*types.Volume, error) {
	// args := m.Called(ctx, name, poolID, description, size, fastVPTieringPolicy, hostIOLimitID, isThinEnabled, isDataReductionEnabled)
	// var r0 *types.Volume
	// return r0, args.Error(1)

	ret := _m.Called(ctx, name, poolID, description, size, fastVPTieringPolicy, hostIOLimitID, isThinEnabled, isDataReductionEnabled)

	var r0 *types.Volume
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, uint64, int, string, bool, bool) *types.Volume); ok {
		r0 = rf(ctx, name, poolID, description, size, fastVPTieringPolicy,
			hostIOLimitID, isThinEnabled, isDataReductionEnabled)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Volume)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, uint64, int, string, bool, bool) error); ok {
		r1 = rf(ctx, name, poolID, description, size, fastVPTieringPolicy,
			hostIOLimitID, isThinEnabled, isDataReductionEnabled)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

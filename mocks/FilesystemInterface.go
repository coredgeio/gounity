// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	gounity "github.com/dell/gounity"
	mock "github.com/stretchr/testify/mock"

	types "github.com/dell/gounity/types"
)

// FilesystemInterface is an autogenerated mock type for the FilesystemInterface type
type FilesystemInterface struct {
	mock.Mock
}

// CreateFilesystem provides a mock function with given fields: ctx, name, storagepool, description, nasServer, size, tieringPolicy, hostIOSize, supportedProtocol, isThinEnabled, isDataReductionEnabled, isReplicationDestination
func (_m *FilesystemInterface) CreateFilesystem(ctx context.Context, name string, storagepool string, description string, nasServer string, size uint64, tieringPolicy int, hostIOSize int, supportedProtocol int, isThinEnabled bool, isDataReductionEnabled bool, isReplicationDestination bool) (*types.Filesystem, error) {
	ret := _m.Called(ctx, name, storagepool, description, nasServer, size, tieringPolicy, hostIOSize, supportedProtocol, isThinEnabled, isDataReductionEnabled, isReplicationDestination)

	var r0 *types.Filesystem
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string, uint64, int, int, int, bool, bool, bool) *types.Filesystem); ok {
		r0 = rf(ctx, name, storagepool, description, nasServer, size, tieringPolicy, hostIOSize, supportedProtocol, isThinEnabled, isDataReductionEnabled, isReplicationDestination)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Filesystem)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, string, uint64, int, int, int, bool, bool, bool) error); ok {
		r1 = rf(ctx, name, storagepool, description, nasServer, size, tieringPolicy, hostIOSize, supportedProtocol, isThinEnabled, isDataReductionEnabled, isReplicationDestination)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateNFSShare provides a mock function with given fields: ctx, name, path, filesystemID, nfsShareDefaultAccess
func (_m *FilesystemInterface) CreateNFSShare(ctx context.Context, name string, path string, filesystemID string, nfsShareDefaultAccess gounity.NFSShareDefaultAccess) (*types.Filesystem, error) {
	ret := _m.Called(ctx, name, path, filesystemID, nfsShareDefaultAccess)

	var r0 *types.Filesystem
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, gounity.NFSShareDefaultAccess) *types.Filesystem); ok {
		r0 = rf(ctx, name, path, filesystemID, nfsShareDefaultAccess)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Filesystem)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, gounity.NFSShareDefaultAccess) error); ok {
		r1 = rf(ctx, name, path, filesystemID, nfsShareDefaultAccess)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateNFSShareFromSnapshot provides a mock function with given fields: ctx, name, path, snapshotID, nfsShareDefaultAccess
func (_m *FilesystemInterface) CreateNFSShareFromSnapshot(ctx context.Context, name string, path string, snapshotID string, nfsShareDefaultAccess gounity.NFSShareDefaultAccess) (*types.NFSShare, error) {
	ret := _m.Called(ctx, name, path, snapshotID, nfsShareDefaultAccess)

	var r0 *types.NFSShare
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, gounity.NFSShareDefaultAccess) *types.NFSShare); ok {
		r0 = rf(ctx, name, path, snapshotID, nfsShareDefaultAccess)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.NFSShare)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, gounity.NFSShareDefaultAccess) error); ok {
		r1 = rf(ctx, name, path, snapshotID, nfsShareDefaultAccess)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteFilesystem provides a mock function with given fields: ctx, filesystemID
func (_m *FilesystemInterface) DeleteFilesystem(ctx context.Context, filesystemID string) error {
	ret := _m.Called(ctx, filesystemID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, filesystemID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteNFSShare provides a mock function with given fields: ctx, filesystemID, nfsShareID
func (_m *FilesystemInterface) DeleteNFSShare(ctx context.Context, filesystemID string, nfsShareID string) error {
	ret := _m.Called(ctx, filesystemID, nfsShareID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, filesystemID, nfsShareID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteNFSShareCreatedFromSnapshot provides a mock function with given fields: ctx, nfsShareID
func (_m *FilesystemInterface) DeleteNFSShareCreatedFromSnapshot(ctx context.Context, nfsShareID string) error {
	ret := _m.Called(ctx, nfsShareID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, nfsShareID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ExpandFilesystem provides a mock function with given fields: ctx, filesystemID, newSize
func (_m *FilesystemInterface) ExpandFilesystem(ctx context.Context, filesystemID string, newSize uint64) error {
	ret := _m.Called(ctx, filesystemID, newSize)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, uint64) error); ok {
		r0 = rf(ctx, filesystemID, newSize)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindFilesystemByID provides a mock function with given fields: ctx, filesystemID
func (_m *FilesystemInterface) FindFilesystemByID(ctx context.Context, filesystemID string) (*types.Filesystem, error) {
	ret := _m.Called(ctx, filesystemID)

	var r0 *types.Filesystem
	if rf, ok := ret.Get(0).(func(context.Context, string) *types.Filesystem); ok {
		r0 = rf(ctx, filesystemID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Filesystem)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, filesystemID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindFilesystemByName provides a mock function with given fields: ctx, filesystemName
func (_m *FilesystemInterface) FindFilesystemByName(ctx context.Context, filesystemName string) (*types.Filesystem, error) {
	ret := _m.Called(ctx, filesystemName)

	var r0 *types.Filesystem
	if rf, ok := ret.Get(0).(func(context.Context, string) *types.Filesystem); ok {
		r0 = rf(ctx, filesystemName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Filesystem)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, filesystemName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindNASServerByID provides a mock function with given fields: ctx, nasServerID
func (_m *FilesystemInterface) FindNASServerByID(ctx context.Context, nasServerID string) (*types.NASServer, error) {
	ret := _m.Called(ctx, nasServerID)

	var r0 *types.NASServer
	if rf, ok := ret.Get(0).(func(context.Context, string) *types.NASServer); ok {
		r0 = rf(ctx, nasServerID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.NASServer)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, nasServerID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindNFSShareByID provides a mock function with given fields: ctx, nfsShareID
func (_m *FilesystemInterface) FindNFSShareByID(ctx context.Context, nfsShareID string) (*types.NFSShare, error) {
	ret := _m.Called(ctx, nfsShareID)

	var r0 *types.NFSShare
	if rf, ok := ret.Get(0).(func(context.Context, string) *types.NFSShare); ok {
		r0 = rf(ctx, nfsShareID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.NFSShare)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, nfsShareID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindNFSShareByName provides a mock function with given fields: ctx, nfsSharename
func (_m *FilesystemInterface) FindNFSShareByName(ctx context.Context, nfsSharename string) (*types.NFSShare, error) {
	ret := _m.Called(ctx, nfsSharename)

	var r0 *types.NFSShare
	if rf, ok := ret.Get(0).(func(context.Context, string) *types.NFSShare); ok {
		r0 = rf(ctx, nfsSharename)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.NFSShare)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, nfsSharename)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetFilesystemIDFromResID provides a mock function with given fields: ctx, filesystemResID
func (_m *FilesystemInterface) GetFilesystemIDFromResID(ctx context.Context, filesystemResID string) (string, error) {
	ret := _m.Called(ctx, filesystemResID)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, string) string); ok {
		r0 = rf(ctx, filesystemResID)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, filesystemResID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ModifyNFSShareCreatedFromSnapshotHostAccess provides a mock function with given fields: ctx, nfsShareID, hostIDs, accessType
func (_m *FilesystemInterface) ModifyNFSShareCreatedFromSnapshotHostAccess(ctx context.Context, nfsShareID string, hostIDs []string, accessType gounity.AccessType) error {
	ret := _m.Called(ctx, nfsShareID, hostIDs, accessType)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, []string, gounity.AccessType) error); ok {
		r0 = rf(ctx, nfsShareID, hostIDs, accessType)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ModifyNFSShareHostAccess provides a mock function with given fields: ctx, filesystemID, nfsShareID, hostIDs, accessType
func (_m *FilesystemInterface) ModifyNFSShareHostAccess(ctx context.Context, filesystemID string, nfsShareID string, hostIDs []string, accessType gounity.AccessType) error {
	ret := _m.Called(ctx, filesystemID, nfsShareID, hostIDs, accessType)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, []string, gounity.AccessType) error); ok {
		r0 = rf(ctx, filesystemID, nfsShareID, hostIDs, accessType)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewFilesystemInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewFilesystemInterface creates a new instance of FilesystemInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewFilesystemInterface(t mockConstructorTestingTNewFilesystemInterface) *FilesystemInterface {
	mock := &FilesystemInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

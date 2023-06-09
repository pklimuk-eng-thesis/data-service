// Code generated by mockery v2.23.2. DO NOT EDIT.

package service

import (
	domain "github.com/pklimuk-eng-thesis/data-service/pkg/domain"
	mock "github.com/stretchr/testify/mock"
)

// MockDeviceService is an autogenerated mock type for the DeviceService type
type MockDeviceService struct {
	mock.Mock
}

type MockDeviceService_Expecter struct {
	mock *mock.Mock
}

func (_m *MockDeviceService) EXPECT() *MockDeviceService_Expecter {
	return &MockDeviceService_Expecter{mock: &_m.Mock}
}

// AddNewRecordToDeviceTable provides a mock function with given fields: isEnabled
func (_m *MockDeviceService) AddNewRecordToDeviceTable(isEnabled bool) error {
	ret := _m.Called(isEnabled)

	var r0 error
	if rf, ok := ret.Get(0).(func(bool) error); ok {
		r0 = rf(isEnabled)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockDeviceService_AddNewRecordToDeviceTable_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddNewRecordToDeviceTable'
type MockDeviceService_AddNewRecordToDeviceTable_Call struct {
	*mock.Call
}

// AddNewRecordToDeviceTable is a helper method to define mock.On call
//   - isEnabled bool
func (_e *MockDeviceService_Expecter) AddNewRecordToDeviceTable(isEnabled interface{}) *MockDeviceService_AddNewRecordToDeviceTable_Call {
	return &MockDeviceService_AddNewRecordToDeviceTable_Call{Call: _e.mock.On("AddNewRecordToDeviceTable", isEnabled)}
}

func (_c *MockDeviceService_AddNewRecordToDeviceTable_Call) Run(run func(isEnabled bool)) *MockDeviceService_AddNewRecordToDeviceTable_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(bool))
	})
	return _c
}

func (_c *MockDeviceService_AddNewRecordToDeviceTable_Call) Return(_a0 error) *MockDeviceService_AddNewRecordToDeviceTable_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDeviceService_AddNewRecordToDeviceTable_Call) RunAndReturn(run func(bool) error) *MockDeviceService_AddNewRecordToDeviceTable_Call {
	_c.Call.Return(run)
	return _c
}

// GetLatestDeviceRecordsLimitN provides a mock function with given fields: limit
func (_m *MockDeviceService) GetLatestDeviceRecordsLimitN(limit int) ([]domain.DeviceData, error) {
	ret := _m.Called(limit)

	var r0 []domain.DeviceData
	var r1 error
	if rf, ok := ret.Get(0).(func(int) ([]domain.DeviceData, error)); ok {
		return rf(limit)
	}
	if rf, ok := ret.Get(0).(func(int) []domain.DeviceData); ok {
		r0 = rf(limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.DeviceData)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDeviceService_GetLatestDeviceRecordsLimitN_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLatestDeviceRecordsLimitN'
type MockDeviceService_GetLatestDeviceRecordsLimitN_Call struct {
	*mock.Call
}

// GetLatestDeviceRecordsLimitN is a helper method to define mock.On call
//   - limit int
func (_e *MockDeviceService_Expecter) GetLatestDeviceRecordsLimitN(limit interface{}) *MockDeviceService_GetLatestDeviceRecordsLimitN_Call {
	return &MockDeviceService_GetLatestDeviceRecordsLimitN_Call{Call: _e.mock.On("GetLatestDeviceRecordsLimitN", limit)}
}

func (_c *MockDeviceService_GetLatestDeviceRecordsLimitN_Call) Run(run func(limit int)) *MockDeviceService_GetLatestDeviceRecordsLimitN_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int))
	})
	return _c
}

func (_c *MockDeviceService_GetLatestDeviceRecordsLimitN_Call) Return(_a0 []domain.DeviceData, _a1 error) *MockDeviceService_GetLatestDeviceRecordsLimitN_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockDeviceService_GetLatestDeviceRecordsLimitN_Call) RunAndReturn(run func(int) ([]domain.DeviceData, error)) *MockDeviceService_GetLatestDeviceRecordsLimitN_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockDeviceService interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockDeviceService creates a new instance of MockDeviceService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockDeviceService(t mockConstructorTestingTNewMockDeviceService) *MockDeviceService {
	mock := &MockDeviceService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

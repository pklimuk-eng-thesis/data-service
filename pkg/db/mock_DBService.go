// Code generated by mockery v2.23.2. DO NOT EDIT.

package db

import (
	domain "github.com/pklimuk-eng-thesis/data-service/pkg/domain"
	mock "github.com/stretchr/testify/mock"
)

// MockDBService is an autogenerated mock type for the DBService type
type MockDBService struct {
	mock.Mock
}

type MockDBService_Expecter struct {
	mock *mock.Mock
}

func (_m *MockDBService) EXPECT() *MockDBService_Expecter {
	return &MockDBService_Expecter{mock: &_m.Mock}
}

// AddNewRecordToACTable provides a mock function with given fields: tableName, isEnabled, temperature, humidity
func (_m *MockDBService) AddNewRecordToACTable(tableName string, isEnabled bool, temperature float32, humidity float32) error {
	ret := _m.Called(tableName, isEnabled, temperature, humidity)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, bool, float32, float32) error); ok {
		r0 = rf(tableName, isEnabled, temperature, humidity)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockDBService_AddNewRecordToACTable_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddNewRecordToACTable'
type MockDBService_AddNewRecordToACTable_Call struct {
	*mock.Call
}

// AddNewRecordToACTable is a helper method to define mock.On call
//   - tableName string
//   - isEnabled bool
//   - temperature float32
//   - humidity float32
func (_e *MockDBService_Expecter) AddNewRecordToACTable(tableName interface{}, isEnabled interface{}, temperature interface{}, humidity interface{}) *MockDBService_AddNewRecordToACTable_Call {
	return &MockDBService_AddNewRecordToACTable_Call{Call: _e.mock.On("AddNewRecordToACTable", tableName, isEnabled, temperature, humidity)}
}

func (_c *MockDBService_AddNewRecordToACTable_Call) Run(run func(tableName string, isEnabled bool, temperature float32, humidity float32)) *MockDBService_AddNewRecordToACTable_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(bool), args[2].(float32), args[3].(float32))
	})
	return _c
}

func (_c *MockDBService_AddNewRecordToACTable_Call) Return(_a0 error) *MockDBService_AddNewRecordToACTable_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDBService_AddNewRecordToACTable_Call) RunAndReturn(run func(string, bool, float32, float32) error) *MockDBService_AddNewRecordToACTable_Call {
	_c.Call.Return(run)
	return _c
}

// AddNewRecordToDeviceTable provides a mock function with given fields: tableName, isEnabled
func (_m *MockDBService) AddNewRecordToDeviceTable(tableName string, isEnabled bool) error {
	ret := _m.Called(tableName, isEnabled)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, bool) error); ok {
		r0 = rf(tableName, isEnabled)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockDBService_AddNewRecordToDeviceTable_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddNewRecordToDeviceTable'
type MockDBService_AddNewRecordToDeviceTable_Call struct {
	*mock.Call
}

// AddNewRecordToDeviceTable is a helper method to define mock.On call
//   - tableName string
//   - isEnabled bool
func (_e *MockDBService_Expecter) AddNewRecordToDeviceTable(tableName interface{}, isEnabled interface{}) *MockDBService_AddNewRecordToDeviceTable_Call {
	return &MockDBService_AddNewRecordToDeviceTable_Call{Call: _e.mock.On("AddNewRecordToDeviceTable", tableName, isEnabled)}
}

func (_c *MockDBService_AddNewRecordToDeviceTable_Call) Run(run func(tableName string, isEnabled bool)) *MockDBService_AddNewRecordToDeviceTable_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(bool))
	})
	return _c
}

func (_c *MockDBService_AddNewRecordToDeviceTable_Call) Return(_a0 error) *MockDBService_AddNewRecordToDeviceTable_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDBService_AddNewRecordToDeviceTable_Call) RunAndReturn(run func(string, bool) error) *MockDBService_AddNewRecordToDeviceTable_Call {
	_c.Call.Return(run)
	return _c
}

// AddNewRecordToSensorTable provides a mock function with given fields: tableName, isEnabled, detected
func (_m *MockDBService) AddNewRecordToSensorTable(tableName string, isEnabled bool, detected bool) error {
	ret := _m.Called(tableName, isEnabled, detected)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, bool, bool) error); ok {
		r0 = rf(tableName, isEnabled, detected)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockDBService_AddNewRecordToSensorTable_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddNewRecordToSensorTable'
type MockDBService_AddNewRecordToSensorTable_Call struct {
	*mock.Call
}

// AddNewRecordToSensorTable is a helper method to define mock.On call
//   - tableName string
//   - isEnabled bool
//   - detected bool
func (_e *MockDBService_Expecter) AddNewRecordToSensorTable(tableName interface{}, isEnabled interface{}, detected interface{}) *MockDBService_AddNewRecordToSensorTable_Call {
	return &MockDBService_AddNewRecordToSensorTable_Call{Call: _e.mock.On("AddNewRecordToSensorTable", tableName, isEnabled, detected)}
}

func (_c *MockDBService_AddNewRecordToSensorTable_Call) Run(run func(tableName string, isEnabled bool, detected bool)) *MockDBService_AddNewRecordToSensorTable_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(bool), args[2].(bool))
	})
	return _c
}

func (_c *MockDBService_AddNewRecordToSensorTable_Call) Return(_a0 error) *MockDBService_AddNewRecordToSensorTable_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDBService_AddNewRecordToSensorTable_Call) RunAndReturn(run func(string, bool, bool) error) *MockDBService_AddNewRecordToSensorTable_Call {
	_c.Call.Return(run)
	return _c
}

// GetLatestACDataByTableNameLimitN provides a mock function with given fields: tableName, limit
func (_m *MockDBService) GetLatestACDataByTableNameLimitN(tableName string, limit int) ([]domain.ACData, error) {
	ret := _m.Called(tableName, limit)

	var r0 []domain.ACData
	var r1 error
	if rf, ok := ret.Get(0).(func(string, int) ([]domain.ACData, error)); ok {
		return rf(tableName, limit)
	}
	if rf, ok := ret.Get(0).(func(string, int) []domain.ACData); ok {
		r0 = rf(tableName, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.ACData)
		}
	}

	if rf, ok := ret.Get(1).(func(string, int) error); ok {
		r1 = rf(tableName, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDBService_GetLatestACDataByTableNameLimitN_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLatestACDataByTableNameLimitN'
type MockDBService_GetLatestACDataByTableNameLimitN_Call struct {
	*mock.Call
}

// GetLatestACDataByTableNameLimitN is a helper method to define mock.On call
//   - tableName string
//   - limit int
func (_e *MockDBService_Expecter) GetLatestACDataByTableNameLimitN(tableName interface{}, limit interface{}) *MockDBService_GetLatestACDataByTableNameLimitN_Call {
	return &MockDBService_GetLatestACDataByTableNameLimitN_Call{Call: _e.mock.On("GetLatestACDataByTableNameLimitN", tableName, limit)}
}

func (_c *MockDBService_GetLatestACDataByTableNameLimitN_Call) Run(run func(tableName string, limit int)) *MockDBService_GetLatestACDataByTableNameLimitN_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(int))
	})
	return _c
}

func (_c *MockDBService_GetLatestACDataByTableNameLimitN_Call) Return(_a0 []domain.ACData, _a1 error) *MockDBService_GetLatestACDataByTableNameLimitN_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockDBService_GetLatestACDataByTableNameLimitN_Call) RunAndReturn(run func(string, int) ([]domain.ACData, error)) *MockDBService_GetLatestACDataByTableNameLimitN_Call {
	_c.Call.Return(run)
	return _c
}

// GetLatestDeviceDataByTableNameLimitN provides a mock function with given fields: tableName, limit
func (_m *MockDBService) GetLatestDeviceDataByTableNameLimitN(tableName string, limit int) ([]domain.DeviceData, error) {
	ret := _m.Called(tableName, limit)

	var r0 []domain.DeviceData
	var r1 error
	if rf, ok := ret.Get(0).(func(string, int) ([]domain.DeviceData, error)); ok {
		return rf(tableName, limit)
	}
	if rf, ok := ret.Get(0).(func(string, int) []domain.DeviceData); ok {
		r0 = rf(tableName, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.DeviceData)
		}
	}

	if rf, ok := ret.Get(1).(func(string, int) error); ok {
		r1 = rf(tableName, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDBService_GetLatestDeviceDataByTableNameLimitN_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLatestDeviceDataByTableNameLimitN'
type MockDBService_GetLatestDeviceDataByTableNameLimitN_Call struct {
	*mock.Call
}

// GetLatestDeviceDataByTableNameLimitN is a helper method to define mock.On call
//   - tableName string
//   - limit int
func (_e *MockDBService_Expecter) GetLatestDeviceDataByTableNameLimitN(tableName interface{}, limit interface{}) *MockDBService_GetLatestDeviceDataByTableNameLimitN_Call {
	return &MockDBService_GetLatestDeviceDataByTableNameLimitN_Call{Call: _e.mock.On("GetLatestDeviceDataByTableNameLimitN", tableName, limit)}
}

func (_c *MockDBService_GetLatestDeviceDataByTableNameLimitN_Call) Run(run func(tableName string, limit int)) *MockDBService_GetLatestDeviceDataByTableNameLimitN_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(int))
	})
	return _c
}

func (_c *MockDBService_GetLatestDeviceDataByTableNameLimitN_Call) Return(_a0 []domain.DeviceData, _a1 error) *MockDBService_GetLatestDeviceDataByTableNameLimitN_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockDBService_GetLatestDeviceDataByTableNameLimitN_Call) RunAndReturn(run func(string, int) ([]domain.DeviceData, error)) *MockDBService_GetLatestDeviceDataByTableNameLimitN_Call {
	_c.Call.Return(run)
	return _c
}

// GetLatestSensorDataByTableNameLimitN provides a mock function with given fields: tableName, limit
func (_m *MockDBService) GetLatestSensorDataByTableNameLimitN(tableName string, limit int) ([]domain.SensorData, error) {
	ret := _m.Called(tableName, limit)

	var r0 []domain.SensorData
	var r1 error
	if rf, ok := ret.Get(0).(func(string, int) ([]domain.SensorData, error)); ok {
		return rf(tableName, limit)
	}
	if rf, ok := ret.Get(0).(func(string, int) []domain.SensorData); ok {
		r0 = rf(tableName, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.SensorData)
		}
	}

	if rf, ok := ret.Get(1).(func(string, int) error); ok {
		r1 = rf(tableName, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDBService_GetLatestSensorDataByTableNameLimitN_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLatestSensorDataByTableNameLimitN'
type MockDBService_GetLatestSensorDataByTableNameLimitN_Call struct {
	*mock.Call
}

// GetLatestSensorDataByTableNameLimitN is a helper method to define mock.On call
//   - tableName string
//   - limit int
func (_e *MockDBService_Expecter) GetLatestSensorDataByTableNameLimitN(tableName interface{}, limit interface{}) *MockDBService_GetLatestSensorDataByTableNameLimitN_Call {
	return &MockDBService_GetLatestSensorDataByTableNameLimitN_Call{Call: _e.mock.On("GetLatestSensorDataByTableNameLimitN", tableName, limit)}
}

func (_c *MockDBService_GetLatestSensorDataByTableNameLimitN_Call) Run(run func(tableName string, limit int)) *MockDBService_GetLatestSensorDataByTableNameLimitN_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(int))
	})
	return _c
}

func (_c *MockDBService_GetLatestSensorDataByTableNameLimitN_Call) Return(_a0 []domain.SensorData, _a1 error) *MockDBService_GetLatestSensorDataByTableNameLimitN_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockDBService_GetLatestSensorDataByTableNameLimitN_Call) RunAndReturn(run func(string, int) ([]domain.SensorData, error)) *MockDBService_GetLatestSensorDataByTableNameLimitN_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockDBService interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockDBService creates a new instance of MockDBService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockDBService(t mockConstructorTestingTNewMockDBService) *MockDBService {
	mock := &MockDBService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

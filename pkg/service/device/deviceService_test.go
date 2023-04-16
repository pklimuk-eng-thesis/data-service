package service

import (
	"fmt"
	"testing"
	"time"

	"github.com/pklimuk-eng-thesis/data-service/pkg/db"
	"github.com/pklimuk-eng-thesis/data-service/pkg/domain"
	"github.com/stretchr/testify/assert"
)

func TestGetLatestDeviceDataByTableNameLimitN_Success(t *testing.T) {
	tests := []struct {
		name     string
		device   domain.Device
		db       db.DBService
		limit    int
		expected []domain.DeviceData
	}{
		{
			name:   "GetLatestDeviceDataByTableNameLimitN_Success_limit_10",
			device: domain.Device{TableName: "test_table"},
			db:     new(db.MockDBService),
			limit:  10,
			expected: []domain.DeviceData{
				{ID: 1, CreatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC), IsEnabled: true},
				{ID: 2, CreatedAt: time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC), IsEnabled: false},
			},
		},
		{
			name:   "GetLatestDeviceDataByTableNameLimitN_Success_limit_1",
			device: domain.Device{TableName: "test_table"},
			db:     new(db.MockDBService),
			limit:  1,
			expected: []domain.DeviceData{
				{ID: 1, CreatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC), IsEnabled: true},
			},
		},
		{
			name:     "GetLatestDeviceDataByTableNameLimitN_Success_limit_0",
			device:   domain.Device{TableName: "test_table"},
			db:       new(db.MockDBService),
			limit:    0,
			expected: []domain.DeviceData{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			dbService := test.db.(*db.MockDBService)
			dbService.EXPECT().GetLatestDeviceDataByTableNameLimitN(test.device.TableName, test.limit).Return(test.expected, nil)
			deviceService := NewDeviceService(&test.device, dbService)
			data, err := deviceService.GetLatestDeviceRecordsLimitN(test.limit)
			assert.NoError(t, err)
			assert.Equal(t, test.expected, data)
		})
	}
}

func TestGetLatestDeviceDataByTableNameLimitN_DBFailure(t *testing.T) {
	device := domain.Device{TableName: "test_table"}
	dbService := new(db.MockDBService)
	limit := 10
	dbService.EXPECT().GetLatestDeviceDataByTableNameLimitN(device.TableName, limit).Return(nil, fmt.Errorf("test error"))
	deviceService := NewDeviceService(&device, dbService)
	_, err := deviceService.GetLatestDeviceRecordsLimitN(limit)
	assert.Error(t, err)
}

func TestAddNewRecordToDeviceTable_Success(t *testing.T) {
	device := domain.Device{TableName: "test_table"}
	dbService := new(db.MockDBService)
	isEnabled := true
	dbService.EXPECT().AddNewRecordToDeviceTable(device.TableName, isEnabled).Return(nil)
	deviceService := NewDeviceService(&device, dbService)
	err := deviceService.AddNewRecordToDeviceTable(isEnabled)
	assert.NoError(t, err)
}

func TestAddNewRecordToDeviceTable_DBFailure(t *testing.T) {
	device := domain.Device{TableName: "test_table"}
	dbService := new(db.MockDBService)
	isEnabled := true
	dbService.EXPECT().AddNewRecordToDeviceTable(device.TableName, isEnabled).Return(fmt.Errorf("test error"))
	deviceService := NewDeviceService(&device, dbService)
	err := deviceService.AddNewRecordToDeviceTable(isEnabled)
	assert.Error(t, err)
}

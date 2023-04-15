package service

import (
	"fmt"
	"testing"
	"time"

	"github.com/pklimuk-eng-thesis/data-service/pkg/db"
	"github.com/pklimuk-eng-thesis/data-service/pkg/domain"
	"github.com/stretchr/testify/assert"
)

func TestGetLatestSensorDataByTableNameLimitN_Success(t *testing.T) {
	tests := []struct {
		name     string
		sensor   domain.Sensor
		db       db.DBService
		limit    int
		expected []domain.SensorData
	}{
		{
			name:   "GetLatestSensorDataByTableNameLimitN_Success_limit_10",
			sensor: domain.Sensor{TableName: "test_table"},
			db:     new(db.MockDBService),
			limit:  10,
			expected: []domain.SensorData{
				{ID: 1, CreatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC), IsEnabled: true, Detected: false},
				{ID: 2, CreatedAt: time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC), IsEnabled: false, Detected: false},
			},
		},
		{
			name:   "GetLatestSensorDataByTableNameLimitN_Success_limit_1",
			sensor: domain.Sensor{TableName: "test_table"},
			db:     new(db.MockDBService),
			limit:  1,
			expected: []domain.SensorData{
				{ID: 1, CreatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC), IsEnabled: true, Detected: false},
			},
		},
		{
			name:     "GetLatestSensorDataByTableNameLimitN_Success_limit_0",
			sensor:   domain.Sensor{TableName: "test_table"},
			db:       new(db.MockDBService),
			limit:    0,
			expected: []domain.SensorData{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			dbService := test.db.(*db.MockDBService)
			dbService.EXPECT().GetLatestSensorDataByTableNameLimitN(test.sensor.TableName, test.limit).Return(test.expected, nil)
			sensorService := NewSensorService(&test.sensor, dbService)
			data, err := sensorService.GetLatestSensorRecordsLimitN(test.limit)
			assert.NoError(t, err)
			assert.Equal(t, test.expected, data)
		})
	}
}

func TestGetLatestSensorDataByTableNameLimitN_DBFailure(t *testing.T) {
	sensor := domain.Sensor{TableName: "test_table"}
	dbService := new(db.MockDBService)
	limit := 10
	dbService.EXPECT().GetLatestSensorDataByTableNameLimitN(sensor.TableName, limit).Return(nil, fmt.Errorf("test error"))
	sensorService := NewSensorService(&sensor, dbService)
	_, err := sensorService.GetLatestSensorRecordsLimitN(limit)
	assert.Error(t, err)
}

func TestAddNewRecordToSensorTable_Success(t *testing.T) {
	sensor := domain.Sensor{TableName: "test_table"}
	dbService := new(db.MockDBService)
	dbService.EXPECT().AddNewRecordToSensorTable(sensor.TableName, true, false).Return(nil)
	sensorService := NewSensorService(&sensor, dbService)
	err := sensorService.AddNewRecordToSensorTable(true, false)
	assert.NoError(t, err)
}

func TestAddNewRecordToSensorTable_DBFailure(t *testing.T) {
	sensor := domain.Sensor{TableName: "test_table"}
	dbService := new(db.MockDBService)
	dbService.EXPECT().AddNewRecordToSensorTable(sensor.TableName, true, false).Return(fmt.Errorf("test error"))
	sensorService := NewSensorService(&sensor, dbService)
	err := sensorService.AddNewRecordToSensorTable(true, false)
	assert.Error(t, err)
}

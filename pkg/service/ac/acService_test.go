package service

import (
	"fmt"
	"testing"
	"time"

	"github.com/pklimuk-eng-thesis/data-service/pkg/db"
	"github.com/pklimuk-eng-thesis/data-service/pkg/domain"
	"github.com/stretchr/testify/assert"
)

func TestGetLatestACRecordsLimitN_Success(t *testing.T) {
	tests := []struct {
		name     string
		ac       domain.AC
		db       db.DBService
		limit    int
		expected []domain.ACData
	}{
		{
			name:  "GetLatestACRecordsLimitN_Success_limit_10",
			ac:    domain.AC{TableName: "test_table"},
			db:    new(db.MockDBService),
			limit: 10,
			expected: []domain.ACData{
				{ID: 1, CreatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC), IsEnabled: true, Temperature: 20.0, Humidity: 50.0},
				{ID: 2, CreatedAt: time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC), IsEnabled: false, Temperature: 25.0, Humidity: 40.0},
			},
		},
		{
			name:  "GetLatestACRecordsLimitN_Success_limit_1",
			ac:    domain.AC{TableName: "test_table"},
			db:    new(db.MockDBService),
			limit: 1,
			expected: []domain.ACData{
				{ID: 1, CreatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC), IsEnabled: true, Temperature: 20.0, Humidity: 50.0},
			},
		},
		{
			name:     "GetLatestACRecordsLimitN_Success_limit_0",
			ac:       domain.AC{TableName: "test_table"},
			db:       new(db.MockDBService),
			limit:    0,
			expected: []domain.ACData{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			dbService := test.db.(*db.MockDBService)
			dbService.EXPECT().GetLatestACDataByTableNameLimitN(test.ac.TableName, test.limit).Return(test.expected, nil)
			acService := NewACService(&test.ac, dbService)
			data, err := acService.GetLatestACRecordsLimitN(test.limit)
			assert.NoError(t, err)
			assert.Equal(t, test.expected, data)
		})
	}
}

func TestGetLatestACRecordsLimitN_DBFailure(t *testing.T) {
	ac := domain.AC{TableName: "test_table"}
	dbService := new(db.MockDBService)
	limit := 10
	dbService.EXPECT().GetLatestACDataByTableNameLimitN(ac.TableName, limit).Return(nil, fmt.Errorf("test error"))
	acService := NewACService(&ac, dbService)
	_, err := acService.GetLatestACRecordsLimitN(10)
	assert.Error(t, err)
}

func TestAddNewRecordToACTable_Success(t *testing.T) {
	ac := domain.AC{TableName: "test_table"}
	dbService := new(db.MockDBService)
	dbService.EXPECT().AddNewRecordToACTable(ac.TableName, true, float32(20.0), float32(50.0)).Return(nil)
	acService := NewACService(&ac, dbService)
	err := acService.AddNewRecordToACTable(true, float32(20.0), float32(50.0))
	assert.NoError(t, err)
}

func TestAddNewRecordToACTable_DBFailure(t *testing.T) {
	ac := domain.AC{TableName: "test_table"}
	dbService := new(db.MockDBService)
	dbService.EXPECT().AddNewRecordToACTable(ac.TableName, true, float32(20.0), float32(50.0)).Return(fmt.Errorf("test error"))
	acService := NewACService(&ac, dbService)
	err := acService.AddNewRecordToACTable(true, float32(20.0), float32(50.0))
	assert.Error(t, err)
}

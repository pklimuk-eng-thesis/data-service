package db

import (
	"fmt"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/pklimuk-eng-thesis/data-service/pkg/domain"
	"github.com/stretchr/testify/assert"
)

func TestGetLatestSensorDataByTableNameLimitN_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	dbx := sqlx.NewDb(db, "sqlmock")

	rows := sqlmock.NewRows([]string{"id", "created_at", "is_enabled", "detected"}).
		AddRow(1, "2023-01-01T00:00:00Z", true, false).
		AddRow(2, "2023-01-02T00:00:00Z", false, true)
	mock.ExpectQuery("SELECT \\* FROM smart_home.test_table ORDER BY created_at DESC LIMIT \\$1").
		WithArgs(2).
		WillReturnRows(rows)

	service := &dbService{DB: dbx}
	data, err := service.GetLatestSensorDataByTableNameLimitN("test_table", 2)
	assert.NoError(t, err)

	expected := []domain.SensorData{
		{ID: 1, CreatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC), IsEnabled: true, Detected: false},
		{ID: 2, CreatedAt: time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC), IsEnabled: false, Detected: true},
	}

	assert.Equal(t, expected, data)
}

func TestGetLatestSensorDataByTableNameLimitN_DBFailure(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	dbx := sqlx.NewDb(db, "sqlmock")

	mock.ExpectQuery("SELECT \\* FROM smart_home.test_table ORDER BY created_at DESC LIMIT \\$1").
		WithArgs(2).
		WillReturnError(fmt.Errorf("test error"))

	service := &dbService{DB: dbx}
	_, err = service.GetLatestSensorDataByTableNameLimitN("test_table", 2)
	assert.Error(t, err)
}

func TestGetLatestSensorDataByTableNameLimitN_TimeParsingFailure(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	dbx := sqlx.NewDb(db, "sqlmock")

	rows := sqlmock.NewRows([]string{"id", "created_at", "is_enabled", "detected"}).
		AddRow(1, "2023-01-01T00:00:00Z", true, false).
		AddRow(2, "Wrong data", false, true)
	mock.ExpectQuery("SELECT \\* FROM smart_home.test_table ORDER BY created_at DESC LIMIT \\$1").
		WithArgs(2).
		WillReturnRows(rows)

	service := &dbService{DB: dbx}
	_, err = service.GetLatestSensorDataByTableNameLimitN("test_table", 2)
	assert.Error(t, err)
}

func TestParseTime(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		tm, err := parseTime("2023-01-01T00:00:00Z")
		assert.NoError(t, err)
		assert.Equal(t, time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC), tm)
	})

	t.Run("Failure", func(t *testing.T) {
		_, err := parseTime("Wrong data")
		assert.Error(t, err)
	})
}

func TestAddNewRecordToSensorTable_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	dbx := sqlx.NewDb(db, "sqlmock")

	mock.ExpectExec("INSERT INTO smart_home.test_table \\(is_enabled, detected\\) VALUES \\(\\$1, \\$2\\)").
		WithArgs(true, false).
		WillReturnResult(sqlmock.NewResult(1, 1))

	service := &dbService{DB: dbx}
	err = service.AddNewRecordToSensorTable("test_table", true, false)
	assert.NoError(t, err)
}

func TestAddNewRecordToSensorTable_DBFailure(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	dbx := sqlx.NewDb(db, "sqlmock")

	mock.ExpectExec("INSERT INTO smart_home.test_table \\(is_enabled, detected\\) VALUES \\(\\$1, \\$2\\)").
		WithArgs(true, false).
		WillReturnError(fmt.Errorf("test error"))

	service := &dbService{DB: dbx}
	err = service.AddNewRecordToSensorTable("test_table", true, false)
	assert.Error(t, err)
}

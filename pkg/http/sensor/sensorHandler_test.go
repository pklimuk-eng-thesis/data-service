package http

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pklimuk-eng-thesis/data-service/pkg/domain"
	service "github.com/pklimuk-eng-thesis/data-service/pkg/service/sensor"
	"github.com/stretchr/testify/assert"
)

func TestGetLatestSensorRecordsLimitN(t *testing.T) {
	sensorService := new(service.MockSensorService)
	sensorService.EXPECT().GetLatestSensorRecordsLimitN(2).Return([]domain.SensorData{
		{ID: 1, CreatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC), IsEnabled: true, Detected: false},
		{ID: 2, CreatedAt: time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC), IsEnabled: false, Detected: false},
	}, nil)
	sensorHandler := NewSensorHandler(sensorService)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest(http.MethodGet, "/?limit=2", nil)
	sensorHandler.GetLatestSensorRecordsLimitN(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `[
		{
			"id": 1,
			"created_at": "2023-01-01T00:00:00Z",
			"is_enabled": true,
			"detected": false
		},
		{
			"id": 2,
			"created_at": "2023-01-02T00:00:00Z",
			"is_enabled": false,
			"detected": false
		}]`, w.Body.String())
}

func TestGetLatestSensorRecordsLimitN_InvalidLimit(t *testing.T) {
	sensorService := new(service.MockSensorService)
	sensorHandler := NewSensorHandler(sensorService)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest(http.MethodGet, "/?limit=invalid", nil)
	sensorHandler.GetLatestSensorRecordsLimitN(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "Invalid limit parameter", w.Body.String())
}

func TestGetLatestSensorRecordsLimitN_DBFailure(t *testing.T) {
	sensorService := new(service.MockSensorService)
	sensorService.EXPECT().GetLatestSensorRecordsLimitN(2).Return(nil, errors.New("DB Failure"))
	sensorHandler := NewSensorHandler(sensorService)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest(http.MethodGet, "/?limit=2", nil)
	sensorHandler.GetLatestSensorRecordsLimitN(c)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, "DB Failure", w.Body.String())
}

func TestAddNewRecordToSensorTable(t *testing.T) {
	sensorService := new(service.MockSensorService)
	sensorService.EXPECT().AddNewRecordToSensorTable(true, false).Return(nil)
	sensorHandler := NewSensorHandler(sensorService)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest(http.MethodPost, "/", strings.NewReader(`{"enabled": true, "detected": false}`))
	c.Request.Header.Set("Content-Type", "application/json")
	sensorHandler.AddNewRecordToSensorTable(c)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestAddNewRecordToSensorTable_InvalidBody(t *testing.T) {
	sensorService := new(service.MockSensorService)
	sensorHandler := NewSensorHandler(sensorService)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest(http.MethodPost, "/", strings.NewReader(`{"enabled": true, "detected": false`))
	c.Request.Header.Set("Content-Type", "application/json")
	sensorHandler.AddNewRecordToSensorTable(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "Invalid request body", w.Body.String())
}

func TestAddNewRecordToSensorTable_DBFailure(t *testing.T) {
	sensorService := new(service.MockSensorService)
	sensorService.EXPECT().AddNewRecordToSensorTable(true, false).Return(errors.New("DB Failure"))
	sensorHandler := NewSensorHandler(sensorService)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest(http.MethodPost, "/", strings.NewReader(`{"enabled": true, "detected": false}`))
	c.Request.Header.Set("Content-Type", "application/json")
	sensorHandler.AddNewRecordToSensorTable(c)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, "DB Failure", w.Body.String())
}

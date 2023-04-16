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
	service "github.com/pklimuk-eng-thesis/data-service/pkg/service/device"
	"github.com/stretchr/testify/assert"
)

func TestGetLatestDeviceRecordsLimitN_Success(t *testing.T) {
	deviceService := new(service.MockDeviceService)
	deviceService.EXPECT().GetLatestDeviceRecordsLimitN(2).Return([]domain.DeviceData{
		{ID: 1, CreatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC), IsEnabled: true},
		{ID: 2, CreatedAt: time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC), IsEnabled: false},
	}, nil)
	deviceHandler := NewDeviceHandler(deviceService)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest(http.MethodGet, "/?limit=2", nil)
	deviceHandler.GetLatestDeviceRecordsLimitN(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `[
		{
			"id": 1,
			"created_at": "2023-01-01T00:00:00Z",
			"is_enabled": true
		},
		{
			"id": 2,
			"created_at": "2023-01-02T00:00:00Z",
			"is_enabled": false
		}]`, w.Body.String())
}

func TestGetLatestDeviceRecordsLimitN_InvalidLimit(t *testing.T) {
	deviceService := new(service.MockDeviceService)
	deviceHandler := NewDeviceHandler(deviceService)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest(http.MethodGet, "/?limit=invalid", nil)
	deviceHandler.GetLatestDeviceRecordsLimitN(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "Invalid limit parameter", w.Body.String())
}

func TestGetLatestDeviceRecordsLimitN_DBFailure(t *testing.T) {
	deviceService := new(service.MockDeviceService)
	deviceService.EXPECT().GetLatestDeviceRecordsLimitN(2).Return([]domain.DeviceData{}, errors.New("DB Failure"))
	deviceHandler := NewDeviceHandler(deviceService)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest(http.MethodGet, "/?limit=2", nil)
	deviceHandler.GetLatestDeviceRecordsLimitN(c)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, "DB Failure", w.Body.String())
}

func TestAddNewRecordToDeviceTable_Success(t *testing.T) {
	deviceService := new(service.MockDeviceService)
	deviceService.EXPECT().AddNewRecordToDeviceTable(true).Return(nil)
	deviceHandler := NewDeviceHandler(deviceService)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest(http.MethodPost, "/", strings.NewReader(`{"enabled": true}`))
	c.Request.Header.Set("Content-Type", "application/json")
	deviceHandler.AddNewRecordToDeviceTable(c)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestAddNewRecordToDeviceTable_InvalidBody(t *testing.T) {
	deviceService := new(service.MockDeviceService)
	deviceHandler := NewDeviceHandler(deviceService)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest(http.MethodPost, "/", strings.NewReader(`{"enabled": true`))
	c.Request.Header.Set("Content-Type", "application/json")
	deviceHandler.AddNewRecordToDeviceTable(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "Invalid request body", w.Body.String())
}

func TestAddNewRecordToDeviceTable_DBFailure(t *testing.T) {
	deviceService := new(service.MockDeviceService)
	deviceService.EXPECT().AddNewRecordToDeviceTable(true).Return(errors.New("DB Failure"))
	deviceHandler := NewDeviceHandler(deviceService)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest(http.MethodPost, "/", strings.NewReader(`{"enabled": true}`))
	c.Request.Header.Set("Content-Type", "application/json")
	deviceHandler.AddNewRecordToDeviceTable(c)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, "DB Failure", w.Body.String())
}

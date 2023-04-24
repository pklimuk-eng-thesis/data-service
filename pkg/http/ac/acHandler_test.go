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
	service "github.com/pklimuk-eng-thesis/data-service/pkg/service/ac"
	"github.com/stretchr/testify/assert"
)

func TestGetLatestACRecordsLimitN_Success(t *testing.T) {
	acService := new(service.MockACService)
	acService.EXPECT().GetLatestACRecordsLimitN(2).Return([]domain.ACData{
		{ID: 1, CreatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC), IsEnabled: true, Temperature: 20.0, Humidity: 50.0},
		{ID: 2, CreatedAt: time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC), IsEnabled: false, Temperature: 25.0, Humidity: 40.0},
	}, nil)
	acHandler := NewACHandler(acService)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest(http.MethodGet, "/?limit=2", nil)
	acHandler.GetLatestACRecordsLimitN(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `[
		{
			"id":1,
			"created_at":"2023-01-01T00:00:00Z",
			"is_enabled":true,
			"temperature":20,
			"humidity":50
		},
		{
			"id":2,
			"created_at":"2023-01-02T00:00:00Z",
			"is_enabled":false,
			"temperature":25,
			"humidity":40
		}]`, w.Body.String())
}

func TestGetLatestACRecordsLimitN_InvalidLimit(t *testing.T) {
	acService := new(service.MockACService)
	acHandler := NewACHandler(acService)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest(http.MethodGet, "/?limit=invalid", nil)
	acHandler.GetLatestACRecordsLimitN(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "Invalid limit parameter", w.Body.String())
}

func TestGetLatestACRecordsLimitN_DBFailure(t *testing.T) {
	acService := new(service.MockACService)
	acService.EXPECT().GetLatestACRecordsLimitN(2).Return(nil, errors.New("DB Failure"))
	acHandler := NewACHandler(acService)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest(http.MethodGet, "/?limit=2", nil)
	acHandler.GetLatestACRecordsLimitN(c)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, "DB Failure", w.Body.String())
}

func TestAddNewRecordToACTable_Success(t *testing.T) {
	acService := new(service.MockACService)
	acService.EXPECT().AddNewRecordToACTable(true, float32(20.0), float32(50.0)).Return(nil)
	acHandler := NewACHandler(acService)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest(http.MethodPost, "/", strings.NewReader(`{"enabled":true,"temperature":20,"humidity":50}`))
	acHandler.AddNewRecordToACTable(c)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestAddNewRecordToACTable_InvalidBody(t *testing.T) {
	acService := new(service.MockACService)
	acHandler := NewACHandler(acService)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest(http.MethodPost, "/", strings.NewReader(`{"enabled":true,"temperature":20`))
	acHandler.AddNewRecordToACTable(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "Invalid request body", w.Body.String())
}

func TestAddNewRecordToACTable_DBFailure(t *testing.T) {
	acService := new(service.MockACService)
	acService.EXPECT().AddNewRecordToACTable(true, float32(20.0), float32(50.0)).Return(errors.New("DB Failure"))
	acHandler := NewACHandler(acService)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest(http.MethodPost, "/", strings.NewReader(`{"enabled":true,"temperature":20,"humidity":50}`))
	acHandler.AddNewRecordToACTable(c)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, "DB Failure", w.Body.String())
}

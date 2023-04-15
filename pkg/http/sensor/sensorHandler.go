package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pklimuk-eng-thesis/data-service/pkg/domain"
	service "github.com/pklimuk-eng-thesis/data-service/pkg/service/sensor"
)

type SensorHandler struct {
	service service.SensorService
}

func NewSensorHandler(service service.SensorService) *SensorHandler {
	return &SensorHandler{service: service}
}

func (h *SensorHandler) GetLatestSensorRecordsLimitN(c *gin.Context) {
	limitStr := c.Query("limit")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid limit parameter")
		return
	}
	sensorData, err := h.service.GetLatestSensorRecordsLimitN(limit)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, sensorData)
}

func (h *SensorHandler) AddNewRecordToSensorTable(c *gin.Context) {
	var sensorInfo domain.SensorInfo
	err := c.BindJSON(&sensorInfo)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid request body")
		return
	}
	err = h.service.AddNewRecordToSensorTable(sensorInfo.Enabled, sensorInfo.Detected)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.Status(http.StatusOK)
}

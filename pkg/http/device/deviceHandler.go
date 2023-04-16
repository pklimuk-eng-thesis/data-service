package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pklimuk-eng-thesis/data-service/pkg/domain"
	service "github.com/pklimuk-eng-thesis/data-service/pkg/service/device"
)

type DeviceHandler struct {
	service service.DeviceService
}

func NewDeviceHandler(service service.DeviceService) *DeviceHandler {
	return &DeviceHandler{service: service}
}

func (h *DeviceHandler) GetLatestDeviceRecordsLimitN(c *gin.Context) {
	limitStr := c.Query("limit")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid limit parameter")
		return
	}
	deviceData, err := h.service.GetLatestDeviceRecordsLimitN(limit)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, deviceData)
}

func (h *DeviceHandler) AddNewRecordToDeviceTable(c *gin.Context) {
	var deviceInfo domain.DeviceInfo
	err := c.BindJSON(&deviceInfo)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid request body")
		return
	}
	err = h.service.AddNewRecordToDeviceTable(deviceInfo.Enabled)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.Status(http.StatusOK)
}

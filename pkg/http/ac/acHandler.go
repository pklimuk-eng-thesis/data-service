package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pklimuk-eng-thesis/data-service/pkg/domain"
	service "github.com/pklimuk-eng-thesis/data-service/pkg/service/ac"
)

type ACHandler struct {
	service service.ACService
}

func NewACHandler(service service.ACService) *ACHandler {
	return &ACHandler{service: service}
}

func (h *ACHandler) GetLatestACRecordsLimitN(c *gin.Context) {
	limitStr := c.Query("limit")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid limit parameter")
		return
	}
	acData, err := h.service.GetLatestACRecordsLimitN(limit)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, acData)
}

func (h *ACHandler) AddNewRecordToACTable(c *gin.Context) {
	var acInfo domain.ACInfo
	err := c.BindJSON(&acInfo)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid request body")
		return
	}
	err = h.service.AddNewRecordToACTable(acInfo.Enabled, acInfo.Temperature, acInfo.Humidity)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.Status(http.StatusOK)
}

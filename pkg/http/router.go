package http

import (
	"github.com/gin-gonic/gin"
	device "github.com/pklimuk-eng-thesis/data-service/pkg/http/device"
	sensor "github.com/pklimuk-eng-thesis/data-service/pkg/http/sensor"
)

func SetupSensorRouter(r *gin.Engine, sensorHandler *sensor.SensorHandler, groupName string) {
	route := r.Group(groupName)
	route.GET("/latest", sensorHandler.GetLatestSensorRecordsLimitN)
	route.POST("/add", sensorHandler.AddNewRecordToSensorTable)
}

func SetupDeviceRouter(r *gin.Engine, deviceHandler *device.DeviceHandler, groupName string) {
	route := r.Group(groupName)
	route.GET("/latest", deviceHandler.GetLatestDeviceRecordsLimitN)
	route.POST("/add", deviceHandler.AddNewRecordToDeviceTable)
}

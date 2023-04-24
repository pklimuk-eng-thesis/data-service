package http

import (
	"github.com/gin-gonic/gin"
	ac "github.com/pklimuk-eng-thesis/data-service/pkg/http/ac"
	device "github.com/pklimuk-eng-thesis/data-service/pkg/http/device"
	sensor "github.com/pklimuk-eng-thesis/data-service/pkg/http/sensor"
)

var latestEndpoint = "/latest"
var addEndpoint = "/add"

func SetupSensorRouter(r *gin.Engine, sensorHandler *sensor.SensorHandler, groupName string) {
	route := r.Group(groupName)
	route.GET(latestEndpoint, sensorHandler.GetLatestSensorRecordsLimitN)
	route.POST(addEndpoint, sensorHandler.AddNewRecordToSensorTable)
}

func SetupDeviceRouter(r *gin.Engine, deviceHandler *device.DeviceHandler, groupName string) {
	route := r.Group(groupName)
	route.GET(latestEndpoint, deviceHandler.GetLatestDeviceRecordsLimitN)
	route.POST(addEndpoint, deviceHandler.AddNewRecordToDeviceTable)
}

func SetupACRouter(r *gin.Engine, acHandler *ac.ACHandler, groupName string) {
	route := r.Group(groupName)
	route.GET(latestEndpoint, acHandler.GetLatestACRecordsLimitN)
	route.POST(addEndpoint, acHandler.AddNewRecordToACTable)
}

package service

import (
	"github.com/pklimuk-eng-thesis/data-service/pkg/db"
	"github.com/pklimuk-eng-thesis/data-service/pkg/domain"
)

type DeviceService interface {
	GetLatestDeviceRecordsLimitN(limit int) ([]domain.DeviceData, error)
	AddNewRecordToDeviceTable(isEnabled bool) error
}

type deviceService struct {
	device *domain.Device
	DB     db.DBService
}

func NewDeviceService(device *domain.Device, db db.DBService) DeviceService {
	return &deviceService{device: device, DB: db}
}

func (s *deviceService) GetLatestDeviceRecordsLimitN(limit int) ([]domain.DeviceData, error) {
	return s.DB.GetLatestDeviceDataByTableNameLimitN(s.device.TableName, limit)
}

func (s *deviceService) AddNewRecordToDeviceTable(isEnabled bool) error {
	return s.DB.AddNewRecordToDeviceTable(s.device.TableName, isEnabled)
}

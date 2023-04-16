package db

import (
	"fmt"

	dbUtils "github.com/pklimuk-eng-thesis/data-service/pkg/db/utils"
	"github.com/pklimuk-eng-thesis/data-service/pkg/domain"
)

func (s *dbService) GetLatestDeviceDataByTableNameLimitN(tableName string, limit int) ([]domain.DeviceData, error) {
	query := fmt.Sprintf("SELECT * FROM smart_home.%s ORDER BY created_at DESC LIMIT $1", tableName)
	rows, err := s.DB.Queryx(query, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var devices []domain.DeviceData
	for rows.Next() {
		var id int
		var createdAt string
		var isEnabled bool
		err = rows.Scan(&id, &createdAt, &isEnabled)
		if err != nil {
			return nil, err
		}
		createdAtTime, err := dbUtils.ParseTime(createdAt)
		if err != nil {
			return nil, err
		}
		devices = append(devices, domain.DeviceData{ID: id, CreatedAt: createdAtTime, IsEnabled: isEnabled})
	}
	return devices, nil
}

func (s *dbService) AddNewRecordToDeviceTable(tableName string, isEnabled bool) error {
	query := fmt.Sprintf("INSERT INTO smart_home.%s (is_enabled) VALUES ($1)", tableName)
	_, err := s.DB.Exec(query, isEnabled)
	if err != nil {
		return err
	}
	return nil
}

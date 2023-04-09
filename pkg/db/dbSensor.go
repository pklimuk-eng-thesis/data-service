package db

import (
	"fmt"
	"time"

	"github.com/pklimuk-eng-thesis/data-service/pkg/domain"
)

func (s *dbService) GetLatestSensorDataByTableNameLimitN(tableName string, limit int) ([]domain.SensorData, error) {
	query := fmt.Sprintf("SELECT * FROM smart_home.%s ORDER BY created_at DESC LIMIT $1", tableName)
	rows, err := s.DB.Queryx(query, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sensors []domain.SensorData
	for rows.Next() {
		var id int
		var createdAt string
		var isEnabled bool
		var detected bool
		err = rows.Scan(&id, &createdAt, &isEnabled, &detected)
		if err != nil {
			return nil, err
		}
		createdAtTime, err := parseTime(createdAt)
		if err != nil {
			return nil, err
		}
		sensors = append(sensors, domain.SensorData{ID: id, CreatedAt: createdAtTime,
			IsEnabled: isEnabled, Detected: detected})
	}
	return sensors, nil
}

func parseTime(t string) (time.Time, error) {
	return time.Parse(time.RFC3339, t)
}

func (s *dbService) AddNewRecordToSensorTable(tableName string, isEnabled bool, detected bool) error {
	query := fmt.Sprintf("INSERT INTO smart_home.%s (is_enabled, detected) VALUES ($1, $2)", tableName)
	_, err := s.DB.Exec(query, isEnabled, detected)
	if err != nil {
		return err
	}
	return nil
}

package domain

import "time"

type Sensor struct {
	TableName string `json:"table_name"`
}

type SensorData struct {
	ID        int       `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	IsEnabled bool      `json:"is_enabled" db:"is_enabled"`
	Detected  bool      `json:"detected"  db:"detected"`
}

type SensorInfo struct {
	Enabled  bool `json:"enabled"`
	Detected bool `json:"detected"`
}

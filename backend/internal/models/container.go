package models

import "time"

// Container - модель контейнера
type Container struct {
	ID           string     `db:"id" json:"id"`
	Name         string     `db:"name" json:"name"`
	Status       string     `db:"status" json:"status"`
	IP           string     `db:"ip" json:"ip"`
	LastPingTime *time.Time `db:"last_ping_time" json:"last_ping_time"`
	IsActive     bool       `db:"is_active" json:"is_active"`
	CreatedAt    time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time  `db:"updated_at" json:"updated_at"`
}

package repository

import (
	"github.com/jmoiron/sqlx"
	"pinger/internal/models"
	"time"
)

type ContainerRepository struct {
	db *sqlx.DB
}

func NewContainerRepository(db *sqlx.DB) *ContainerRepository {
	return &ContainerRepository{db: db}
}

// GetAll получение всех контейнеров
func (r *ContainerRepository) GetAll() ([]models.Container, error) {
	var containers []models.Container
	query := `SELECT * FROM containers ORDER BY created_at DESC`
	err := r.db.Select(&containers, query)
	return containers, err
}

// GetByID получение контейнера по ID
func (r *ContainerRepository) GetByID(id string) (models.Container, error) {
	var container models.Container
	query := `SELECT * FROM containers WHERE id = $1`
	err := r.db.Get(&container, query, id)
	return container, err
}

// Create создание нового контейнера
func (r *ContainerRepository) Create(container *models.Container) error {
	query := `
        INSERT INTO containers (id, name, status, ip, last_ping_time, is_active)
        VALUES ($1, $2, $3, $4, $5, $6)
        RETURNING created_at, updated_at
    `
	return r.db.QueryRow(
		query,
		container.ID,
		container.Name,
		container.Status,
		container.IP,
		container.LastPingTime,
		container.IsActive,
	).Scan(&container.CreatedAt, &container.UpdatedAt)
}

// Update обновление контейнера
func (r *ContainerRepository) Update(container *models.Container) error {
	query := `
        UPDATE containers 
        SET name = $2, 
            status = $3, 
            ip = $4, 
            last_ping_time = $5, 
            is_active = $6,
            updated_at = CURRENT_TIMESTAMP
        WHERE id = $1
        RETURNING updated_at
    `
	return r.db.QueryRow(
		query,
		container.ID,
		container.Name,
		container.Status,
		container.IP,
		container.LastPingTime,
		container.IsActive,
	).Scan(&container.UpdatedAt)
}

// UpdateStatus обновление статуса контейнера
func (r *ContainerRepository) UpdateStatus(id string, status string, isActive bool) error {
	query := `
        UPDATE containers 
        SET status = $2,
            is_active = $3,
            last_ping_time = $4,
            updated_at = CURRENT_TIMESTAMP
        WHERE id = $1
    `
	_, err := r.db.Exec(query, id, status, isActive, time.Now())
	return err
}

// Delete удаление контейнера
func (r *ContainerRepository) Delete(id string) error {
	query := `DELETE FROM containers WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

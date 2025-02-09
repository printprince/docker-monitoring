package repository

import (
	"docker-monitoring/backend/internal/models"
	"github.com/jmoiron/sqlx"
)

// ContainerRepository - репозиторий для работы с контейнерами
type ContainerRepository struct {
	DB *sqlx.DB
}

// NewContainerRepository - конструктор репозитория
func NewContainerRepository(db *sqlx.DB) *ContainerRepository {
	return &ContainerRepository{DB: db}
}

// GetAll - получение всех контейнеров
func (r *ContainerRepository) GetAll() ([]models.Container, error) {
	var containers []models.Container
	err := r.DB.Select(&containers, "SELECT * FROM containers")
	return containers, err
}

// GetByID - получение контейнера по ID
func (r *ContainerRepository) GetByID(id string) (models.Container, error) {
	var container models.Container
	err := r.DB.Get(&container, "SELECT * FROM containers WHERE id=$1", id)
	return container, err
}

// Create - создание контейнера
func (r *ContainerRepository) Create(container *models.Container) error {
	_, err := r.DB.Exec(`
		INSERT INTO containers (id, name, status, ip, last_ping_time, is_active, created_at, updated_at)
		VALUES ($1, $2, $3, $4, COALESCE($5, NOW()), $6, NOW(), NOW())
		ON CONFLICT (id) DO UPDATE 
		SET status = EXCLUDED.status, 
		    ip = EXCLUDED.ip, 
		    last_ping_time = COALESCE(EXCLUDED.last_ping_time, containers.last_ping_time),
		    is_active = EXCLUDED.is_active,
		    updated_at = NOW()`,
		container.ID, container.Name, container.Status, container.IP, container.LastPingTime, container.IsActive)
	return err
}

// Update - обновление контейнера
func (r *ContainerRepository) Update(container *models.Container) error {
	_, err := r.DB.Exec(`
		UPDATE containers SET 
		    name=$2, 
		    status=$3, 
		    ip=$4, 
		    last_ping_time=COALESCE($5, last_ping_time), 
		    is_active=$6, 
		    updated_at=NOW() 
		WHERE id=$1`,
		container.ID, container.Name, container.Status, container.IP, container.LastPingTime, container.IsActive)
	return err
}

// Delete - удаление контейнера по ID
func (r *ContainerRepository) Delete(id string) error {
	_, err := r.DB.Exec("DELETE FROM containers WHERE id=$1", id)
	return err
}

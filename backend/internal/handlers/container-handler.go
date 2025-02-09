package handlers

import (
	"docker-monitoring/backend/internal/models"
	"docker-monitoring/backend/internal/repository"
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

// ContainerHandler - обработчик контейнеров
type ContainerHandler struct {
	Repo *repository.ContainerRepository
}

// NewContainerHandler - конструктор обработчика контейнеров
func NewContainerHandler(repo *repository.ContainerRepository) *ContainerHandler {
	return &ContainerHandler{Repo: repo}
}

// Универсальная функция для JSON-ответов
func jsonResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

// GetContainers - обработчик получения всех контейнеров
func (h *ContainerHandler) GetContainers(w http.ResponseWriter, r *http.Request) {
	log.Println("📥 Запрос на получение всех контейнеров")
	containers, err := h.Repo.GetAll()
	if err != nil {
		log.Printf("❌ Ошибка получения контейнеров: %v\n", err)
		jsonResponse(w, http.StatusInternalServerError, map[string]string{"error": "Ошибка получения контейнеров"})
		return
	}
	jsonResponse(w, http.StatusOK, containers)
}

// GetContainer - обработчик получения контейнера по ID
func (h *ContainerHandler) GetContainer(w http.ResponseWriter, r *http.Request) {
	log.Println("📥 Запрос на получение контейнера по ID")

	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 3 || pathParts[2] == "" {
		log.Println("❌ Ошибка: ID контейнера не указан")
		jsonResponse(w, http.StatusBadRequest, map[string]string{"error": "Не указан ID контейнера"})
		return
	}

	containerID := pathParts[2] // 🔥 ID теперь строка
	log.Printf("🔍 Ищем контейнер с ID=%s\n", containerID)

	container, err := h.Repo.GetByID(containerID) // 🔥 Используем string
	if err != nil {
		log.Printf("❌ Контейнер не найден: %v\n", err)
		jsonResponse(w, http.StatusNotFound, map[string]string{"error": "Контейнер не найден"})
		return
	}

	jsonResponse(w, http.StatusOK, container)
}

// CreateContainer - обработчик добавления контейнера
func (h *ContainerHandler) CreateContainer(w http.ResponseWriter, r *http.Request) {
	log.Println("📥 Запрос на создание контейнера")
	var container models.Container
	err := json.NewDecoder(r.Body).Decode(&container)
	if err != nil {
		log.Printf("❌ Ошибка декодирования JSON: %v\n", err)
		jsonResponse(w, http.StatusBadRequest, map[string]string{"error": "Ошибка декодирования JSON"})
		return
	}

	log.Printf("🆕 Добавление контейнера: %+v\n", container)
	err = h.Repo.Create(&container)
	if err != nil {
		log.Printf("❌ Ошибка сохранения в БД: %v\n", err)
		jsonResponse(w, http.StatusInternalServerError, map[string]string{"error": "Ошибка сохранения в БД"})
		return
	}

	jsonResponse(w, http.StatusCreated, container)
}

// UpdateContainer - обработчик обновления контейнера
func (h *ContainerHandler) UpdateContainer(w http.ResponseWriter, r *http.Request) {
	log.Println("📥 Запрос на обновление контейнера")
	var container models.Container
	err := json.NewDecoder(r.Body).Decode(&container)
	if err != nil {
		log.Printf("❌ Ошибка декодирования JSON: %v\n", err)
		jsonResponse(w, http.StatusBadRequest, map[string]string{"error": "Ошибка декодирования JSON"})
		return
	}

	log.Printf("♻️ Обновление контейнера: %+v\n", container)
	err = h.Repo.Update(&container)
	if err != nil {
		log.Printf("❌ Ошибка обновления контейнера: %v\n", err)
		jsonResponse(w, http.StatusInternalServerError, map[string]string{"error": "Ошибка обновления контейнера"})
		return
	}

	jsonResponse(w, http.StatusOK, container)
}

// DeleteContainer - обработчик удаления контейнера
func (h *ContainerHandler) DeleteContainer(w http.ResponseWriter, r *http.Request) {
	log.Println("📥 Запрос на удаление контейнера")

	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 3 || pathParts[2] == "" {
		log.Println("❌ Ошибка: ID контейнера не указан")
		jsonResponse(w, http.StatusBadRequest, map[string]string{"error": "Не указан ID контейнера"})
		return
	}

	containerID := pathParts[2]
	log.Printf("🗑️ Удаление контейнера: ID=%s\n", containerID)

	err := h.Repo.Delete(containerID)
	if err != nil {
		log.Printf("❌ Ошибка удаления контейнера: %v\n", err)
		jsonResponse(w, http.StatusInternalServerError, map[string]string{"error": "Ошибка удаления контейнера"})
		return
	}

	jsonResponse(w, http.StatusOK, map[string]string{"message": "Контейнер удален"})
}

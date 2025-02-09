package main

import (
	"docker-monitoring/backend/internal/database"
	"docker-monitoring/backend/internal/handlers"
	"docker-monitoring/backend/internal/repository"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
)

// Middleware для CORS
func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func jsonResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func main() {
	// Загружаем конфигурацию БД из переменных окружения
	dbConfig := &database.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  "disable",
	}

	// Подключаемся к БД
	db, err := database.NewConnection(dbConfig)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	containerRepo := repository.NewContainerRepository(db)
	containerHandler := handlers.NewContainerHandler(containerRepo)

	mux := http.NewServeMux()

	// Роутер API
	mux.HandleFunc("/api/containers/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		path := strings.TrimPrefix(r.URL.Path, "/api/containers")
		switch {
		case path == "" || path == "/":
			switch r.Method {
			case http.MethodGet:
				containerHandler.GetContainers(w, r)
			case http.MethodPost:
				containerHandler.CreateContainer(w, r)
			default:
				jsonResponse(w, http.StatusMethodNotAllowed, map[string]string{"error": "Метод не доступен"})
			}
		default:
			switch r.Method {
			case http.MethodGet:
				containerHandler.GetContainer(w, r)
			case http.MethodPut:
				containerHandler.UpdateContainer(w, r)
			case http.MethodDelete:
				containerHandler.DeleteContainer(w, r)
			default:
				jsonResponse(w, http.StatusMethodNotAllowed, map[string]string{"error": "Метод не доступен"})
			}
		}
	})

	// Роутер для пингера
	mux.HandleFunc("/api/pinger", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		log.Println("📥 Запрос на /api/pinger") // 🔥 Лог для проверки
		containerHandler.CreateContainer(w, r)
	})

	// Запуск сервера
	handlerWithCORS := enableCORS(mux)

	log.Println("✅ Сервер запущен на порту :8082")
	if err := http.ListenAndServe(":8082", handlerWithCORS); err != nil {
		log.Fatal(err)
	}
}

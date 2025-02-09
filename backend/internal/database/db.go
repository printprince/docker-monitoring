package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// Структура конфигурации БД
type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// Глобальная переменная для БД
var DB *sqlx.DB

// Функция подключения к БД
func NewConnection(cfg *Config) (*sqlx.DB, error) {
	// Проверяем переменную окружения DATABASE_URL
	dsn := os.Getenv("DATABASE_URL")

	// Если переменная не установлена, то используем конфигурацию
	if dsn == "" {
		dsn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName)
	}

	log.Printf("🔍 Подключение к БД: %s\n", dsn)

	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("❌ Ошибка подключения к БД: %w", err)
	}

	// Пытаемся подключиться несколько раз (5 попыток)
	for i := 1; i <= 5; i++ {
		err = db.Ping()
		if err == nil {
			log.Println("✅ Подключение к базе данных успешно установлено")
			break
		}
		log.Printf("⏳ Ошибка подключения (Попытка %d): %v\n", i, err)
		time.Sleep(5 * time.Second)
	}

	if err != nil {
		return nil, fmt.Errorf("❌ Не удалось подключиться к базе данных: %w", err)
	}

	// Создание таблицы, если её нет
	err = createTable(db)
	if err != nil {
		return nil, err
	}

	log.Println("✅ База данных готова к использованию")
	DB = db
	return db, nil
}

// Функция создания таблицы containers
func createTable(db *sqlx.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS containers (
		id TEXT PRIMARY KEY,
		name TEXT,
		status TEXT,
		ip TEXT,
		last_ping_time TIMESTAMP,
		is_active BOOLEAN DEFAULT true,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`
	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("❌ Ошибка создания таблицы containers: %w", err)
	}

	log.Println("✅ Таблица containers проверена или создана")
	return nil
}

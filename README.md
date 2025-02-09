# Docker Monitoring System

## 📌 Описание
Docker Monitoring System — это веб-приложение для мониторинга контейнеров в среде Docker.
Система позволяет отображать активные и завершённые контейнеры с их статусами, IP-адресами и другими параметрами в реальном времени.

## ⚙️ Функциональность
- 📊 **Мониторинг контейнеров** — отображает список активных контейнеров.
- 📜 **История контейнеров** — показывает ранее завершённые контейнеры.
- 🚦 **Статусы** — контейнеры отображаются с различными статусами (running, exited).
- 🔄 **Обновление данных в реальном времени** — периодическое обновление информации.
- 🎨 **Современный интерфейс** — адаптивный и минималистичный дизайн.
- 🔧 **Настройка через переменные окружения** — возможность гибкой настройки системы.

## 🛠️ Технологии
- **Backend:** Golang, Gin, PostgreSQL, Docker
- **Frontend:** React (TypeScript), TailwindCSS, Vite, Fetch API
- **База данных:** PostgreSQL
- **Docker:** Docker Compose для развертывания всей системы

## 📦 Установка и запуск

### 🔹 1. Клонирование репозитория
```sh
git clone https://github.com/your-repository/docker-monitoring.git
cd docker-monitoring
```

### 🔹 2. Запуск с помощью Docker Compose
```sh
docker-compose up --build -d
```

### 🔹 3. Открытие в браузере
Перейдите по адресу:
```
http://localhost
```

_(Порт можно изменить в `docker-compose.yml`)_

##  API-эндпоинты
### 📌 Получение списка контейнеров
```
GET /api/containers
```
_Возвращает список всех контейнеров с их статусами._

### 📌 Добавление контейнера
```
POST /api/pinger
```
_Добавляет новый контейнер в список мониторинга._

## 🛠 Возможные ошибки и их решение
1. **Ошибка подключения к базе данных**
    - Проверьте переменные окружения в `.env` файле.
    - Убедитесь, что контейнер с базой данных запущен (`docker ps`).

2. **Не загружается фронтенд**
    - Очистите кэш браузера.
    - Перезапустите контейнер фронтенда:
      ```sh
      docker-compose restart frontend
      ```

3. **Неправильное количество контейнеров в списке**
    - Проверьте в базе данных актуальные записи с помощью:
      ```sh
      docker exec -it docker-monitoring-db psql -U postgres -d docker_monitoring -c "SELECT * FROM containers;"
      ```
    - Перезапустите backend:
      ```sh
      docker-compose restart backend
      ```

##  TODO / Планы на будущее 
-  Улучшение UI
-  Оптимизация кода
-  Интеграция с Grafana
-  Расширенный анализ контейнеров

## Contacts
- **[Qayrolla Adilet]** - Автор
- **[Telegram]** - @princccceee
- **[Email]** - kairollaadilet@gmail.com
- **[LinkedIn]** - https://www.linkedin.com/in/adiletj/




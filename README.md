# student-crud-practice

# 🎓 Student CRUD API

Простое REST API для управления студентами на Go + Gin + GORM + PostgreSQL.  
Поддерживает операции CRUD и пагинацию.

---

## 🚀 Возможности
- Создание студента
- Получение списка студентов (с пагинацией)
- Получение студента по ID
- Обновление данных студента
- Удаление студента

---

## 🛠 Технологии
- [Go](https://golang.org/) (Gin, GORM)
- [PostgreSQL](https://www.postgresql.org/)
- [Docker](https://www.docker.com/) + Docker Compose
- Миграции через [golang-migrate](https://github.com/golang-migrate/migrate)

---

## 📦 Запуск проекта

```bash
1. Клонировать репозиторий
git clone https://github.com/yourname/student-crud.git
cd student-crud



2. Создать файл .env

Пример в .env.example:

# Сервер
SERVER_HOST=0.0.0.0
SERVER_PORT=3000
READ_TIMEOUT=10
WRITE_TIMEOUT=10

# База данных
DB_DRIVER=postgres
DB_HOST=db
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=students

3. Запустить проект
docker-compose up --build

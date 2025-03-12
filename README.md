# Описание проекта

Этот проект предназначен для ведения учёта задач. Он запускает сервер, который по адресу `http://localhost` (при стандартных настройках) предоставляет пользователю доступ к удобному интерфейсу управления задачами.

---

## Структура проекта

```plaintext
- /cmd
  - /todolist
    - main.go         - Точка входа в приложение.
- /internal
  - /handlers        - Логика HTTP-обработчиков.
    - handlers.go    - Основные обработчики.
  - /models
    - models.go       - Модели данных, используемых проектом.
  - /repository      - Работа с базой данных.
    - connect.go      - Инициализация базы данных.
    - tasks.go       - Методы для взаимодействия с базой данных.
- /migration
  - 001_init.sql    - SQL-скрипт для инициализации базы данных.
- .env               - Переменные окружения проекта.
- .gitignore         - Игнорируемые файлы для Git.
- go.mod             - Файл модулей Go.
- go.sum             - Контрольная сумма зависимостей.
```

## Локальный запуск
База данных запускается при помощи `docker-compose up -d` из локальной папки проекта.
Используйте .env файл в котором будете хранить значения:
  1) TODO_DBUSER=tasks (имя пользователя для подключения к БД)
  2) TODO_DBPASS=tasks (пароль для подключения)
  3) TODO_DBHOST=localhost (хост для подключения к БД)
  4) TODO_DBPORT=5432 (порт на котором запускается PostgreSQL)
  5) TODO_DBNAME=dbname (имя бд инициализированное в docker-compose.yml)
Вызывайте localhost или 127.0.0.1 и через двоеточие указывайте выбранный порт.

## Инструкция по сборке Docker и запуску контейнера

Так как хорошей практикой является реализацией двух отдельных контейнеров, один для приложения и второй для базы данных. Мне не хватило знаний для качественной реализации этой части задания.

## Скриншоты работоспособности


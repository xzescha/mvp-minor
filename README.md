
# 🎯 Telegram Task & Gift Bot (MVP)

Простой MVP Telegram-бота на Go, который выдаёт пользователю задания из банка, позволяет отмечать их выполнение и выбирать подарок из списка.

## 📦 Структура проекта

```

.
├── cmd
│   └── bot
│       └── main.go          # Точка входа: инициализация бота и подключения к БД
├── internal
│   ├── app
│   │   ├── bot.go           # Логика запуска бота, чтение апдейтов
│   │   └── handlers.go      # Обработчики /start, /tasks и callback-запросов
│   └── db
│       ├── db.go            # Работа с PostgreSQL
│       └── migrations
│           ├── 001\_create\_tables.sql
│           ├── 002\_insert\_tasks.sql
│           └── 003\_insert\_gifts.sql
├── go.mod
├── go.sum
├── Dockerfile
└── docker-compose.yml

```

## 🚀 Функционал

- `/start` — приветственное сообщение
- `/tasks` — получить список заданий
- Inline-кнопки в списке заданий:
  - Отметить задание как выполненное
  - После выполнения – выбрать подарок
- Все данные (задания и подарки) создаются дефолтно при миграции БД

## 🛠 Технологии

- Язык: Go (1.20+)
- Telegram API: [go-telegram-bot-api/v5](https://github.com/go-telegram-bot-api/telegram-bot-api)
- База данных: PostgreSQL
- Контейнеризация: Docker & Docker Compose

## 🔧 Установка и запуск

1. Клонировать репозиторий:
```bash
git clone https://github.com/xzescha/mvp-minor.git
cd telegram-task-bot
```

2. Создать файл `.env` (или задать переменные в `docker-compose.yml`):

```dotenv
TELEGRAM_BOT_TOKEN=ваш_токен
POSTGRES_USER=user
POSTGRES_PASSWORD=password
POSTGRES_DB=appdb
```

3. Убедиться, что в `docker-compose.yml` совпадают эти переменные:

```yaml
services:
  db:
   environment:
      - POSTGRES_USER=${POSTGRES_USER}
      - POSTGRES_PASSWORD=${POSTGRES_PASSWORD}
      - POSTGRES_DB=${POSTGRES_DB}
   volumes:
      - ./internal/db/migrations:/docker-entrypoint-initdb.d/
   bot:
   environment:
      - TELEGRAM_BOT_TOKEN=${TELEGRAM_BOT_TOKEN}
      - POSTGRES_CONN=postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@db:5432/${POSTGRES_DB}?sslmode=disable
```

4. Поднять контейнеры:

```bash
docker-compose down -v
docker-compose up --build
```

5. Проверить, что БД инициализировалась (миграции из `internal/db/migrations`):

```bash
docker-compose logs db
```

6. Открыть Telegram, найти бота по токену и отправить:

```
/start
/tasks
```

## ⚙️ Миграции

Все `.sql`-файлы из папки `internal/db/migrations/` автоматически выполняются при первом старте контейнера `db`.
Если вы вносите изменения в миграции — удаляйте том данных и перезапускайте:

```bash
docker-compose down -v
docker-compose up --build
```

## 📝 Примеры запросов

* Получить задания:

  ```
  /tasks
  ```
* Отметить задание:
  \<Inline-кнопка “Прочитать 3 главы книги “Гарри Поттер””>
* Выбрать подарок:
  \<Inline-кнопка “Набор Lego “Гарри Поттер””>

## Дальнейшее развитие

* Админ-панель для добавления/удаления заданий и подарков
* Хранение прогресса в платёжной системе
* Рассылка уведомлений и напоминаний
* Платёжная интеграция (покупка виртуальных бонусов)

---

© 2025 Ed-Tech MVP Bot


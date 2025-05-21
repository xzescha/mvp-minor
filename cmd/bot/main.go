package main

import (
	"log"
	"os"

	"main/internal/app"
	"main/internal/db"
)

func main() {
	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	if token == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN не задан")
	}

	pgConn := os.Getenv("POSTGRES_CONN")
	if pgConn == "" {
		log.Fatal("POSTGRES_CONN не задан")
	}

	database, err := db.New(pgConn)
	log.Println("Подключаемся к БД с:", pgConn)
	if err != nil {
		log.Fatalf("Ошибка подключения к базе: %v", err)
	}

	bot, err := app.NewBot(token, database)
	if err != nil {
		log.Fatalf("Ошибка создания бота: %v", err)
	}

	log.Println("Бот запущен")
	bot.Run()
}

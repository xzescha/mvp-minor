package app

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"main/internal/db"
)

type Bot struct {
	tgBot *tgbotapi.BotAPI
	db    *db.DB
}

func NewBot(token string, database *db.DB) (*Bot, error) {
	tgBot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}
	return &Bot{tgBot: tgBot, db: database}, nil
}

func (b *Bot) Run() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := b.tgBot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			b.handleMessage(update.Message)
		} else if update.CallbackQuery != nil {
			b.handleCallback(update.CallbackQuery)
		}
	}
}

func StartBot() {
	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	dbConn := os.Getenv("POSTGRES_CONN")

	database, err := db.New(dbConn)
	if err != nil {
		log.Fatalf("DB connection error: %v", err)
	}
	bot, err := NewBot(token, database)
	if err != nil {
		log.Fatalf("Bot init error: %v", err)
	}

	log.Println("Bot is up and running")
	bot.Run()
}

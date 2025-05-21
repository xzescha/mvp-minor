package app

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// handleMessage обрабатывает текстовые команды
func (b *Bot) handleMessage(msg *tgbotapi.Message) {
	switch msg.Text {
	case "/start":
		b.sendWelcome(msg.Chat.ID)
	case "/tasks":
		b.sendTasks(msg.Chat.ID)
	default:
		reply := tgbotapi.NewMessage(msg.Chat.ID,
			"Неизвестная команда. Используйте /tasks для списка заданий.")
		b.tgBot.Send(reply)
	}
}

func (b *Bot) sendWelcome(chatID int64) {
	msg := tgbotapi.NewMessage(chatID,
		"Привет! Я бот для выдачи заданий.\n"+
			"Команды:\n"+
			"/start — приветствие\n"+
			"/tasks — список заданий")
	b.tgBot.Send(msg)
}

func (b *Bot) sendTasks(chatID int64) {
	tasks, err := b.db.GetTasks()
	if err != nil {
		b.tgBot.Send(tgbotapi.NewMessage(chatID, "Ошибка при получении заданий"))
		return
	}
	if len(tasks) == 0 {
		b.tgBot.Send(tgbotapi.NewMessage(chatID, "Заданий нет"))
		return
	}

	msg := tgbotapi.NewMessage(chatID, "Выберите задание:")
	var rows [][]tgbotapi.InlineKeyboardButton
	for _, t := range tasks {
		rows = append(rows,
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(t.Description, fmt.Sprintf("complete_task_%d", t.ID)),
			))
	}
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(rows...)
	b.tgBot.Send(msg)
}

func (b *Bot) handleCallback(query *tgbotapi.CallbackQuery) {
	data := query.Data
	chatID := query.Message.Chat.ID
	userID := query.From.ID

	// Ответ на callback, чтобы убрать индикатор “…” в клиенте
	if _, err := b.tgBot.Request(tgbotapi.NewCallback(query.ID, "")); err != nil {
		log.Println("Callback answer failed:", err)
	}

	switch {
	case strings.HasPrefix(data, "complete_task_"):
		id, _ := strconv.Atoi(strings.TrimPrefix(data, "complete_task_"))
		if err := b.db.MarkTaskCompleted(int64(userID), id); err != nil {
			b.tgBot.Send(tgbotapi.NewMessage(chatID, "Не удалось отметить задание"))
			return
		}
		b.tgBot.Send(tgbotapi.NewMessage(chatID, "Задание выполнено! Выберите подарок:"))
		b.sendGifts(chatID)

	case strings.HasPrefix(data, "select_gift_"):
		id, _ := strconv.Atoi(strings.TrimPrefix(data, "select_gift_"))
		if err := b.db.AssignGiftToUser(int64(userID), id); err != nil {
			b.tgBot.Send(tgbotapi.NewMessage(chatID, "Не удалось выбрать подарок"))
			return
		}
		b.tgBot.Send(tgbotapi.NewMessage(chatID, "Подарок выбран! Спасибо 😊"))

	default:
		log.Printf("Unknown callback data: %q", data)
	}
}

func (b *Bot) sendGifts(chatID int64) {
	gifts, err := b.db.GetGifts()
	if err != nil {
		b.tgBot.Send(tgbotapi.NewMessage(chatID, "Ошибка при получении подарков"))
		return
	}

	msg := tgbotapi.NewMessage(chatID, "Список подарков:")
	var rows [][]tgbotapi.InlineKeyboardButton
	for _, g := range gifts {
		rows = append(rows,
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(g.Name, fmt.Sprintf("select_gift_%d", g.ID)),
			))
	}
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(rows...)
	b.tgBot.Send(msg)
}

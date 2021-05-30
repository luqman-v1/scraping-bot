package telegram

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func Send(message string) error {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TOKEN_BOT"))
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	msg := tgbotapi.NewMessageToChannel(os.Getenv("USERNAME_CHANNEL"), message)
	msg.ParseMode = "Markdown"
	_, err = bot.Send(msg)
	if err != nil {
		log.Println("error send telegram", err)
		return err
	}
	return nil
}

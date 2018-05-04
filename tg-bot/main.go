package main

import (
	"fmt"
	"log"
	"net/http"

	api "github.com/go-telegram-bot-api/telegram-bot-api"
	configuration "github.com/usemam/ledger-cloud/tg-bot/configuration"
)

func processUpdate(update api.Update, bot *api.BotAPI) error {
	if update.Message == nil {
		return nil
	}

	var reply string
	if update.Message.IsCommand() {
		reply = fmt.Sprintf("Command = %s, Args = %s", update.Message.Command(), update.Message.CommandArguments())
	} else {
		reply = "Please, use commands only."
	}

	message := api.NewMessage(update.Message.Chat.ID, reply)
	_, err := bot.Send(message)
	return err
}

func fail(msg string, err error) {
	if err != nil {
		log.Fatal(fmt.Sprintf("%s - %v", msg, err))
	}
}

func main() {
	cfg := configuration.New()

	bot, err := api.NewBotAPI(cfg.BotToken)
	fail("Failed to initialize bot", err)

	_, err = bot.SetWebhook(api.NewWebhook(cfg.BotURL + cfg.BotToken))
	fail("Failed to set webhook", err)

	updates := bot.ListenForWebhook("/")
	go http.ListenAndServe(":"+cfg.Port, nil)

	for update := range updates {
		err = processUpdate(update, bot)
		if err != nil {
			log.Println(err)
		}
	}
}

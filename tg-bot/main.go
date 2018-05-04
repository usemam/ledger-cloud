package main

import (
	"fmt"
	"log"
	"net/http"

	api "github.com/go-telegram-bot-api/telegram-bot-api"
	cmd "github.com/usemam/ledger-cloud/tg-bot/commands"
	configuration "github.com/usemam/ledger-cloud/tg-bot/configuration"
)

func processUpdate(update api.Update, bot *api.BotAPI) error {
	if update.Message == nil {
		return nil
	}

	var reply string
	if update.Message.IsCommand() {
		command, err := cmd.CreateCommand(update.Message.Command())
		if err != nil {
			reply = fmt.Sprintf("Error: %v", err)
			log.Println(err)
		} else {
			reply, err = command.Execute(update.Message.CommandArguments())
			if err != nil {
				reply = fmt.Sprintf("Error: %v", err)
				log.Println(err)
			}
		}
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

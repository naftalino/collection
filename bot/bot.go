package commands

import (
	commands "app/bot/commands/user"
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/go-telegram/bot"
)

func StartBot() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	b, err := bot.New("7874709875:AAFAIh9WH61GCnpGtYDLdjHkr1rZONymILY")
	if err != nil {
		panic(err)
	}

	b.RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypeContains, commands.StartCommand)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/help", bot.MatchTypePrefix, commands.HelpCommand)

	log.Println("[!] Bot rodando")

	b.Start(ctx)
}

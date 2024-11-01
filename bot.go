package main

import (
	commands "app/commands/user"
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
	log.Println("[!] Bot rodando")

	b.Start(ctx)
}

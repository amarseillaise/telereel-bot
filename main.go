package main

import (
	"log"
	"os"

	"github.com/amarseillaise/instareels_to_telegram/bot"
	"gopkg.in/telebot.v4"

	env "github.com/joho/godotenv"
)

func main() {
	initEnv()
	token := os.Getenv("TELETOKEN")
	b, err := bot.InitBot(&token)
	if err != nil {
		log.Fatal(err)
	}
	b.Handle(telebot.OnText, bot.OnTextHandler)
	b.Handle(telebot.OnQuery, bot.OnQueryHandler)

	b.Start()
}

func initEnv() {
	if err := env.Load(); err != nil {
		log.Fatal("No .env file found")
	}
}

package main

import (
	"log"
	"vkIntern/internal/bot"
)

func main() {
	b := bot.NewBot()
	if err := b.Start(); err != nil {
		log.Fatal(err)
	}
}

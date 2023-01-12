package main

import (
	eventConsumer "TgBot/clients/consumer/event-consumer"
	tgClient "TgBot/clients/telegram"
	"TgBot/events/telegram"
	"TgBot/storage/files"
	"flag"
	"log"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "storage"
	batchSize   = 100
)

func main() {
	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, mustToken()),
		files.New(storagePath),
	)

	log.Print("Service started!")

	consumer := eventConsumer.New(eventsProcessor, eventsProcessor, batchSize)
	if err := consumer.Start(); err != nil {
		log.Fatal("Service is stopped", err)
	}

}
func mustToken() string {
	token := flag.String("token-bot-token",
		"",
		"toke for access to telegram bot",
	)
	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}
	return *token
}

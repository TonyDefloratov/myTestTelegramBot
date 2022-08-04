package telegram

import "TgBot/clients/telegram"

type Processor struct {
	tg     *telegram.Client
	offset int
	//storage
}

func New(client *telegram.Client) {

}

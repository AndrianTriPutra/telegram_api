package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	telebot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	flag.Usage = func() {
		fmt.Printf("Usage:\n")
		fmt.Printf("      go run . token \n")
		flag.PrintDefaults()
	}

	flag.Parse()
	if len(flag.Args()) != 1 {
		flag.Usage()
		os.Exit(1)
	}

	token := flag.Args()[0]
	log.Printf("token:%s", token)

	bot, err := telebot.NewBotAPI(token)
	if err != nil {
		log.Fatalf("NewBotAPI:" + err.Error())
	}
	bot.Debug = false
	log.Printf("Authorized on account:" + bot.Self.UserName)

	u := telebot.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for {
		for update := range updates {
			if update.Message != nil {
				log.Printf("from   :" + update.Message.From.UserName)
				log.Printf("msg    :" + update.Message.Text)

				response := ""
				switch update.Message.Text {
				case "hai":
					response = "hallo"
				case "hallo":
					response = "hai"
				default:
					response = "please say hai"
				}

				msg := telebot.NewMessage(update.Message.Chat.ID, response)
				msg.ReplyToMessageID = update.Message.MessageID

				log.Printf("chatID :" + strconv.Itoa(int(update.Message.Chat.ID)))
				log.Printf("msgID  :" + strconv.Itoa(update.Message.MessageID))

				bot.Send(msg)
				fmt.Println()
			}
		}
	}

}

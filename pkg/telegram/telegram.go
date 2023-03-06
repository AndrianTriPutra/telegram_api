package telegram

import (
	"api-telegram/pkg/utils/logger"
	"context"
	"time"

	telebot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Setting struct {
	Timezone string
	Token    string
	ChatID   int64
}

type teleRepo struct {
	setting Setting
	bot     *telebot.BotAPI
}

type RepositoryI interface {
	Send(ctx context.Context, data string) error
}

func Newtelegram(ctx context.Context, setting Setting) RepositoryI {
	bot, err := telebot.NewBotAPI(setting.Token)
	if err != nil {
		logger.Level("fatal", "[Newtelegram] ", "NewBotAPI:"+err.Error())
	}
	bot.Debug = false
	logger.Level("info", "[Newtelegram] ", "Authorized on account::"+bot.Self.UserName)

	u := telebot.NewUpdate(0)
	u.Timeout = 60

	return teleRepo{
		setting: setting,
		bot:     bot,
	}
}

func (tel teleRepo) Send(ctx context.Context, data string) error {
	loc, _ := time.LoadLocation(tel.setting.Timezone)
	now := time.Now().In(loc)
	timestamp := now.Format(time.RFC3339)
	message := "[" + timestamp + "] "
	message += data
	msg := telebot.NewMessage(tel.setting.ChatID, message)
	_, err := tel.bot.Send(msg)
	if err != nil {
		return err
	}

	return nil
}

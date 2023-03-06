package app

import (
	"api-telegram/pkg/telegram"
	"context"
	"time"

	"github.com/labstack/echo/v4"
)

type Setting struct {
	Server   Server
	Telegram telegram.Setting
}

type Server struct {
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}

type appStruct struct {
	setting   Setting
	echoNew   *echo.Echo
	errServer chan error
}

type Application interface {
	Start(ctx context.Context) error
	Run(ctx context.Context)
}

func NewApp(setting Setting) Application {
	return &appStruct{
		setting: setting,
	}
}

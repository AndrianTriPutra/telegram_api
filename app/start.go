package app

import (
	"api-telegram/app/endpoint/docker"
	"api-telegram/pkg/telegram"
	"api-telegram/pkg/utils/logger"
	"context"
	"crypto/tls"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (a *appStruct) Start(ctx context.Context) error {
	//repository
	teleRepo := telegram.Newtelegram(ctx, a.setting.Telegram)

	a.echoNew = echo.New()
	a.echoNew.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus:    true,
		LogMethod:    true,
		LogURI:       true,
		LogUserAgent: true,
		LogLatency:   true,
		LogValuesFunc: func(c echo.Context, values middleware.RequestLoggerValues) error {
			logEcho := fmt.Sprintf("{status:%v} {method:%v} {latency:%v} {uri:%v} {user_agent:%v}", values.Status, values.Method, values.Latency, values.URI, values.UserAgent)
			if values.Status != 200 {
				logger.Level("error", "[logEcho] ", logEcho)
			} else {
				logger.Level("info", "[logEcho] ", logEcho)
			}
			return nil
		},
	}))

	//endpoint
	docker.NewHandler(a.echoNew, "/docker", teleRepo)

	a.errServer = make(chan error)
	cfg := &tls.Config{
		MinVersion:               tls.VersionTLS12,
		CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
		PreferServerCipherSuites: true,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_RSA_WITH_AES_256_CBC_SHA,
		},
	}

	a.echoNew.Server.TLSConfig = cfg
	port := ":" + a.setting.Server.Port
	a.echoNew.Server.Addr = port
	//optional
	a.echoNew.Server.ReadTimeout = a.setting.Server.ReadTimeout
	a.echoNew.Server.WriteTimeout = a.setting.Server.WriteTimeout
	a.echoNew.Server.IdleTimeout = a.setting.Server.IdleTimeout

	return nil
}

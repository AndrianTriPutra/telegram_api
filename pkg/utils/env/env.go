package env

import (
	"api-telegram/app"
	"api-telegram/pkg/telegram"
	"errors"
	"time"

	"api-telegram/pkg/utils/logger"

	"github.com/spf13/viper"
)

type Setting struct {
	General  General
	Server   app.Server
	Telegram telegram.Setting
}

type General struct {
	Core int
	Log  string
}

func ReadEnv(path, filename, format string) (int, app.Setting, error) {
	var core int
	var apps app.Setting
	var err error

	viper.SetConfigType(format)
	viper.AddConfigPath(path)
	viper.SetConfigName(filename)
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		newErr := errors.New("error ReadInConfig :" + err.Error())
		return core, apps, newErr
	}

	var config Setting
	err = viper.Unmarshal(&config)
	if err != nil {
		newErr := errors.New("error Unmarshal :" + err.Error())
		return core, apps, newErr
	}

	core = config.General.Core
	logger.Newlogger(config.General.Log)
	logger.Level("trace", "", "============= [general] =============")
	logger.Level("trace", "Core         :", config.General.Core)
	logger.Level("trace", "Log Level    :", config.General.Log)

	config.Server.ReadTimeout = config.Server.ReadTimeout * time.Second
	config.Server.WriteTimeout = config.Server.WriteTimeout * time.Second
	config.Server.IdleTimeout = config.Server.IdleTimeout * time.Second
	logger.Level("trace", "", "============= [server] =============")
	logger.Level("trace", "Port          :", config.Server.Port)
	logger.Level("trace", "ReadTimeout   :", config.Server.ReadTimeout)
	logger.Level("trace", "WriteTimeout  :", config.Server.WriteTimeout)
	logger.Level("trace", "IdleTimeout   :", config.Server.IdleTimeout)

	logger.Level("trace", "", "============= [telegram] =============")
	logger.Level("trace", "Timezone      :", config.Telegram.Timezone)
	logger.Level("trace", "Token         :", config.Telegram.Token)
	logger.Level("trace", "ChatID        :", config.Telegram.ChatID)

	apps = app.Setting{
		Server:   config.Server,
		Telegram: config.Telegram,
	}

	return core, apps, nil
}

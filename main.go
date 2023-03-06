package main

import (
	"context"
	"os"
	"path/filepath"
	"runtime"

	"api-telegram/app"
	"api-telegram/pkg/utils/env"
	"api-telegram/pkg/utils/logger"
)

// ENV=dev go run .

func main() {
	ctx := context.Background()

	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)
	server := os.Getenv("ENV")
	if len(server) == 0 {
		server = "prod"
	}
	filename := "." + server

	core, apps, err := env.ReadEnv(basepath+"/config/", filename, "yaml")
	if err != nil {
		logger.Level("fatal", "main", "error on ReadEnv:"+err.Error())
	}

	runtime.GOMAXPROCS(core)

	application := app.NewApp(apps)
	err = application.Start(ctx)
	if err != nil {
		logger.Level("fatal", "[main] ", "Error on Start:"+err.Error())
	}
	application.Run(ctx)

}

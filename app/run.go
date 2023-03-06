package app

import (
	"api-telegram/pkg/utils/logger"
	"context"
	"time"
)

func (a *appStruct) Run(ctx context.Context) {
	// start server
	runServer := func() {
		logger.Level("info", "[Run] ", "server running on port:"+a.setting.Server.Port)
		a.errServer <- a.echoNew.Server.ListenAndServe()
	}

	go runServer()

	for {
		select {
		case <-ctx.Done():
			ctxShutDown, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			go func(ctx context.Context) {
				defer cancel()

				// shutdown server
				if err := a.echoNew.Shutdown(ctxShutDown); err != nil {
					logger.Level("fatal", "[Run] ", "server shutdown failed:"+err.Error())
				}
				logger.Level("fatal", "[Run] ", "server exited properly")
			}(ctx)

		case err := <-a.errServer:
			logger.Level("fatal", "[Run] ", "server error got:"+err.Error())
		}
	}
}

package server

import (
	"context"
	"fmt"
	"go-microservice-skeleton/config"
	"go-microservice-skeleton/utils"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func RunServer(handler http.Handler) {
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", config.GetConfig().ServerPort),
		Handler: handler,
	}

	go func() {
		utils.LogInfo(fmt.Sprintf("server started on port: %s", config.GetConfig().ServerPort))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			utils.FailOnError(err, "error start server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	utils.LogInfo("shutdown server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		utils.FailOnError(err, "error shutdown server")
	}

	utils.LogInfo("server exiting")
}

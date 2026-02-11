package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"test-new/configs"
	"test-new/controllers"
	v1Controllers "test-new/controllers/v1"
	"test-new/helpers"
	"test-new/routers"
	"test-new/scheduler"
	v1Usecases "test-new/usecases/v1"

	
	
	
)

func main() {
	configs.LoadConfig()
	logger := helpers.InitializeZeroLogs()

	// Initialize Database (Composite)
	

	

	v1Usecase := v1Usecases.InitializeV1Usecase(
		
		
		logger,
	)
	v1Controller := v1Controllers.InitializeV1Controller(v1Usecase, logger)

	controller := controllers.InitializeController(v1Controller, logger)

	
	app := routers.InitializeRouter(controller, logger)
	

	

	

	

	

	ctx := context.Background()
	logger.Info().Ctx(ctx).Msg("Finish Initializing")

	
	// Start Server (Backend and Publisher usually expose an API)
	go func() {
		if err := app.Listen(fmt.Sprintf("%s:%s", configs.Cfg.Server.Address, configs.Cfg.Server.Port)); err != nil {
			logger.Error().Ctx(ctx).Err(err).Msg("error while running api")
		}
	}()
	

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	logger.Info().Ctx(ctx).Msg("Shutting down server...")

	
	if err := app.Shutdown(); err != nil {
		logger.Error().Ctx(ctx).Err(err).Msg("Server forced to shutdown")
	}
	

	
}

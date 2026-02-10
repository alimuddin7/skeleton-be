package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"test-project/configs"
	"test-project/controllers"
	v1Controllers "test-project/controllers/v1"
	"test-project/helpers"
	"test-project/routers"
	"test-project/scheduler"
	v1Usecases "test-project/usecases/v1"

	
	"test-project/databases"
	
	
	postgre "test-project/databases/postgre"
	
	
	
	
	
		
	
	
	
	redis "test-project/databases/redis"
	
	
	
	
		
	
	
	
	
)

func main() {
	configs.LoadConfig()
	logger := helpers.InitializeZeroLogs()

	// Initialize Database (Composite)
	
	db := databases.InitializeDatabase(
		
		
		postgre.InitializePostgreDatabase(postgre.ConnectPostgre(logger), logger),
		
		
		
		
		
		
		
		
		redis.InitializeRedisDatabase(redis.ConnectRedis(logger), logger),
		
		
		
		
		
		logger,
	)
	

	

	v1Usecase := v1Usecases.InitializeV1Usecase(
		db,
		
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

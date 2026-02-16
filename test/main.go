package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"test/configs"
	"test/controllers"
	v1Controllers "test/controllers/v1"
	"test/helpers"
	"test/routers"
	"test/scheduler"
	v1Usecases "test/usecases/v1"
	"test/databases"
	mysql "test/databases/mysql"
	postgre "test/databases/postgre"
	redis "test/databases/redis"
	rediscluster "test/databases/redis_cluster"
	kafka "test/databases/kafka"
	nats "test/databases/nats"
	minio "test/databases/minio"
	grpcserver "test/grpc/server"
)

func main() {
	configs.LoadConfig()
	logger := helpers.InitializeZeroLogs()

	// Initialize Database (Composite)
	db := databases.InitializeDatabase(
		mysql.InitializeMysqlDatabase(mysql.ConnectMysql(logger), logger),
		postgre.InitializePostgreDatabase(postgre.ConnectPostgre(logger), logger),
		redis.InitializeRedisDatabase(redis.ConnectRedis(logger), logger),
		rediscluster.InitializeRedisCluster(rediscluster.ConnectRedisCluster(logger), logger),
		kafka.InitializeKafkaDatabase(kafka.ConnectKafkaReader(logger), kafka.ConnectKafkaWriter(logger), logger),
		nats.InitializeNatsDatabase(nats.ConnectNats(logger), logger),
		minio.InitializeMinioDatabase(minio.ConnectMinio(logger), logger),
		logger,
	)

	v1Usecase := v1Usecases.InitializeV1Usecase(db,
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

package databases

import (
	"github.com/rs/zerolog"
	
	"test/databases/mysql"
	
	
	
	
	
	
	
	
	"test/databases/postgre"
	
	
	
	
	
	
	
	
	"test/databases/redis"
	
	
	
	
	
	
	
	
	
	
	
	rediscluster "test/databases/redis_cluster"
	
	
	
	
	"test/databases/kafka"
	
	
	
	
	
	
	
	
	"test/databases/nats"
	
	
	
	
	
	
	
	
	"test/databases/minio"
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
)

type Database interface {
	
	GetMysql() mysql.MysqlDatabase
	
	
	
	
	
	
	
	
	GetPostgre() postgre.PostgreDatabase
	
	
	
	
	
	
	
	
	GetRedis() redis.RedisDatabase
	
	
	
	
	
	
	
	
	
	
	
	GetRedisCluster() rediscluster.RedisCluster
	
	
	
	
	GetKafka() kafka.KafkaDatabase
	
	
	
	
	
	
	
	
	GetNats() nats.NatsDatabase
	
	
	
	
	
	
	
	
	GetMinio() minio.MinioDatabase
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
}

type database struct {
	
	Mysql mysql.MysqlDatabase
	
	
	
	
	
	
	
	
	Postgre postgre.PostgreDatabase
	
	
	
	
	
	
	
	
	Redis redis.RedisDatabase
	
	
	
	
	
	
	
	
	
	
	
	RedisCluster rediscluster.RedisCluster
	
	
	
	
	Kafka kafka.KafkaDatabase
	
	
	
	
	
	
	
	
	Nats nats.NatsDatabase
	
	
	
	
	
	
	
	
	Minio minio.MinioDatabase
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
}

func InitializeDatabase(
	
	mysqlDb mysql.MysqlDatabase,
	
	
	
	
	
	
	
	
	postgreDb postgre.PostgreDatabase,
	
	
	
	
	
	
	
	
	redisDb redis.RedisDatabase,
	
	
	
	
	
	
	
	
	
	
	
	redisClusterDb rediscluster.RedisCluster,
	
	
	
	
	kafkaDb kafka.KafkaDatabase,
	
	
	
	
	
	
	
	
	natsDb nats.NatsDatabase,
	
	
	
	
	
	
	
	
	minioDb minio.MinioDatabase,
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	l zerolog.Logger,
) Database {
	return &database{
		
		Mysql: mysqlDb,
		
		
		
		
		
		
		
		
		Postgre: postgreDb,
		
		
		
		
		
		
		
		
		Redis: redisDb,
		
		
		
		
		
		
		
		
		
		
		
		RedisCluster: redisClusterDb,
		
		
		
		
		Kafka: kafkaDb,
		
		
		
		
		
		
		
		
		Nats: natsDb,
		
		
		
		
		
		
		
		
		Minio: minioDb,
		
		
		
		
		
		
		
		
		
		
		
		
		
		
		
		
		
		
	}
}



func (d *database) GetMysql() mysql.MysqlDatabase {
	return d.Mysql
}










func (d *database) GetPostgre() postgre.PostgreDatabase {
	return d.Postgre
}










func (d *database) GetRedis() redis.RedisDatabase {
	return d.Redis
}













func (d *database) GetRedisCluster() rediscluster.RedisCluster {
	return d.RedisCluster
}






func (d *database) GetKafka() kafka.KafkaDatabase {
	return d.Kafka
}










func (d *database) GetNats() nats.NatsDatabase {
	return d.Nats
}










func (d *database) GetMinio() minio.MinioDatabase {
	return d.Minio
}




















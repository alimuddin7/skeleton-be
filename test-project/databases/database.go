package databases

import (
	"github.com/rs/zerolog"
	
	
	"test-project/databases/postgre"
	
	
	
	
	
	
	
	
	"test-project/databases/redis"
	
	
	
	
	
)

type Database interface {
	
	
	GetPostgre() postgre.PostgreDatabase
	
	
	
	
	
	
	
	
	GetRedis() redis.RedisDatabase
	
	
	
	
	
}

type database struct {
	
	
	Postgre postgre.PostgreDatabase
	
	
	
	
	
	
	
	
	Redis redis.RedisDatabase
	
	
	
	
	
}

func InitializeDatabase(
	
	
	postgreDb postgre.PostgreDatabase,
	
	
	
	
	
	
	
	
	redisDb redis.RedisDatabase,
	
	
	
	
	
	l zerolog.Logger,
) Database {
	return &database{
		
		
		Postgre: postgreDb,
		
		
		
		
		
		
		
		
		Redis: redisDb,
		
		
		
		
		
	}
}




func (d *database) GetPostgre() postgre.PostgreDatabase {
	return d.Postgre
}










func (d *database) GetRedis() redis.RedisDatabase {
	return d.Redis
}







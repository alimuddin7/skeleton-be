package configs

import (
	"log"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	General struct {
		ServiceName string `envconfig:"SERVICE_NAME" required:"true"`
	}
	Server struct {
		Address    string `envconfig:"ROUTER_SERVER_ADDRESS" default:"0.0.0.0"`
		Port       string `envconfig:"ROUTER_PORT" default:"3000"`
		MainPath   string `envconfig:"PATH_MAIN" default:"/api"`
		SetMode    string `envconfig:"ROUTER_SETMODE" default:"debug"`
		MaxTimeout string `envconfig:"MAX_TIMEOUT" default:"30s"`
	}
	Hosts     Hosts     `envconfig:"HOSTS"`
	
	
	
}

type Hosts struct {
	
}



var Cfg Config

func LoadConfig() {
	godotenv.Load(".env")
	if err := envconfig.Process("", &Cfg); err != nil {
		log.Fatalf("Failed to load from ENV: %v", err)
	}
}

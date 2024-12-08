package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	ConfigPath     string `env:"CONFIG_PATH" env-default:"config/config.yaml"`
	Database       Database
	Grpc           Grpc
	HttpServer     HttpServer
	AllowedOrigins []string `env:"ALLOWED_ORIGINS" env-default:"*"`
}

type Database struct {
	DbHost         string `env:"DB_HOST" env-required:"true"`
	DbPort         int    `env:"DB_PORT" env-required:"true"`
	DbUser         string `env:"DB_USER" env-required:"true"`
	DbPass         string `env:"DB_PASS" env-required:"true"`
	DbName         string `env:"DB_NAME" env-required:"true"`
	MaxConnections int    `env:"DB_MAX_CONNECTIONS" env-required:"true"`
}

type Grpc struct {
	UserPort       int `env:"USER_PORT" env-default:"50052"`
	AttractionPort int `env:"ATTRACTIONS_PORT" env-default:"50051"`
	TripPort       int `env:"TRIPS_PORT" env-default:"50053"`
	SurveyPort     int `env:"SURVEYS_PORT" env-default:"50054"`
	GatewayPort    int `env:"GATEWAY_PORT" env-default:"8080"`

	UserContainerIp       string `env:"USER_CONTAINER_IP" env-default:"users"`
	AttractionContainerIp string `env:"ATTRACTION_CONTAINER_IP" env-default:"attractions"`
	TripContainerIp       string `env:"TRIP_CONTAINER_IP" env-default:"trips"`
	SurveyContainerIp     string `env:"SURVEY_CONTAINER_IP" env-default:"survey"`
}

type HttpServer struct {
	Address      string        `yaml:"address" yaml-default:"8080"`
	IdleTimeout  time.Duration `yaml:"idle_timeout" yaml-default:"60s"`
	ReadTimeout  time.Duration `yaml:"read_timeout" yaml-default:"10s"`
	WriteTimeout time.Duration `yaml:"write_timeout" yaml-default:"10s"`
}

func Load() *Config {
	var cfg Config

	if err := cleanenv.ReadEnv(&cfg); err != nil {
		log.Printf("cannot read .env file: %s\n (fix: you need to put .env file in main dir)", err)
		os.Exit(1)
	}

	if err := cleanenv.ReadConfig(cfg.ConfigPath, &cfg); err != nil {
		log.Printf("config loader cannot read %s: %v", cfg.ConfigPath, err)
		os.Exit(1)
	}

	return &cfg
}

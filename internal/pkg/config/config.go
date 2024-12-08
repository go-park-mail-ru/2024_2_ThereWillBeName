package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"time"
)

type Config struct {
	ConfigPath     string `env:"CONFIG_PATH" env-default:"config/config.yaml"`
	Database       Database
	Grpc           Grpc
	HttpServer     HttpServer `yaml:"HttpServer"`
	AllowedOrigins []string   `env:"ALLOWED_ORIGINS" env-default:"*"`
	Metric         Metric
}

type Database struct {
	DbHost string `env:"DB_HOST" env-required:"true"`
	DbPort int    `env:"DB_PORT" env-required:"true"`
	DbUser string `env:"DB_USER" env-required:"true"`
	DbPass string `env:"DB_PASS" env-required:"true"`
	DbName string `env:"DB_NAME" env-required:"true"`
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
	Address      int           `yaml:"Address" yaml-default:"8081"`
	IdleTimeout  time.Duration `yaml:"idle_timeout" yaml-default:"60s"`
	ReadTimeout  time.Duration `yaml:"read_timeout" yaml-default:"10s"`
	WriteTimeout time.Duration `yaml:"write_timeout" yaml-default:"10s"`
}

type Metric struct {
	UserPort       int `env:"USER_METRIC_PORT" env-default:"8093"`
	AttractionPort int `env:"ATTRACTIONS_METRIC_PORT" env-default:"8091"`
	TripPort       int `env:"TRIPS_METRIC_PORT" env-default:"8092"`
	SurveyPort     int `env:"SURVEYS_METRIC_PORT" env-default:"8095"`
	GatewayPort    int `env:"GATEWAY_METRIC_PORT" env-default:"8094"`
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

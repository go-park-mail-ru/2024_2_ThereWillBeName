package models

type Config struct {
	Port          int
	Env           string
	AllowedOrigin string
	ConnStr       string
	GRPCPort      int
}

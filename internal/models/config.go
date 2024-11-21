package models

type Config struct {
	Port          int
	Env           string
	AllowedOrigin string
	ConnStr       string
	GRPCPort      int
}

type ConfigGrpc struct {
	Port    int
	ConnStr string
}

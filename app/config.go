package app

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	RedisAddress string
	ServerPort   uint16
}

func LoadConfig() Config {
	cfg := Config{
		RedisAddress: "localhost:6379",
		ServerPort:   3000,
	}

	updateServerPort(&cfg)

	if redisAddr, exists := os.LookupEnv("REDIS_ENDPOINT"); exists {
		cfg.RedisAddress = redisAddr
	}

	return cfg
}

func updateServerPort(cfg *Config) error {
	serverPort, exists := os.LookupEnv("SERVER_PORT")
	if !exists {
		return nil
	}

	port, err := strconv.ParseUint(serverPort, 10, 16)
	if err != nil {
		return fmt.Errorf("invalid SERVER_PORT: %v", err)
	}

	cfg.ServerPort = uint16(port)
	return nil
}

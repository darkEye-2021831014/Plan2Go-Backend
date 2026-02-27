package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/lpernett/godotenv"
)

type Config struct {
	Version           string
	ServiceName       string
	HttpPort          int
	Jwt_SecretKey     string
	OpenWeatherApiKey string
}

var configuration *Config

func LoadConfig() {
	// Load .env file if exists (for local dev)
	err := godotenv.Load()
	if err != nil {
		fmt.Println("No .env file found, using system environment variables")
	}

	// Helper to read env or default
	getEnv := func(key, defaultVal string) string {
		val := os.Getenv(key)
		if val == "" {
			return defaultVal
		}
		return val
	}

	version := getEnv("VERSION", "1.0.0")
	serviceName := getEnv("SERVICE_NAME", "MyService")
	jwtSecretKey := getEnv("JWT_SECRETKEY", "defaultsecret")
	openWeatherApiKey := getEnv("OPENWEATHER_API_KEY", "")

	// Railway requires this env variable for the port
	portStr := os.Getenv("PORT")
	if portStr == "" {
		// fallback to HTTP_PORT for local dev
		portStr = getEnv("HTTP_PORT", "8080")
	}

	httpPort, err := strconv.Atoi(portStr)
	if err != nil {
		fmt.Println("Invalid port, defaulting to 8080")
		httpPort = 8080
	}

	configuration = &Config{
		Version:           version,
		ServiceName:       serviceName,
		HttpPort:          httpPort,
		Jwt_SecretKey:     jwtSecretKey,
		OpenWeatherApiKey: openWeatherApiKey,
	}

	fmt.Printf("Config Loaded → Version:%s, Service:%s, Port:%d\n", version, serviceName, httpPort)
}

func GetConfig() *Config {
	if configuration == nil {
		LoadConfig()
	}
	return configuration
}

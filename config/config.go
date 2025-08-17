package config

import (
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Env             string
	Port            string
	DBHost          string
	DBPort          string
	DBUser          string
	DBPassword      string
	DBName          string
	SSLMode         string            // e.g. "disable" for local, "require" for cloud
	AllowedOrigins  []string          // CORS
	RequestTimeout  time.Duration     // e.g. 3s
	ShutdownTimeout time.Duration     // e.g. 15s
}

func LoadConfig() *Config {
	_ = godotenv.Load() // ok if missing in prod

	return &Config{
		Env:             getenv("ENV", "dev"),
		Port:            getenv("PORT", "8080"),
		DBHost:          getenv("DB_HOST", "localhost"),
		DBPort:          getenv("DB_PORT", "5432"),
		DBUser:          getenv("DB_USER", "postgres"),
		DBPassword:      getenv("DB_PASSWORD", "password"),
		DBName:          getenv("DB_NAME", "postgres"),
		SSLMode:         getenv("DB_SSLMODE", "disable"),
		AllowedOrigins:  splitCSV(getenv("CORS_ORIGINS", "http://localhost:5173")),
		RequestTimeout:  mustDur(getenv("REQUEST_TIMEOUT", "3s")),
		ShutdownTimeout: mustDur(getenv("SHUTDOWN_TIMEOUT", "15s")),
	}
}

func getenv(k, def string) string {
	if v := os.Getenv(k); v != "" { return v }
	return def
}
func splitCSV(s string) []string {
	parts := strings.Split(s, ",")
	for i := range parts { parts[i] = strings.TrimSpace(parts[i]) }
	return parts
}
func mustDur(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil { panic(err) }
	return d
}

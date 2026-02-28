package config

import "os"

type Config struct {
	APIPort     string
	RuachURL    string
	ModelName   string
	DBPath      string
	RecentLimit int
	MaxFileSize int64
	MaxDuration float64
}

func Load() *Config {
	return &Config{
		APIPort:     getEnv("API_PORT", "3000"),
		RuachURL:    getEnv("RUACH_URL", "http://localhost:8000"),
		ModelName:   getEnv("MODEL_NAME", "whisper-medium-am-v1-47wer-v2"),
		DBPath:      getEnv("DB_PATH", "./data/ruach.db"),
		RecentLimit: 10,
		MaxFileSize: 2 << 20, // 2MB
		MaxDuration: 30.0,    // 30 seconds
	}
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}

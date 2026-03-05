package config

import (
	"os"
	"strconv"
)

type Config struct {
	APIPort     string
	RuachURL    string
	ModelName   string
	DBPath      string
	RecentLimit int
	MaxFileSize int64
	MaxDuration float64

	// Telegram
	BotToken string
	BotDev   bool
}

func Load() *Config {
	return &Config{
		APIPort:     getEnv("API_PORT", "3000"),
		RuachURL:    getEnv("RUACH_URL", "http://localhost:8000"),
		ModelName:   getEnv("MODEL_NAME", "whisper-medium-am-v1-47wer-v2"),
		DBPath:      getEnv("DB_PATH", "./data/ruach.db"),
		RecentLimit: 10,
		MaxFileSize: 2 << 20, // 2MB
		MaxDuration: float64(getEnvInt("MAX_DURATION", 30)),

		BotToken: getEnv("BOT_TOKEN", ""),
		BotDev:   getEnv("BOT_DEV", "false") == "true",
	}
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}

func getEnvInt(key string, fallback int) int {
	if val := os.Getenv(key); val != "" {
		if i, err := strconv.Atoi(val); err == nil {
			return i
		}
	}
	return fallback
}

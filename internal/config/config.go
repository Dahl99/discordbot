package config

import (
	"github.com/joho/godotenv"
	"log/slog"
	"os"
)

type configuration struct {
	AppEnvironment      string
	BotPrefix           string
	BotStatus           string
	BotGuildJoinMessage string
	DiscordToken        string
	YoutubeApiKey       string
	DatabaseDatabase    string
	DatabasePort        string
	DatabaseUsername    string
	DatabasePassword    string
}

const APP_ENVIRONMENT_LOCAL string = "LOCAL"
const APP_ENVIRONMENT_TEST string = "TEST"
const APP_ENVIRONMENT_PRODUCTION string = "PRODUCTION"

// config contains all environment variables that should be included in .env
var config *configuration

// Load loads the environment variables from the .env file
func Load() {
	err := godotenv.Load()
	if err != nil {
		slog.Error("failed to load environment variables")
	}

	config = &configuration{
		AppEnvironment:      os.Getenv("APP_ENVIRONMENT"),
		BotPrefix:           os.Getenv("BOT_PREFIX"),
		BotStatus:           os.Getenv("BOT_STATUS"),
		BotGuildJoinMessage: os.Getenv("BOT_GUILD_JOIN_MESSAGE"),
		DiscordToken:        os.Getenv("DISCORD_TOKEN"),
		YoutubeApiKey:       os.Getenv("YOUTUBE_API_KEY"),
		DatabaseDatabase:    os.Getenv("DB_DATABASE"),
		DatabasePort:        os.Getenv("DB_PORT"),
		DatabaseUsername:    os.Getenv("DB_USERNAME"),
		DatabasePassword:    os.Getenv("DB_PASSWORD"),
	}
}

func GetAppEnvironment() string {
	return config.AppEnvironment
}

func IsAppEnvironment(environments ...string) bool {
	if len(environments) == 0 {
		return config.AppEnvironment == environments[0]
	}

	for _, environment := range environments {
		if config.AppEnvironment == environment {
			return true
		}
	}

	return false
}

func GetBotPrefix() string {
	return config.BotPrefix
}

func GetBotStatus() string {
	return config.BotStatus
}

func GetBotGuildJoinMessage() string {
	return config.BotGuildJoinMessage
}

func GetDiscordToken() string {
	return config.DiscordToken
}

func GetYoutubeApiKey() string {
	return config.YoutubeApiKey
}

func GetDatabaseDatabase() string {
	return config.DatabaseDatabase
}

func GetDatabasePort() string {
	return config.DatabasePort
}

func GetDatabaseUsername() string {
	return config.DatabaseUsername
}

func GetDatabasePassword() string {
	return config.DatabasePassword
}

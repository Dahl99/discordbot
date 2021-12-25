package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type configuration struct {
	AppEnvironment      string
	BotPrefix           string
	BotStatus           string
	BotGuildJoinMessage string
	DiscordToken        string
	YoutubeKey          string
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
		log.Fatalln("Error loading .env file")
	}

	config = &configuration{
		AppEnvironment:      os.Getenv("APP_ENVIRONMENT"),
		BotPrefix:           os.Getenv("BOT_PREFIX"),
		BotStatus:           os.Getenv("BOT_STATUS"),
		BotGuildJoinMessage: os.Getenv("BOT_GUILD_JOIN_MESSAGE"),
		DiscordToken:        os.Getenv("DISCORD_TOKEN"),
		YoutubeKey:          os.Getenv("YOUTUBE_KEY"),
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

func GetPrefix() string {
	return config.BotPrefix
}

func GetStatusText() string {
	return config.BotStatus
}

func GetGuildJoinMessage() string {
	return config.BotGuildJoinMessage
}

func GetDiscordToken() string {
	return config.DiscordToken
}

func GetYoutubeKey() string {
	return config.YoutubeKey
}

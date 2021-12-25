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

const APP_ENVIRONMENT_LOCAL = "LOCAL"
const APP_ENVIRONMENT_TEST = "TEST"
const APP_ENVIRONMENT_PRODUCTION = "PRODUCTION"

// Config is a global struct object containing bot config
var config *configuration

// Load reads the data the bot needs from the provided JSON file
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
	for _, environment := range environments {
		return config.AppEnvironment == environment
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

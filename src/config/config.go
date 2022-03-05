package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type configuration struct {
	AppEnvironment         string
	BotPrefix              string
	BotStatus              string
	BotGuildJoinMessage    string
	DiscordToken           string
	YoutubeKey             string
	SentryDsn              string
	SentryEnvironment      string
	SentryTracesSampleRate float64
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

	tracesSampleRate, _ := strconv.ParseFloat(os.Getenv("SENTRY_TRACES_SAMPLE_RATE"), 64)

	config = &configuration{
		AppEnvironment:         os.Getenv("APP_ENVIRONMENT"),
		BotPrefix:              os.Getenv("BOT_PREFIX"),
		BotStatus:              os.Getenv("BOT_STATUS"),
		BotGuildJoinMessage:    os.Getenv("BOT_GUILD_JOIN_MESSAGE"),
		DiscordToken:           os.Getenv("DISCORD_TOKEN"),
		YoutubeKey:             os.Getenv("YOUTUBE_KEY"),
		SentryDsn:              os.Getenv("SENTRY_DSN"),
		SentryEnvironment:      os.Getenv("SENTRY_ENVIRONMENT"),
		SentryTracesSampleRate: tracesSampleRate,
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

func GetYoutubeKey() string {
	return config.YoutubeKey
}

func GetSentryDsn() string {
	return config.SentryDsn
}

func GetSentryEnvironment() string {
	return config.SentryEnvironment
}

func GetSentryTracesSampleRate() float64 {
	return config.SentryTracesSampleRate
}

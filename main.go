package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/getsentry/sentry-go"

	"discordbot/src/bot"
	"discordbot/src/config"
	"discordbot/src/database"
	"discordbot/src/database/migrations"
)

func main() {
	// Load environment variables
	config.Load()
	if config.IsAppEnvironment(config.APP_ENVIRONMENT_TEST) {
		fmt.Println("App environment is test, aborting startup")
		return
	}

	// Connect to database and run migrations
	database.Connect()
	migrations.AutoMigrate()

	// Start the bot
	bot.Start()

	// Initialize Sentry for logging
	err := sentry.Init(sentry.ClientOptions{
		Dsn:              config.GetSentryDsn(),
		Environment:      config.GetSentryEnvironment(),
		TracesSampleRate: config.GetSentryTracesSampleRate(),
	})

	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}

	// Flush buffered events before the program terminates.
	defer sentry.Flush(2 * time.Second)

	fmt.Println("Bot is running. Press Ctrl + C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	bot.Stop()
}

package main

import (
	"os"
	"strconv"

	"private/gentools/tgbot/logs"
)

const (
	// DefaultUpdateTimeout - default messages updates timeout.
	DefaultUpdateTimeout = 3
)

// Config - config.
type Config struct {
	Token      string
	UpdateTout int
	DebugLevel int
	BotDebug   bool
}

// configFromEnv - fill config from environment vars.
func configFromEnv() Config {
	var debug bool

	token := os.Getenv("BOT_TOKEN")               // Telegram token
	updateTout := os.Getenv("BOT_UPDATE_TIMEOUT") // DefaultUpdateTimeout = 3
	debugLevel := os.Getenv("DEBUG")              // number now, 0 is silent
	botDebug := os.Getenv("BOT_DEBUG")            // st now, is debug

	tout, _ := strconv.Atoi(updateTout)
	if tout <= 0 {
		tout = DefaultUpdateTimeout
	}

	dbg, _ := strconv.Atoi(debugLevel)
	if dbg > int(logs.LevelDebug) {
		dbg = int(logs.LevelDebug)
	}

	if botDebug != "0" && botDebug != "" {
		debug = true
	}

	return Config{
		Token:      token,
		UpdateTout: tout,
		DebugLevel: dbg,
		BotDebug:   debug,
	}
}

package main

import (
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/vpngen/wordsgens/tgbot/logs"
)

func main() {
	cfg := configFromEnv()

	// set logs
	logs.SetLogLevel(int32(cfg.DebugLevel))

	// create a bot
	bot, err := createBot(cfg.Token, cfg.BotDebug)
	if err != nil {
		log.Panic(err)
	}

	waitGroup := &sync.WaitGroup{}
	stop := make(chan struct{})

	// run the bot
	waitGroup.Add(1)

	go runBot(waitGroup, stop, bot, cfg.UpdateTout, cfg.DebugLevel)

	// catch exit signals
	kill := make(chan os.Signal, 1)
	signal.Notify(kill, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT)

	for range kill {
		logs.Criticln("[-] Main: Stop signal was received")
		// avoid message loosing
		bot.StopReceivingUpdates()
		time.Sleep(time.Second * time.Duration(cfg.UpdateTout))
		// stop!
		close(stop)
		// bye bye!
		break
	}

	// stop app
	waitGroup.Wait()
	logs.Criticln("[-] Main routine was finished")
}

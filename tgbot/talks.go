package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"private/gentools/namesgenerator"
	"private/gentools/seedgenerator"
	"private/gentools/tgbot/logs"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Security chat settings - autodelete ranges.
const (
	secondsInTheDay  = 24 * 3600
	minSecondsToLive = secondsInTheDay
	maxSecondsToLive = 2 * secondsInTheDay
)

// Dialog stages.
const (
	stageStart int = iota //nolint
	stageWait4Choice
	stageWait4Bill
	stageWait4Decision
	stageCleanup
)

// SlowAnswerTimeout - timeout befor each our answer.
const SlowAnswerTimeout = 3 * time.Second

const seedPrefix = "карамба"

// handlers options.
type hOpts struct {
	wg    *sync.WaitGroup
	bot   *tgbotapi.BotAPI
	debug int
}

// Handling messages (opposed callback).
func messageHandler(opts hOpts, update tgbotapi.Update) {
	defer opts.wg.Done()

	ecode := genEcode() // unique e-code

	defer func() {
		if err := RemoveMsg(opts.bot, update.Message.Chat.ID, update.Message.MessageID); err != nil {
			// we don't want to handle this
			logs.Errf("[!:%s] remove: %s\n", ecode, err)
		}
	}()

	if update.Message.ForwardFrom != nil ||
		update.Message.ForwardFromChat != nil {
		SendMessage(opts.bot, update.Message.Chat.ID, WarnForbidForwards, ecode)

		return
	}

	// check all dialog conditions.
	if !auth(opts, update.Message.Chat.ID, update.Message.Date, ecode) {
		return
	}

	if update.Message.IsCommand() {
		err := handleCommands(opts, update.Message, ecode)
		if err != nil {
			stWrong(opts.bot, update.Message.Chat.ID, ecode, fmt.Errorf("command: %s: %w", update.Message.Command(), err))
		}

		return
	}

	_, err := SendMessage(opts.bot, update.Message.Chat.ID, MsgHelp, ecode)
	if err != nil {
		stWrong(opts.bot, update.Message.Chat.ID, ecode, fmt.Errorf("help msg: %w", err))
	}
}

// handleCommands - .
func handleCommands(opts hOpts, Message *tgbotapi.Message, ecode string) error {
	switch Message.Command() {
	case "start":
		_, err := SendMessage(opts.bot, Message.Chat.ID, MsgHelp, ecode)

		return err
	case "physics":
		fullname, person, err := namesgenerator.PhysicsAwardee()
		if err != nil {
			return fmt.Errorf("gen physics: %w", err)
		}

		text := fmt.Sprintf("*%s*\n\nИмя: %s\nПрисуждение премии по физике: _%s_\n%s\n\n",
			fullname,
			person.Name,
			person.Desc,
			tgbotapi.EscapeText(tgbotapi.ModeMarkdown, person.URL),
		)

		_, err = SendMessage(opts.bot, Message.Chat.ID, text, ecode)

		return err

	case "peace":
		fullname, person, err := namesgenerator.PeaceAwardee()
		if err != nil {
			return fmt.Errorf("gen peace: %w", err)
		}

		text := fmt.Sprintf("*%s*\n\nИмя: %s\nПрисуждение премии мира: _%s_\n%s\n\n",
			fullname,
			person.Name,
			person.Desc,
			tgbotapi.EscapeText(tgbotapi.ModeMarkdown, person.URL),
		)

		_, err = SendMessage(opts.bot, Message.Chat.ID, text, ecode)

		return err

	case "seed12":
		mnemo, seed, err := seedgenerator.Seed(seedgenerator.ENT128, seedPrefix)
		if err != nil {
			return fmt.Errorf("gen seed12: %w", err)
		}

		text := fmt.Sprintf("_%s_\nseed (128): %x\n", mnemo, seed)

		_, err = SendMessage(opts.bot, Message.Chat.ID, text, ecode)

		return err

	case "seed15":
		mnemo, seed, err := seedgenerator.Seed(seedgenerator.ENT160, seedPrefix)
		if err != nil {
			return fmt.Errorf("gen seed15: %w", err)
		}

		text := fmt.Sprintf("_%s_\nseed (160): %x\n", mnemo, seed)

		_, err = SendMessage(opts.bot, Message.Chat.ID, text, ecode)

		return err

	case "seed18":
		mnemo, seed, err := seedgenerator.Seed(seedgenerator.ENT192, seedPrefix)
		if err != nil {
			return fmt.Errorf("gen seed18: %w", err)
		}

		text := fmt.Sprintf("_%s_\nseed (192): %x\n", mnemo, seed)

		_, err = SendMessage(opts.bot, Message.Chat.ID, text, ecode)

		return err

	case "seed21":
		mnemo, seed, err := seedgenerator.Seed(seedgenerator.ENT224, seedPrefix)
		if err != nil {
			return fmt.Errorf("gen seed21: %w", err)
		}

		text := fmt.Sprintf("_%s_\nseed (224): %x\n", mnemo, seed)

		_, err = SendMessage(opts.bot, Message.Chat.ID, text, ecode)

		return err

	case "seed24":
		mnemo, seed, err := seedgenerator.Seed(seedgenerator.ENT256, seedPrefix)
		if err != nil {
			return fmt.Errorf("gen seed24: %w", err)
		}

		text := fmt.Sprintf("_%s_\nseed (256): %x\n", mnemo, seed)

		_, err = SendMessage(opts.bot, Message.Chat.ID, text, ecode)

		return err

	default:
		_, err := SendMessage(opts.bot, Message.Chat.ID, WarnUnknownCommand, ecode)

		return err
	}
}

// genEcode - generate some uniq e-code.
func genEcode() string {
	return fmt.Sprintf("%04x", rand.Int31()) //nolint
}

// RemoveMsg - emoving message.
func RemoveMsg(bot *tgbotapi.BotAPI, chatID int64, msgID int) error {
	remove := tgbotapi.NewDeleteMessage(chatID, msgID)
	if _, err := bot.Request(remove); err != nil {
		return fmt.Errorf("request: %w", err)
	}

	return nil
}

// Something wrong handling.
func stWrong(bot *tgbotapi.BotAPI, chatID int64, ecode string, err error) {
	text := fmt.Sprintf("%s: код %s", FatalSomeThingWrong, ecode)

	logs.Debugf("[!:%s] %s\n", ecode, err)
	SendMessage(bot, chatID, text, ecode)
}

// Check autodelete chat option.
func checkChatAutoDeleteTimer(bot *tgbotapi.BotAPI, chatID int64) (bool, error) {
	chat, err := bot.GetChat(
		tgbotapi.ChatInfoConfig{
			ChatConfig: tgbotapi.ChatConfig{
				ChatID: chatID,
			},
		},
	)
	if err != nil {
		return false, fmt.Errorf("get chat: %w", err)
	}

	if chat.MessageAutoDeleteTime < minSecondsToLive || chat.MessageAutoDeleteTime > maxSecondsToLive {
		return false, nil
	}

	return true, nil
}

// authentificate for dilog.
func auth(opts hOpts, chatID int64, ut int, ecode string) bool {
	adSet, err := checkChatAutoDeleteTimer(opts.bot, chatID)
	if err != nil {
		stWrong(opts.bot, chatID, ecode, fmt.Errorf("check autodelete: %w", err))

		return false
	}

	if !adSet {
		SendMessage(opts.bot, chatID, FatalUnwellSecurity, ecode)

		return false
	}

	return true
}

// SendMessage - send common message.
func SendMessage(bot *tgbotapi.BotAPI, chatID int64, text, ecode string) (*tgbotapi.Message, error) {
	logs.Debugf("[!:%s] send message\n", ecode)

	msg := tgbotapi.NewMessage(chatID, text)
	msg.ParseMode = tgbotapi.ModeMarkdown
	msg.ProtectContent = true

	newMsg, err := bot.Send(msg)
	if err != nil {
		logs.Errf("[!:%s] send message: %s\n", ecode, err)

		return nil, err
	}

	return &newMsg, nil
}

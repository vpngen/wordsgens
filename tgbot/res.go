package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	// MsgHelp - welcome message.
	MsgHelp = `/physics - имя на основе лауреатов Нобелевской премии по физике
	/peace - имя на основе лауреатов Нобелевской премии мира
	/seed12 - мнемоника ключа 12 слов
	/seed15 - мнемоника ключа 15 слов
	/seed18 - мнемоника ключа 18 слов
	/seed21 - мнемоника ключа 21 слов
	/seed24 - мнемоника ключа 24 слов
	`
	// WarnGroupsNotAllowed - this bot is only private.
	WarnGroupsNotAllowed = `Извини, в группах бот не работает`

	// WarnForbidForwards - this bot is only private.
	WarnForbidForwards = `Извини, мы не работаем с пересылками`

	// WarnUnknownCommand - unknown command.
	WarnUnknownCommand = `Извини, мы не обрабатываем эту команду`

	// FatalUnwellSecurity - we wont security
	FatalUnwellSecurity = `Привет!

Рекомендую установить автоудаление сообщений в этом чате *через 1 день*.`

	// FatalSomeThingWrong - something wrong happened.
	FatalSomeThingWrong = `Что-то пошло не так`
)

// ContinueKeyboard - continue keyboard.
var ContinueKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Продолжить", "continue"),
	),
) //nolint

package telebot

import "regexp"

// Telegram API URL.
const telegramApiBaseUrl string = "https://api.telegram.org/bot"

// API endpoints
const answerCallbackQueryEndpoint string = "/answerCallbackQuery"
const deleteMessageEndpoint string = "/deleteMessage"
const deleteWebhookEndpoint string = "/deleteWebhook"
const editMessageReplyMarkupEndpoint string = "/editMessageReplyMarkup"
const editMessageTextEndpoint string = "/editMessageText"
const getUpdatesEndpoint string = "/getUpdates"
const kickChatMemberEndpoint string = "/kickChatMember"
const setMyCommandsEndpoint string = "/setMyCommands"
const sendDiceEndpoint string = "/sendDice"
const sendMessageEndpoint string = "/sendMessage"
const setWebhookEndpoint string = "/setWebhook"
const unbanChatMemberEndpoint string = "/unbanChatMember"

//
// Events
//

// Match commands (i.e. when text starts with the filter but can contain more text)
var ONCOMMAND = Event{
	Identifier: "oncommand",
	Checker: func(toCheck string, filter string) bool {

		var match bool

		if len(toCheck) > 0 && toCheck[0:1] == "/" {
			match, _ = regexp.MatchString("^"+toCheck+".*", filter)
		} else {
			match, _ = regexp.MatchString("^/"+toCheck+".*", filter)
		}
		return match
	},
}

// Match text (i.e. exact match between text and filter)
var ONTEXT = Event{
	Identifier: "ontext",
	Checker: func(toCheck string, filter string) bool {
		return toCheck == filter
	},
}

// Match a CallbackQuery with exact data match
var ONCALLBACK = Event{
	Identifier: "oncallback",
	Checker: func(toCheck string, filter string) bool {
		return toCheck == filter
	},
}

// Match a CallbackQuery with data starting with filter
var ONPAYLOAD = Event{
	Identifier: "onpayload",
	Checker: func(toCheck string, filter string) bool {
		match, _ := regexp.MatchString("^"+toCheck+".*", filter)
		return match
	},
}

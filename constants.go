package telebot

import "regexp"

// Telegram API URL.
const telegramApiBaseUrl string = "https://api.telegram.org/bot"

// API endpoints
const deleteWebhookEndpoint string = "/deleteWebhook"
const getUpdatesEndpoint string = "/getUpdates"
const kickChatMemberEndpoint string = "/kickChatMember"
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
		match, _ := regexp.MatchString(toCheck+".*", filter)
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

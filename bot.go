// Telebot is a Go wrapper for the Telegram Bot API.
// It provides a convenient way to write a Telegram in Go whitout diving in Telegram API.
// Updates are receveived from Telegram either with Telegram Webhook or with calls to Telegram /getUpdates API endpoint.
// In order to build a bot, follow the steps :
// 	* Create a bot with the appropriate config
//	* Write handlers and link it to your bot with the OnText function
//	* And last but not least, start the bot
//
// In addition to update reception, telebot has some functions designed to make your bot send content. You can use it in your handlers.

package telebot

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"time"
)

// Create a bog with the appropriate config.
func CreateBot(apiToken string, config map[string]string) Bot {

	// handerMap is a map to make the correspondance between events and handlers.
	handlerMap := make(map[string]map[string]func(u *Update))

	// Create the bot.
	return Bot{apiToken: apiToken, config: config, handlerMap: handlerMap}
}

// Start the bot.
func (b *Bot) Start() {

	// Determine the type of the bot.
	isWebook := b.config != nil

	// If the bot uses webhook to get Updates.
	if isWebook {

		// Set up webhook.
		_, err := b.setWebhook()

		if err != nil {
			panic(err)
		}

		// Start a webhook bot and handle incoming updates.
		u, err := url.Parse(b.config["WebhookUrl"])

		if err != nil {
			panic(err)
		}

		http.HandleFunc(u.Path+b.apiToken, func(w http.ResponseWriter, r *http.Request) {
			b.handleTelegramWebHook(w, r)
		})
		http.ListenAndServeTLS(":8443", b.config["SslCertificate"], b.config["SslPrivkey"], nil)

		// Else, the bot uses getUpdates.
	} else {

		// Deactivate previous webhooks if exists.
		b.deleteWebhook()

		// Start with no offset.
		offset := 0
		for {
			// Fetch updates twice a second.
			time.Sleep(500 * time.Millisecond)
			// Get and dispatch updates. Set a new offset.
			offset = b.getUpdates(offset)
		}

	}
}

// Call the handler corresponding to a pair (event, filter) id it exists.
func (b *Bot) dispatchEvent(event Event, filter string, u *Update) {

	eventMap := b.handlerMap[event.Identifier]

	// Get all keys of the eventMap
	keys := make([]string, len(eventMap))

	i := 0
	for k := range eventMap {
		keys[i] = k
		i += 1
	}

	// Dispatch handler if checker is true for each key
	for _, k := range keys {

		if event.Checker(k, filter) {

			eventMap[k](u)
		}
	}

}

// Dispatch an update to the corresponding handler based on the detected event.
func (b *Bot) dispatchUpdate(u *Update) {

	// Find the corresponding event and dispatch it.
	switch {
	case u.Message.Text != "":
		b.dispatchEvent(ONCOMMAND, u.Message.Text, u)
		b.dispatchEvent(ONTEXT, u.Message.Text, u)
	case u.CallbackQuery.Data != "":
		b.dispatchEvent(ONCALLBACK, u.CallbackQuery.Data, u)
		b.dispatchEvent(ONPAYLOAD, u.CallbackQuery.Data, u)
	}
}

// Register the handler corresponding to the pair (event, filter)
func (b *Bot) registerHandler(event Event, filter string, handler func(u *Update)) {
	// Check if event is already registered.
	_, exists := b.handlerMap[event.Identifier]

	// If the event doesn't exist, create a new eventMap and register the handler.
	if !exists {
		eventMap := make(map[string]func(u *Update))
		eventMap[filter] = handler
		b.handlerMap[event.Identifier] = eventMap

		// Otherwise, register the handler.
	} else {
		b.handlerMap[event.Identifier][filter] = handler
	}
}

// Register a command in the bot commands
func (b *Bot) registerInCommands(command string, description string) {

	// Only append commands with description >= 3 (otherwise Telegram will ignore it)
	if len(description) >= 3 {
		b.commands = append(b.commands, BotCommand{command, description})
	}
}

// Set the bot commands with Telegram API
func (b *Bot) SetCommands() {

	// If no commands are specified, reset bot commands
	commands := "[]"

	if len(b.commands) > 0 {
		jsonCommands, err := json.Marshal(b.commands)

		if err != nil {
			log.Println(err)
		}

		commands = string(jsonCommands)
	}

	val := url.Values{
		"commands": {commands},
	}

	_, err := b.makeAPICall(setMyCommandsEndpoint, val)

	if err != nil {
		log.Println(err)
	}

}

//
// Below are defined the module API functions used to link handlers to events.
//

// Trigger handler if the text of the update matches the variable text.
func (b *Bot) OnText(text string, handler func(u *Update)) {

	event := ONTEXT

	// Register handler.
	b.registerHandler(event, text, handler)
}

// Match commands (i.e. when text starts with the filter but can contain more text)
func (b *Bot) OnCommand(text string, description string, handler func(u *Update)) {

	event := ONCOMMAND

	// Register handler.
	b.registerHandler(event, text, handler)

	// Register the command in the commandMap
	b.registerInCommands(text, description)

}

// Match CallbackQuery
func (b *Bot) OnCallback(data string, handler func(u *Update)) {

	event := ONCALLBACK

	// Register handler.
	b.registerHandler(event, data, handler)
}

// Match CallbackQuery with payload
func (b *Bot) OnPayload(data string, handler func(u *Update)) {

	event := ONPAYLOAD

	// Register handler.
	b.registerHandler(event, data, handler)
}

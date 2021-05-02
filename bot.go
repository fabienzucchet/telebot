// Telebot is a Go wrapper for the Telegram Bot API.
// It provides a convenient way to write a Telegram in Go whitout diving in Telegram API.
// Updates are receveived from Telegram either with Telegram Webhook or with calls to Telegram /getUpdates API endpoint.
// In order to build a bot, follow the steps :
// 	* Create a bot with the appropriate config
//
//		const apiToken = "<your token>"
//
//		// Ignore config if you don't want to use Webhook
// 		config := make(map[string]string)
//
// 		config["WebhookUrl"] = "<public url of your bot>"
// 		config["IpAddr"] = "<ip of your bot>"
// 		config["SslCertificate"] = "<path to .crt SSL cert file>"
// 		config["SslPrivkey"] = "<path to .key SSL cert file>"
//
//		// Replace config with nil if you want to fetch Updates with the /getUpdates endpoint.
// 		bot := telebot.CreateBot(apiToken, config)
//
//	* Write handlers and link it to your bot with the OnText function
//
//		bot.OnText("/test", func(u telebot.Update) {
//		const text = "I hear you loud and clear !"
//		const chatId = u.Message.From
//		_, err := telebot.SendTextToTelegramChat(chatId, text)
//
//		if err != nil {
//			log.Printf("Error sending message: %s", err.Error())
//			}
//		})
//
//	* And last but not least, start the bot
//
//		bot.Start()
//
//
// In addition to update reception, telebot has some functions designed to make your bot send content. You can use it in your handlers.

package telebot

import (
	"log"
	"net/http"
	"time"
)

// Create a bog with the appropriate config.
func CreateBot(apiToken string, config map[string]string) Bot {

	// handerMap is a map to make the correspondance between events and handlers.
	handlerMap := make(map[string]map[string]func(u Update))

	// Create the bot.
	return Bot{apiToken: apiToken, config: config, handlerMap: handlerMap}
}

// Start the bot.
func (b Bot) Start() {

	// Determine the type of the bot.
	isWebook := b.config != nil

	// If the bot uses webhook to get Updates.
	if isWebook {

		// Set up webhook.
		_, err := b.setWebhook()

		if err != nil {
			log.Printf("Error initalising the webhook: %s", err.Error())
		}

		// Start a webhook bot and handle incoming updates.
		http.HandleFunc("/bot", func(w http.ResponseWriter, r *http.Request) {
			handleTelegramWebHook(w, r, b)
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
func (b Bot) dispatchEvent(event string, filter string, u *Update) {
	handler, exists := b.handlerMap[event][filter]

	// If a corresponding event exists, call the handler.
	if exists {
		handler(*u)
	}
}

// Dispatch an update to the corresponding handler based on the detected event.
func (b Bot) dispatchUpdate(u *Update) {

	// Find the corresponding event and dispatch it.
	switch {
	// Dispatch ONTEXT events.
	case u.Message.Text != "":
		b.dispatchEvent(ONTEXT, u.Message.Text, u)
	}
}

// Register the handler corresponding to the pair (event, filter)
func (b Bot) registerHandler(event string, filter string, handler func(u Update)) {
	// Check if event is already registered.
	_, exists := b.handlerMap[event]

	// If the event doesn't exist, create a new eventMap and register the handler.
	if !exists {
		eventMap := make(map[string]func(u Update))
		eventMap[filter] = handler
		b.handlerMap[event] = eventMap

		// Otherwise, register the handler.
	} else {
		b.handlerMap[event][filter] = handler
	}
}

//
// Below are defined the module API functions used to link handlers to events.
//

// Trigger handler if the text of the update matches the variable text.
func (b Bot) OnText(text string, handler func(u Update)) {

	event := ONTEXT
	filter := text

	// Register handler.
	b.registerHandler(event, filter, handler)
}

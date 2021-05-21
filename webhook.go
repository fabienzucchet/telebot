package telebot

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// Set the webhook according to the bot config.
func (b *Bot) setWebhook() (string, error) {

	// Set the webhook with Telegram /setWebhook API endpoint.
	res, err := http.PostForm(
		telegramApiBaseUrl+b.apiToken+setWebhookEndpoint,
		url.Values{
			"url":        {b.config["WebhookUrl"] + b.apiToken},
			"ip_address": {b.config["IPAddress"]},
		})

	if err != nil {
		log.Printf("Error when posting to webhook endpoint: %s", err.Error())
		return "", err
	}

	defer res.Body.Close()

	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("Error when parsing Telegram response: %s", err.Error())
		return "", nil
	}

	bodyString := string(bodyBytes)

	return bodyString, nil
}

// Delete Bot webhook with the Telegram /deleteWebhook API endpoint.
func (b *Bot) deleteWebhook() (string, error) {

	res, err := http.PostForm(
		telegramApiBaseUrl+b.apiToken+deleteWebhookEndpoint,
		url.Values{},
	)

	if err != nil {
		log.Printf("Error when delete webhook: %s", err.Error())
		return "", err
	}

	defer res.Body.Close()

	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("Error when parsing Telegram response: %s", err.Error())
		return "", nil
	}

	bodyString := string(bodyBytes)

	log.Printf("Body of Telegram response: %s", bodyString)

	return bodyString, nil
}

// Parse the request body of the Telegram webhook.
func parseTelegramWebhookRequest(r *http.Request) (*Update, error) {

	var update Update
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		log.Printf("Could not decode incoming update %s", err.Error())
		return nil, err
	}

	return &update, nil
}

// Handle the webhook http request from Telegram.
func (b *Bot) handleTelegramWebHook(w http.ResponseWriter, r *http.Request) {

	// parse Update object
	update, err := parseTelegramWebhookRequest(r)
	if err != nil {
		log.Printf("Error parsing update, %s", err.Error())
		return
	}

	// Dispatch update.
	b.dispatchUpdate(update)

}

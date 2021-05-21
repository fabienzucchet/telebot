package telebot

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// Helper to call Telegram API on the endpoint passed as parameter
func (b *Bot) makeAPICall(endpoint string, v url.Values) (string, error) {
	// Try to send message with telegram API /sendMessage endpoint.
	response, err := http.PostForm(
		telegramApiBaseUrl+b.apiToken+endpoint,
		v,
	)

	if err != nil {
		log.Printf("Error when posting text to the chat: %s", err.Error())
		return "", err
	}

	defer response.Body.Close()

	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Error when parsing Telegram response: %s", err.Error())
		return "", nil
	}

	bodyString := string(bodyBytes)

	return bodyString, nil
}

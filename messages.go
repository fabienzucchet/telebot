package telebot

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

// Send the message text in the chat chatId
func (b Bot) SendTextMessage(chatId int, text string) (string, error) {

	// Try to send message with telegram API /sendMessage endpoint.
	response, err := http.PostForm(
		telegramApiBaseUrl+b.apiToken+sendMessageEndpoint,
		url.Values{
			"chat_id": {strconv.Itoa(chatId)},
			"text":    {text},
		})

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

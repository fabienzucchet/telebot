package telebot

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

// Helper to call /sendMessage API to send a message
func (b Bot) sendMessageAPICall(v url.Values) (string, error) {
	// Try to send message with telegram API /sendMessage endpoint.
	response, err := http.PostForm(
		telegramApiBaseUrl+b.apiToken+sendMessageEndpoint,
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

// Send the message text in the chat chatId.
func (b Bot) SendTextMessage(chatId int, text string, options SendMessageOptions) (string, error) {

	// Mandatory arguments.
	val := url.Values{
		"chat_id":                     {strconv.Itoa(chatId)},
		"text":                        {text},
		"disable_web_page_preview":    {strconv.FormatBool(options.DisableWebPagePreview)},
		"disable_notification":        {strconv.FormatBool(options.DisableNotification)},
		"allow_sending_without_reply": {strconv.FormatBool(options.AllowSendingWithoutReply)},
	}

	// Parse mode
	if options.ParseMode != "" {
		val["parse_mode"] = []string{options.ParseMode}
	}

	// Reply to message
	if options.ReplyToMessageId != 0 {
		val["reply_to_message_id"] = []string{strconv.Itoa(options.ReplyToMessageId)}
	}

	return b.sendMessageAPICall(val)

}

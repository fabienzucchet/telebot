package telebot

import (
	"encoding/json"
	"log"
	"net/url"
	"strconv"
)

// Send the message text in the chat chatId.
func (b *Bot) SendTextMessage(chatId int, text string, options SendMessageOptions) (string, error) {

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

	return b.makeAPICall(sendMessageEndpoint, val)

}

// Send a text message with a ReplyKeyboardMarkup keyboard
func (b *Bot) SendReplyKeyboardMarkupTextMessage(chatId int, text string, keyboard ReplyKeyboardMarkup, options SendMessageOptions) (string, error) {

	jsonStr, err := json.Marshal(keyboard)

	if err != nil {
		log.Println(err)
	}

	log.Println(string(jsonStr))

	// Mandatory arguments.
	val := url.Values{
		"chat_id":                     {strconv.Itoa(chatId)},
		"text":                        {text},
		"disable_web_page_preview":    {strconv.FormatBool(options.DisableWebPagePreview)},
		"disable_notification":        {strconv.FormatBool(options.DisableNotification)},
		"allow_sending_without_reply": {strconv.FormatBool(options.AllowSendingWithoutReply)},
		"reply_markup":                {string(jsonStr)},
	}

	// Parse mode
	if options.ParseMode != "" {
		val["parse_mode"] = []string{options.ParseMode}
	}

	// Reply to message
	if options.ReplyToMessageId != 0 {
		val["reply_to_message_id"] = []string{strconv.Itoa(options.ReplyToMessageId)}
	}

	return b.makeAPICall(sendMessageEndpoint, val)

}

// Send a text message with a ReplyKeyboardRemove keyboard
func (b *Bot) SendReplyKeyboardRemoveTextMessage(chatId int, text string, selective bool, options SendMessageOptions) (string, error) {

	keyboard := ReplyKeyboardRemove{RemoveKeyboard: true, Selective: selective}

	jsonStr, err := json.Marshal(keyboard)

	if err != nil {
		log.Println(err)
	}

	log.Println(string(jsonStr))

	// Mandatory arguments.
	val := url.Values{
		"chat_id":                     {strconv.Itoa(chatId)},
		"text":                        {text},
		"disable_web_page_preview":    {strconv.FormatBool(options.DisableWebPagePreview)},
		"disable_notification":        {strconv.FormatBool(options.DisableNotification)},
		"allow_sending_without_reply": {strconv.FormatBool(options.AllowSendingWithoutReply)},
		"reply_markup":                {string(jsonStr)},
	}

	// Parse mode
	if options.ParseMode != "" {
		val["parse_mode"] = []string{options.ParseMode}
	}

	// Reply to message
	if options.ReplyToMessageId != 0 {
		val["reply_to_message_id"] = []string{strconv.Itoa(options.ReplyToMessageId)}
	}

	return b.makeAPICall(sendMessageEndpoint, val)

}

// Send a text message with an inline keyboard
func (b *Bot) SendInlineKeyboardMarkupTextMessage(chatId int, text string, keyboard InlineKeyboardMarkup, options SendMessageOptions) (string, error) {

	jsonStr, err := json.Marshal(keyboard)

	if err != nil {
		log.Println(err)
	}

	log.Println(string(jsonStr))

	// Mandatory arguments.
	val := url.Values{
		"chat_id":                     {strconv.Itoa(chatId)},
		"text":                        {text},
		"disable_web_page_preview":    {strconv.FormatBool(options.DisableWebPagePreview)},
		"disable_notification":        {strconv.FormatBool(options.DisableNotification)},
		"allow_sending_without_reply": {strconv.FormatBool(options.AllowSendingWithoutReply)},
		"reply_markup":                {string(jsonStr)},
	}

	// Parse mode
	if options.ParseMode != "" {
		val["parse_mode"] = []string{options.ParseMode}
	}

	// Reply to message
	if options.ReplyToMessageId != 0 {
		val["reply_to_message_id"] = []string{strconv.Itoa(options.ReplyToMessageId)}
	}

	return b.makeAPICall(sendMessageEndpoint, val)

}

// Send a dice
func (b *Bot) SendDice(chatId int, options SendMessageOptions) (string, error) {
	val := url.Values{
		"chat_id":                     {strconv.Itoa(chatId)},
		"disable_notification":        {strconv.FormatBool(options.DisableNotification)},
		"allow_sending_without_reply": {strconv.FormatBool(options.AllowSendingWithoutReply)},
	}

	// Reply to message
	if options.ReplyToMessageId != 0 {
		val["reply_to_message_id"] = []string{strconv.Itoa(options.ReplyToMessageId)}
	}

	return b.makeAPICall(sendDiceEndpoint, val)
}

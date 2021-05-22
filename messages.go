package telebot

import (
	"encoding/json"
	"log"
	"math/rand"
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

	jsonKeyboard, err := json.Marshal(keyboard)

	if err != nil {
		log.Println(err)
	}

	// Mandatory arguments.
	val := url.Values{
		"chat_id":                     {strconv.Itoa(chatId)},
		"text":                        {text},
		"disable_web_page_preview":    {strconv.FormatBool(options.DisableWebPagePreview)},
		"disable_notification":        {strconv.FormatBool(options.DisableNotification)},
		"allow_sending_without_reply": {strconv.FormatBool(options.AllowSendingWithoutReply)},
		"reply_markup":                {string(jsonKeyboard)},
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

// Edit a text message.
func (b *Bot) EditTextMessage(chatId int, newText string, messageId int, options SendMessageOptions) (string, error) {

	// Mandatory arguments.
	val := url.Values{
		"chat_id":                  {strconv.Itoa(chatId)},
		"message_id":               {strconv.Itoa(messageId)},
		"text":                     {newText},
		"disable_web_page_preview": {strconv.FormatBool(options.DisableWebPagePreview)},
	}

	// Parse mode
	if options.ParseMode != "" {
		val["parse_mode"] = []string{options.ParseMode}
	}

	return b.makeAPICall(editMessageTextEndpoint, val)
}

// Edit a text message with InlineKeyboardMarkup
func (b *Bot) EditInlineKeyboardTextMessage(chatId int, newText string, messageId int, newKeyboard InlineKeyboardMarkup, options SendMessageOptions) (string, error) {

	jsonKeyboard, err := json.Marshal(newKeyboard)

	if err != nil {
		log.Println(err)
	}

	// Mandatory arguments.
	val := url.Values{
		"chat_id":                  {strconv.Itoa(chatId)},
		"message_id":               {strconv.Itoa(messageId)},
		"text":                     {newText},
		"disable_web_page_preview": {strconv.FormatBool(options.DisableWebPagePreview)},
		"reply_markup":             {string(jsonKeyboard)},
	}

	// Parse mode
	if options.ParseMode != "" {
		val["parse_mode"] = []string{options.ParseMode}
	}

	return b.makeAPICall(editMessageTextEndpoint, val)
}

// Edit the inline keyboard of a message
func (b *Bot) EditMessageInlineKeyboardMarkup(chatId int, messageId int, newKeyboard InlineKeyboardMarkup) (string, error) {

	jsonKeyboard, err := json.Marshal(newKeyboard)

	if err != nil {
		log.Println(err)
	}

	// Mandatory arguments.
	val := url.Values{
		"chat_id":      {strconv.Itoa(chatId)},
		"message_id":   {strconv.Itoa(messageId)},
		"reply_markup": {string(jsonKeyboard)},
	}

	return b.makeAPICall(editMessageReplyMarkupEndpoint, val)

}

// Delete a message
func (b *Bot) DeleteMessage(chatId int, messageId int) (string, error) {

	// Mandatory arguments.
	val := url.Values{
		"chat_id":    {strconv.Itoa(chatId)},
		"message_id": {strconv.Itoa(messageId)},
	}

	return b.makeAPICall(deleteMessageEndpoint, val)

}

// send a dice
func (b *Bot) SendDice(chatId int, options SendMessageOptions) (string, error) {

	return b.SendDiceEmoji(chatId, "", options)
}

// send a random dice
func (b *Bot) SendRandomDice(chatId int, options SendMessageOptions) (string, error) {

	emojiList := []string{"🎲", "🎯", "🏀", "⚽", "🎳", "🎰"}

	emoji := emojiList[rand.Intn(len(emojiList))]

	return b.SendDiceEmoji(chatId, emoji, options)

}

// Send a dice Emoji (Supported emojis : “🎲”, “🎯”, “🏀”, “⚽”, “🎳”, or “🎰”. Default is “🎲”. )
func (b *Bot) SendDiceEmoji(chatId int, emoji string, options SendMessageOptions) (string, error) {
	val := url.Values{
		"chat_id":                     {strconv.Itoa(chatId)},
		"disable_notification":        {strconv.FormatBool(options.DisableNotification)},
		"allow_sending_without_reply": {strconv.FormatBool(options.AllowSendingWithoutReply)},
	}

	// Emoji (Supported emojis : “🎲”, “🎯”, “🏀”, “⚽”, “🎳”, or “🎰”. Default is “🎲”. )
	if emoji != "" {
		val["emoji"] = []string{emoji}
	}

	// Reply to message
	if options.ReplyToMessageId != 0 {
		val["reply_to_message_id"] = []string{strconv.Itoa(options.ReplyToMessageId)}
	}

	return b.makeAPICall(sendDiceEndpoint, val)
}

package telebot

import (
	"math/rand"
	"net/url"
	"strconv"
)

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

package telebot

import (
	"net/url"
	"strconv"
)

// Kick an user from a group.
func (b *Bot) KickChatMember(chatId int, userId int) (string, error) {

	val := url.Values{
		"chat_id": {strconv.Itoa(chatId)},
		"user_id": {strconv.Itoa(userId)},
	}

	return b.makeAPICall(kickChatMemberEndpoint, val)
}

// Unban a member from a group.
func (b *Bot) UnbanChatMember(chatId int, userId int) (string, error) {

	val := url.Values{
		"chat_id": {strconv.Itoa(chatId)},
		"user_id": {strconv.Itoa(userId)},
	}

	return b.makeAPICall(unbanChatMemberEndpoint, val)
}

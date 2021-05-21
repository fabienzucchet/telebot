package telebot

import (
	"net/url"
	"strconv"
)

// Answer a callback query without notification
func (b Bot) AnswerCallbackQuery(callbackQueryId string) (string, error) {

	val := url.Values{
		"callback_query_id": {callbackQueryId},
	}

	return b.makeAPICall(answerCallbackQueryEndpoint, val)

}

// Answer a callback query with notification
func (b Bot) AnswerCallbackQueryNotification(callbackQueryId string, text string, showAlert bool) (string, error) {

	val := url.Values{
		"callback_query_id": {callbackQueryId},
		"text":              {text},
		"show_alert":        {strconv.FormatBool(showAlert)},
	}

	return b.makeAPICall(answerCallbackQueryEndpoint, val)

}

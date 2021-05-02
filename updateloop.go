package telebot

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

// Parse Telegram API response body to get the slice of updates.
func parseTelegramGetUpdateResponse(r *http.Response) ([]Update, error) {

	var getUpdateResponse GetUpdateResponse

	// Try to parse responde body.
	if err := json.NewDecoder(r.Body).Decode(&getUpdateResponse); err != nil {
		log.Printf("Could not decode incoming response %s", err.Error())
		return nil, err
	}

	if !getUpdateResponse.Ok {
		log.Printf("Telegram response was not ok")
		return nil, errors.New("telegram response was not ok")
	}

	return getUpdateResponse.Result, nil
}

//Fetch and dispatch last updates for a Bot with Telegram /getUpdates API endpoint.
func (b Bot) getUpdates(offset int) int {

	// Get Updates with Telegram /getUpdates API.
	res, err := http.PostForm(
		telegramApiBaseUrl+b.apiToken+getUpdatesEndpoint,
		url.Values{
			"offset": {strconv.Itoa(offset)},
		},
	)

	// In case of connection error.
	if err != nil {
		log.Printf("Error fetching updates: %s", err.Error())
		return offset
	}

	defer res.Body.Close()

	// Parse Telegram response.
	updates, err := parseTelegramGetUpdateResponse(res)

	// Handle parsing errors.
	if err != nil {
		log.Printf("Error parsing update, %s", err.Error())
		return offset
	}

	// Dispatch fetched updates to handlers.
	for _, update := range updates {
		b.dispatchUpdate(&update)
	}

	// If Updates were received, we must update offset
	if len(updates) > 0 {
		newOffset := updates[len(updates)-1].UpdateId + 1
		return newOffset
	}

	// If not offset stay the same
	return offset

}

package telebot

// Bot object definition.
type Bot struct {
	apiToken   string
	config     map[string]string
	handlerMap map[string]map[string]func(u Update)
}

// Paths to SSL certificate .key and .crt file
type Cert struct {
	Privkey     string
	Certificate string
}

// Structure of the /getUpdates response body.
type GetUpdateResponse struct {
	Ok     bool     `json:"ok"`
	Result []Update `json:"result"`
}

//
// Below are defined the types corresponding to Telegram API Objects
//

// Chat type corresponding to the interesting part of the Chat Object in the Telegram API.
type Chat struct {
	Id int `json:"id"`
}

// Message type corresponding to the interesting part of the Message Object in the Telegram API.
type Message struct {
	Text string `json:"text"`
	From User   `json:"from"`
	Chat Chat   `json:"chat"`
}

// Update type corresponding to the interesting part of the Update Object in the Telegram API.
type Update struct {
	UpdateId int     `json:"update_id"`
	Message  Message `json:"message"`
}

// User type corresponding to the interesting part of the User Object in the Telegram API.
type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

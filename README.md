# Telebot

Telebot is a Go wrapper for the Telegram Bot API. It provides a convenient way to write a Telegram in Go whitout diving in Telegram API.
Updates are receveived from Telegram either with Telegram Webhook or with calls to Telegram /getUpdates API endpoint.

## How to build a Go Telegram bot with telebot ?

In order to build a bot, follow the steps :

* Create a bot with the appropriate config

```Go
    const apiToken = "<your token>"
    // Ignore config if you don't want to use Webhook
    config := make(map[string]string)

    config["WebhookUrl"] = "<public url of your bot>"
    config["IpAddr"] = "<ip of your bot>"
    config["SslCertificate"] = "<path to .crt SSL cert file>"
    config["SslPrivkey"] = "<path to .key SSL cert file>"

    // Replace config with nil if you want to fetch Updates with the /getUpdates endpoint.
    bot := telebot.CreateBot(apiToken, config)
```

* Write handlers and link it to your bot with the OnText function

```Go
    bot.OnText("/test", func(u telebot.Update) {
    const text = "I hear you loud and clear !"
    const chatId = u.Message.Chat.Id
    _, err := telebot.SendTextMessage(chatId, text)

    if err != nil {
        log.Printf("Error sending message: %s", err.Error())
        }
    })

    bot.OnCommand("/repeat", func(u telebot.Update) {
        payload := u.Message.Text[len("/repeat"):]
        chatId := u.Message.Chat.Id

        _, err := bot.SendTextMessage(chatId, payload)

        if err != nil {
            log.Printf("Error sending message: %s", err.Error())
        }
    })
```

* And last but not least, start the bot

```Go
    bot.Start()
```

### List of events available

The Events are defined in [constants.go](constants.go).

* **ONCOMMAND**: Match messages starting with a command (Ex : In `/repeat Hello World !`, the command is `/repeat` and the payload is `Hello World !`)

Below is an example making the bot repeat the payload:

```Go
bot.OnCommand("/repeat", func(u telebot.Update) {
    payload := u.Message.Text[len("/repeat"):]
    chatId := u.Message.Chat.Id

    _, err := bot.SendTextMessage(chatId, payload)

    if err != nil {
        log.Printf("Error sending message: %s", err.Error())
    }
})
```

* **ONTEXT**: Match exactly the text of the message

Below is an example matching the message `/hello` but not `/hello you`:

```Go
bot.OnText("/hello", func(u telebot.Update) {
    text := "Hello World !"
    chatId := u.Message.Chat.Id
    _, err := bot.SendTextMessage(chatId, text)

    if err != nil {
        log.Printf("Error sending message: %s", err.Error())
    }
})
```

## How to make the bot send content to Telegram chat with telebot ?

In addition to update reception, telebot has some functions designed to make your bot send content. You can use it in your handlers. Such methods are defined in [messages.go](messages.go).

### List of reply methods available

* **SendTextMessage** : Sends a text message.

```Go
bot.SendTextMessage(chatId int, text string)
```

## Example bot

Below is an example of a simple bot that you can use to experiment with telebot.

```Go
package main

import (
    "log"

    "github.com/fabienzucchet/telebot"
)

func main() {

    // Define Telegram API token.
    const apiToken = "changeme"

    // If you want to use webhook, define config.
    config := make(map[string]string)

    config["WebhookUrl"] = "changeme"
    config["IpAddr"] = "changeme"
    config["SslCertificate"] = "changeme"
    config["SslPrivkey"] = "changeme"

    // Replace config with nil below to use an update loop instead of a webhook.
    bot := telebot.CreateBot(apiToken, config)

    // Bind a handler to the message /text.
    bot.OnText("/test", func(u telebot.Update) {
        text := "I hear you loud and clear !"
        chatId := u.Message.Chat.Id

        _, err := bot.SendTextMessage(chatId, text)

        if err != nil {
            log.Printf("Error sending message: %s", err.Error())
        }
    })

    bot.OnCommand("/repeat", func(u telebot.Update) {
        payload := u.Message.Text[len("/repeat"):]
        chatId := u.Message.Chat.Id

        _, err := bot.SendTextMessage(chatId, payload)

        if err != nil {
            log.Printf("Error sending message: %s", err.Error())
        }
    })

    // Start the bot.
    bot.Start()

}
```

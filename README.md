# Telebot

[![Go Reference](https://pkg.go.dev/badge/github.com/fabienzucchet/telebot.svg)](https://pkg.go.dev/github.com/fabienzucchet/telebot)

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
    bot.OnText("/test", func(u *telebot.Update) {
    const text = "I hear you loud and clear !"
    const chatId = u.Message.Chat.Id
    _, err := telebot.SendTextMessage(chatId, text, telebot.SendMessageOptions{})

    if err != nil {
        log.Printf("Error sending message: %s", err.Error())
        }
    })

    bot.OnCommand("/repeat", func(u *telebot.Update) {
        payload := u.Message.Text[len("/repeat"):]
        chatId := u.Message.Chat.Id

        _, err := bot.SendTextMessage(chatId, payload, telebot.SendMessageOptions{})

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
bot.OnCommand("/repeat", func(u *telebot.Update) {
    payload := u.Message.Text[len("/repeat"):]
    chatId := u.Message.Chat.Id

    _, err := bot.SendTextMessage(chatId, payload, telebot.SendMessageOptions{})

    if err != nil {
        log.Printf("Error sending message: %s", err.Error())
    }
})
```

* **ONTEXT**: Match exactly the text of the message

Below is an example matching the message `/hello` but not `/hello you`:

```Go
bot.OnText("/hello", func(u *telebot.Update) {
    text := "Hello World !"
    chatId := u.Message.Chat.Id
    _, err := bot.SendTextMessage(chatId, text, telebot.SendMessageOptions{})

    if err != nil {
        log.Printf("Error sending message: %s", err.Error())
    }
})
```

* **ONCALLBACK**: Match a CallbackQuery with exact data match

Below is an example matching a callback with `Yes` as data and deleting the message linked to the callback.

```Go
bot.OnCallback("Yes", func(u *telebot.Update) {
    chatId := u.CallbackQuery.Message.Chat.Id
    messageId := u.CallbackQuery.Message.Id

    res, err := bot.DeleteMessage(chatId, messageId)

    if err != nil {
        log.Printf("Error deleting message: %s", err.Error())
    }
})
```

* **ONPAYLOAD**: Match a CallbackQuery with data starting with the filter

Below is an example matching a callback event with a payload starting with `Yes` and delete the source message.

```Go
bot.OnPayload("Yes", func(u *telebot.Update) {
    chatId := u.CallbackQuery.Message.Chat.Id
    messageId := u.CallbackQuery.Message.Id

    res, err := bot.DeleteMessage(chatId, messageId)

    if err != nil {
        log.Printf("Error deleting message: %s", err.Error())
    }
})
```

## How to make the bot send content to Telegram chat with telebot ?

In addition to update reception, telebot has some functions designed to make your bot send content. You can use it in your handlers.

### List of message methods available

Methods aiming at sending messages are defined in [messages.go](messages.go).

* **SendTextMessage**: Sends a text message.

```Go
bot.SendTextMessage(chatId int, text string, options telebot.SendMessageOptions)
```

```Go
    bot.OnCommand("/repeat", func(u *telebot.Update) {
        payload := u.Message.Text[len("/repeat"):]
        chatId := u.Message.Chat.Id

        _, err := bot.SendTextMessage(chatId, payload, telebot.SendMessageOptions{ReplyToMessageId: u.Message.Id, AllowSendingWithoutReply: true, DisableWebPagePreview: true})

        if err != nil {
            log.Printf("Error sending message: %s", err.Error())
        }
    })
```

Options are a struct with type and supports telegram API options described [here](https://core.telegram.org/bots/api#sendmessage) with the same name.

```Go
type SendMessageOptions struct {
    ParseMode                string
    DisableWebPagePreview    bool
    DisableNotification      bool
    ReplyToMessageId         int
    AllowSendingWithoutReply bool
}
```

* **SendReplyKeyboardMarkupTextMessage**: Send a text message with a ReplyKeyboardMarkup keyboard.

```Go
bot.SendReplyKeyboardMarkupTextMessage(chatId int, text string, keyboard ReplyKeyboardMarkup, options SendMessageOptions)
```

You can create a custom reply keyboard and send the message with the following code snippet.

```Go
text := "Your text here"
yesButton := telebot.KeyboardButton{"Yes"}
noButton := telebot.KeyboardButton{"No"}
firstRow := []telebot.KeyboardButton{yesButton, noButton}
keyboard := telebot.ReplyKeyboardMarkup{Keyboard: [][]telebot.KeyboardButton{firstRow}}

res, err := bot.SendReplyKeyboardMarkupTextMessage(chatId, text, keyboard, telebot.SendMessageOptions{})
```

* **SendReplyKeyboardRemoveTextMessage**: Send a text message with a ReplyKeyboardRemove keyboard

```Go
bot.SendReplyKeyboardRemoveTextMessage(chatId int, text string, selective bool, options SendMessageOptions)
```

* **SendInlineKeyboardMarkupTextMessage**: Send a text message with an inline keyboard

```Go
bot.SendInlineKeyboardMarkupTextMessage(chatId int, text string, keyboard InlineKeyboardMarkup, options SendMessageOptions)
```

You can define a custom inline keyboard the same way as below.

```Go
bot.OnText("/hello", func(u *telebot.Update) {
    chatId := u.Message.Chat.Id
    text := "Hello ?"

    yesButton := telebot.InlineKeyboardButton{"Hello !", "Hello"}
    noButton := telebot.InlineKeyboardButton{"WHo are you", "WhoAreYou"}

    firstRow := []telebot.InlineKeyboardButton{yesButton, noButton}
    keyboard := telebot.InlineKeyboardMarkup{[][]telebot.InlineKeyboardButton{firstRow}}

    res, err := bot.SendInlineKeyboardMarkupTextMessage(chatId, text, keyboard, telebot.SendMessageOptions{})

    if err != nil {
            log.Printf("Error sending message: %s", err.Error())
        }

})
```

* **EditTextMessage**: Edit a text message

```Go
bot.EditTextMessage(chatId int, newText string, messageId int, options SendMessageOptions)
```

Some message options are available. Specify the options using `SendMessageOptions` like if you were using `SendTextMessage` to send a message.

* **EditInlineKeyboardTextMessage**: Edit a text message with InlineKeyboardMarkup

```Go
bot.EditInlineKeyboardTextMessage(chatId int, newText string, messageId int, newKeyboard InlineKeyboardMarkup, options SendMessageOptions)
```

* **EditMessageInlineKeyboardMarkup**: Edit the inline keyboard of a message

```Go
bot.EditMessageInlineKeyboardMarkup(chatId int, messageId int, newKeyboard InlineKeyboardMarkup)
```

* **DeleteMessage**: Delete a message

```Go
bot.DeleteMessage(chatId int, messageId int)
```

### List of callback methods available

Some methods defined in `callback.go` can be used to answer CallbackQueries.

* **AnswerCallbackQuery**: Answer a callback query without notification

```Go
bot.AnswerCallbackQuery(callbackQueryId string)
```

* **AnswerCallbackQueryNotification**: Answer a callback query with notification

```Go
bot.AnswerCallbackQueryNotification(callbackQueryId string, text string, showAlert bool)
```

### List of chat methods available

You can use the methods defined in `chat.go` to manage a channel (kick, unban users)

* **KickChatMember**: Kick an user from a group

```Go
bot.KickChatMember(chatId int, userId int)
```

* **UnbanChatMember**: Unban an user from a group

```Go
bot.UnbanChatMember(chatId int, userId int)
```

### List of dice methods

The methods defined in `dice.go` allow your bot to send random animated emoji such as a dice toss

* **SendDice**: Sends a dice.

```Go
bot.SendDice(chatId int, options telebot.SendMessageOptions)
```

```Go
    bot.OnText("/dice", func(u *telebot.Update) {
        chatId := u.Message.Chat.Id

        _, err := bot.SendDice(chatId, telebot.SendMessageOptions{ReplyToMessageId: u.Message.Id, AllowSendingWithoutReply: true})

        if err != nil {
            log.Printf("Error sending message: %s", err.Error())
        }
    })
```

* **SendRandomDice**: send a random dice

```Go
bot.SendRandomDice(chatId int, options SendMessageOptions)
```

* **SendDiceEmoji**: Send a dice Emoji (Supported emojis : “🎲”, “🎯”, “🏀”, “⚽”, “🎳”, or “🎰”. Default is “🎲”. )

```Go
bot.SendDiceEmoji(chatId int, emoji string, options SendMessageOptions)
```

Check the [Telegram API documentation](https://core.telegram.org/bots/api#senddice) to see the options of Telegram sendDice API function supported (defined in telebot.SendMessageOptions).

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
    bot.OnText("/test", func(u *telebot.Update) {
        text := "I hear you <strong>loud and clear </strong> !"
        chatId := u.Message.Chat.Id
        parseMode := "HTML"

        _, err := bot.SendTextMessage(chatId, text, telebot.SendMessageOptions{ParseMode: parseMode})

        if err != nil {
            log.Printf("Error sending message: %s", err.Error())
        }
    })

    // Bin a handler to the command /repeat
    bot.OnCommand("/repeat", func(u *telebot.Update) {
        payload := u.Message.Text[len("/repeat"):]
        chatId := u.Message.Chat.Id

        _, err := bot.SendTextMessage(chatId, payload, telebot.SendMessageOptions{ReplyToMessageId: u.Message.Id, AllowSendingWithoutReply: true, DisableWebPagePreview: true})

        if err != nil {
            log.Printf("Error sending message: %s", err.Error())
        }
    })

    // Start the bot.
    bot.Start()

}
```

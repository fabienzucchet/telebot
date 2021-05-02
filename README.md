# Telebot

Telebot is a Go wrapper for the Telegram Bot API. It provides a convenient way to write a Telegram in Go whitout diving in Telegram API.
Updates are receveived from Telegram either with Telegram Webhook or with calls to Telegram /getUpdates API endpoint.

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
    const chatId = u.Message.From
    _, err := telebot.SendTextToTelegramChat(chatId, text)

    if err != nil {
        log.Printf("Error sending message: %s", err.Error())
        }
    })
```

* And last but not least, start the bot

```Go
    bot.Start()
```

In addition to update reception, telebot has some functions designed to make your bot send content. You can use it in your handlers.

package main

import (
	"os"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
	"github.com/mamal72/golyrics"
)

func lyrics(input string) string {
	if input == "help" || input == "/help" {
		return "Type your song name like this example: Blackfield:Some Day"
	}
	suggestions, err := golyrics.SearchTrack(input)
	if err != nil || len(suggestions) == 0 {
		return "Couldn't find your requested song"
	}

	track := suggestions[0]
	err = track.FetchLyrics()

	if err != nil {
		panic(err)
	}
	result := "===" + track.Name + " by " + track.Artist + "===" + "\n\n " + track.Lyrics
	return result
}

func main() {
	err := godotenv.Load()
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TG_TOKEN"))
	if err != nil {
		panic(err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, lyrics(update.Message.Text))
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}

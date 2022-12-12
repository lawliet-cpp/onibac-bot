package main

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var numericKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("/math"),
		tgbotapi.NewKeyboardButton("/physics"),
		tgbotapi.NewKeyboardButton("/science"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("/arabic"),
		tgbotapi.NewKeyboardButton("/english"),
		tgbotapi.NewKeyboardButton("/french"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("/hisgeo"),
		tgbotapi.NewKeyboardButton("/islamics"),
	),
)

func main() {
	const TELEGRAM_API_TOKEN = "5496513460:AAG4cuLo2am7b8V5eoC1Z6B2WcVVld-E698"
	bot, err := tgbotapi.NewBotAPI(TELEGRAM_API_TOKEN)
	if err != nil {
		panic(err)
	}
	fmt.Println("Tht bot " + bot.Self.UserName + " Is Ready")
	bot.Debug = true

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)
	for update := range updates {

		if update.Message == nil {
			continue
		}

		if !update.Message.IsCommand() {
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		switch update.Message.Command() {

		case "math":
			link := GetRandomMathExam()
			msg.Text = link + "\n Random Maths exam"
		case "physics":
			link := GetRandomPhysicsExam()
			msg.Text = link + "\n  Random physics Exam"

		case "arabic":
			link := GetRandomArabicExam()
			msg.Text = link + "\n  Random arabic Exam"
		case "hisgeo":
			link := GetRandomHistoryExam()
			msg.Text = link + "\n  Random hisgeo Exam"
		case "french":
			link := GetRandomFrenchExam()
			msg.Text = link + "\n  Random french Exam"

		case "english":

			link := GetRandomEnglishExam()
			msg.Text = link + "\n  Random english Exam"
		case "sceince":
			link := GetRandomScienceExam()
			msg.Text = link + "\n  Random science Exam"
		case "islamics":
			link := GetRandomIslamicsExam()
			msg.Text = link + "\n  Random islamics Exam"
		case "start":
			msg.ReplyMarkup = numericKeyboard
			msg.Text = "Choose a Subject"
		default:
			msg.Text = "I don't know this command"
		}

		if _, err := bot.Send(msg); err != nil {
			panic(err)
		}

	}

}

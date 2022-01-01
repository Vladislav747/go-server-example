package main

import (
	"encoding/json"
	"fmt"
	tgbotapi "gopkg.in/telegram-bot-api.v4"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

//  Для вендоринга используется GB
//  сборка проекта осуществляется с помощью gb build
//  установка зависимостей - gb vendor fetch gopkg.in/telegram-bot-api.v4
//  установка зависимотей из манифеста - gb vendor restore

type Joke struct {
	ID   uint32 `json:"id"`
	Joke string `json:"joke"`
}

type JokeResponse struct {
	Type  string `json:"type"`
	Value Joke   `json:"value"`
}

const WebhookURL = "https://msu-go-20177.herokuapp.com"
const APIKey = "510741428:AAGaezQQbrw8NiIxDtpSp5B9VnnXXYMauAg"

//Это те кнопки которые показываются внизу
var buttons = []tgbotapi.KeyboardButton{
	tgbotapi.KeyboardButton{Text: "Get Joke"},
}

func getJoke() string {
	c := http.Client{}
	resp, err := c.Get("http://api.icndb.com/jokes/random?limitTo=[nerdy]")
	if err != nil {
		return "jokes API not responding"
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	joke := JokeResponse{}

	err = json.Unmarshal(body, &joke)
	if err != nil {
		return "Joke error"
	}

	return joke.Value.Joke
}

func main() {
	//Heorku прокидывает порт для приложения
	port := os.Getenv("PORT")
	bot, err := tgbotapi.NewBotAPI(APIKey)

	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	//Устанавливаем вебхук
	_, err = bot.SetWebhook(tgbotapi.NewWebhook(WebhookURL))
	if err != nil {
		log.Fatal(err)
	}

	updates := bot.ListenForWebhook("/")
	go http.ListenAndServe(":"+port, nil)

	//получаем все обновления из канала updates
	for update := range updates {
		var message tgbotapi.MessageConfig
		fmt.Println("received text: ", update.Message.Text)

		switch update.Message.Text {
		case "Get Joke":
			//Если пользователь нажал на кнопку, то придет сообщение "Get Joke"
			message = tgbotapi.NewMessage(update.Message.Chat.ID, getJoke())
		default:
			message = tgbotapi.NewMessage(update.Message.Chat.ID, "Press 'Get Joke' to receive joke")

			//в ответном сообщение просим бота показать клавиатуру
			message.ReplyMarkup = tgbotapi.NewReplyKeyboard(buttons)

			bot.Send(message)
		}

	}

}

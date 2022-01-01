package main

//

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	tgbotapi "gopkg.in/telegram-bot-api.v4"
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

const WebhookURL = "https://msu-go-2017.herokuapp.com"
const APIKey = "510741428:AAGaezQQbrw8NiIxDtpSp5B9VnnXXYMauAg"

func getJoke() string {
	c := http.Client{}
	resp, err := c.Get("http://api.icndb.com/jokes/random?limitTi=[nerdy]")
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
	log.Printf("Authorized on account %s", bot.Self.Username)

	//Устанавливаем вебхук
	_, err = bot.SetWebhook(tgbotapi.NewWebhook(WebhookURL))
	if err != nil {
		log.Fatal(err)
	}
}

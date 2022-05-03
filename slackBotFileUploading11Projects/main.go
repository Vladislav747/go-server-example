package main

import (
	"context"
	"fmt"
	"github.com/shomali11/slacker"
	"github.com/slack-go/slack"
	"log"
	"os"
	"strconv"
)

//func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
//	for event := range analyticsChannel {
//		fmt.Println("Command Evenets")
//		fmt.Println(event.Timestamp)
//		fmt.Println(event.Command)
//		fmt.Println(event.Parameters)
//		fmt.Println(event.Event)
//		fmt.Println()
//	}
//}

func main() {
	//API Slackbot
	os.Setenv("SLACK_BOT_TOKEN", "")
	os.Setenv("CHANNEL_ID", "")

	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	chanelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{""}

	for i := 0; i < len(fileArr); i++ {
		params := slack.FileUploadParameters{
			Channels: chanelArr,
			File:     fileArr[i],
		}
		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}
		fmt.Printf("Name: %s, Url:%s\n", file.Name, file.URL)
	}

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	//Горутина обработки команд
	go printCommandEvents(bot.CommandEvents())

	bot.Command("My job is to...", &slacker.CommandDefinition{
		Description: "you calcualtor",
		Example:     "my yob is 2020",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			//year of birth - yob
			yob, err := strconv.Atoi(year)
			if err != nil {
				println("error")
			}
			age := 2021 - yob
			r := fmt.Sprintf("age is %d", age)
			response.Reply(r)
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

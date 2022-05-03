package main

import (
	"fmt"
	"github.com/slack-go/slack"
	"os"
)

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
}

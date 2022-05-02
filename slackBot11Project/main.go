package main

import (
	"github.com/showmali11/slacker"
	"os"
)

func main() {
	//API Slackbot
	os.Setenv("SLACK_BOT_TOKEN", "")
	os.Setenv("SLACK_APP_TOKEN", "")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"))
}

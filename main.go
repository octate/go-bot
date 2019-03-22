package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/shomali11/slacker"
)

func checkWebsiteStatus(website string) string {
	timeout := time.Duration(1 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	resp, err := client.Get(website)
	if err != nil {
		return ("ERROR:\n" + err.Error())
	}
	return (string(resp.StatusCode) + resp.Status)
}

func main() {
	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"))

	bot.Command("Hello", &slacker.CommandDefinition{
		Handler: func(request slacker.Request, response slacker.ResponseWriter) {
			response.Reply("Hello, I am go-bot\n\nI can do the following:\n- [wstatus] check https://www.mrenmajozi.com status")
		},
	})

	bot.Command("wstatus", &slacker.CommandDefinition{
		Handler: func(request slacker.Request, response slacker.ResponseWriter) {
			response.Reply(checkWebsiteStatus(os.Getenv("MY_WEBSITE")))
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

package main

// go get "github.com/slack-go/slack"
// go mod tidy

// xoxb-3379758857668-3414360301952-jotysIKGwXzYvnYI84P1V6Gt
// C03B3C5P2TC

// channels on app.slack.com/client
// right click the channel required (general) and press Channel Details. At the bottom-most is the channel id

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-3379758857668-3414360301952-jotysIKGwXzYvnYI84P1V6Gt")
	os.Setenv("CHANNEL_ID", "C03B3C5P2TC")
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{"IEEEHackHub2022.jpg"}

	for i := 0; i < len(fileArr); i++ {
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File: fileArr[i],
		}
		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Printf("%s\n", err)
		}
		fmt.Printf("Name: %s, URL: %s\n", file.Name, file.URL)
	}
}
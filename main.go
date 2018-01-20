package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/urfave/cli"
)

var (
	channel  string
	emoji    string
	message  string
	userName string
	webhook  string
)

func checkError(err error) {
	if err != nil {
		log.Fatal("Fatal error:", err.Error())
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "Slack Push Message"
	app.Usage = "Enjoy sending messages through the terminal"
	app.UsageText = os.Args[0] + " [global options] [arguments...]"
	app.Version = "0.1.2"
	app.Authors = []cli.Author{
		cli.Author{
			Name: "Boris Bakshiyev",
		},
	}
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "channel, c",
			Usage:       "The `channel` to send the message too",
			Destination: &channel,
		},
		cli.StringFlag{
			Name:        "emoji, e",
			Usage:       "The emoji that will apear next to the username",
			Value:       ":sos:",
			Destination: &emoji,
		},
		cli.StringFlag{
			Name:        "message, m",
			Usage:       "The `message` to be sent",
			Destination: &message,
		},
		cli.StringFlag{
			Name:        "username, u",
			Usage:       "The `username` which will be shown as the sender",
			Destination: &userName,
		},
		cli.StringFlag{
			Name:        "webhook, w",
			Usage:       "The slack `webhook`",
			Destination: &webhook,
		},
	}

	app.Action = func(c *cli.Context) error {
		if len(os.Args) == 1 {
			errorMessage := os.Args[0] + ": missing arguments\nTry '" + os.Args[0] + " --help' for more information."
			return cli.NewExitError(errorMessage, 1)
		}

		channelToUse := "#" + channel
		payLoad := []byte("{\"channel\": \"" + channelToUse + "\", \"username\": \"" + userName + "\", \"text\": \"" + message + "\", \"icon_emoji\": \"" + emoji + "\"}")

		req, err := http.NewRequest("POST", webhook, bytes.NewBuffer(payLoad))
		checkError(err)

		client := &http.Client{}
		resp, err := client.Do(req)
		checkError(err)
		defer resp.Body.Close()

		if resp.Status != "200 OK" {
			fmt.Println(resp.Status)
			fmt.Println(resp.Header)
			os.Exit(2)
		} else {
			fmt.Println("Successfully sent message")
		}
		return nil
	}

	app.Run(os.Args)
}

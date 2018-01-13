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
	message  string
	webhook  string
	userName string
	channel  string
)

func main() {
	app := cli.NewApp()
	app.Name = "Slack Push Message"
	app.Usage = ""
	app.UsageText = os.Args[0] + " -m <MESSAGE> -w <WEBHOOK> -u <USERNAME> -c <CHANNEL>"
	app.Version = "0.1.1"
	app.Authors = []cli.Author{
		cli.Author{
			Name: "Boris Bakshiyev",
		},
	}
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "message, m",
			Usage:       "`Message` to be sent",
			Destination: &message,
		},
		cli.StringFlag{
			Name:        "webhook, w",
			Usage:       "The slack `webhook`",
			Destination: &webhook,
		},
		cli.StringFlag{
			Name:        "username, u",
			Usage:       "The `username` which will be shown as the sender",
			Destination: &userName,
		},
		cli.StringFlag{
			Name:        "channel, c",
			Usage:       "The `channel` to send the message too",
			Destination: &channel,
		},
	}

	app.Action = func(c *cli.Context) error {
		emoji := ":sos:"
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


func checkError(err error) {
	if err != nil {
		log.Fatal("Fatal error:", err.Error())
	}
}

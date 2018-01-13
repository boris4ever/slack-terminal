package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	message  = kingpin.Flag("message", "Message to be sent").Short('m').Required().String()
	webhook  = kingpin.Flag("webhook", "The slack webhook").Short('w').Required().String()
	userName = kingpin.Flag("username", "The username which will be shown as the sender").Short('u').Required().String()
	channel  = kingpin.Flag("channel", "The channel to send the message too").Short('c').Required().String()
	help     = kingpin.CommandLine.HelpFlag.Short('h')
)

func main() {
	kingpin.Parse()
	kingpin.Version("0.1.1")

	emoji := ":sos:"
	channelToUse := "#" + *channel
	payLoad := []byte("{\"channel\": \"" + channelToUse + "\", \"username\": \"" + *userName + "\", \"text\": \"" + *message + "\", \"icon_emoji\": \"" + emoji + "\"}")

	req, err := http.NewRequest("POST", *webhook, bytes.NewBuffer(payLoad))
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
	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		log.Fatal("Fatal error:", err.Error())
	}
}

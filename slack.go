package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	msg := flag.String("msg", "", "Message to be sent\n")
	url := flag.String("url", "", "The slack webhook\n")
	userName := flag.String("username", "", "The username which will send the message\n")
	channel := flag.String("channel", "", "The channel to send the message too\n")

	flag.Parse()

	if len(os.Args) != 5 {
		fmt.Println("You forgot to use variables, user -h flag for usage")
		os.Exit(2)
	}

	emoji := ":sos:"
	channelToUse := "#" + *channel
	payLoad := []byte("{\"channel\": \"" + channelToUse + "\", \"username\": \"" + *userName + "\", \"text\": \"" + *msg + "\", \"icon_emoji\": \"" + emoji + "\"}")

	req, err := http.NewRequest("POST", *url, bytes.NewBuffer(payLoad))
	checkError(err)

	client := &http.Client{}
	resp, err := client.Do(req)
	checkError(err)
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Body:", string(body))
}

func checkError(err error) {
	if err != nil {
		log.Fatal("Fatal error:", err.Error())
	}
}

package main

import (
	"fmt"
	"log"
	"os"
)

//xoxb-512933145905-512933530481-ig4w9AG3tPYojs7jsvUCTCV0
func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: mybot slack-bot-token\n")
		os.Exit(1)
	}

	// start a websocket-based Real Time API session
	ws, _ := slackConnect(os.Args[1])
	fmt.Println("mybot ready, ^C exits")

	for {
		// read each incoming message
		m, err := getMessage(ws)
		if err != nil {
			log.Fatal(err)
		}
		// see if we're mentioned
		if m.Type == "message" {
			if m.Text == "вопрос" {
				m.Text = getAnswer()
				postMessage(ws, m)
			} else {
				// huh?
				m.Text = fmt.Sprintf("Для получения информации отправь слово вопрос\n")
				postMessage(ws, m)
			}
		}
	}
}

func getAnswer() string {
	return fmt.Sprintf("ответ\n")
}

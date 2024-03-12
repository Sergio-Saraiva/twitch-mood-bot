package main

import (
	"fmt"
	"log"

	"github.com/Sergio-Saraiva/twitch-mood-bot/pkg/chat"
	"github.com/Sergio-Saraiva/twitch-mood-bot/pkg/gpt"
)

// Import the missing package

var Messages []chat.ChatMsg

func main() {
	c, err := chat.FromChatProgram("./bot.js")
	if err != nil {
		log.Fatal("Error creating chat program", err)
	}

	for msg := range c.Chat {

		fmt.Printf("Got a message: %v\n", msg)
		Messages = append(Messages, msg)

		if len(Messages) > 30 {
			fmt.Printf("messages got bigger than 100 %d\n", len(Messages))
			res, err := gpt.ModerateChat(Messages)
			if err != nil {
				fmt.Printf("error moderating chat %v", err)
			}
			for _, v := range res.Choices {
				fmt.Println(v)
			}
			Messages = []chat.ChatMsg{}
		}
	}
}

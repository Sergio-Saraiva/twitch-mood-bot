package gpt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/Sergio-Saraiva/twitch-mood-bot/pkg/chat"
)

var apiKey = ""

func ModerateChat(messages []chat.ChatMsg) (*GPTResponse, error) {

	var bodyString string
	for _, message := range messages {
		bodyString += strings.Join([]string{message.Name + ":" + message.Msg}, "\n")
	}

	requestBody := GPTRequest{
		Model:       "gpt-3.5-turbo-instruct",
		Prompt:      "You are an assistant that helps to moderate a chat. Please infer the mood, if is there any aggressive, racist, sexist message based on this log. Also highlight the main topics that are being discussed. Based on this log:" + bodyString,
		MaxTokens:   300,
		Temperature: 0,
		// Messages: []GPTRequestMessage{
		// 	{
		// 		Role:    "system",
		// 		Content: "You are an assistant that helps to moderate a chat.",
		// 	},
		// 	{
		// 		Role:    "user",
		// 		Content: "You are an assistant that helps to moderate a chat. Please infer the mood, if is there any aggressive, racist, sexist message based on this log" + bodyString,
		// 	},
		// },
	}

	gptBodyMarshal, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Printf("error marshaling gptBody %v", err)
		return nil, err
	}

	request, err := http.NewRequest("POST", "https://api.openai.com/v1/completions", bytes.NewBuffer(gptBodyMarshal))

	if err != nil {
		fmt.Printf("error creating http request %v", err)
		return nil, err
	}
	request.Header.Add("Authorization", "Bearer "+apiKey)
	request.Header.Add("Content-Type", "application/json")

	client := http.Client{}
	res, err := client.Do(request)
	if err != nil {
		fmt.Printf("error calling gpt service %v", err)
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		var gptErrorResponse GPTErrorResposne
		err := json.NewDecoder(res.Body).Decode(&gptErrorResponse)
		if err != nil {
			fmt.Printf("error decoding response %v", err)
			return nil, err
		}
		fmt.Print(res.Body)
		fmt.Printf("unexptected response %v %v", res.StatusCode, gptErrorResponse)
		return nil, err
	}

	var gptModerationResponse GPTResponse
	err = json.NewDecoder(res.Body).Decode(&gptModerationResponse)
	if err != nil {
		fmt.Printf("error decoding response %v", err)
		return nil, err
	}

	fmt.Printf("%v", gptModerationResponse)

	return &gptModerationResponse, nil
}

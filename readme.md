# TWITCH CHAT MOOD BOT

## Description

This is a bot that will analyze the chat of a twitch stream and determine the mood of the chat.

It also analyze a chat log and checks if theres any agressive or toxic messages. And, display the main topics of discussion.

Built using The Primeagen's [vim-with-me](https://github.com/ThePrimeagen/vim-with-me) code as a base. See files bot.js, main.go and chat.go.

## Tools

- Go Lang 1.22.1
- Javascript
- Node
- Tmi.js
- gpt-3.5-turbo-instruct

## How to use

1. Clone the repository
2. Run `npm install` to install the dependencies
3. Replace the `apiKey` variable in `pkg/gpt/gpt.go` with your own twitch keys
4. Replace `channels` in `bot.js` with the channels you want to analyze
5. Run `go run cmd/main.go` to start the bot

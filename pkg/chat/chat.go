package chat

import (
	"bufio"
	"errors"
	"os/exec"
	"strconv"
	"strings"
)

type ChatMsg struct {
	Name string
	Msg  string
	Bits int
}

type Chat struct {
	Chat chan ChatMsg
}

func parseFromStdout(msg string) (*ChatMsg, error) {
	parts := strings.SplitN(msg, ":", 2)
	if len(parts) != 2 {
		return nil, errors.New("malformed message")
	}

	messageType := parts[0]
	switch messageType {
	case "message":
		msgParts := strings.SplitN(parts[1], ":", 2)
		return &ChatMsg{
			Name: msgParts[0],
			Msg:  msgParts[1],
			Bits: 0,
		}, nil

	case "bits":
		msgParts := strings.SplitN(parts[1], ":", 3)

		bits, err := strconv.Atoi(msgParts[1])

		if err != nil {
			return nil, err
		}

		return &ChatMsg{
			Name: msgParts[0],
			Bits: bits,
			Msg:  msgParts[2],
		}, nil
	}

	return nil, errors.New("unknown message type")
}

func FromChatProgram(path string) (*Chat, error) {

	// spawn program and read from the standard out

	cmd := exec.Command(path)
	stdout, err := cmd.StdoutPipe()

	if err != nil {
		return nil, err
	}

	if err := cmd.Start(); err != nil {
		return nil, err
	}

	chat := make(chan ChatMsg)

	go func() {
		buf_reader := bufio.NewReader(stdout)
		for {
			line, _, err := buf_reader.ReadLine()
			if err != nil {
				// TODO: how to close the thing properly
				close(chat)
				break
			}
			msg, err := parseFromStdout(string(line))
			if err != nil {
				continue
			}

			chat <- *msg
		}
	}()

	return &Chat{
		Chat: chat,
	}, nil
}

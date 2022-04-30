package commands

import (
	"fmt"
	"strings"
	"github.com/meehighlov/celebot/telegram"
)


type StartCommand struct {}


func (handler StartCommand) OnEvent(event telegram.Event) {
	event.SendMessage(
		event.Message.GetChatIdStr(),
		"Hello, i'm celebot! Tell me about your friends birthdays and i will remind you about it ;)",
	)
}


type AddPersonCommand struct {}


func (handler AddPersonCommand) OnEvent(event telegram.Event) {
	params := getCommandParams(event.Message.Text)
	name := params["name"]
	bd := params["bd"]

	text := fmt.Sprintf("Added new person: %s birth date: %s", name, bd)
	event.SendMessage(
		event.Message.GetChatIdStr(),
		text,
	)
}

func getCommandParams (text string) map[string]string {
	// command syntax: command param1=value1 param2=value2
	fmt.Println("raw message text:", text)
	trancatedCommand := strings.Fields(text)

	var preparedParams = map[string]string{}

	params := trancatedCommand[1:]

	for _, param := range params {
		splitedParam := strings.Split(param, "=")
		if len(splitedParam) > 1 {
			paramName := splitedParam[0]
			paramValue := splitedParam[1]
			preparedParams[paramName] = paramValue
		}
	}

	fmt.Println("prepared params:", preparedParams)

	return preparedParams
}
package commands

import (
	"fmt"
	"strings"
	"github.com/meehighlov/celebot/telegram"
	"github.com/meehighlov/celebot/app/db"
)


type StartCommand struct {}
type AddPersonCommand struct {}
type RandomCongratulationCommand struct {}
type AddUserCommand struct {}


func (handler StartCommand) Handle(c telegram.Context) {
	c.SendMessage(
		c.Message.GetChatIdStr(),
		"Hello, i'm celebot! Tell me about your friends birthdays and i will remind you about it ;)",
	)
}

func (handler AddPersonCommand) Handle(c telegram.Context) {
	params := getCommandParams(c.Message.Text)
	name := params["name"]
	bd := params["bd"]

	text := fmt.Sprintf("Added new person: %s birth date: %s", name, bd)
	c.SendMessage(
		c.Message.GetChatIdStr(),
		text,
	)
}

func (handler RandomCongratulationCommand) Handle(c telegram.Context) {
	c.SendMessage(
		c.Message.GetChatIdStr(),
		"i don't know any congratulations yet, may be you would like add one?:)",
	)
}

func (handler AddUserCommand) Handle(c telegram.Context) {
	params := getCommandParams(c.Message.Text)
	user := db.User {
		Name: params["name"],
		BirthDate: params["bd"],
		UserLinks: []db.UserLink{},
	}

	user.Save([]db.UserLink{})
	text := fmt.Sprintf("Added new person: %s birth date: %s", user.Name, user.BirthDate)
	c.SendMessage(
		c.Message.GetChatIdStr(),
		text,
	)
}


func getCommandParams(text string) map[string]string {
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

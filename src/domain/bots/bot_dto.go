package bots

import (
	"github.com/ferralucho/chatter-stock/src/rest_errors"
	"strings"
)

type Bot struct {
	Id int64 `json:"id"`
}
type Enum string

type CommandType Enum

// Command tyoe
const (
	CommandTypeStock CommandType = "STOCK"
	CommandTypeChat  CommandType = "CHAT"
)

type Command struct {
	Id          int64       `json:"id"`
	Message     string      `json:"message"`
	DateCreated string      `json:"date_created"`
	UserId      int64       `json:"user_id"`
	CommandType CommandType `json:"command_type"`
}

func (command *Command) Validate() rest_errors.RestErr {
	command.Message = strings.TrimSpace(strings.ToLower(command.Message))
	if command.Message == "" {
		return rest_errors.NewBadRequestError("invalid message")
	}

	if command.UserId == 0 {
		return rest_errors.NewBadRequestError("invalid user id")
	}
	return nil
}

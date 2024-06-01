package api

import (
	"context"
)

type Command any

type CommandResult any

type CommandHandler interface {
	Handle(context.Context, Command) (CommandResult, error)
}

type Event any

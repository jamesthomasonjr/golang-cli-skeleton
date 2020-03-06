package app

import (
    "skeleton/command"
)

func init() {
	application.AddCommand(command.Version)
}

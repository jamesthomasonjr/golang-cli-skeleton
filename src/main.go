package main

import (
    "github.com/jamesthomasonjr/golang-cli-skeleton/app"
    "github.com/jamesthomasonjr/golang-cli-skeleton/command"
)

func main() {
    app.Register(command.Version)

    app.Execute()
}

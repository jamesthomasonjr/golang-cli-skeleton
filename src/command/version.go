package command

import (
    "fmt"

    "github.com/spf13/cobra"

    "github.com/jamesthomasonjr/golang-cli-skeleton/version"
)

var (
    Version = &cobra.Command{
        Use:   "version",
        Short: "Display version information",
        Long:  fmt.Sprintf("version shows the version details for the %s application.", version.ApplicationName()),
        Run: func(cmd *cobra.Command, args []string) {
            fmt.Println(version.MainVersionLine())
            fmt.Println(version.GitVersionLine())
            fmt.Println(version.BuildVersionLine())
        },
    }
)

package app

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"

    "github.com/jamesthomasonjr/golang-cli-skeleton/config"
    "github.com/jamesthomasonjr/golang-cli-skeleton/version"
)

var (
    configurationFile string
    profileName       string

    application = &cobra.Command{}
)

func Execute() {
    if err := application.Execute(); err != nil {
        logError(err)
    }
}

func Register(cmd *cobra.Command) {
    application.AddCommand(cmd)
}

func init() {
    cobra.OnInitialize(initializeConfig)

    executableName, err := version.ExecutableName()
    if err != nil {
        logError(err)
    }

    application.Use = executableName
    application.Short = version.ShortDescription()
    application.Long =  version.LongDescription()

    //@TODO: Don't do this, maybe? The default value for OS X if $XDG_CONFIG_HOME isn't set sucks. It's ~/Library/Preferences instead of ~/.config
    defaultConfigurationMessage := fmt.Sprintf("config file (default is $XDG_CONFIG_HOME/%s)", version.ApplicationName())
    defaultProfileMessage := "configuration profile to use (default is 'default')"

    application.PersistentFlags().StringVar(&configurationFile, "config", "", defaultConfigurationMessage)
    application.PersistentFlags().StringVar(&profileName, "profile", "", defaultProfileMessage)

    config.Main.BindPFlag("profile", application.PersistentFlags().Lookup("profile"))
    config.Main.SetDefault("profile", "default")
}

func initializeConfig() {
    config.Initialize(configurationFile)
}

func logError(msg interface{}) {
    fmt.Errorf("Error:", msg)
    os.Exit(1)
}

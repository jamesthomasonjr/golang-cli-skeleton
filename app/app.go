package app

import (
    "fmt"
    "os"

    "skeleton/version"

    "github.com/spf13/cobra"
    config "github.com/spf13/viper"

    "github.com/adrg/xdg"
)

var (
    defaults map[string]interface{}
    application = &cobra.Command{
        Use: version.ExecutableName(),
        Short: version.ShortDescription(),
        Long: version.LongDescription(),
        Version: version.VersionTemplate(),
    }
)

func Run() {
    if err := application.Execute(); err != nil {
        fmt.Errorf("Error:", err)
        os.Exit(1)
    }
}

func init() {
    initializeDefaults()
    initializeFlags()
    initializeBindings()
    readConfigurationFile()
}

func initializeDefaults() {
    defaults = make(map[string]interface{})

    //@TODO: Don't do this, maybe? The default value for OS X if $XDG_CONFIG_HOME isn't set sucks. It's ~/Library/Preferences instead of ~/.config
    configurationFile, err := xdg.ConfigFile(version.ApplicationName())
    if err != nil {
        fmt.Errorf("Error determining default config file: %s", err)
    }

    defaults["configuration_file"] = configurationFile
}

func initializeFlags() {
    defaultConfigurationFile := defaults["configuration_file"].(string)

    application.PersistentFlags().StringP("config", "c", defaultConfigurationFile, "configuration file to use")
}

func initializeBindings() {
    config.BindEnv("config", "CONFIG_FILE")
    config.BindPFlag("config", application.PersistentFlags().Lookup("config"))

    config.BindEnv("profile", "PROFILE")
    config.BindPFlag("profile", application.PersistentFlags().Lookup("profile"))
}

func readConfigurationFile() {
    configurationFile := config.GetString("config")

    config.SetConfigFile(configurationFile)
    config.SetConfigType("hcl")

    if err := config.ReadInConfig(); err != nil {
        if _, ok := err.(config.ConfigFileNotFoundError); ok {
            // Configuration file not found, continue
        } else {
            // Configuration found but couldn't be loaded, err
            fmt.Printf("Error using config file: %s\n", configurationFile)
        }
    } else {
        fmt.Printf("Using configuration file: %s\n", configurationFile)
    }
}

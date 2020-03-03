package config

import (
    "fmt"

    "github.com/spf13/viper"
    "github.com/adrg/xdg"

    "github.com/jamesthomasonjr/golang-cli-skeleton/version"
)

var (
    Main    *viper.Viper
    Profile *viper.Viper
)

func init () {
    Main = viper.New()

    Main.SetEnvPrefix(version.ApplicationName())
    Main.AutomaticEnv()
}

func Initialize(configurationFile string) {
    if configurationFile != "" {
        viper.SetConfigFile(configurationFile)
    } else {
        //@TODO: Don't do this, maybe? The default value for OS X if $XDG_CONFIG_HOME isn't set sucks. It's ~/Library/Preferences instead of ~/.config
        configurationFile, err := xdg.ConfigFile(version.ApplicationName())
        if err != nil {
            fmt.Errorf("Error reading default config file: %s", err)
        }

        viper.SetConfigFile(configurationFile)
        viper.SetConfigType("ini")
    }

    if err := viper.ReadInConfig(); err != nil {
        fmt.Printf("Using config file: %s\n", viper.ConfigFileUsed())
    }

    profileName := Main.GetString("profile")
    Profile = viper.Sub(profileName)

    if Profile == nil {
        Profile = viper.New()
    }

    Profile.SetEnvPrefix(version.ApplicationName())
    Profile.AutomaticEnv()

    Profile.SetDefault("profile_name", profileName)
}

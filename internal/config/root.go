package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type ServerConfig struct {
	ConfigFile  string `mapstructure:"config_file" json:"config_file"`
	Environment string `mapstructure:"environment" json:"environment"`

	TwitchClientId     string `mapstructure:"twitch_client_id" json:"twitch_client_id"`
	TwitchClientSecret string `mapstructure:"twitch_client_secret" json:"twitch_client_secret"`
	TwitchRedirectURI  string `mapstructure:"twitch_redirect_uri" json:"twitch_redirect_uri"`

	TwitchEventSubSecret   string `mapstructure:"twitch_eventsub_secret" json:"twitch_eventsub_secret"`
	TwitchEventSubCallback string `mapstructure:"twitch_eventsub_callback" json:"twitch_eventsub_callback"`

	VRCUsername string `mapstructure:"vrc_username" json:"vrc_username"`

	MongoURI string `mapstructure:"mongo_uri" json:"mongo_uri"`
}

var defaultConfig = ServerConfig{
	ConfigFile: "config.yaml",
}

var Config = viper.New()

func init() {
	Config.SetConfigFile(Config.GetString("config_file"))
	Config.SetConfigType("yaml")
	Config.AddConfigPath("./")

	err := Config.ReadInConfig()
	if err != nil {
		fmt.Println("Fatal error config file: default \n", err)
		os.Exit(1)
	}

	// Environment
	Config.AutomaticEnv()
	Config.SetEnvPrefix("GGAPI")
	Config.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	Config.AllowEmptyEnv(true)
}

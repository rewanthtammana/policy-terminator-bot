package config

import (
	"github.com/rewanthtammana/policy-terminator/utils"
	"github.com/spf13/viper"
)

type Config struct {
	POLICY_TERMINATOR_SLACK_BOT_TOKEN string `mapstructure:"POLICY_TERMINATOR_SLACK_BOT_TOKEN"`
	CHANNELID                         string `mapstructure:"CHANNELID"`
}

func LoadValues(path string) (config Config, err error) {
	viper.AddConfigPath(path)

	// Read env values from config.env
	viper.SetConfigName("config")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	utils.CheckIfError(err)

	err = viper.Unmarshal(&config)
	return
}

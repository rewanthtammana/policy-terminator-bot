package slack

import (
	"fmt"

	"github.com/rewanthtammana/policy-terminator/config"
	"github.com/rewanthtammana/policy-terminator/utils"
	"github.com/slack-go/slack"
)

// Notifies customer's operation channel on policy-termination actions
func NotifyUser(message string) {

	// Loads config values
	config, err := config.LoadValues(".")
	utils.CheckIfError(err)

	api := slack.New(config.POLICY_TERMINATOR_SLACK_BOT_TOKEN)
	channelId := config.CHANNELID

	// Posts message to respective slack channel
	channelId, timestamp, err := api.PostMessage(channelId, slack.MsgOptionText(message, false))
	utils.CheckIfError(err)

	fmt.Printf("Message successfully sent to channel %s at %s\n", channelId, timestamp)
}

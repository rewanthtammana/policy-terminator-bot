package slack

import (
	"fmt"

	"github.com/rewanthtammana/policy-terminator/config"
	"github.com/rewanthtammana/policy-terminator/utils"
	"github.com/slack-go/slack"
)

func NotifyUser(message string) {

	config, err := config.LoadValues(".")
	utils.CheckIfError(err)

	api := slack.New(config.POLICY_TERMINATOR_SLACK_BOT_TOKEN)
	channelId := config.CHANNELID

	fmt.Println("channelid = ", channelId)
	fmt.Println("token = ", config.POLICY_TERMINATOR_SLACK_BOT_TOKEN)

	channelId, timestamp, err := api.PostMessage(channelId, slack.MsgOptionText(message, false))
	utils.CheckIfError(err)

	fmt.Printf("Message successfully sent to channel %s at %s\n", channelId, timestamp)
}

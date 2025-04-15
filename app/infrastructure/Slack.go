
package infrastructure


import (
	"github.com/slack-go/slack"
)


type Slack struct {
	Channels struct {
		Test string
	}
	Connection *slack.Client
}


func NewSlack() *Slack {
	c := NewConfig()
	return newSlack(c)
}


func newSlack(c *Config) *Slack {

	s := &Slack{
		Connection: slack.New(c.Slack.APIToken),
	}

	s.Channels.Test = "bot_test"

	return s
}


func (s *Slack) SendMessage(message string) (channel string, timestamp string, text string, err error) {
	// func (api *Client) SendMessage(channel string, options ...MsgOption) (string, string, string, error)
	// func MsgOptionUsername(username string) MsgOption
	// func MsgOptionIconURL(iconEmoji string) MsgOption
	// func MsgOptionText(text string, escape bool) MsgOption
	return s.Connection.SendMessage(
		s.Channels.Test,
		slack.MsgOptionUsername("Mr. Bot man"),
		slack.MsgOptionIconURL("https://github.githubassets.com/images/modules/site/integrators/slackhq.png"),
		slack.MsgOptionText(message, true))
}

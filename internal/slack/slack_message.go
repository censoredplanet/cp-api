package slack

import "encoding/json"

type slackMessage struct {
	Text      string `json:"text"`
	IconEmoji string `json:"emoji"`
}

func messageData(text, iconEmoji string) ([]byte, error) {
	return json.Marshal(slackMessage{
		Text:      text,
		IconEmoji: iconEmoji,
	})
}

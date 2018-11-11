package pkg

import (
	"encoding/json"
	"net/http"
	"strings"
)

type SlackMessage struct {
	Text      string `json:"text"`
	Username  string `json:"username"`
	IconEmoji string `json:"icon_emoji"`
}

const (
	text      = "Happy Monday MegaMakers! :mm:\nOur weekly <{{.DiscoursePostURL}}|WAYWOTW post is up!>\nCheckout what we are all working on!"
	username  = "JonBot"
	iconemoji = ":robot_face:"
)

func BuildSlackRequest(apiURL, discoursePostURL string) (*http.Request, error) {

	msg := strings.Replace(text, "{{.DiscoursePostURL}}", discoursePostURL, -1)

	sMsg := SlackMessage{
		Text:      msg,
		Username:  username,
		IconEmoji: iconemoji,
	}

	return createJSONHTTPRequest(apiURL, sMsg)
}

func createJSONHTTPRequest(url string, msg SlackMessage) (*http.Request, error) {
	payload, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}

	// todo: better way to convert []byte -> reader?
	req, err := http.NewRequest("POST", url, strings.NewReader(string(payload)))
	if err != nil {
		return nil, err
	}
	req.Header.Add("content-type", "application/json")
	req.Header.Add("cache-control", "no-cache")

	return req, nil
}

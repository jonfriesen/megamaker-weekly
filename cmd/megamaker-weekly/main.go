package main

import (
	"fmt"
	"log"

	"github.com/jonfriesen/megamaker-weekly/pkg"
)

func main() {
	// Load up our environment variables
	// - slack URL (/w api key)
	// - discourse api key
	slackKeyURL, err := pkg.GetENV("SLACK_API_URL", "")
	if err != nil {
		log.Fatalln(err)
	}
	discourseAPIKey, err := pkg.GetENV("DISCOURSE_API_KEY", "")
	if err != nil {
		log.Fatalln(err)
	}

	// Create our discourse HTTP request
	discourseReq, err := pkg.BuildDiscourseRequest(discourseAPIKey)
	if err != nil {
		log.Fatalf("Failed to create Discourse request: %+v", err)
	}

	// execute our discourse requests
	discourseResp, err := pkg.DoPost(discourseReq)
	if err != nil {
		log.Fatalf("Failed to execute Discourse request: %+v", err)
	}

	// get slug from discourse resp
	slug := pkg.GetSlug(discourseResp)

	// Create our slack HTTP request (/w URL from the discourse response)
	slackReq, err := pkg.BuildSlackRequest(slackKeyURL, fmt.Sprintf("%s/t/%s", pkg.DiscourseURL, slug))
	if err != nil {
		log.Fatalf("Failed to create Slack request: %+v", err)
	}

	// execute our slack req
	_, err = pkg.DoPost(slackReq)
	if err != nil {
		log.Fatalf("Failed to execute Slack request: %+v", err)
	}

	// call it a day
}

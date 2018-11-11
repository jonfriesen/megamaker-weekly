package pkg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"time"
)

const DiscourseURL = "https://club.megamaker.co"

func BuildDiscourseRequest(apiKey string) (*http.Request, error) {
	url := DiscourseURL + "/posts"

	// d := buildDate()
	d := time.Now().String()

	params := make(map[string]string)

	params["api_key"] = apiKey
	params["title"] = fmt.Sprintf("What are you working on this week? %v", d)
	params["api_username"] = "jon"
	params["raw"] = createPostBody()
	params["category"] = "23"

	return createMultipartFormRequest(url, params)
}

func createMultipartFormRequest(url string, params map[string]string) (*http.Request, error) {

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	for k, v := range params {
		_ = writer.WriteField(k, v)
	}
	err := writer.Close()
	if err != nil {
		return nil, err
	}

	r, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return nil, err
	}

	r.Header.Add("content-type", writer.FormDataContentType())
	r.Header.Add("cache-control", "no-cache")

	return r, nil
}

func GetSlug(resp *http.Response) string {
	b := getStringBodyFromResponse(resp)

	m := handleJSONBody(b)

	return (*m)["topic_slug"]
}

func handleJSONBody(b string) *map[string]string {
	m := make(map[string]string)
	json.Unmarshal([]byte(b), &m)
	return &m
}

func createPostBody() string {
	return `
What are your goals for the week? What awesome stuff did you do last week? 

**Feel free to comment on others goals** 

Here's a template:

	**Last week**
	- [*] Thing 1
	- [*] Thing 2

	**This week**
	- [] Thing 1
	- [] Thing 2
`
}

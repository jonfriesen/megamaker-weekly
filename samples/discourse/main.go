package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
)

func main() {

	url := "https://club.megamaker.co/posts"

	params := make(map[string]string)

	params["api_key"] = "API Key"
	params["title"] = "Jon is doing something crazy again"
	params["api_username"] = "jon"
	params["raw"] = "this is my body, not's not great but it'll do"
	params["category"] = "23"

	req, err := createDiscourseRequest(url, params)
	if err != nil {
		log.Fatalf("Failed to create Discourse request %+v", err)
	}

	log.Printf("%+v", req.Header)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}

func createDiscourseRequest(url string, params map[string]string) (*http.Request, error) {

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

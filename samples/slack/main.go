package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	url := "url /w key"

	payload := strings.NewReader("{\n\t\"text\": \"Happy Monday MegaMakers! :mm:\\nOur weekly <https://example.com|WAYWOTW post is up!>\\nCheckout what we are all working on!\",\n\t\"username\": \"JonBot\",\n\t\"icon_emoji\": \":robot_face:\"\n\t\n}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/json")
	req.Header.Add("cache-control", "no-cache")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}

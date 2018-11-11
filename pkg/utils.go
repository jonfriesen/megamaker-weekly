package pkg

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

// DoPost
// posts stuff
func DoPost(req *http.Request) (*http.Response, error) {
	log.Printf("%+v", req.Header)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	// defer res.Body.Close()
	// body, _ := ioutil.ReadAll(res.Body)

	// fmt.Println(res)
	// fmt.Println(string(body))

	return res, nil
}

// GetENV accepts a key and a default
// if the default is empty, returns error on unfound envars
func GetENV(k, d string) (string, error) {
	v := os.Getenv(k)
	if v == "" {
		if d != "" {
			return d, nil
		}

		return "", fmt.Errorf("Environment variable %s not found", k)
	}

	return v, nil
}

func buildDate() string {
	t := time.Now()
	return fmt.Sprintf("%v %v, %v", t.Month().String(), t.Day(), t.Year())
}

func getStringBodyFromResponse(res *http.Response) string {
	defer res.Body.Close()
	// todo handle this error
	body, _ := ioutil.ReadAll(res.Body)

	return string(body)
}

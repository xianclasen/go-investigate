package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

var key string = os.Getenv("INVESTIGATE_KEY")
var base_url string = "https://investigate.api.umbrella.com"
var domain = os.Args[1]
var getCategory = Endpoint{"GET", base_url + "/domains/categorization/" + domain}

type Endpoint struct {
	method string
	url    string
}

func call(E Endpoint) (int, string) {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest(E.method, E.url, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Authorization", "Bearer "+key)
	req.Header.Set("User-Agent", "Go app")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	bodyStatus := resp.StatusCode
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)
	defer resp.Body.Close()
	return bodyStatus, bodyString
}

func main() {
	status, body := call(getCategory)
	if status == 200 {
		bodyBytes := []byte(body)
		jsonMap := make(map[string](interface{}))
		err := json.Unmarshal([]byte(bodyBytes), &jsonMap)
		if err != nil {
			fmt.Printf(err.Error())
		}
		responseMap := jsonMap["facebook.com"].(map[string]interface{})
		fmt.Printf("\nContent Categories: %s\n", responseMap["content_categories"])
	} else {
		fmt.Print(status)
	}
}

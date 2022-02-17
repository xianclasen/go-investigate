package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

var key string = os.Getenv("INVESTIGATE_KEY")
var base_url string = "https://investigate.api.umbrella.com/"

func call(url, method string) (int, string) {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Println(err)
	}
	req.Header.Set("Authorization", "Bearer "+key)
	req.Header.Set("User-Agent", "Go app")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	bodyStatus := resp.StatusCode
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)
	defer resp.Body.Close()
	return bodyStatus, bodyString
}

func getCategory(domain string) (int, string) {
	endpoint := "domains/categorization/"
	status, body := call(base_url+endpoint+domain, "GET")
	return status, body
}

func getVolume(domain string) (int, string) {
	endpoint := "domains/volume/"
	status, body := call(base_url+endpoint+domain, "GET")
	return status, body
}

func getCoocurances(domain string) (int, string) {
	endpoint := "recommendations/name/"
	status, body := call(base_url+endpoint+domain+".json", "GET")
	return status, body
}

func getPdnsByName(domain string) (int, string) {
	endpoint := "pdns/name/"
	status, body := call(base_url+endpoint+domain, "GET")
	return status, body
}

func getPdnsByDomain(domain string) (int, string) {
	endpoint := "pdns/domain/"
	status, body := call(base_url+endpoint+domain, "GET")
	return status, body
}

func getPdnsByIp(domain string) (int, string) {
	endpoint := "pdns/ip/"
	status, body := call(base_url+endpoint+domain, "GET")
	return status, body
}

func getRawPdns(domain string) (int, string) {
	endpoint := "pdns/raw/"
	status, body := call(base_url+endpoint+domain, "GET")
	return status, body
}

func getRelatedDomaims(domain string) (int, string) {
	endpoint := "links/name/"
	status, body := call(base_url+endpoint+domain, "GET")
	return status, body
}

func getSecurityInfo(domain string) (int, string) {
	endpoint := "security/name/"
	status, body := call(base_url+endpoint+domain, "GET")
	return status, body
}

func getRiskScore(domain string) (int, string) {
	endpoint := "domains/risk-score"
	status, body := call(base_url+endpoint+domain, "GET")
	return status, body
}

func getAsInfoByIp(domain string) (int, string) {
	endpoint := "bgp_routes/ip"
	status, body := call(base_url+endpoint+domain, "GET")
	return status, body
}

func getAsInfoByAsn(domain string) (int, string) {
	endpoint := "bgp_routes/asn"
	status, body := call(base_url+endpoint+domain, "GET")
	return status, body
}

func main() {
	domain := os.Args[1]
	status, body := getCategory(domain)
	if status == 200 {
		fmt.Printf(body)
	} else {
		fmt.Sprint(status)
	}
}

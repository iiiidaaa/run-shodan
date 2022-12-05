package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const LEAKIX_URL = "https://leakix.net/search"

type LeakixClient struct {
	key string
}

type LeakixResult struct {
	Time      string
	IP        string        `json:"ip"`
	Network   LeakixNetwork `json:"network"`
	Service   LeakixService `json:"service"`
	Protocol  string
	Transport []string
	Port      string
	Summary   string `json:"summary"`
}

type LeakixNetwork struct {
	OrganizationName string `json:"organization_name"`
}

type LeakixService struct {
	Software LeakixSoftware `json:"software"`
}

type LeakixSoftware struct {
	Name string `json:"name"`
}

func (c LeakixClient) Search(keyword string) []LeakixResult {
	fmt.Println("Leakix Start")
	var results []LeakixResult
	d := c.SearchLeakix(keyword)
	err := json.Unmarshal(d, &results)
	if err != nil {
		log.Printf("json unmarshal error when scanning leakix, error: %v", err)
	}
	fmt.Println("Leakix Finish")
	return results
}

func (c LeakixClient) SearchLeakix(keyword string) []byte {
	client := http.Client{}
	req, err := http.NewRequest("GET", LEAKIX_URL, nil)
	if err != nil {
		log.Fatalf(err.Error())
	}
	q := req.URL.Query()
	q.Add("scope", keyword)
	//	q.Add("page", strconv.Itoa(pageCount))
	req.Header.Set("x-api-key", c.key)
	req.Header.Set("Accept", "application/json")
	req.URL.RawQuery = q.Encode()
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf(err.Error())
	}
	return body
}

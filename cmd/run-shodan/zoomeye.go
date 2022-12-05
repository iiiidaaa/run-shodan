package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

const ZOOMEYE_URL = "https://api.zoomeye.org/host/search"

type ZoomeyeClient struct {
	key string
}

type ZoomeyeIPData struct {
	Available int
	Results   []ZoomeyeResult `json:"matches"`
}

type ZoomeyeResult struct {
	Timestamp string
	IP        string          `json:"ip"`
	PortInfo  ZoomeyePortInfo `json:"portinfo"`
	Protocol  ZoomeyeProtocol `json:"protocol"`
}

type ZoomeyeProtocol struct {
	Transport string
}

type ZoomeyePortInfo struct {
	Service string
	Port    int
	Banner  string
	App     string
}

func (c ZoomeyeClient) Search(keyword string) []ZoomeyeResult {
	fmt.Println("Zoomeye Start")
	pageCount := 0
	var results []ZoomeyeResult
	for {
		var data ZoomeyeIPData
		d := c.SearchZoomeye(keyword, pageCount)
		err := json.Unmarshal(d, &data)
		if err != nil {
			log.Printf("json unmarshal error when scanning zoomeye, error: %v", err)
		}
		results = append(results, data.Results...)
		if (pageCount+1)*20 > data.Available {
			break
		}
	}
	fmt.Println("Zoomeye Finish")
	return results
}

func (c ZoomeyeClient) SearchZoomeye(keyword string, pageCount int) []byte {
	client := http.Client{}
	req, err := http.NewRequest("GET", ZOOMEYE_URL, nil)
	if err != nil {
		log.Fatalf(err.Error())
	}
	q := req.URL.Query()
	q.Add("query", keyword)
	q.Add("page", strconv.Itoa(pageCount))
	req.Header.Set("API-KEY", c.key)
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

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

const CRIMINALIP_URL = "https://api.criminalip.io/v1/banner/search"

type CriminalIPClient struct {
	key string
}

type CriminalIPResponse struct {
	Data CriminalIPData `json:"data"`
}

type CriminalIPData struct {
	Count   int
	Results []CriminalIPResult `json:"result"`
}

type CriminalIPResult struct {
	ScanDtime   string `json:"scan_dtime"`
	IPAddress   string `json:"ip_address"`
	CloudSource string `json:"cloud_source"`
	OrgName     string `json:"org_name"`
	Product     string `json:"product"`
	ServiceName string `json:"service_name"`
	SocketType  string `json:"socket_type"`
	OpenPortNo  int    `json:"open_port_no"`
	Banner      string
	RawData     string
}

func (c CriminalIPClient) Search(keyword string) []CriminalIPResult {
	fmt.Println("Criminal IP Start")
	offset := 0
	var results []CriminalIPResult
	for {
		var resp CriminalIPResponse
		d := c.SearchCriminalIP(keyword, offset)
		err := json.Unmarshal(d, &resp)
		if err != nil {
			log.Printf("json unmarshal error when scanning criminalip, error: %v", err)
		}
		data := resp.Data
		results = append(results, data.Results...)
		if offset+100 > data.Count {
			break
		}
		offset = offset + 100
	}
	fmt.Println("Criminal IP Finish")
	return results
}

func (c CriminalIPClient) SearchCriminalIP(keyword string, offset int) []byte {
	client := http.Client{}
	req, err := http.NewRequest("GET", CRIMINALIP_URL, nil)
	if err != nil {
		log.Fatalf(err.Error())
	}
	q := req.URL.Query()
	q.Add("query", keyword)
	q.Add("offset", strconv.Itoa(offset))
	req.Header.Set("x-api-key", c.key)
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

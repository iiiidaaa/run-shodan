package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ns3777k/go-shodan/v4/shodan"
	// go modules required
)

func main() {
	args := os.Args
	if len(args) != 2 {
		log.Fatalf("Usage: go run main.go {target}")
	}
	target := args[1]
	dbConfig := InitConf()
	dbClient, err := NewClient(&dbConfig)
	if err != nil {
		log.Fatalf(err.Error())
	}

	shodanClient := shodan.NewEnvClient(nil)
	var shodanResults *shodan.HostMatch
	if os.Getenv("SHODAN_KEY") == "" {
		fmt.Println("SHODAN_KEY is not set. Shodan Scan is skipped.")
	} else {
		shodanResults, err = searchHost(shodanClient, target)
	}
	if err != nil {
		log.Fatalf("error occured: %+v", err)
	}
	var zoomeyeResults []ZoomeyeResult
	if os.Getenv("ZOOMEYE_KEY") == "" {
		fmt.Println("ZOOMEYE_KEY is not set. Zoomeye Scan is skipped.")
	} else {
		zcli := ZoomeyeClient{
			key: os.Getenv("ZOOMEYE_KEY"),
		}
		zoomeyeResults = zcli.Search(target)
	}
	var criminalipResults []CriminalIPResult
	if os.Getenv("CRIMINALIP_KEY") == "" {
		fmt.Println("CRIMINALIP_KEY is not set. CriminalIP Scan is skipped.")
	} else {
		ccli := CriminalIPClient{
			key: os.Getenv("CRIMINALIP_KEY"),
		}
		criminalipResults = ccli.Search(target)
	}
	var leakixResults []LeakixResult
	if os.Getenv("LEAKIX_KEY") == "" {
		fmt.Println("LEAKIX_KEY is not set. Leakix Scan is skipped.")
	} else {
		lcli := LeakixClient{
			key: os.Getenv("LEAKIX_KEY"),
		}
		leakixResults = lcli.Search(target)
	}

	var results []Result
	results = append(results, ShodanToResult(shodanResults)...)
	results = append(results, ZoomeyeToResult(zoomeyeResults)...)
	results = append(results, CriminalIPToResult(criminalipResults)...)
	results = append(results, leakixToResult(leakixResults)...)
	for _, r := range results {
		dbClient.InsertResult(context.TODO(), &r)
	}

}

// func printJsonGetHostResult(hostMatch *shodan.HostMatch) error {
// 	j, err := json.Marshal(hostMatch)
// 	if err != nil {
// 		return err
// 	}
// 	fmt.Println(string(j))
// 	return nil
// }

// func printTextGetHostResult(hostMatch *shodan.HostMatch) {

// 	for _, m := range hostMatch.Matches {
// 		fmt.Printf("product: %+v\n", m.Product)
// 		fmt.Printf("Hostnames: %+v\n", m.Hostnames)
// 		fmt.Printf("IP: %+v\n", m.IP)
// 		fmt.Printf("Port: %+v\n", m.Port)
// 		fmt.Printf("ISP: %+v\n", m.ISP)
// 		fmt.Printf("CPE: %+v\n", m.CPE)
// 		fmt.Printf("ASN: %+v\n", m.ASN)
// 		fmt.Printf("Banner: %+v\n", m.Banner)
// 		fmt.Printf("Link: %+v\n", m.Link)
// 		fmt.Printf("Transport: %+v\n", m.Transport)
// 		fmt.Printf("Domains: %+v\n", m.Domains)
// 		fmt.Printf("DeviceType: %+v\n", m.DeviceType)
// 		fmt.Printf("Location: %+v\n", m.Location)
// 		fmt.Printf("ShodanData: %+v\n\n", m.ShodanData)
// 	}
// 	fmt.Printf("total: %v\n", hostMatch.Total)
// }

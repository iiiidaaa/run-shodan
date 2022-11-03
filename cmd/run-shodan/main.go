package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/ns3777k/go-shodan/v4/shodan" // go modules required
)

func main() {
	args := os.Args
	if len(args) != 3 {
		log.Fatalf("Usage: go run main.go {text|json}, {target}")
	}
	if !(args[1] == "text") && !(args[1] == "json") {
		log.Fatalf("The format must be text or json")
	}
	client := shodan.NewEnvClient(nil)
	found, err := searchHost(client, args[2])
	if err != nil {
		log.Fatalf("error occured: %+v", err)
	}
	switch args[1] {
	case "text":
		printTextGetHostResult(found)
	case "json":
		printJsonGetHostResult(found)
	}

}

func printJsonGetHostResult(hostMatch *shodan.HostMatch) error {
	j, err := json.Marshal(hostMatch)
	if err != nil {
		return err
	}
	fmt.Println(string(j))
	return nil
}

func printTextGetHostResult(hostMatch *shodan.HostMatch) {

	for _, m := range hostMatch.Matches {
		fmt.Printf("product: %+v\n", m.Product)
		fmt.Printf("Hostnames: %+v\n", m.Hostnames)
		fmt.Printf("IP: %+v\n", m.IP)
		fmt.Printf("Port: %+v\n", m.Port)
		fmt.Printf("ISP: %+v\n", m.ISP)
		fmt.Printf("CPE: %+v\n", m.CPE)
		fmt.Printf("ASN: %+v\n", m.ASN)
		fmt.Printf("Banner: %+v\n", m.Banner)
		fmt.Printf("Link: %+v\n", m.Link)
		fmt.Printf("Transport: %+v\n", m.Transport)
		fmt.Printf("Domains: %+v\n", m.Domains)
		fmt.Printf("DeviceType: %+v\n", m.DeviceType)
		fmt.Printf("Location: %+v\n", m.Location)
		fmt.Printf("ShodanData: %+v\n\n", m.ShodanData)
	}
	fmt.Printf("total: %v\n", hostMatch.Total)
}

func searchHost(client *shodan.Client, query string) (*shodan.HostMatch, error) {
	options := &shodan.HostQueryOptions{
		Query:  query,
		Facets: "",
		Minify: false,
		Page:   0,
	}
	found, err := client.GetHostsForQuery(context.Background(), options)
	if err != nil {
		return nil, err
	}

	return found, nil
}

func getDNS(client *shodan.Client, hostnames []string) {
	dns, err := client.GetDNSResolve(context.Background(), hostnames)

	if err != nil {
		log.Panic(err)
	}
	log.Println(dns["google.com"])
}

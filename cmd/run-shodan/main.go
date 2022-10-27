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
	if len(args) != 2 {
		log.Fatalf("target must be set")

	}
	client := shodan.NewEnvClient(nil)
	searchHost(client, args[1])

}

func searchHost(client *shodan.Client, query string) {
	options := &shodan.HostQueryOptions{
		Query:  query,
		Facets: "",
		Minify: false,
		Page:   0,
	}
	found, err := client.GetHostsForQuery(context.Background(), options)
	if err != nil {
		log.Panic(err)
	}
	j, err := json.Marshal(found)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(string(j))

}

func getDNS(client *shodan.Client, hostnames []string) {
	dns, err := client.GetDNSResolve(context.Background(), hostnames)

	if err != nil {
		log.Panic(err)
	}
	log.Println(dns["google.com"])
}

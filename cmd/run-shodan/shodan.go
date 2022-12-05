package main

import (
	"context"

	"github.com/ns3777k/go-shodan/v4/shodan"
)

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

package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/ns3777k/go-shodan/v4/shodan"
)

type Result struct {
	ID          uint32
	Source      string
	ScanAt      time.Time
	IPAddress   string
	CloudSource string
	Orgnization string
	Product     string
	Service     string
	Socket      string
	Port        int
	Banner      string
	RawData     string
	Invisible   bool
}

func ShodanToResult(match *shodan.HostMatch) []Result {
	var results []Result
	for _, m := range match.Matches {
		time, err := time.Parse("2006-01-02T15:04:05", m.Timestamp)
		if err != nil {
			log.Fatalf("time parsing error, err:%v\n", err)
		}
		results = append(results, Result{
			Source:      "Shodan",
			ScanAt:      time,
			IPAddress:   string(m.IP),
			Orgnization: m.Organization,
			Product:     m.Product,
			Socket:      m.Transport,
			Port:        m.Port,
			Banner:      m.Banner,
			RawData:     fmt.Sprint(m.ShodanData),
		})
	}
	return results
}

func CriminalIPToResult(criminalipResults []CriminalIPResult) []Result {
	var results []Result
	for _, r := range criminalipResults {
		time, err := time.Parse("2006-01-02 15:04:05", r.ScanDtime)
		if err != nil {
			log.Fatalf("time parsing error, err:%v\n", err)
		}
		results = append(results, Result{
			Source:      "CriminalIP",
			ScanAt:      time,
			IPAddress:   r.IPAddress,
			Orgnization: r.OrgName,
			Product:     r.Product,
			Socket:      r.SocketType,
			Port:        r.OpenPortNo,
			Banner:      r.Banner,
			RawData:     fmt.Sprint(r),
		})
	}
	return results
}

func ZoomeyeToResult(zoomeyeResults []ZoomeyeResult) []Result {
	var results []Result
	for _, r := range zoomeyeResults {
		time, err := time.Parse("2006-01-02T15:04:05", r.Timestamp)
		if err != nil {
			log.Fatalf("time parsing error, err:%v\n", err)
		}
		results = append(results, Result{
			Source:    "Zoomeye",
			ScanAt:    time,
			IPAddress: r.IP,
			Product:   r.PortInfo.App,
			Service:   r.PortInfo.Service,
			Socket:    r.Protocol.Transport,
			Port:      r.PortInfo.Port,
			Banner:    r.PortInfo.Banner,
		})
	}
	return results
}

func leakixToResult(leakixResults []LeakixResult) []Result {
	var results []Result
	for _, r := range leakixResults {
		time, err := time.Parse(time.RFC3339Nano, r.Time)
		if err != nil {
			log.Fatalf("time parsing error, err:%v\n", err)
		}
		port, err := strconv.Atoi(r.Port)
		if err != nil {
			log.Fatalf("port parsing error,port:%v, err:%v\n", r.Port, err)
		}
		results = append(results, Result{
			Source:      "Leakix",
			ScanAt:      time,
			IPAddress:   r.IP,
			Product:     r.Service.Software.Name,
			Service:     r.Protocol,
			Orgnization: r.Network.OrganizationName,
			Socket:      strings.Join(r.Transport, ","),
			Port:        port,
			RawData:     r.Summary,
		})
	}
	return results
}

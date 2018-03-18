package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/publicgov/spain-boe-reader/net"
	"github.com/publicgov/spain-boe-reader/params"
	"github.com/publicgov/spain-boe-reader/summary"
)

var currentDate string

func main() {
	// parse command line argument
	flag.StringVar(&currentDate, "date", defaultTime(), "BOE publication date in format YYYYMMDD")
	flag.Parse()

	// create the URL for the day
	p := params.Params{
		SummaryType: "BOE",
		ItemType:    "S",
		Date:        currentDate,
	}

	// make the network request
	client := net.New(p)
	summary := client.MakeRequest()

	// print basic info
	log.Println(showBasicInfo(summary))
}

func defaultTime() string {
	time := time.Now().UTC()
	time.Format("2006-01-02")
	return fmt.Sprintf("%d%02d%02d", time.Year(), time.Month(), time.Day())
}

func showBasicInfo(b summary.BoeSummary) string {
	return fmt.Sprintf("Date(%s) Found %d diaries with %d sections",
		b.Meta.PublicationDate, len(b.Diaries), b.SectionsSize())
}

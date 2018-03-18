package main

import (
	"fmt"
	"log"

	"github.com/publicgov/spain-boe-reader/net"
	"github.com/publicgov/spain-boe-reader/params"
	"github.com/publicgov/spain-boe-reader/summary"
)

func main() {
	// create the URL for the day
	p := params.Params{
		SummaryType: "BOE",
		ItemType:    "S",
		Date:        "20170926",
	}

	// make the network request
	client := net.New(p)
	summary := client.MakeRequest()

	// print basic info
	log.Println(showBasicInfo(summary))
}

func showBasicInfo(b summary.BoeSummary) string {
	return fmt.Sprintf("Date(%s) Found %d diaries with %d sections",
		b.Meta.PublicationDate, len(b.Diaries), b.SectionsSize())
}

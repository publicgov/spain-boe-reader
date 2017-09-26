package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/publicgov/spain-boe-reader/params"
	"github.com/publicgov/spain-boe-reader/summary"
	"golang.org/x/net/html/charset"
)

func main() {
	// create the URL for the day
	p := params.Params{
		SummaryType: "BOE",
		ItemType:    "S",
		Date:        "20170926",
	}

	log.Println(p.ToString())

	// make the network request
	resp, err := http.Get(p.ToString())
	if err != nil {
		log.Fatalln("Can't make the network request", err)
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("Error reading body", err)
	}

	// parse the result
	var summary summary.BoeSummary
	decoder := xml.NewDecoder(bytes.NewReader(content))
	decoder.CharsetReader = charset.NewReaderLabel
	err = decoder.Decode(&summary)

	// print basic info
	log.Println(showBasicInfo(summary))
}

func showBasicInfo(b summary.BoeSummary) string {
	return fmt.Sprintf("Date(%s) Found %d diaries with %d sections",
		b.Meta.PublicationDate, len(b.Diaries), b.SectionsSize())
}

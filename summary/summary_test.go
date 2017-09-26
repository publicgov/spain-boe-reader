package summary

import (
	"encoding/xml"
	"log"
	"os"
	"testing"

	"golang.org/x/net/html/charset"
)

func TestSectionsSize(t *testing.T) {
	boe := readXML()
	expected := 7
	actual := boe.SectionsSize()
	if actual != expected {
		t.Error("Expected", expected, "but found", actual)
	}
}

func TestMetadata(t *testing.T) {
	summary := readXML()
	bulletinType := summary.Meta.BulletinType
	lastDate := summary.Meta.LastPublicationDate
	nextDate := summary.Meta.NextPublicationDate
	pubDate := summary.Meta.PublicationDate
	if bulletinType != "BOE" {
		t.Error("Expected bulletin type BOE, got", bulletinType)
	}
	if lastDate != "17/10/2014" {
		t.Error("Expected last data 18/10/2014, got", lastDate)
	}
	if nextDate != "20/10/2014" {
		t.Error("Expected last data 20/10/2014, got", nextDate)
	}
	if pubDate != "18/10/2014" {
		t.Error("Expected last data 18/10/2014, got", pubDate)
	}
}

func readXML() BoeSummary {
	xmlFile, err := os.Open("example.xml")
	if err != nil {
		log.Fatalln("Error opening file:", err)
	}
	defer xmlFile.Close()

	var summary BoeSummary
	decoder := xml.NewDecoder(xmlFile)
	decoder.CharsetReader = charset.NewReaderLabel
	err = decoder.Decode(&summary)

	return summary
}

package summary

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestBoeSummary(t *testing.T) {
	summary := readXML()
	bulletinType := summary.Meta.BulletinType
	if bulletinType != "BOE" {
		t.Error("Expected bulletin type BOE, got", bulletinType)
	}
}

func readXML() BoeSummary {
	xmlFile, err := os.Open("example.xml")
	if err != nil {
		log.Fatalln("Error opening file:", err)
	}
	defer xmlFile.Close()
	content, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		log.Fatalln("Error reading content from", err)
	}
	var summary BoeSummary
	xml.Unmarshal(content, &summary)
	return summary
}

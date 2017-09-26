package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/publicgov/spain-boe-reader/summary"
)

func main() {
	summary := readXML()
	bulletinType := summary.Meta.BulletinType
	fmt.Println(bulletinType)
}

func readXML() summary.BoeSummary {
	file := ""
	xmlFile, err := os.Open(file)
	if err != nil {
		log.Fatalln("Error opening file:", err)
	}
	defer xmlFile.Close()
	content, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		log.Fatalln("Error reading content from", err)
	}
	var summary summary.BoeSummary
	xml.Unmarshal(content, &summary)
	return summary
}

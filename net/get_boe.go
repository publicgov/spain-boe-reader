package net

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/publicgov/spain-boe-reader/params"
	"github.com/publicgov/spain-boe-reader/summary"
	"golang.org/x/net/html/charset"
)

// Client struct to make the network request
type Client struct {
	// url to request
	url params.Params
}

// New creates a new client to make network requests
func New(param params.Params) *Client {
	return &Client{param}
}

// MakeRequest executes the network request and return the struct with all info.
// See the BoeSummary struct
func (c *Client) MakeRequest() summary.BoeSummary {
	log.Println("Request URL:", c.url.ToString())
	resp, err := http.Get(c.url.ToString())
	if err != nil {
		log.Fatalln("Can't make the network request", err)
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("Error reading body", err)
	}

	// parse the result and return it
	var summary summary.BoeSummary
	decoder := xml.NewDecoder(bytes.NewReader(content))
	decoder.CharsetReader = charset.NewReaderLabel
	err = decoder.Decode(&summary)
	return summary
}

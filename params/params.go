package params

import "fmt"

// baseURL is the base URL to request. The variable format of this URL is
// id=PUB-I-AAAAMMDD, where PUB is the type of the summary, always BOE. The
// I is the type of the item, always S. The last element is the date, where
// AAAA is the year, MM is the month number and DD is the day number
const baseURL string = "http://boe.es/diario_boe/xml.php?id=%s-%s-%s"

// Params is the variable format of the URL
type Params struct {
	summaryType string
	itemType    string
	date        string
}

// ToString convert the Param struct to a string
func (p *Params) ToString() string {
	return fmt.Sprintf(baseURL, p.summaryType, p.itemType, p.date)
}

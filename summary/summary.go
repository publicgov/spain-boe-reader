package summary

// BoeSummary is a struct to document the different dispositions
// and announcements that have been published on a certain date
type BoeSummary struct {
	Meta    Metadata `xml:"meta"`
	Diaries []Diary  `xml:"diario"`
}

// Metadata is used to provide metainformation that has been published
// in a certain date
type Metadata struct {
}

// Diary contains the information of the dispositions that conform
// each of the bulletins on a certain date
type Diary struct {
}

package summary

// BoeSummary is a struct to document the different dispositions
// and announcements that have been published on a certain date
type BoeSummary struct {
	Meta    Metadata `xml:"meta"`
	Diaries []Diary  `xml:"diario"`
}

// Metadata is used to provide metainformation that has been published
// in a certain date. Note that all dates from this struct has the
// format dd/mm/yyyy
type Metadata struct {
	BulletinType        string `xml:"pub"`
	PublicationDate     string `xml:"fecha"`
	LastPublicationDate string `xml:"fechaAnt"`
	NextPublicationDate string `xml:"fechaSig"`
}

// Diary contains the information of the dispositions that conform
// each of the bulletins on a certain date
type Diary struct {
	BulletinNumber string     `xml:"nbo,attr"`
	Summary        PdfSummary `xml:"sumario_nbo"`
}

// PdfSummary contains information of the pdf summary associated to the bulletin
type PdfSummary struct {
	ID   string  `xml:"id,attr"`
	Info PdfInfo `xml:"urlPdf"`
}

// PdfInfo contains file info, like number of pages and size
type PdfInfo struct {
	SizeBytes  string `xml:"szBytes,attr"`
	SizeKBytes string `xml:"szKBytes,attr"`
	Pages      string `xml:"numPages,attr"`
}

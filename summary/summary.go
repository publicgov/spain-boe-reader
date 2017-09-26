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
	Sections       []Section  `xml:"seccion"`
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

// Section is the set os sections that compose the diary
type Section struct {
	Number      string       `xml:"num,attr"`
	Name        string       `xml:"nombre,attr"`
	Departments []Department `xml:"departamento"`
}

// Department contains the relative data to the various departments  which have
// been published dispositions belonging the current section
type Department struct {
	Name      string     `xml:"nombre,attr"`
	Epigraphs []Epigraph `xml:"epigrafe,omitempty"`
	Items     []Item     `xml:"item,omitempty"`
}

// Epigraph is the list in which the dispositions are grouped corresponding
// to the department. Only available for sections 1, 2 and 3
type Epigraph struct {
	Name  string `xml:"nombre,attr"`
	Items []Item `xml:"item,omitempty"`
}

// Item is the list of dispositions corresponding to the section. This element
// can be inside the epigraph or without it. If this element exists without
// the epigrap element, only apply for section 4 and 5
type Item struct {
	ID         string     `xml:"id,attr"`
	Title      string     `xml:"titulo"`
	PdfInfo    PdfInfo    `xml:"urlPdf,omitempty"`
	HTMLURL    string     `xml:"urlHtm,omitempty"`
	XMLURL     string     `xml:"urlXml,omitempty"`
	Supplement Supplement `xml:"suplemento,omitempty"`
}

// Supplement contain information necessary to obtain those documents which, in some times,
// are provided along with the bulletin of the day
type Supplement struct {
	Description string  `xml:"descripcion"`
	PdfInfo     PdfInfo `xml:"urlPdf"`
}

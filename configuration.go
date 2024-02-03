package fsxml

import "encoding/xml"

type FSConfiguration struct {
	XMLName     xml.Name  `xml:"configuration"`
	Name        string    `xml:"name,attr"`
	Description string    `xml:"description,attr,omitempty"`
	Settings    []Setting `xml:"settings"`
}

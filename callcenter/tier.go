package callcenter

import "encoding/xml"

type Tier struct {
	XMLName  xml.Name `xml:"tier"`
	Agent    string   `xml:"agent,attr"`
	Queue    string   `xml:"queue,attr"`
	Level    string   `xml:"level,attr"`
	Position string   `xml:"position,attr"`
}

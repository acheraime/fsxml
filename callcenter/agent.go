package callcenter

import "encoding/xml"

type Agent struct {
	XMLName         xml.Name `xml:"agent"`
	Name            string   `xml:"name,attr"`
	Type            string   `xml:"type,attr"`
	Contact         string   `xml:"contact,attr"`
	Status          string   `xml:"status,attr"`
	MaxNoAnswer     string   `xml:"max-no-answer,attr"`
	WrapUpTime      string   `xml:"wrap-up-time,attr"`
	RejectDelayTime string   `xml:"reject-delay-time,attr"`
	BusyDelayTime   string   `xml:"busy-delay-time,attr"`
}

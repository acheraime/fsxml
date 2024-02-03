package fsxml

import (
	"encoding/xml"
	"errors"
)

type Dialplan struct {
	XMLName    xml.Name `xml:"context"`
	Name       string   `xml:"name,attr"`
	Extensions []*Extension
}

type Extension struct {
	XMLName   xml.Name `xml:"extension"`
	Name      string   `xml:"name,attr"`
	Condition []*Condition
}

type Condition struct {
	XMLName    xml.Name `xml:"condition"`
	Field      string   `xml:"field,attr"`
	Expression string   `xml:"expression,attr"`
	Break      string   `xml:"break,attr,omitempty"`
	Actions    []*Action
}

type Action struct {
	XMLName     xml.Name `xml:"action"`
	Application string   `xml:"application,attr"`
	Data        string   `xml:"data,attr,omitempty"`
}

func NewDialplan(dpName string, extensions []*Extension) (*FSDocument, error) {
	if extensions == nil {
		return nil, errors.New("extensions cannot be nil")
	}
	doc := NewDocument("dialplan")
	dp := Dialplan{
		Name:       dpName,
		Extensions: extensions,
	}
	doc.SetChildElement(dp)

	return &doc, nil
}

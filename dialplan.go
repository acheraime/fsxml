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

func NewDialPlan(dpName string, extensions []*Extension) (*FSDocument, error) {
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

func (d *Dialplan) AddExtension(extension Extension) {
	extensions := d.Extensions
	extensions = append(extensions, &extension)

	d.Extensions = extensions
}

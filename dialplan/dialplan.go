package dialplan

import (
	"encoding/xml"
	"errors"

	"github.com/acheraime/fsxml"
)

type Dialplan struct {
	XMLName    xml.Name `xml:"context"`
	Name       string   `xml:"name,attr"`
	Extensions []*Extension
}

func New(dpName string, extensions []*Extension) (*fsxml.FSDocument, error) {
	if extensions == nil {
		return nil, errors.New("extensions cannot be nil")
	}
	doc := fsxml.NewDocument("dialplan")
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

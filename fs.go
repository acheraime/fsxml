package fsxml

import (
	"encoding/xml"
	"strings"
	"sync"
)

const (
	freeswitchDocType = "freeswitch/xml"
	XmlHeader         = "<?xml version=\"1.0\" encoding=\"UTF-8\" standalone=\"no\"?>\n"
)

var pmux sync.Mutex
var vmux sync.Mutex

// FSDocument represents the top root xml structure
// of any freeswitch xml document
type FSDocument struct {
	XMLName xml.StartElement `xml:"document"`
	Type    string           `xml:"type,attr"`
	Section DocumentSection
}

func (f *FSDocument) SetChildElement(el interface{}) {
	f.Section.ChildElement = el
}

func (f *FSDocument) MarshalIndent() ([]byte, error) {
	return xml.MarshalIndent(f, "", "  ")
}

func (f *FSDocument) Marshal() ([]byte, error) {
	return xml.Marshal(f)
}

// NewDocument instantiate and return a
// freeswitch xml document. The doctType
// parameter set the value of the type attribute
// of the document node
func NewDocument(docType string) FSDocument {
	docSection := DocumentSection{Name: docType}
	return FSDocument{
		Type:    freeswitchDocType,
		Section: docSection,
	}
}

type DocumentSection struct {
	XMLName      xml.Name `xml:"section"`
	Name         string   `xml:"name,attr"`
	Description  string   `xml:"description,attr,omitempty"`
	ChildElement interface{}
}

type Param struct {
	XMLName xml.Name `xml:"param"`
	Name    string   `xml:"name,attr"`
	Value   string   `xml:"value,attr"`
}

type Params []Param

func (p Params) Exists(name string) (bool, int) {
	name = strings.ToLower(name)
	for i, f := range p {
		if strings.ToLower(f.Name) == name {
			return true, i
		}
	}

	return false, 0
}

func (p *Params) Add(param Param) {
	pmux.Lock()
	if exits, index := p.Exists(param.Name); exits {
		(*p)[index] = param
	} else {
		(*p) = append((*p), param)
	}
	pmux.Unlock()
}

func (p *Params) Set(name, value string) {
	p.Add(Param{Name: name, Value: value})
}

type Variable struct {
	XMLName xml.Name `xml:"variable"`
	Name    string   `xml:"name,attr"`
	Value   string   `xml:"value,attr"`
}

type Variables []Variable

func (v Variables) Exists(name string) (bool, int) {
	name = strings.ToLower(name)
	for i, f := range v {
		if strings.ToLower(f.Name) == name {
			return true, i
		}
	}

	return false, 0
}

func (v *Variables) Add(variable Variable) {
	vmux.Lock()
	if exits, index := v.Exists(variable.Name); exits {
		(*v)[index] = variable
	} else {
		(*v) = append((*v), variable)
	}
	vmux.Unlock()
}

func (v *Variables) Set(name, value string) {
	v.Add(Variable{Name: name, Value: value})
}

type Setting struct {
	XMLName xml.Name `xml:"settings"`
	Params  Params   `xml:"param"`
}

type fsNotFound struct {
	XMLName xml.Name `xml:"result"`
	Status  string   `xml:"status,attr"`
}

func FSNotFound() FSDocument {
	doc := NewDocument("result")
	doc.SetChildElement(
		fsNotFound{
			Status: "not found",
		},
	)

	return doc
}

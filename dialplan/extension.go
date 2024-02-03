package dialplan

import (
	"encoding/xml"
	"strconv"
)

type Extension struct {
	XMLName    xml.Name `xml:"extension"`
	Name       string   `xml:"name,attr"`
	Conditions []*Condition
	Continue   string `xml:"continue,attr,omitempty"`
}

func NewExtension(name string, conditions []*Condition) Extension {
	return Extension{
		Name:       name,
		Conditions: conditions,
	}
}

func (e *Extension) SetContinue(flag bool) {
	e.Continue = strconv.FormatBool(flag)
}

func (e *Extension) AddCondition(condition Condition) {
	conditions := e.Conditions
	conditions = append(conditions, &condition)

	e.Conditions = conditions
}

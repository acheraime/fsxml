package fsxml

import (
	"encoding/xml"
	"strconv"
)

type Condition struct {
	XMLName    xml.Name `xml:"condition"`
	Field      string   `xml:"field,attr"`
	Expression string   `xml:"expression,attr"`
	Break      string   `xml:"break,attr,omitempty"`
	Actions    []*Action
}

func NewCondition(field, expression string, actions []*Action) Condition {
	return Condition{
		Field:      field,
		Expression: expression,
		Actions:    actions,
	}
}

func (c *Condition) SetBreak(flag bool) {
	c.Break = strconv.FormatBool(flag)
}

func (c *Condition) AddAction(action Action) {
	actions := c.Actions
	actions = append(actions, &action)

	c.Actions = actions
}

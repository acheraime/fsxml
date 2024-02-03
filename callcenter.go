package fsxml

import (
	"encoding/xml"
)

type CallCenter struct {
	FSConfiguration
	Queues []Queue `xml:"queues>queue"`
	Agents []Agent `xml:"agents>agent,omitempty"`
	Tiers  []Tier  `xml:"tiers>tier,omitempty"`
}

func (f *CallCenter) MarshalIndent() ([]byte, error) {
	return xml.MarshalIndent(f, "", "  ")
}

func (f *CallCenter) Marshal() ([]byte, error) {
	return xml.Marshal(f)
}

func NewCallCenter(settings []Setting, queues []Queue, agents []Agent, tiers []Tier) (*FSDocument, error) {
	cc := &CallCenter{
		Queues: queues,
		Agents: agents,
		Tiers:  tiers,
	}

	cc.Name = "callcenter.conf"
	cc.Description = "CallCenter"
	cc.Settings = settings

	document := NewDocument("configuration")
	document.SetChildElement(cc)
	return &document, nil
}

package fsxml

import (
	"encoding/xml"
	"errors"
)

type Directory struct {
	XMLName    xml.Name    `xml:"domain"`
	DomainName string      `xml:"name,attr"`
	Alias      bool        `xml:"alias,attr"`
	Params     []*Param    `xml:"params>param"`
	Variables  []*Variable `xml:"variables>variable,omitempty"`
	Groups     []*Group    `xml:"groups>group"`
}

type Group struct {
	XMLName xml.Name `xml:"group"`
	Name    string   `xml:"name,attr"`
	Users   []*User  `xml:"users>user"`
}

type User struct {
	XMLName   xml.Name    `xml:"user"`
	ID        string      `xml:"id,attr"`
	Mailbox   string      `xml:"mailbox,omitempty"`
	Params    []*Param    `xml:"params>param,omitempty"`
	Variables []*Variable `xml:"variables>variable,omitempty"`
}

func NewDirectory(fsDomain string, fsGroups []*Group, alias bool) (*FSDocument, error) {
	if fsGroups == nil {
		return nil, errors.New("a valid group with users is required")
	}
	doc := NewDocument("directory")
	dir := Directory{
		DomainName: fsDomain,
		Alias:      alias,
		Params: []*Param{
			{
				Name:  "dial-string",
				Value: "{presence_id=${dialed_user}@${dialed_domain}}${sofia_contact(${dialed_user}@${dialed_domain})}",
			},
		},
		Groups: fsGroups,
	}
	doc.SetChildElement(dir)

	return &doc, nil

}

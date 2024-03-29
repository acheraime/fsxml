package fsxml

import (
	"encoding/xml"
)

type Directory struct {
	XMLName    xml.Name  `xml:"domain"`
	DomainName string    `xml:"name,attr"`
	Alias      bool      `xml:"alias,attr"`
	Params     Params    `xml:"params>param"`
	Variables  Variables `xml:"variables>variable,omitempty"`
	Groups     []*Group  `xml:"groups>group"`
}

func (d *Directory) AddGroup(group Group) {
	groups := d.Groups
	groups = append(groups, &group)

	d.Groups = groups
}

type Group struct {
	XMLName xml.Name `xml:"group"`
	Name    string   `xml:"name,attr"`
	Users   []*User  `xml:"users>user"`
}

func NewGroup(name string, users []*User) Group {
	return Group{Name: name, Users: users}
}

func (g *Group) AddUser(user User) {
	users := g.Users
	users = append(users, &user)

	g.Users = users
}

func NewDirectory(fsDomain string, fsGroups []*Group, alias bool) (*FSDocument, error) {
	doc := NewDocument("directory")
	dir := Directory{
		DomainName: fsDomain,
		Alias:      alias,
		Groups:     fsGroups,
		Params: Params{
			{
				Name:  "dial-string",
				Value: "{presence_id=${dialed_user}@${dialed_domain}}${sofia_contact(${dialed_user}@${dialed_domain})}",
			},
		},
	}
	doc.SetChildElement(dir)

	return &doc, nil

}

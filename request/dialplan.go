package request

import (
	"net/http"

	"github.com/pkg/errors"
)

type DialplanRequest struct {
	request
	Section          string `fs:"section"`
	Tag              string `fs:"tag_name"`
	KeyValue         string `fs:"key_value"`
	KeyName          string `fs:"key_name"`
	Domain           string `fs:"variable_requested_domain_name"`
	CallerDomain     string `fs:"variable_domain_name"`
	Extension        string `fs:"Caller-Destination-Number"`
	CallerContext    string `fs:"Caller-Context"`
	Context          string `fs:"variable_user_context"`
	SipToUser        string `fs:"variable_sip_to_user"`
	CallDirection    string `fs:"Call-Direction"`
	ChannelVariables map[string]string
	CallerData       map[string]string
	HuntData         map[string]string
	EventData        map[string]string
}

func NewDialplanRequest(r *http.Request) (*DialplanRequest, error) {
	if err := r.ParseForm(); err != nil {
		return nil, err
	}

	req := new(DialplanRequest)
	req.requestForm = r.Form

	// Parse the request and populated map fields
	if err := req.parse(); err != nil {
		return nil, err
	}

	// Marshall will populate fields with fs tags
	if err := req.Marshal(req); err != nil {
		return nil, err
	}

	// Validation
	if err := req.validate(); err != nil {
		return nil, err
	}

	return req, nil
}

func (d *DialplanRequest) parse() error {
	d.ChannelVariables = d.ParseVariables()
	d.CallerData = d.ParseCallerData()
	d.HuntData = d.ParseHuntData()
	d.EventData = d.ParseEventData()

	return nil
}

func (d DialplanRequest) validate() error {
	if d.Section != "dialplan" {
		return errors.New("invalid request. dialplan section is expected")
	}
	return nil
}

func (d DialplanRequest) Fields() interface{} {
	return d
}

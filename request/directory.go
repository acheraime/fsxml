package request

import (
	"net/http"

	"github.com/pkg/errors"
)

type DirectoryRequest struct {
	request
	Section           string `fs:"section"`
	Tag               string `fs:"tag_name"`
	KeyValue          string `fs:"key_value"`
	KeyName           string `fs:"key_name"`
	RegDomain         string `fs:"variable_requested_domain_name"`
	CallerDomain      string `fs:"variable_domain_name"`
	Extension         string `fs:"Caller-Destination-Number"`
	CallerContext     string `fs:"Caller-Context"`
	Context           string `fs:"variable_user_context"`
	SipToUser         string `fs:"variable_sip_to_user"`
	CallDirection     string `fs:"Call-Direction"`
	SipReqUser        string `fs:"variable_sip_req_user"`
	DefaultGateway    string `fs:"variable_default_gateway"`
	RequestedUserName string `fs:"variable_requested_user_name"`
	FreeswitchHost    string `fs:"hostname"`
	Purpose           string `fs:"purpose"`
	Profile           string `fs:"sip_profile"`
	User              string `fs:"user"`
	Action            string `fs:"action"`
	Domain            string `fs:"domain"`
	AuthUser          string `fs:"sip_auth_username"`
	TenantID          string `fs:"variable_tenant_uuid"`
	SipAuthRealm      string `fs:"sip_auth_realm"`
	ChannelVariables  map[string]string
	CallerData        map[string]string
	HuntData          map[string]string
	EventData         map[string]string
}

func NewDirectoryRequest(r *http.Request) (*DirectoryRequest, error) {
	if err := r.ParseForm(); err != nil {
		return nil, err
	}

	req := new(DirectoryRequest)
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

func (d *DirectoryRequest) parse() error {
	d.ChannelVariables = d.ParseVariables()
	d.CallerData = d.ParseCallerData()
	d.HuntData = d.ParseHuntData()
	d.EventData = d.ParseEventData()

	return nil
}

func (d *DirectoryRequest) validate() error {
	if d.Section != "directory" {
		return errors.New("invalid request. directory section is expected")
	}
	return nil
}

func (d DirectoryRequest) Fields() interface{} {
	return d
}

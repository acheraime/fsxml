package request

import (
	"net/http"

	"github.com/pkg/errors"
)

type ConfigurationRequest struct {
	request
	Section          string `fs:"section"`
	Tag              string `fs:"tag_name"`
	KeyValue         string `fs:"key_value"`
	KeyName          string `fs:"key_name"`
	ChannelVariables map[string]string
	EventData        map[string]string
}

func NewConfigurationRequest(r *http.Request) (*ConfigurationRequest, error) {
	if err := r.ParseForm(); err != nil {
		return nil, err
	}

	req := new(ConfigurationRequest)
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

func (d *ConfigurationRequest) parse() error {
	d.ChannelVariables = d.ParseVariables()
	d.EventData = d.ParseEventData()

	return nil
}

func (d ConfigurationRequest) validate() error {
	if d.Section != "configuration" {
		return errors.New("invalid request. configuration section is expected")
	}
	return nil
}

func (d ConfigurationRequest) Fields() interface{} {
	return d
}

func (d ConfigurationRequest) IsCallcenter() bool {
	return d.KeyValue == "callcenter.conf"
}

func (d ConfigurationRequest) AsCallcenter() (*CallCenter, error) {
	if !d.IsCallcenter() {
		return nil, errors.New("this is not valid callcenter configuration request")
	}
	var c CallCenter
	if err := d.Marshal(&c); err != nil {
		return nil, err
	}

	return &c, nil
}

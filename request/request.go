package request

import (
	"fmt"
	"net/url"
	"reflect"
	"strings"
	"sync"

	"github.com/pkg/errors"
)

const (
	tagName          = "fs"
	variablePrefix   = "variable_"
	callerDataPrefix = "Caller-"
	huntDataPrefix   = "Hunt-"
	eventDataPrefix  = "Event-"
)

type FSRequest interface {
	parse() error
	validate() error
	Fields() interface{}
}

type request struct {
	requestForm url.Values
	mu          sync.Mutex
}

func (s request) parseVariable(prefix string) map[string]string {
	vars := map[string]string{}
	s.mu.Lock()
	for k, v := range s.requestForm {
		if strings.HasPrefix(k, prefix) {
			vars[strings.TrimPrefix(k, prefix)] = v[0]
		}
	}
	s.mu.Unlock()

	return vars
}

func (s request) ParseVariables() map[string]string {
	return s.parseVariable(variablePrefix)
}

func (s request) ParseCallerData() map[string]string {
	return s.parseVariable(callerDataPrefix)
}

func (s request) ParseHuntData() map[string]string {
	return s.parseVariable(huntDataPrefix)
}

func (s request) ParseEventData() map[string]string {
	return s.parseVariable(eventDataPrefix)
}

func (s *request) Marshal(in interface{}) error {
	r := reflect.Indirect(reflect.ValueOf(in))
	for i := 0; i < r.NumField(); i++ {
		// Get the field
		vf := r.Type().Field(i)
		tag := vf.Tag.Get(tagName)
		if tag == "-" {
			continue
		}

		field := r.FieldByName(vf.Name)
		if !field.IsValid() {
			return errors.Errorf("field not valid: %s", field)
		}

		if field.CanSet() && tag != "" {
			field.SetString(s.requestForm.Get(tag))
		}
	}

	return nil
}

func (d request) Dump() {
	for k, v := range d.requestForm {
		fmt.Printf("%s => %s\n", k, v)
	}
}

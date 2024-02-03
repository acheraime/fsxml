package event

import (
	"encoding/json"
	"reflect"
	"strconv"
	"strings"

	"github.com/acheraime/fsxml"
	"github.com/pkg/errors"
)

const tagName = "fs"

type EventSubClass string

const (
	CallCenterInfoClass       EventSubClass = "callcenter::info"
	KonektemCallCenterPublish EventSubClass = "callcenter::publish"
)

type EventType string

const (
	QueueCountEvent            EventType = "cc_queue_count"
	AgentUpdateEvent           EventType = "cc_agent_update"
	QueueOriginateToAgentEvent EventType = "outbound_agent_thread_run"
	CallCenterFunctionEvent    EventType = "callcenter_function"
)

type Event map[string]interface{}

func NewEvent(data []byte) (Event, error) {
	var e Event

	if err := json.Unmarshal(data, &e); err != nil {
		return nil, err
	}

	return e, nil
}

func (e Event) Get(field string) string {
	value := e.ValueOf(field)
	if v, ok := value.(string); ok {
		return v
	}
	return ""
}

func (e Event) ValueOf(field string) interface{} {
	value, ok := e[field]
	if !ok {
		return nil
	}

	return value
}

func (s Event) Marshal(in interface{}) error {
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
			field.SetString(s.Get(tag))
		}
	}

	return nil
}

func (e Event) Raw() ([]byte, error) {
	return json.Marshal(e)
}

func (e Event) Name() (*string, error) {
	name := fsxml.String(e.Get("Event-Name"))
	if name == nil {
		return nil, errors.New("Event-Name field does not exist")
	}

	return name, nil
}

func (e Event) SubClass() (*string, error) {
	sub := fsxml.String(e.Get("Event-Subclass"))
	if sub == nil {
		return nil, errors.New("Event-Subclass field does not exist")
	}

	return sub, nil
}

func (e Event) CallingFunc() string {
	return e.Get("Event-Calling-Function")
}

func (e Event) TimeStamp() int {
	ts := e.Get("Event-Date-Timestamp")
	its, err := strconv.Atoi(ts)
	if err != nil {
		return 0
	}

	return its
}

func (e Event) IsCallcenter() bool {
	sub, err := e.SubClass()
	if err != nil {
		return false
	}

	if sub == nil {
		return false
	}

	return *sub == string(CallCenterInfoClass)
}

func (e Event) IsFromKonektem() bool {
	sub, err := e.SubClass()
	if err != nil {
		return false
	}
	if sub == nil {
		return false
	}

	eventName, err := e.Name()
	if err != nil {
		return false
	}

	return *sub == string(KonektemCallCenterPublish) && *eventName == "KONEKTEM"
}

func (e Event) CallcenterAction() string {
	return e.Get("CC-Action")
}

func (e Event) KonektemAction() string {
	return e.Get("action")
}

func (e Event) AsMemberCount() (*MemberCount, error) {
	var memberCount eMemberCount
	if !e.IsCallcenter() && e.CallingFunc() != string(QueueCountEvent) &&
		e.CallcenterAction() != QueueMemberCountAction {
		return nil, errors.New("not a callcenter member count event")
	}

	if err := e.Marshal(&memberCount); err != nil {
		return nil, err
	}
	mc := memberCount.ToMemberCount()

	return &mc, nil
}

func (e Event) AsAgentUpdate() (*AgentUpdate, error) {
	var agentUpdate AgentUpdate
	if !e.IsCallcenter() || e.CallingFunc() != string(AgentUpdateEvent) ||
		!e.IsAgentStateUpdate() {

		return nil, errors.New("not a callcenter agent update event")
	}

	if err := e.Marshal(&agentUpdate); err != nil {
		return nil, err
	}

	return &agentUpdate, nil
}

func (e Event) AsAgentStatusUpdate() (*AgentStatusUpdate, error) {
	var agentStatusUpdate AgentStatusUpdate
	if !e.IsCallcenter() || e.CallingFunc() != string(AgentUpdateEvent) ||
		!e.IsAgentStatusUpdate() {
		return nil, errors.New("not a callcenter agent update event")
	}

	if err := e.Marshal(&agentStatusUpdate); err != nil {
		return nil, err
	}

	return &agentStatusUpdate, nil
}

func (e Event) AsQueueToAgent() (*QueueToAgent, error) {
	var qToAgent QueueToAgent

	if !e.IsCallcenter() || e.CallingFunc() != string(QueueOriginateToAgentEvent) ||
		!e.IsQueueToAgent() {

		return nil, errors.New("not a callcenter queue to agent originate event")
	}

	if err := e.Marshal(&qToAgent); err != nil {
		return nil, err
	}

	return &qToAgent, nil
}

func (e Event) AsMemberToQueue() (*MemberToQueue, error) {
	var mToQueue MemberToQueue

	if !e.IsCallcenter() || e.CallingFunc() != string(CallCenterFunctionEvent) ||
		!e.IsMemberToQueue() {

		return nil, errors.New("not a callcenter function event")
	}

	if err := e.Marshal(&mToQueue); err != nil {
		return nil, err
	}

	return &mToQueue, nil
}

func (e Event) IsQueueToAgent() bool {
	return e.CallcenterAction() == string(AgentOfferingAction) ||
		e.CallcenterAction() == string(BridgeAgentFailAction) ||
		e.CallcenterAction() == string(BridgeAgentEndAction) ||
		e.CallcenterAction() == string(BridgeAgentStartAction)
}

func (e Event) IsMemberToQueue() bool {
	return e.CallcenterAction() == string(MemberQueueStartAction) ||
		e.CallcenterAction() == string(MemberQueueEndAction)
}

func (e Event) IsAgentStateUpdate() bool {
	return e.Get("CC-Action") == string(AgentStateChangeAction) && e.Get("CC-Agent-State") != ""
}

func (e Event) IsAgentStatusUpdate() bool {
	return e.Get("CC-Action") == string(AgentStatusChangeAction) && e.Get("CC-Agent-Status") != ""
}

func (e Event) AsKonektemQueue() (*KQueue, error) {
	var kQueue KQueue

	if !e.IsFromKonektem() || !strings.HasPrefix(e.KonektemAction(), "queue") {
		return nil, errors.New("not a konektem queue related event")
	}
	data, err := e.Raw()
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &kQueue); err != nil {
		return nil, err
	}

	return &kQueue, nil
}

func (e Event) AsKonektemAgent() (*KAgent, error) {
	var kAgent KAgent

	if !e.IsFromKonektem() || !strings.HasPrefix(e.KonektemAction(), "agent") {
		return nil, errors.New("not a konektem agent related event")
	}
	data, err := e.Raw()
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &kAgent); err != nil {
		return nil, err
	}

	return &kAgent, nil
}

func (e Event) AsKonektemTier() (*KTier, error) {
	var kTier KTier

	if !e.IsFromKonektem() || !strings.HasPrefix(e.KonektemAction(), "tier") {
		return nil, errors.New("not a konektem tier related event")
	}
	data, err := e.Raw()
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &kTier); err != nil {
		return nil, err
	}

	return &kTier, nil
}

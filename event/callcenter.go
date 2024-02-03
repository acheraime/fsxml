package event

import (
	"strconv"
	"strings"

	"github.com/acheraime/fsxml"
)

type QueueToAgentAction string

const (
	AgentOfferingAction    QueueToAgentAction = "agent-offering"
	BridgeAgentStartAction QueueToAgentAction = "bridge-agent-start"
	BridgeAgentEndAction   QueueToAgentAction = "bridge-agent-end"
	BridgeAgentFailAction  QueueToAgentAction = "bridge-agent-fail"
)

type MemberToQueueAction string

const (
	MemberQueueStartAction MemberToQueueAction = "member-queue-start"
	MemberQueueEndAction   MemberToQueueAction = "member-queue-end"
)

const QueueMemberCountAction = "members-count"

type AgentUpdateAction string

const (
	AgentStateChangeAction  AgentUpdateAction = "agent-state-change"
	AgentStatusChangeAction AgentUpdateAction = "agent-status-change"
)

type eMemberCount struct {
	QueueName string `fs:"CC-Queue"`
	Action    string `fs:"CC-Action"`
	Members   string `fs:"CC-Count"`
	Selection string `fs:"CC-Selection"`
}

type MemberCount struct {
	QueueName   string
	Action      string
	Members     int
	Selection   string
	QueuePrefix string
	TenantID    string
}

func (e *eMemberCount) ToMemberCount() MemberCount {
	m := MemberCount{
		QueueName: e.QueueName,
		Action:    e.Action,
		Selection: e.Selection,
	}

	cc, err := strconv.Atoi(e.Members)
	if err != nil {
		cc = 0
	}

	m.Members = cc
	if strings.Contains(e.QueueName, "@") {
		m.QueuePrefix = strings.Split(e.QueueName, "@")[0]
		m.TenantID = strings.Split(e.QueueName, "@")[1]
	}

	return m
}

type AgentUpdate struct {
	Agent  string `fs:"CC-Agent"`
	Action string `fs:"CC-Action"`
	State  string `fs:"CC-Agent-State"`
}

func (a AgentUpdate) NamePrefix() string {
	return strings.Split(a.Agent, "@")[0]
}

func (a AgentUpdate) OrgID() string {
	return strings.Split(a.Agent, "@")[1]
}

var agentCallState = map[QueueToAgentAction]string{
	AgentOfferingAction:    "Ringing",
	BridgeAgentStartAction: "In Progress",
	BridgeAgentEndAction:   "Ended",
}

type QueueToAgent struct {
	TenantID             string `fs:"variable_tenant_id"`
	QueueCustomID        string `fs:"variable_queue_custom_uuid"`
	QueueName            string `fs:"CC-Queue"`
	AgentName            string `fs:"CC-Agent"`
	AgentType            string `fs:"CC-Agent-Type"`
	Action               string `fs:"CC-Action"`
	AgentUUID            string `fs:"CC-Agent-UUID"`
	AgentCalledTime      string `fs:"CC-Agent-Called-Time"`
	AgentAnsweredTime    string `fs:"CC-Agent-Answered-Time"`
	MemberJoinedTime     string `fs:"CC-Member-Joined-Time"`
	MemberUUID           string `fs:"CC-Member-UUID"`
	MemberSessionUUID    string `fs:"CC-Member-Session-UUID"`
	MemberCallerIDName   string `fs:"CC-Member-CID-Name"`
	MemberCallerIDNumber string `fs:"CC-Member-CID-Number"`
	HangupCause          string `fs:"CC-Hangup-Cause"`
	AgentAbortedTime     string `fs:"CC-Agent-Aborted-Time"`
	Cause                string `fs:"CC-Cause"`
	BridgeTerminatedTime string `fs:"CC-Bridge-Terminated-Time"`
}

func (q QueueToAgent) State() string {
	var callState string
	if st, ok := agentCallState[QueueToAgentAction(q.Action)]; ok {
		callState = st
	}

	return callState
}

func (q QueueToAgent) MemberUniquePrefix() string {
	var prefix string
	if q.MemberCallerIDNumber != "" {
		prefix = strings.TrimPrefix(q.MemberCallerIDNumber, "+")
	}

	return prefix
}

func (q QueueToAgent) AgentNamePrefix() string {
	return strings.Split(q.AgentName, "@")[0]
}

func (q QueueToAgent) QueueNamePrefix() string {
	return strings.Split(q.QueueName, "@")[0]
}

func (q QueueToAgent) OrgID() string {
	if q.TenantID != "" && fsxml.UUIDFromString(q.TenantID) != nil {
		return q.TenantID
	}

	return strings.Split(q.QueueName, "@")[1]
}

type MemberToQueue struct {
	TenantID             string `fs:"variable_tenant_id"`
	QueueCustomID        string `fs:"variable_queue_custom_uuid"`
	QueueName            string `fs:"CC-Queue"`
	Action               string `fs:"CC-Action"`
	HangupCause          string `fs:"CC-Hangup-Cause"`
	Cause                string `fs:"CC-Cause"`
	CancelReason         string `fs:"CC-Cancel-Reason"`
	AgentCalledTime      string `fs:"CC-Agent-Called-Time"`
	AgentAnsweredTime    string `fs:"CC-Agent-Answered-Time"`
	AgentJoinedTime      string `fs:"CC-Member-Joined-Time"`
	MemberJoinedTime     string `fs:"CC-Member-Joined-Time"`
	MemberLeavingTime    string `fs:"CC-Member-Leaving-Time"`
	MemberUUID           string `fs:"CC-Member-UUID"`
	MemberSessionUUID    string `fs:"CC-Member-Session-UUID"`
	MemberCallerIDName   string `fs:"CC-Member-CID-Name"`
	MemberCallerIDNumber string `fs:"CC-Member-CID-Number"`
}

func (m MemberToQueue) QueueNamePrefix() string {
	if m.nameIsValid() {
		return strings.Split(m.QueueName, "@")[0]
	}

	return ""
}

func (m MemberToQueue) OrgID() string {
	if m.TenantID != "" && fsxml.UUIDFromString(m.TenantID) != nil {
		return m.TenantID
	}

	if m.nameIsValid() {
		return strings.Split(m.QueueName, "@")[1]
	}

	return ""
}

func (m MemberToQueue) nameIsValid() bool {
	return strings.Contains(m.QueueName, "@")
}

type AgentStatusUpdate struct {
	AgentName string `fs:"CC-Agent"`
	Status    string `fs:"CC-Agent-Status"`
	Action    string `fs:"CC-Action"`
}

func (a AgentStatusUpdate) NamePrefix() string {
	return strings.Split(a.AgentName, "@")[0]
}

func (a AgentStatusUpdate) OrgID() string {
	return strings.Split(a.AgentName, "@")[1]
}

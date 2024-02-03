package event

import (
	"sort"
	"strconv"
)

type CallerProfile struct {
	UserName          string `json:"username"`
	DialPlan          string `json:"dialplan"`
	CallerIDName      string `json:"caller_id_name"`
	CallerIDNumber    string `json:"caller_id_number"`
	ANI               string `json:"ani"`
	NetworkAddr       string `json:"network_addr"`
	RDNIS             string `json:"rdnis"`
	DestinationNumber string `json:"destination_number"`
	UUID              string `json:"uuid"`
	Source            string `json:"source"`
	Context           string `json:"context"`
	ChannelName       string `json:"chan_name"`
}

type Variables struct {
	UUID                      string `json:"uuid"`
	EndpointDisposition       string `json:"endpoint_disposition"`
	CurrentApplication        string `json:"current_application"`
	CurrentApplicationData    string `json:"current_application_data"`
	CCSide                    string `json:"cc_side"`
	CCMemberUUID              string `json:"cc_member_uuid"`
	CCMemberSessionUUID       string `json:"cc_member_session_uuid"`
	CCQueueJoinedEpoch        string `json:"cc_queue_joined_epoch"`
	CCQueue                   string `json:"cc_queue"`
	CCQueueCanceledEpoch      string `json:"cc_queue_canceled_epoch"`
	CCCause                   string `json:"cc_cause"`
	CCCancelReason            string `json:"cc_cancel_reason"`
	CCAgentBridged            string `json:"cc_agent_bridged"`
	CCAgentType               string `json:"cc_agent_type"`
	CCAgent                   string `json:"cc_agent"`
	HangupCause               string `json:"hangup_cause"`
	HangupCauseQ850           string `json:"hangup_cause_q850"`
	Duration                  string `json:"duration"`
	BillSec                   string `json:"billsec"`
	DigitsDialed              string `json:"digits_dialed"`
	LastHoldEpoch             string `json:"last_hold_epoch"`
	HoldAccumulationseconds   string `json:"hold_accum_seconds"`
	SipUserAgent              string `json:"sip_user_agent"`
	WaitSecond                string `json:"waitsec"`
	AnswerSecond              string `json:"answersec"`
	ProgressSecond            string `json:"progresssec"`
	OriginalDestinationNumber string `json:"original_destination_number"`
	OriginalCallerIDNumber    string `json:"original_caller_id_number"`
	OriginalCallerIDName      string `json:"original_caller_id_name"`
	LastSentCalleeIDNumber    string `json:"last_sent_callee_id_number"`
	LastSentCalleeIDName      string `json:"last_sent_callee_id_name"`
}

type Timing struct {
	CreatedTime        string `json:"created_time"`
	ProfileCreatedTime string `json:"profile_created_time"`
	ProgressTime       string `json:"progress_time"`
	ProgressMediaTime  string `json:"progress_media_time"`
	AnsweredTme        string `json:"answered_time"`
	BridgedTime        string `json:"bridged_time"`
	LastHoldTime       string `json:"last_hold_time"`
	LastHoldAccumTime  string `json:"hold_accum_time"`
	HangupTime         string `json:"hangup_time"`
	RessurectTime      string `json:"resurrect_time"`
	TransferTime       string `json:"transfer_time"`
}

type CallFlow struct {
	ProfileIndex  string                 `json:"profile_index"`
	Dialplan      string                 `json:"dialplan"`
	CurrentApp    string                 `json:"current_app"`
	Extension     map[string]interface{} `json:"extension"`
	CallerProfile CallerProfile          `json:"caller_profile"`
	Times         Timing                 `json:"times"`
}

type CDR struct {
	Variables Variables              `json:"variables"`
	CallFlow  []CallFlow             `json:"callflow"`
	CallStats map[string]interface{} `json:"callStats"`
}

func (c CDR) Source() string {
	var src string
	if c.CallFlow != nil {
		// grab the callerid from the first caller profile
		src = c.sortedCallFlow()[0].CallerProfile.CallerIDNumber
	}

	return src
}

func (c CDR) Destination() string {
	var dst string
	if c.CallFlow != nil {
		// grab the dialed number from the first caller profile
		dst = c.sortedCallFlow()[len(c.CallFlow)-1].CallerProfile.DestinationNumber
	}

	return dst
}

func (c CDR) StartTime() int {
	var ts int
	if c.CallFlow != nil {
		tstring := c.sortedCallFlow()[0].Times.CreatedTime
		if tsi, err := strconv.Atoi(tstring); err == nil {
			ts = tsi / 1000000
		}
	}

	return ts
}

func (c CDR) EndTime() int {
	var et int
	if c.CallFlow != nil {
		estring := c.sortedCallFlow()[len(c.CallFlow)-1].Times.HangupTime
		if eti, err := strconv.Atoi(estring); err == nil {
			et = eti / 1000000
		}
	}

	return et
}

func (c CDR) sortedCallFlow() []CallFlow {
	cf := c.CallFlow

	sort.Slice(cf, func(i, j int) bool {
		iProfileIndex, _ := strconv.Atoi(cf[i].ProfileIndex)
		jProfileIndex, _ := strconv.Atoi(cf[j].ProfileIndex)
		return iProfileIndex < jProfileIndex
	},
	)

	return cf
}

func (c CDR) CallStatus() string {
	if c.Variables.EndpointDisposition == "" {
		return c.Variables.HangupCause
	}

	return c.Variables.EndpointDisposition
}

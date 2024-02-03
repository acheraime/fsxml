package event

import "strings"

var AgentStatusString = map[string]string{
	"log_out":             "Logged Out",
	"available":           "Available",
	"available_on_demand": "Available (On Demand)",
	"on_break":            "On Break",
}

var AgentStateString = map[string]string{
	"idle":            "Idle",
	"waiting":         "Waiting",
	"receiving":       "Receiving",
	"in_a_queue_call": "In a queue call",
}

var AgentTypeString = map[string]string{
	"callback": "Callback",
}

type KQueue struct {
	ID                            string `json:"id"`
	Name                          string `json:"name"`
	SystemName                    string `json:"system_name"`
	Strategy                      string `json:"strategy"`
	MohID                         string `json:"moh_id"`
	TimeBaseScore                 string `json:"time_base_score"`
	TierRulesApply                bool   `json:"tier_rules_apply"`
	TierRuleWaitSecond            int64  `json:"tier_rule_wait_second"`
	TierRuleWaitMultiplyLevel     bool   `json:"tier_rule_wait_multiply_level"`
	TierRuleNoAgentNoWait         bool   `json:"tier_rule_no_agent_no_wait"`
	DiscardAbandonAfter           int64  `json:"discard_abandon_after"`
	AbandonResumeAllowed          bool   `json:"abandon_resume_allowed"`
	MaxWaitTime                   int64  `json:"max_wait_time"`
	MaxWaitTimeNoAgent            int64  `json:"max_wait_time_no_agent"`
	MaxWaitTimeNoAgentTimeReached int64  `json:"max_wait_time_no_agent_time_reached"`
	RingProgessivelyDelay         int64  `json:"ring_progessively_delay"`
	AutoRecord                    bool   `json:"auto_record"`
	Extension                     string `json:"extension"`
	Domain                        string `json:"domain"`
	DisplayName                   string `json:"display_name"`
	ManagerEmail                  string `json:"manager_email"`
	OrganizationID                string `json:"organization_id"`
	CampaignID                    string `json:"campaign_id"`
	CreatedAt                     string `json:"created_at"`
	UpdatedAt                     string `json:"updated_at"`
	Action                        string `json:"action"`
	EventName                     string `json:"Event-Name"`
	EventSubclass                 string `json:"Event-Subclass"`
}

func (q KQueue) SystemNamePrefix() string {
	return strings.Split(q.SystemName, "@")[0]
}

type KAgent struct {
	Status         string `json:"status"`
	AgentType      string `json:"agent_type"`
	WrapupTime     int64  `json:"wrapup_time"`
	RejectDelay    int64  `json:"reject_delay"`
	BusyDelay      int64  `json:"busy_delay"`
	MaxNoAnswer    int64  `json:"max_no_answer"`
	NoAnswerDelay  int64  `json:"no_answer_delay"`
	ReadyTime      int64  `json:"ready_time"`
	UserID         string `json:"user_id"`
	OrganizationID string `json:"organization_id"`
	ID             string `json:"id"`
	SystemName     string `json:"system_name"`
	DisplayName    string `json:"display_name"`
	SofiaContact   string `json:"sofia_contact"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
	Action         string `json:"action"`
	EventName      string `json:"Event-Name"`
	EventSubclass  string `json:"Event-Subclass"`
}

func (a KAgent) SystemNamePrefix() string {
	return strings.Split(a.SystemName, "@")[0]
}

type KTier struct {
	Action          string `json:"action"`
	AgentSystemName string `json:"agent_system_name"`
	QueueSystemName string `json:"queue_system_name"`
	Position        int64  `json:"position"`
	Level           int64  `json:"level"`
	Status          string `json:"status"`
	EventName       string `json:"Event-Name"`
	EventSubclass   string `json:"Event-Subclass"`
}

func (t KTier) OrganizationID() string {
	return strings.Split(t.QueueSystemName, "@")[1]
}

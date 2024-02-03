package fsxml

import "encoding/xml"

type QueueParam string

const (
	Strategy                          QueueParam = "strategy"
	MohSound                          QueueParam = "moh-sound"
	TimeBaseScore                     QueueParam = "time-base-score"
	TierRuleApply                     QueueParam = "tier-rules-apply"
	TierRuleWaitSecond                QueueParam = "tier-rule-wait-second"
	TierRuleWaitMultiplyLevel         QueueParam = "tier-rule-wait-multiply-level"
	TierRuleNoAgentNoWait             QueueParam = "tier-rule-no-agent-no-wait"
	DiscardAbandonedAfter             QueueParam = "discard-abandoned-after"
	MaxWaitTime                       QueueParam = "max-wait-time"
	MaxWaitTimeWithNoAgent            QueueParam = "max-wait-time-with-no-agent"
	MaxWaitTimeWithNoAgentTimeReached QueueParam = "max-wait-time-with-no-agent-time-reached"
	RecordTemplate                    QueueParam = "record-template"
	RingProgessivelyDelay             QueueParam = "ring-progressively-delay"
	AdandonedResumeAllowed            QueueParam = "abandoned-resume-allowed"
)

type DistributionStrategy string

const (
	RingAll                  DistributionStrategy = "ring-all"
	LongestIdleAgent         DistributionStrategy = "longest-idle-agent"
	RoundRobin               DistributionStrategy = "round-robin"
	TopDown                  DistributionStrategy = "top-down"
	AgentWithLeastTalkTime   DistributionStrategy = "agent-with-least-talk-time"
	AgentWithFewestCalls     DistributionStrategy = "agent-with-fewest-calls"
	SequentiallyByAgentOrder DistributionStrategy = "sequentially-by-agent-order"
	Random                   DistributionStrategy = "random"
	RingProgessively         DistributionStrategy = "ring-progressively"
)

type TimeBaseScoreType string

const (
	QueueScore  TimeBaseScoreType = "queue"
	SystemScore TimeBaseScoreType = "system"
)

type Agent struct {
	XMLName         xml.Name `xml:"agent"`
	Name            string   `xml:"name,attr"`
	Type            string   `xml:"type,attr"`
	Contact         string   `xml:"contact,attr"`
	Status          string   `xml:"status,attr"`
	MaxNoAnswer     string   `xml:"max-no-answer,attr"`
	WrapUpTime      string   `xml:"wrap-up-time,attr"`
	RejectDelayTime string   `xml:"reject-delay-time,attr"`
	BusyDelayTime   string   `xml:"busy-delay-time,attr"`
}

type Tier struct {
	XMLName  xml.Name `xml:"tier"`
	Agent    string   `xml:"agent,attr"`
	Queue    string   `xml:"queue,attr"`
	Level    string   `xml:"level,attr"`
	Position string   `xml:"position,attr"`
}

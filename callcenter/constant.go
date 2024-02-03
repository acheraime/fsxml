package callcenter

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

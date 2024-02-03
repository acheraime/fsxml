package callcenter

import (
	"encoding/xml"
	"strconv"

	"github.com/acheraime/fsxml"
)

type Queue struct {
	XMLName xml.Name     `xml:"queue"`
	Name    string       `xml:"name,attr"`
	Params  fsxml.Params `xml:"param"`
}

func (q *Queue) addParam(name QueueParam, value string) {
	pname := string(name)
	param := fsxml.Param{
		Name:  pname,
		Value: value,
	}
	q.Params.Add(param)

}

func (q *Queue) SetStrategy(strategy DistributionStrategy) {
	q.addParam(Strategy, string(strategy))
}

func (q *Queue) SetMohSound(moh string) {
	q.addParam(MohSound, moh)
}

func (q *Queue) SetTimeBaseScore(score TimeBaseScoreType) {
	q.addParam(TimeBaseScore, string(score))
}

func (q *Queue) SetTierRuleApply(flag bool) {
	q.addParam(TierRuleApply, strconv.FormatBool(flag))
}

func (q *Queue) SetTierRuleWaitSecond(wait int) {
	q.addParam(TierRuleWaitSecond, strconv.Itoa(wait))
}

func (q *Queue) SetTierRuleWaitMultiplyLevel(flag bool) {
	q.addParam(TierRuleWaitMultiplyLevel, strconv.FormatBool(flag))
}

func (q *Queue) SetTierRuleNoAgentNoWait(flag bool) {
	q.addParam(TierRuleNoAgentNoWait, strconv.FormatBool(flag))
}

func (q *Queue) SetMaxWaitTimeWithNoAgentTimeReached(wait int) {
	q.addParam(MaxWaitTimeWithNoAgentTimeReached, strconv.Itoa(wait))
}

func (q *Queue) SetDiscardAbandonedAfter(duration int) {
	q.addParam(DiscardAbandonedAfter, strconv.Itoa(duration))
}

func (q *Queue) SetMaxWaitTime(wait int) {
	q.addParam(MaxWaitTime, strconv.Itoa(wait))
}

func (q *Queue) SetMaxWaitTimeWithNoAgent(wait int) {
	q.addParam(MaxWaitTimeWithNoAgentTimeReached, strconv.Itoa(wait))
}

func (q *Queue) SetRecordTemplate(template string) {
	q.addParam(RecordTemplate, template)
}

func (q *Queue) SetRingProgressivelyDelay(delay int) {
	q.addParam(RingProgessivelyDelay, strconv.Itoa(delay))
}

func (q *Queue) SetAbandonedResumeAllowed(flag bool) {
	q.addParam(AdandonedResumeAllowed, strconv.FormatBool(flag))
}

func NewQueue(name string) Queue {
	q := Queue{Name: name}
	// Setting default queue parameters
	q.SetMohSound("$${hold_music}")
	q.SetStrategy(RingAll)
	q.SetTimeBaseScore(SystemScore)
	q.SetTierRuleApply(false)
	q.SetTierRuleWaitSecond(300)
	q.SetTierRuleWaitMultiplyLevel(false)
	q.SetTierRuleNoAgentNoWait(false)
	q.SetDiscardAbandonedAfter(14400)

	return q
}

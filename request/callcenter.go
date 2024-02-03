package request

import "strings"

type CallCenter struct {
	QueueName string `fs:"CC-Queue"`
}

func (c CallCenter) TenantID() string {
	if strings.Contains(c.QueueName, "@") {
		return strings.Split(c.QueueName, "@")[1]
	}

	return ""
}

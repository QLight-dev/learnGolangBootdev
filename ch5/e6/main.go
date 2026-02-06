package main

type notification interface {
	importance() int
}

type directMessage struct {
	senderUsername string
	messageContent string
	priorityLevel  int
	isUrgent       bool
}

type groupMessage struct {
	groupName      string
	messageContent string
	priorityLevel  int
}

type systemAlert struct {
	alertCode      string
	messageContent string
}

func (msg directMessage) importance() int {
	if msg.isUrgent {
		return 50
	} else {
		return msg.priorityLevel
	}
}

func (msg groupMessage) importance() int {
	return msg.priorityLevel
}

func (msg systemAlert) importance() int {
	return 100
}

func processNotification(n notification) (string, int) {
	dm, dm_ok := n.(directMessage)
	gm, gm_ok := n.(groupMessage)
	sa, sa_ok := n.(systemAlert)

	if dm_ok {
		return dm.senderUsername, n.importance()
	} else if gm_ok {
		return gm.groupName, n.importance()
	} else if sa_ok {
		return sa.alertCode, n.importance()
	} else {
		return "", 0
	}
}

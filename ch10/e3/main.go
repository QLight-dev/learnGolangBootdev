package main

type Analytics struct {
	MessagesTotal     int
	MessagesFailed    int
	MessagesSucceeded int
}

type Message struct {
	Recipient string
	Success   bool
}

// don't touch above this line

func analyzeMessage(antPtr *Analytics, msg Message) {
	antPtr.MessagesTotal++
	if !msg.Success {
		antPtr.MessagesFailed++
	} else {
		antPtr.MessagesSucceeded++
	}
}

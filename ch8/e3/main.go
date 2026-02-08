package main

func getMessageCosts(messages []string) []float64 {
	messageCosts := make([]float64, len(messages))
	for i := 0; i < len(messages); i++ {
		// cost = length of message * 0.01
		messageCosts[i] = float64(len(messages[i])) * 0.01
	}
	return messageCosts
}

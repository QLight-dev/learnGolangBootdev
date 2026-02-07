package main

func bulkSend(numMessages int) float64 {
	totalMessageCost := float64(numMessages)
	additionalFee := 0.00
	for i := 0; i < numMessages; i++ {
		totalMessageCost += additionalFee
		additionalFee += 0.01
	}
	return totalMessageCost
}

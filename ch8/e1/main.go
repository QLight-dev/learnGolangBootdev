package main

func getMessageWithRetries(primary, secondary, tertiary string) ([3]string, [3]int) {
	messages := [3]string{primary,secondary, tertiary}
	priceOfFirstMsg := len(primary)
	priceOfSecondMsg := len(primary) + len(secondary)
	priceOfThirdMsg := len(primary) + len(secondary) + len(tertiary)
	return messages, [3]int{priceOfFirstMsg, priceOfSecondMsg, priceOfThirdMsg}
}

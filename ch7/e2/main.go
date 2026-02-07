package main

func maxMessages(thresh int) int {
	fee := 0
	i := 0
	for ; ; i++ {
		if fee+100+i > thresh {
			break
		} else {
			fee += 100 + i
		}
	}
	return i
}

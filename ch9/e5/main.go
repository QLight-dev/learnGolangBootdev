package main

import "strings"

func countDistinctWords(messages []string) int {
	var distinctWordsCount int
	distinctWords := []string{}
	for _, msg := range messages {
		for _, word := range strings.Fields(strings.ToLower(msg)) {
			isDistinct := true
			for _, distinctWord := range distinctWords {
				if word == distinctWord {
					isDistinct = false
					break
				}
			}

			if isDistinct {
				distinctWords = append(distinctWords, word)
				distinctWordsCount++
			}
		}
	}
	return distinctWordsCount
}

package main

func indexOfFirstBadWord(msg []string, badWords []string) int {
	for indexOfWord, word := range msg {
		for _, badWord := range badWords {
			if word == badWord {
				return indexOfWord
			}
		}
	}
	return -1
}

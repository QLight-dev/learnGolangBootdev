package main

func getNameCounts(names []string) map[rune]map[string]int {
	resultMap := map[rune]map[string]int{}
	for _, name := range names {
		// get first rune 
		// avoids getting a chunk of a letter
		var firstRune rune
		for _, v := range name {
			firstRune = v
			break
		}
		// if unique first char does not exist
		if _, ok := resultMap[rune(name[0])]; !ok {
			resultMap[firstRune] = map[string]int{}
		}

		// if name does not exist
		if _, ok := resultMap[rune(name[0])][name]; !ok {
			resultMap[firstRune][name] = 1
			continue
		}

		resultMap[rune(name[0])][name]++
	}
	return resultMap
}

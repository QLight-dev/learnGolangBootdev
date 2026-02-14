package main

func updateCounts(messagedUsers []string, validUsers map[string]int) {
	for _, v := range messagedUsers {
		if _, ok := validUsers[v]; ok {
			validUsers[v] += 1
		}
	}
}

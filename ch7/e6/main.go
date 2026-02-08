package main

func countConnections(groupSize int) int {
	numOfConnections := 0
	for i := groupSize - 1; i > 0; i-- {
		numOfConnections += i
	}
	return numOfConnections
}

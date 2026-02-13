package main

type cost struct {
	day   int
	value float64
}

func getDayCosts(costs []cost, day int) []float64 {
	if len(costs) == 0 {
		return []float64{}
	}
	costsOfDay := []float64{}
	for i := 0; i < len(costs); i++ {
		if costs[i].day == day {
			costsOfDay = append(costsOfDay, costs[i].value)
		}
	}
	return costsOfDay
}

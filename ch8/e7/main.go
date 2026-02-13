package main

func createMatrix(rows, cols int) [][]int {
	matrix := [][]int{}
	for row := 0; row < rows; row++ {
		rowContent := []int{}
		for col := 0; col < cols; col++ {
			rowContent = append(rowContent, row*col)
		}
		matrix = append(matrix, rowContent)
	}
	return matrix
}

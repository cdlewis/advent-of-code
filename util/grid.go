package util

func ValidCoordinate[U any](i int, j int, grid [][]U) bool {
	return i >= 0 && j >= 0 && i < len(grid) && j < len(grid[0])
}

var Directions = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

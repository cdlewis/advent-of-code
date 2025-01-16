package four

import (
	"fmt"
	"strings"

	"github.com/cdlewis/advent-of-code/2024/util/aoc"
)

var testData = `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

func Four() int {
	wordsearch := strings.Split(aoc.GetInput(4, false, testData), "\n")

	count := 0
	for i := 1; i < len(wordsearch)-1; i++ {
		for j := 1; j < len(wordsearch[i])-1; j++ {
			if wordsearch[i][j] != 'A' {
				continue
			}

			topLeft := wordsearch[i-1][j-1]
			bottomRight := wordsearch[i+1][j+1]

			topRight := wordsearch[i-1][j+1]
			bottomLeft := wordsearch[i+1][j-1]

			if validDiagonal(topLeft, bottomRight) && validDiagonal(topRight, bottomLeft) {
				fmt.Println(i, j)
				count++
			}
		}
	}

	return count
}

func validDiagonal(x, y byte) bool {
	return (x == 'M' || x == 'S') && (y == 'M' || y == 'S') && x != y
}

package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/cdlewis/advent-of-code/util"
)

func main() {
	dat, _ := os.ReadFile("./input")
	raw := strings.Split(string(dat), "\n")
	util.Flatten([][]int{})

	fmt.Println(raw)
}

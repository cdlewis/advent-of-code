package five

import (
	"os"
	"strings"

	"github.com/cdlewis/advent-of-code/2022/util"
)

func Five() string {
	dat, _ := os.ReadFile("./input")
	instructions := strings.Split(string(dat), "\n")
	stacks := util.Map([]string{
		"ZJG",
		"QLRPWFVC",
		"FPMCLGR",
		"LFBWPHM",
		"GCFSVQ",
		"WHJZMQTL",
		"HFSBV",
		"FJZS",
		"MCDPFHBT",
	}, func(i string) []byte { return []byte(i) })

	for _, serializedInstruction := range instructions {
		tokens := strings.Split(serializedInstruction, " ")
		count := util.ToInt(tokens[1])
		from := util.ToInt(tokens[3]) - 1
		to := util.ToInt(tokens[5]) - 1

		curr := util.Pops(&stacks[from], count)
		stacks[to] = append(stacks[to], curr...)
	}

	return string(util.Map(stacks, func(s []byte) byte { return s[len(s)-1] }))
}

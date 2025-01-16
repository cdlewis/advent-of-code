package thirteen

import (
	"encoding/json"
	"sort"
	"strings"

	"github.com/cdlewis/advent-of-code/2022/util"
)

type Item struct {
	Value    int
	IsList   bool
	Contents []Item
}

func Thirteen(isTest bool, testData string) int {
	raw := util.Filter(strings.Split(util.GetInput(13, isTest, testData), "\n"), func(s string) bool { return len(s) > 0 })

	messages := util.Map(raw, func(s string) Item {
		var unknownJson any
		json.Unmarshal([]byte(s), &unknownJson)
		return newItem(unknownJson)
	})
	messages = append(messages, Item{Value: 2}, Item{Value: 6})

	sort.Slice(messages, func(i, j int) bool {
		return compareItem(messages[i], messages[j]) > 0
	})

	packet1, packet2 := 0, 0
	for idx, i := range messages {
		if i.Value == 2 {
			packet1 = idx + 1
		}

		if i.Value == 6 {
			packet2 = idx + 1
		}
	}

	return packet1 * packet2
}

func compareItem(x, y Item) int {
	if !x.IsList && !y.IsList {
		return compareNumeric(x.Value, y.Value)
	}

	if x.IsList != y.IsList {
		if !x.IsList {
			x = Item{IsList: true, Contents: []Item{x}}
		}

		if !y.IsList {
			y = Item{IsList: true, Contents: []Item{y}}
		}
	}

	for i := 0; i < len(x.Contents) && i < len(y.Contents); i++ {
		if less := compareItem(x.Contents[i], y.Contents[i]); less != 0 {
			return less
		}
	}

	return compareNumeric(len(x.Contents), len(y.Contents))
}

func compareNumeric(x, y int) int {
	if x < y {
		return 1
	} else if x == y {
		return 0
	} else {
		return -1
	}
}

func newItem(b any) Item {
	switch b.(type) {
	case float64:
		return Item{Value: int(b.(float64))}
	default:
		return Item{
			IsList:   true,
			Contents: util.Map(b.([]interface{}), newItem),
		}
	}
}

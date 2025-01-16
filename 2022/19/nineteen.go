package nineteen

import (
	"strings"

	"github.com/cdlewis/advent-of-code/2022/util"
)

func Nineteen() int {
	blueprints := parseBlueprints()

	score := 1

	for idx, b := range blueprints {
		if idx >= 3 {
			break
		}

		initialPerm := Perm{
			robots:    Robots{OreRobot: 1},
			materials: Materials{},
		}

		q := []Perm{initialPerm}
		steps := 0
		maxSteps := 33
		result := 0

		for len(q) > 0 && steps < maxSteps {
			seen := map[[8]int]bool{}
			newQ := []Perm{}

			for len(q) > 0 {
				curr := q[0]
				q = q[1:]

				if _, ok := seen[curr.serialize()]; ok {
					continue
				}
				seen[curr.serialize()] = true

				// Note: robot construction shouldn't take account of new materials
				newMaterials := curr.materials.Add(Materials{
					Ore:      curr.robots.OreRobot,
					Clay:     curr.robots.ClayRobot,
					Obsidian: curr.robots.ObsidianRobot,
					Geode:    curr.robots.GeodeRobot,
				})

				if curr.materials.Geode > result {
					result = curr.materials.Geode
				}

				if curr.materials.Contains(b.GeodeRobot) {
					newQ = append(newQ, Perm{
						robots:    curr.robots.Add(Robots{GeodeRobot: 1}),
						materials: newMaterials.Subtract(b.GeodeRobot),
					})
				}

				if curr.materials.Contains(b.OreRobot) {
					newQ = append(newQ, Perm{
						robots:    curr.robots.Add(Robots{OreRobot: 1}),
						materials: newMaterials.Subtract(b.OreRobot),
					})
				}

				if curr.materials.Contains(b.ClayRobot) {
					// fmt.Println("Make clay robot")
					newQ = append(newQ, Perm{
						robots:    curr.robots.Add(Robots{ClayRobot: 1}),
						materials: newMaterials.Subtract(b.ClayRobot),
					})
				}

				if curr.materials.Contains(b.ObsidianRobot) {
					newQ = append(newQ, Perm{
						robots:    curr.robots.Add(Robots{ObsidianRobot: 1}),
						materials: newMaterials.Subtract(b.ObsidianRobot),
					})
				}

				newQ = append(newQ, Perm{
					robots:    curr.robots,
					materials: newMaterials,
				})
			}

			q = newQ

			steps++
		}

		score *= result
	}

	return score
}

func parseBlueprints() []Blueprint {
	raw := util.GetInput(19, false, `Blueprint 1: Each ore robot costs 4 ore. Each clay robot costs 2 ore. Each obsidian robot costs 3 ore and 14 clay. Each geode robot costs 2 ore and 7 obsidian.
Blueprint 2: Each ore robot costs 2 ore. Each clay robot costs 3 ore. Each obsidian robot costs 3 ore and 8 clay. Each geode robot costs 3 ore and 12 obsidian.`)
	blueprints := []Blueprint{}

	for _, row := range strings.Split(raw, "\n") {
		b := Blueprint{}

		for jdx, section := range strings.Split(row, ".") {
			m := Materials{}

			tokens := strings.Split(section, " ")
			for i := 1; i < len(tokens); i++ {
				first := tokens[i-1]
				second := tokens[i]

				if first == "Each" {
					continue
				}

				if string(second) == "ore" {
					m.Ore = util.ToInt(first)
				}
				if string(second) == "clay" {
					m.Clay = util.ToInt(first)
				}
				if string(second) == "obsidian" {
					m.Obsidian = util.ToInt(first)
				}
			}

			if jdx == 0 {
				b.OreRobot = m
			}
			if jdx == 1 {
				b.ClayRobot = m
			}
			if jdx == 2 {
				b.ObsidianRobot = m
			}
			if jdx == 3 {
				b.GeodeRobot = m
			}
		}

		blueprints = append(blueprints, b)
	}

	return blueprints
}

type Perm struct {
	robots    Robots
	materials Materials
}

// hack to constrain the search space
var clamp = 16

func (p Perm) serialize() [8]int {
	return [8]int{
		util.Min(p.robots.ClayRobot, clamp),
		util.Min(p.robots.ObsidianRobot, clamp),
		util.Min(p.robots.OreRobot, clamp),
		p.robots.GeodeRobot,
		util.Min(p.materials.Clay, clamp),
		util.Min(p.materials.Obsidian, clamp),
		util.Min(p.materials.Ore, clamp),
		p.materials.Geode,
	}
}

type Blueprint struct {
	OreRobot      Materials
	ClayRobot     Materials
	ObsidianRobot Materials
	GeodeRobot    Materials
}

type Materials struct {
	Ore      int
	Clay     int
	Obsidian int
	Geode    int
}

func (m Materials) Contains(x Materials) bool {
	return m.Ore >= x.Ore && m.Clay >= x.Clay && m.Obsidian >= x.Obsidian
}

func (x Materials) Add(y Materials) Materials {
	return Materials{
		Ore:      x.Ore + y.Ore,
		Clay:     x.Clay + y.Clay,
		Obsidian: x.Obsidian + y.Obsidian,
		Geode:    x.Geode + y.Geode,
	}
}

func (x Materials) Subtract(y Materials) Materials {
	return Materials{
		Ore:      x.Ore - y.Ore,
		Clay:     x.Clay - y.Clay,
		Obsidian: x.Obsidian - y.Obsidian,
		Geode:    x.Geode - y.Geode,
	}
}

type Robots struct {
	OreRobot      int
	ClayRobot     int
	ObsidianRobot int
	GeodeRobot    int
}

func (r Robots) Add(x Robots) Robots {
	return Robots{
		OreRobot:      r.OreRobot + x.OreRobot,
		ClayRobot:     r.ClayRobot + x.ClayRobot,
		ObsidianRobot: r.ObsidianRobot + x.ObsidianRobot,
		GeodeRobot:    r.GeodeRobot + x.GeodeRobot,
	}
}

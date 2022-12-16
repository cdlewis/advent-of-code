package main

import (
	"fmt"
	"strings"

	"github.com/cdlewis/advent-of-code/util"
)

type Actor struct {
	CurrentVertex    string
	RemainingMinutes int
	Score            int
}

func main() {
	raw := strings.Split(util.GetInput(16, false, `Valve AA has flow rate=0; tunnels lead to valves DD, II, BB
	Valve BB has flow rate=13; tunnels lead to valves CC, AA
	Valve CC has flow rate=2; tunnels lead to valves DD, BB
	Valve DD has flow rate=20; tunnels lead to valves CC, AA, EE
	Valve EE has flow rate=3; tunnels lead to valves FF, DD
	Valve FF has flow rate=0; tunnels lead to valves EE, GG
	Valve GG has flow rate=0; tunnels lead to valves FF, HH
	Valve HH has flow rate=22; tunnel leads to valve GG
	Valve II has flow rate=0; tunnels lead to valves AA, JJ
	Valve JJ has flow rate=21; tunnel leads to valve II`), "\n")

	graph := map[string][]string{}
	flowRate := map[string]int{}

	for _, i := range raw {
		tokens := strings.Split(i, " ")

		id := tokens[1]
		f := util.ToInt(tokens[4][5 : len(tokens[4])-1])
		tunnels := util.Map(tokens[9:], func(s string) string {
			return strings.ReplaceAll(s, ",", "")
		})

		graph[id] = tunnels
		flowRate[id] = f
	}

	dist := allVertexShortestPaths(graph)

	var findOptimalScore func(human Actor, elephant Actor, currentScore int, nodeSet map[string]bool) int
	findOptimalScore = func(human Actor, elephant Actor, currentScore int, nodeSet map[string]bool) int {

		bestScore := currentScore

		for i, useI := range nodeSet {
			if !useI {
				continue
			}

			newHuman := Actor{
				RemainingMinutes: human.RemainingMinutes - dist[human.CurrentVertex][i] - 1,
				CurrentVertex:    i,
			}

			newHumanScore := newHuman.RemainingMinutes * flowRate[i]

			if newHuman.RemainingMinutes < 0 {
				continue
			}

			nodeSet[i] = false

			for j, useJ := range nodeSet {
				if !useJ {
					continue
				}

				newElephant := Actor{
					RemainingMinutes: elephant.RemainingMinutes - dist[elephant.CurrentVertex][j] - 1,
					CurrentVertex:    j,
				}

				newElephantScore := newElephant.RemainingMinutes * flowRate[j]

				if newElephant.RemainingMinutes < 0 {
					continue
				}

				nodeSet[j] = false

				bestScore = util.Max(bestScore, findOptimalScore(newHuman, newElephant, currentScore+newHumanScore+newElephantScore, nodeSet))

				nodeSet[j] = true
			}

			nodeSet[i] = true
		}

		return bestScore
	}

	// We can safely exclude nodes with 0 flow rate
	nodeSet := map[string]bool{}
	for k := range graph {
		if flowRate[k] > 0 {
			nodeSet[k] = true
		}
	}

	score := findOptimalScore(
		Actor{
			CurrentVertex:    "AA",
			RemainingMinutes: 26,
		},
		Actor{
			CurrentVertex:    "AA",
			RemainingMinutes: 26,
		},
		0,
		nodeSet)

	fmt.Println("FOUND", score)
}

// Calculate the shortest paths between all vertices in the graph
func allVertexShortestPaths(graph map[string][]string) map[string]map[string]int {
	dist := map[string]map[string]int{}

	for key := range graph {
		q := []string{key}
		seen := map[string]bool{}
		distance := 0

		for len(q) > 0 {
			newKeys := []string{}

			for len(q) > 0 {
				curr := q[0]
				q = q[1:]

				if seen[curr] {
					continue
				}

				seen[curr] = true

				if dist[key] == nil {
					dist[key] = map[string]int{}
				}

				dist[key][curr] = distance

				newKeys = append(newKeys, graph[curr]...)
			}

			q = newKeys
			distance++
		}
	}
	return dist
}

// 2536

package main

import "fmt"

type Graph map[string]map[string]bool

func main() {
	flights := [][]string{{"SFO", "EWR"}}
	//SFO->EWR
	flights = [][]string{{"ATL", "EWR"}, {"SFO", "ATL"}}
	//SFO->ATL->EWR
	flights = [][]string{{"IND", "EWR"}, {"SFO", "ATL"}, {"GSO", "IND"}, {"ATL", "GSO"}}
	//SFO->ATL->GSO->IND->EWR
	flights = [][]string{{"IND", "EWR"}, {"SFO", "ATL"}, {"GSO", "IND"}, {"ATL", "GSO"}, {"ATL", "EWR"}}
	//SFO->ATL->EWR
	//SFO->ATL->GSO->IND->EWR
	flights = [][]string{{"IND", "EWR"}, {"SFO", "ATL"}, {"GSO", "IND"}, {"GSO", "EWR"}, {"ATL", "GSO"}, {"ATL", "EWR"}}
	//SFO->ATL->EWR
	//SFO->ATL->GSO->EWR
	//SFO->ATL->GSO->IND->EWR
	flights = [][]string{{"IND", "EWR"}, {"SFO", "ATL"}, {"GSO", "IND"}, {"GSO", "EWR"}, {"ATL", "GSO"}, {"ATL", "EWR"}, {"ATL", "SFO"}} // with circle
	//SFO->ATL->EWR
	//SFO->ATL->GSO->EWR
	//SFO->ATL->GSO->IND->EWR
	//flights = [][]string{{"IND", "EWR"}, {"SFO", "SPO"}, {"GSO", "IND"}, {"GSO", "EWR"}, {"ATL", "GSO"}, {"ATL", "EWR"}}
	// Zero

	start := "SFO"
	end := "EWR"

	graph := buildGraph(flights)
	fmt.Println(graph)
	paths := findPaths(graph, start, end)
	fmt.Println(paths)
}

func buildGraph(flights [][]string) Graph {
	graph := map[string]map[string]bool{}
	for _, flight := range flights {
		origin, dest := flight[0], flight[1]
		if _, ok := graph[origin]; !ok {
			graph[origin] = map[string]bool{}
		}
		graph[origin][dest] = false
	}
	return graph
}

func findPaths(graph Graph, start, end string) int {
	count := 0
	for edge := range graph[start] {
		if graph[start][edge] { // node already visited, avoid circle
			continue
		}
		graph[start][edge] = true
		if edge == end {
			count++
			continue
		}
		count += findPaths(graph, edge, end)
	}
	return count
}

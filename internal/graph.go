package internal

import "fmt"

type graph = map[string]map[string]bool

type GraphService struct {
	graph graph
}

func NewGraphService(flights [][]string) GraphService {
	graph := map[string]map[string]bool{}
	for _, flight := range flights {
		origin, dest := flight[0], flight[1]
		if _, ok := graph[origin]; !ok {
			graph[origin] = map[string]bool{}
		}
		graph[origin][dest] = false
	}
	return GraphService{graph: graph}
}

func (g GraphService) FindPaths(start, end string) []string {
	matchs := []string{}
	for edge := range g.graph[start] {
		if g.graph[start][edge] { // node already visited, avoid circle
			continue
		}
		g.graph[start][edge] = true
		if edge == end {
			matchs = append(matchs, fmt.Sprintf("%s->%s", start, end))
			continue
		}

		paths := g.FindPaths(edge, end)
		for _, path := range paths {
			matchs = append(matchs, fmt.Sprintf("%s->%s", start, path))
		}
	}
	return matchs
}

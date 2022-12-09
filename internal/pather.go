package internal

import "fmt"

func FindPath(flights [][]string) (string, error) {
	pather := map[string]int{}
	for _, flight := range flights {
		origin, dest := flight[0], flight[1]
		pather[origin]++
		pather[dest]--
	}

	origin := ""
	dest := ""
	for key, count := range pather {
		if count == 1 {
			origin = key
		}
		if count == -1 {
			dest = key
		}
	}
	if origin == "" || dest == "" {
		return "", fmt.Errorf("could not define path")
	}
	return fmt.Sprintf("%s->%s", origin, dest), nil
}

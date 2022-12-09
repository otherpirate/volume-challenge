package flights

import "fmt"

func FindPath(flights [][]string) (string, error) {
	pather := map[string]int{}
	for _, flight := range flights {
		origin, dest := flight[0], flight[1]
		pather[origin]++
		pather[dest]--
		if pather[origin] == 0 {
			delete(pather, origin)
		}
		if pather[dest] == 0 {
			delete(pather, dest)
		}
	}

	if len(pather) > 2 {
		return "", fmt.Errorf("could not define path")
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

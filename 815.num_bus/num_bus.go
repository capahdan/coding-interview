package numbus

func NumBusesToDestination(routes [][]int, S int, T int) int {

	n := len(routes)
	toRoutes := make(map[int]map[int]struct{})
	for i := 0; i < len(routes); i++ {
		for _, j := range routes[i] {
			if _, ok := toRoutes[j]; !ok {
				toRoutes[j] = make(map[int]struct{})
			}
			toRoutes[j][i] = struct{}{}
		}
	}
	bfs := [][]int{{S, 0}}
	seen := make(map[int]struct{})
	seen[S] = struct{}{}
	seenRoutes := make([]bool, n)
	for len(bfs) > 0 {
		stop := bfs[0][0]
		bus := bfs[0][1]
		bfs = bfs[1:]
		if stop == T {
			return bus
		}
		for i := range toRoutes[stop] {
			if seenRoutes[i] {
				continue
			}
			for _, j := range routes[i] {
				if _, ok := seen[j]; !ok {
					seen[j] = struct{}{}
					bfs = append(bfs, []int{j, bus + 1})
				}
			}
			seenRoutes[i] = true
		}
	}
	return -1
}

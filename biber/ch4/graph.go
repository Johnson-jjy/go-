package ch4

// graph将一个字符串类型的key映射到一组相关的字符串集合，它们指向新的graph的key

var graph = make(map[string]map[string]bool)

func addEdge(from, to string) {
	edges := graph[from]
	if edges == nil { // 惰性初始化map，在每个值首词作为key时才初始化
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}

func hasEdge(from, to string) bool {
	return graph[from][to]
}

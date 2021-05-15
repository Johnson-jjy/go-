package template

//UniFind defined
//1.路径压缩+秩优化
type UnionFind struct {
	parent, rank []int
	count int
}

//Init define
func (uf *UnionFind) Init(n int) {
	uf.count = n
	uf.parent = make([]int, n)
	uf.rank = make([]int, n)
	for i := 0; i < n; i++ {
		uf.parent[i] = i
	}
}

//Find define
func (uf *UnionFind) Find(p int) int {
	root := p
	for root != uf.parent[root] {
		root = uf.parent[root]
	}
	//compress path
	for p != root {
		tmp := uf.parent[p]
		uf.parent[p] = root
		p = tmp
	}
	return root
}

//Union define
func (uf *UnionFind) Union(p, q int) {
	proot := uf.Find(p)
	qroot := uf.Find(q)
	if proot == qroot {
		return
	}
	if uf.rank[qroot] > uf.rank[proot] {
		uf.parent[proot] = qroot
	} else {
		uf.parent[qroot] = proot
		if uf.rank[proot] == uf.rank[qroot] {
			uf.rank[proot]++
		}
	}
	uf.count--
}

//TotalCount define
func (uf *UnionFind) TotalCount() int {
	return uf.count
}
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

//UnionFindCount define
//计算每个集合中元素的个数 + 最大集合元素个数
type UnionFindCount struct {
	parent, count []int
	maxUnionCount int
}

//Init define
func (uf *UnionFindCount) Init (n int) {
	uf.parent = make([]int, n)
	uf.count = make([]int, n)
	for i := range uf.parent {
		uf.parent[i] = i
		uf.count[i] = 1
	}
}

//Find define
func (uf *UnionFindCount) Find (p int) int {
	root := p
	for root != uf.parent[root] {
		root = uf.parent[root]
	}
	return root
}

//不进行秩压缩，时间复杂度爆炸
//func (uf *UnionFindCount) Union(p, q int) {
//	proot := uf.Find(p)
//	qroot := uf.Find(q)
//	if proot == qroot {
//		return
//	}
//	if proot != qroot {
//		uf.parent[proot] = qroot
//		uf.count[qroot] += uf.parent[proot]
//	}
//}

// Union define
func (uf *UnionFindCount) Union(p, q int) {
	proot := uf.Find(p)
	qroot := uf.Find(q)
	if proot == qroot {
		return
	}
	if proot == len(uf.parent) - 1 {
		//proot is root
	} else if qroot == len(uf.parent) - 1 {
		//qroot is root, always attach to root
		proot, qroot = qroot, proot
	} else if uf.count[qroot] > uf.count[proot] {
		proot, qroot = qroot, proot
	}

	//set relation[0] as parent
	uf.count[proot] += uf.count[qroot]
	uf.parent[qroot] = proot
	uf.maxUnionCount = max(uf.maxUnionCount, uf.count[proot])
}

//Count define
func (uf *UnionFindCount) Count() []int {
	return uf.count
}

//MaxUnionCount define
func (uf *UnionFindCount) MaxUnionCount() int {
	return uf.maxUnionCount
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
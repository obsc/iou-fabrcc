package graph

type Graph struct {
	nodes map[Id]*Node
}

type Node struct {
	person *Person
	in     []*DiEdge
	out    []*DiEdge
}

type DiEdge struct {
	source *Node
	sink   *Node
	value  int
}

func NewGraph() *Graph {
	return &Graph{make(map[Id]*Node)}
}

func (graph *Graph) NewNode(person *Person, basesize int) {
	var in []*DiEdge = make([]*DiEdge, 0, basesize)
	var out []*DiEdge = make([]*DiEdge, 0, basesize)
	graph.nodes[1] = &Node{person, in, out}
}

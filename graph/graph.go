package graph

// Type representing a graph
type Graph struct {
	nodes map[Id]*Node
}

// Type representing a node
type Node struct {
	person *Person
	in     []*DiEdge
	out    []*DiEdge
}

// Type representing a directed edge with a cost and a reason
type DiEdge struct {
	source *Node
	sink   *Node
	value  int
	reason string
}

// Initializes a new graph
func NewGraph() *Graph {
	return &Graph{make(map[Id]*Node)}
}

// Constructs a new node from a person
func createNode(person *Person, basesize int) *Node {
	var in []*DiEdge = make([]*DiEdge, 0, basesize)
	var out []*DiEdge = make([]*DiEdge, 0, basesize)
	return &Node{person, in, out}
}

// Adds a newly constructed node to the graph from a person
func (graph *Graph) NewPerson(person *Person, basesize int) {
	graph.nodes[person.id] = createNode(person, basesize)
}

// Removes a node from the graph based on person
// Only works if node exists
func (graph *Graph) RemovePerson(person *Person) {
	_, ok := graph.nodes[person.id]
	if ok {
		delete(graph.nodes, person.id)
	}
}

func (graph *Graph) AddEdge(source *Person, sink *Person, value int, reason string) {

}

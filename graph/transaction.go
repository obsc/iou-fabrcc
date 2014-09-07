package graph

type TransId int

// Type representing a transaction
type Trans struct {
	id     TransId
	value  int
	reason string
}

var nextTrans func() int = idGenerator()

// Constructs a new transaction with next available id
func NewTrans(value int, reason string) *Trans {
	return &Trans{TransId(nextTrans()), value, reason}
}

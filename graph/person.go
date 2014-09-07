package graph

type PersonId int

// Type representing a person
type Person struct {
	id   PersonId
	Name string
}

var nextPerson func() int = idGenerator()

// Constructs a new person with the next available id
func NewPerson(name string) *Person {
	return &Person{PersonId(nextPerson()), name}
}

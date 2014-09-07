package graph

// Type of an id
type Id int

// Type representing a person
type Person struct {
	id   Id
	Name string
}

// The next available id
var curId Id = 0

// Returns the next id
func nextId() Id {
	curId += 1
	return curId - 1
}

// Resets Ids to 0
func ClearIds() {
	curId = 0
}

// Constructs a new person with the next available id
func NewPerson(name string) *Person {
	return &Person{nextId(), name}
}

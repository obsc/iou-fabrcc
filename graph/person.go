package graph

type Id int

type Person struct {
	id   Id
	name string
}

var curId Id = 0

func nextId() Id {
	curId += 1
	return curId - 1
}

func newPerson(name string) *Person {
	return &Person{nextId(), name}
}

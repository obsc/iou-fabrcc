package graph

func idGenerator() func() int {
	var curId int = 0
	return func() int {
		curId += 1
		return curId - 1
	}
}

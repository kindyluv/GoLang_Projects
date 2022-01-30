package mapping

type Maps struct {
	Id   int
	Name string
}

func (mapping *Maps) Phoenix(m map[int]string, key int) (id int, details string) {

	for id, details = range m {
		if id == key {
			return id, details
		}
	}
	return 0, "native with this id doesnt exist"
}

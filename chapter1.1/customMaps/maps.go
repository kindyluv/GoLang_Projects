package customMaps

type Map struct {
	Id   int
	Name string
}

func (m *Map) Phoenix(givenMap map[int]string, key int) (id int, details string) {
	for id, details = range givenMap {
		if id == key {
			return id, details
		}
	}
	return 0, "native with this id doesnt exist"
}

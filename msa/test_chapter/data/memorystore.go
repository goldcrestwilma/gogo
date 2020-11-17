package data

var data = []Kitten{
	Kitten{
		Id:     "1",
		Name:   "Felix",
		Weight: 12.3,
	},
	Kitten{
		Id:     "2",
		Name:   "Fat Freddy's Cat",
		Weight: 20.0,
	},
	Kitten{
		Id:     "3",
		Name:   "Garfield",
		Weight: 35.0,
	},
}

type MemoryStore struct {
}

func (m *MemoryStore) Search(name string) []Kitten {
	var kittens []Kitten

	for _, k := range data {
		if k.Name == name {
			kittens = append(kittens, k)
		}
	}

	return kittens
}

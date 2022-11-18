package _map

type OrderedMap[T any] struct {
	Keys []string     `json:"keys"`
	Map  map[string]T `json:"map"`
}

func (m *OrderedMap[T]) Set(key string, value T) {
	if m.Keys == nil {
		m.Keys = make([]string, 0)
	}
	if m.Map == nil {
		m.Map = make(map[string]T, 0)
	}

	if m.isExist(key) {
		m.Map[key] = value
	} else {
		m.Keys = append(m.Keys, key)
		m.Map[key] = value
	}
}

func (m *OrderedMap[T]) Get(key string) (value T) {
	value, _ = m.Map[key]
	return value
}

func (m *OrderedMap[T]) isExist(key string) (res bool) {
	_, res = m.Map[key]
	return res
}

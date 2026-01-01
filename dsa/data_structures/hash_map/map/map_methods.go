package hashmap

import (
	"fmt"
	"maps"
)

func (m *HashMap[K, V]) Size() int     { return len(m.data) }
func (m *HashMap[K, V]) IsEmpty() bool { return len(m.data) == 0 }
func (m *HashMap[K, V]) Clear()        { m.data = make(map[K]V) }

func (m *HashMap[K, V]) Set(key K, value V) {
	m.data[key] = value
}

func (m *HashMap[K, V]) Get(key K) (V, bool) {
	v, ok := m.data[key]
	return v, ok
}

func (m *HashMap[K, V]) Delete(key K) (V, bool) {
	if v, ok := m.data[key]; ok {
		delete(m.data, key)
		return v, true
	}
	var zero V
	return zero, false
}

func (m *HashMap[K, V]) Contains(key K) bool {
	_, ok := m.data[key]
	return ok
}

func (m *HashMap[K, V]) ContainValue(value V) bool {
	for _, v := range m.data {
		if v == value {
			return true
		}
	}
	return false
}

func (m *HashMap[K, V]) ToMap() map[K]V {
	out := make(map[K]V, len(m.data))
	maps.Copy(out, m.data)
	return out
}

func (m *HashMap[K, V]) Keys() []K {
	out := make([]K, 0, len(m.data))
	for k := range m.data {
		out = append(out, k)
	}
	return out
}

func (m *HashMap[K, V]) Values() []V {
	out := make([]V, 0, len(m.data))
	for _, v := range m.data {
		out = append(out, v)
	}
	return out
}

func (m *HashMap[K, V]) Copy() *HashMap[K, V] {
	out := New[K, V]()
	maps.Copy(out.data, m.data)
	return out
}

func (m *HashMap[K, V]) Equal(other *HashMap[K, V]) bool {
	if m == other {
		return true
	}
	if m == nil || other == nil {
		return false
	}
	return maps.Equal(m.data, other.data)
}

func (m *HashMap[K, V]) Traverse(fn func(K, V) bool) {
	for k, v := range m.data {
		if !fn(k, v) {
			return
		}
	}
}

func (m *HashMap[K, V]) String() string {
	return fmt.Sprintf("%v", m.data)
}

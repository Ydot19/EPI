package utils

type Set interface {
	Add(val interface{})
	Remove(val interface{})
	Contains(val interface{}) bool
	//IsSubset(s *Set) bool
}

type HashSetString struct {
	s map[string]struct{}
}

func NewStringHashSet() Set {
	return &HashSetString{
		s: map[string]struct{}{}}

}

func (h *HashSetString) Add(val interface{}) {
	v := val.(string)
	h.s[v] = struct{}{}
}

func (h *HashSetString) Contains(val interface{}) bool {
	v := val.(string)
	_, ok := h.s[v]
	return ok
}

func (h *HashSetString) Remove(val interface{}) {
	v := val.(string)
	delete(h.s, v)
}

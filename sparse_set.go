package ecstatic

const (
	elementNotFound = -1
)

type SparseSet[T any] struct {
	dense    []uint32
	capacity uint32
	n        uint32

	sparse  []uint32
	highest uint32

	payload []T
}

func NewSparseSet[T any](capacity, highest uint32) *SparseSet[T] {
	ss := &SparseSet[T]{
		n:        0,
		capacity: capacity,
		dense:    make([]uint32, capacity),
		payload:  make([]T, capacity),

		highest: highest,
		sparse:  make([]uint32, highest+1),
	}

	return ss
}

func (s *SparseSet[T]) Insert(el uint32, pay T) bool {
	if el > s.highest {
		return false
	}
	if loc, _ := s.Search(el); loc != elementNotFound {
		// Update existing element payload.
		s.payload[loc] = pay
		return true
	}
	if s.n >= s.capacity {
		return false
	}

	s.dense[s.n] = el
	s.payload[s.n] = pay
	s.sparse[el] = s.n
	s.n++

	return true
}

func (s *SparseSet[T]) Delete(el uint32) bool {
	if loc, _ := s.Search(el); loc != elementNotFound {
		// Find last element in dense array.
		replaceEl := s.dense[s.n-1]
		// Overwrite deleted element with last.
		s.dense[loc] = replaceEl
		s.payload[loc] = s.payload[s.n-1]
		// Assign new index to previously last element.
		s.sparse[replaceEl] = uint32(loc)
		// Decrease amount of stored elements.
		s.n--

		return true
	}

	return false
}

func (s *SparseSet[T]) Search(el uint32) (loc int, pay T) {
	loc = elementNotFound
	if el <= s.highest { // Must not be larger than maximum.
		location := s.sparse[el]
		if location < s.n && s.dense[location] == el { // Must be valid.
			loc = int(location)
			pay = s.payload[loc]
		}
	}
	return
}

func (s *SparseSet[T]) Clear() {
	s.n = 0
}

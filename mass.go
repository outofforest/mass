package mass

// New returns new pool.
func New[T any](capacity uint64) *Mass[T] {
	return &Mass[T]{
		pool:     make([]T, capacity),
		capacity: capacity,
	}
}

// Mass maintains a pool of objects to avoid excessive number of heap allocations.
type Mass[T any] struct {
	pool     []T
	capacity uint64
}

// New returns new object from the pool.
func (m *Mass[T]) New() *T {
	if len(m.pool) == 0 {
		m.pool = make([]T, m.capacity)
	}
	o := &m.pool[0]
	m.pool = m.pool[1:]
	return o
}

// NewSlice returns new slice of objects from the pool.
func (m *Mass[T]) NewSlice(n uint64) []T {
	if n > uint64(len(m.pool)) {
		if n2 := n << 1; n2 > m.capacity {
			m.capacity = n2
		}
		m.pool = make([]T, m.capacity)
	}
	o := m.pool[:n]
	m.pool = m.pool[n:]
	return o
}

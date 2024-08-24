package mass

// New returns new pool.
func New[T any](capacity int) *Mass[T] {
	return &Mass[T]{
		pool: make([]T, capacity),
	}
}

// Mass maintains a pool of objects to avoid excessive number of heap allocations.
type Mass[T any] struct {
	pool []T
	pos  int
}

// New returns new object from the pool.
func (m *Mass[T]) New() *T {
	if m.pos == cap(m.pool) {
		m.pool = make([]T, cap(m.pool))
		m.pos = 0
	}
	o := &m.pool[m.pos]
	m.pos++
	return o
}

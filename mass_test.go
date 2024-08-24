package mass

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMass(t *testing.T) {
	m := New[int](10)
	os := make([]*int, 0, 100)
	for i := range cap(os) {
		o := m.New()
		os = append(os, o)
		*o = i
	}
	for i, o := range os {
		require.Equal(t, i, *o)
	}
}

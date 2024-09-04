package mass

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
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

func TestNewSlice(t *testing.T) {
	m := New[int](10)
	os := [][]int{}
	expected := [][]int{}

	k := 0
	for i := range 50 {
		o := m.NewSlice(i)
		e := make([]int, 0, len(o))
		for j := range o {
			o[j] = k
			e = append(e, k)
			k++
		}
		os = append(os, o)
		expected = append(expected, e)
	}
	require.Equal(t, expected, os)
}

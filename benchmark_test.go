package mass

import (
	"fmt"
	"io"
	"testing"
)

// go test -bench=. -run=^$ -cpuprofile profile.out
// go tool pprof -http="localhost:8000" pprofbin ./profile.out

type testType struct {
	Field [4096]byte
}

func BenchmarkNewSingleAllocation(b *testing.B) {
	b.StopTimer()
	b.ResetTimer()

	var o *testType

	b.StartTimer()
	for range b.N {
		o = &testType{}
	}
	b.StopTimer()
	_, _ = fmt.Fprint(io.Discard, o)
}

func BenchmarkNewMassAllocation(b *testing.B) {
	b.StopTimer()
	b.ResetTimer()

	m := New[testType](100)
	var o *testType

	b.StartTimer()
	for range b.N {
		o = m.New()
	}
	b.StopTimer()
	_, _ = fmt.Fprint(io.Discard, o)
}

func BenchmarkNewSliceSingleAllocation(b *testing.B) {
	b.StopTimer()
	b.ResetTimer()

	var o []testType

	b.StartTimer()
	for range b.N {
		o = make([]testType, 10)
	}
	b.StopTimer()
	_, _ = fmt.Fprint(io.Discard, o)
}

func BenchmarkNewSliceMassAllocation(b *testing.B) {
	b.StopTimer()
	b.ResetTimer()

	m := New[testType](1000)
	var o []testType

	b.StartTimer()
	for range b.N {
		o = m.NewSlice(10)
	}
	b.StopTimer()
	_, _ = fmt.Fprint(io.Discard, o)
}

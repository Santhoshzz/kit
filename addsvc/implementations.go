package main

// Add is the abstract definition of what this service does.
// It could easily be an interface type with multiple methods.
type Add func(int64, int64) int64

func pureAdd(a, b int64) int64 { return a + b }

func addVia(r Resource) Add {
	return func(a, b int64) int64 {
		return r.Value(a) + r.Value(b)
	}
}

type Resource interface {
	Value(int64) int64
}

type mockResource struct{}

func (mockResource) Value(i int64) int64 { return i }

// Package doc for a
package a

// This is a's favourite number.
// It returns and integer.
func Number() int {
	return 14902
}

// This is a useful struct
type AStruct struct {
	Num   int
	label string
}

// This is a private struct
type ostruct struct {
	Name string
}

// You should never see this
func secondNumber() int {
	return 42
}

// A very important interface
type Numberer interface {
	// The 'main event'
	Number() int
}

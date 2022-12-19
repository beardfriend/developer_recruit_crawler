package saramin

import (
	"fmt"
	"testing"
)

func TestSaramin(t *testing.T) {
	s := NewSaramin()
	s.get(1, "backend", 1)
}

func TestTotal(t *testing.T) {
	s := NewSaramin()
	a := s.total(1, "backend")
	fmt.Println(a)
}

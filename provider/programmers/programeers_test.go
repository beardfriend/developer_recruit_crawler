package programmers

import (
	"fmt"
	"testing"
)

func TestProgrammers(t *testing.T) {
	p := NewProgrammers()
	res := p.GetRecruitment(1, "backend")
	fmt.Println(res)
}

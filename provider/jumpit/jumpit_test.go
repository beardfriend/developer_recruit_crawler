package jumpit

import (
	"fmt"
	"testing"
)

func TestGetJumpit(t *testing.T) {
	j := &Jumpit{}
	res := j.GetRecruitment(1, "backend")
	fmt.Println(res)
}

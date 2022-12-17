package wanted

import (
	"fmt"
	"testing"
)

func TestGetWanted(t *testing.T) {
	j := &Wanted{}
	res := j.GetRecruitment(1, "backend")
	fmt.Println(res)
}

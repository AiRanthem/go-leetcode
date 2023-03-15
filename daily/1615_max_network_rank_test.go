package daily

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	fmt.Println(maximalNetworkRank(2, [][]int{
		{0, 1},
	}))
}

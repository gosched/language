package sorting

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
)

func TestIsSorted(t *testing.T) {
	items := rand.Perm(100)
	fmt.Println(IsSorted(items))

	sort.Ints(items)
	fmt.Println(IsSorted(items))
}

package sorting

import "sort"

// IsSorted .
func IsSorted(items []int) bool {
	return sort.IntsAreSorted(items)
}

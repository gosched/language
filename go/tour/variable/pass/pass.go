package main

import (
	"fmt"
)

func hello(items []int) {
	fmt.Printf("%p %v\n", items, items) // 0xc000080030 [0 1 2 3 4 5]

	// items[0] = -1

	items = append(items, 0)
	// items = []int{5, 4, 3, 2, 1}

	fmt.Printf("%p %v\n", items, items) // 0xc000060060 [0 1 2 3 4 5 0]
}

func main() {
	var items = []int{0, 1, 2, 3, 4, 5}
	fmt.Printf("%p %v\n", items, items) // 0xc000080030 [0 1 2 3 4 5]
	hello(items)
	fmt.Printf("%p %v\n", items, items) // 0xc000080030 [0 1 2 3 4 5]
}


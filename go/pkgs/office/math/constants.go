package explore

import (
	"fmt"
	"math"
)

// ShowConstants .
func ShowConstants() {
	fmt.Println(math.MaxFloat64)
	fmt.Println(math.MaxFloat32)

	fmt.Println(math.MaxInt64)
	fmt.Println(math.MaxInt32)
	fmt.Println(math.MaxInt16)
	fmt.Println(math.MaxInt8)

	fmt.Println(math.MinInt64)
	fmt.Println(math.MinInt32)
	fmt.Println(math.MinInt16)
	fmt.Println(math.MinInt8)

	fmt.Println(math.Pi)
	fmt.Println(math.E)
}

# Tour

- https://tour.golang.org/list
- https://tour.go-zh.org/welcome/1
- https://go-tour-zh-tw.appspot.com/welcome/1

## Packages, variables, and functions.
- the basic components of any Go program.

### package
- import
- exported names (public variable or private variable)

### functions
- return any number of results
- named return values

### variables
- var
- :=

### basic type
- bool
- string
- int
- int8
- int16
- int32
- int64
- uint
- uint8
- uint16
- uint32
- uint64
- uintptr
- byte (alias for uint8)
- rune (alias for int32) (represents a Unicode code point)
- float32
- float64
- complex64
- complex128

### zero value
- 0
- false
- ""

### type conversions
- float64(101)

### type inference
- the variable's type is inferred from the value

### constants
- const Pi = 3.14
- const Big = 1 << 100
- const Small = Big >> 99

### number literal
```go
const b = 0b1010
const o = 0o12
const x = 0xa

fmt.Printf("%b %b %b\n", b, o, x)
fmt.Printf("%v %v %v\n", b, o, x)
fmt.Printf("%x %x %x\n", b, o, x)
```

## Flow control statements: for, if, else, switch and defer

### for (while)
- the init statement: executed before the first iteration
- the condition expression: evaluated before every iteration
- the post statement: executed at the end of every iteration

### if, else
```go
if v := math.Pow(x, n); v < limit {
    return v
} else {

}
```

### switch
```go
switch os := runtime.GOOS; os {
case "darwin":
    // MacOS
case "linux":
    // Linux
default:
    // Windows...
    // fmt.Printf("%s.\n", os)
}

today := time.Now().Weekday()
switch time.Monday {
case today + 0:
    fmt.Println("Today.")
case today + 1:
    fmt.Println("Tomorrow.")
default:
    fmt.Println("Too far away.")
}

t := time.Now()
switch {
case t.Hour() < 12:
    fmt.Println("morning")
case t.Hour() < 17:
    fmt.Println("afternoon")
default:
    fmt.Println("evening")
}
```

### defer
- defers the execution of a function until the surrounding function returns
```go
fmt.Println("start")
for i := 0; i < 10; i++ {
    defer fmt.Println(i)
}
fmt.Println("end")
```

## More types: structs, slices, and maps.
- https://tour.golang.org/moretypes/1

### pointers

## Methods and interfaces
- https://tour.golang.org/methods/1
- define objects and their behavior

## Concurrency
- https://tour.golang.org/concurrency/1

## Unchecked

### copy
```go
items := []int{1, 2, 3, 4, 5}

items1 := []int{}
items2 := make([]int, len(items))
items3 := make([]int, len(items), len(items))
items4 := make([]int, len(items), cap(items))
items5 := make([]int, 0         , cap(items))

copy(items1, items)
copy(items2, items)
copy(items3, items)
copy(items4, items)
copy(items5, items)

fmt.Println(items1) // []
fmt.Println(items2) // [1 2 3 4 5]
fmt.Println(items3) // [1 2 3 4 5]
fmt.Println(items4) // [1 2 3 4 5]
fmt.Println(items5) // []

m := map[int]int{
    0: 1,
    1: 2,
    2: 3,
}

m1 := map[int]int{}
m2 := make(map[int]int, len(m))

for k, v := range m {
    m1[k] = v
}

for k, v := range m {
    m2[k] = v
}

fmt.Println(m1) // map[0:1 1:2 2:3]
fmt.Println(m2) // map[0:1 1:2 2:3]
```
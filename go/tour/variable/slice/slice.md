# Slice

- https://github.com/golang/go/wiki/SliceTricks

## Index

```go
s := []int{1, 2, 3, 4, 5}
s2 := &[]int{1, 2, 3, 4, 5}

s[0]
(*s2)[0]
```

## From

```go
s2 := s[i:j]
s2 := s[i:i+l]
s2 := s[i:]
s2 := s[:j]
```

## Copy

### shallow copy
```go
s := []int{1, 2, 3, 4, 5}
s2 := s[:]
```

### deep copy
```go
// error
// s1 := []int{1, 2, 3, 4, 5}
// s2 := make([]int, 0, cap(s1))
// copy(s2, s1)
// s2 == []int{}

s1 := []int{1, 2, 3, 4, 5}
s2 := make([]int, len(s1))
copy(s2, s1)

s3 := []int{1, 2, 3, 4, 5}
s4 := append([]int(nil), s3...)

s5 := []int{1, 2, 3, 4, 5}
s6 := append(s5[:0:0], s5...)
```
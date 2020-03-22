package explore

import (
	"fmt"
	"testing"
)

func TestMockPathEscape(t *testing.T) {
	input := []string{
		" test test",
		" test ",
	}

	output := MockPathEscape(input)
	for _, str := range output {
		fmt.Println(str)
	}
}

func TestMockPathUnEscape(t *testing.T) {
	input := []string{
		"%20test%20test",
		"%20test%20",
		"%20%20%20",
		"+test+test",
	}

	output, err := MockPathUnescape(input)
	if err != nil {
		t.Fatalf("%v\n", err)
	}

	for _, str := range output {
		fmt.Println(str)
	}
}

func TestMockQueryEscape(t *testing.T) {
	input := []string{
		" test test",
		" test ",
		" ",
		"%",
	}

	output := MockQueryEscape(input)
	for _, str := range output {
		fmt.Println(str)
	}
}

func TestMockQueryUnescape(t *testing.T) {
	input := []string{
		"%20test%20test",
		"%20test%20",
		"%20%20%20",
		"+test+test",
		"+++",
	}

	output, err := MockQueryUnescape(input)
	if err != nil {
		t.Fatalf("%v\n", err)
	}

	for _, str := range output {
		fmt.Println(str)
	}
}

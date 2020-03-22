package explore

import (
	"net/url"
)

// MockPathEscape .
func MockPathEscape(input []string) []string {
	output := []string{}
	for _, str := range input {
		output = append(output, url.PathEscape(str))
	}
	return output
}

// MockPathUnescape .
func MockPathUnescape(input []string) ([]string, error) {
	output := []string{}
	for _, str := range input {
		temp, err := url.PathUnescape(str)
		if err != nil {
			return nil, err
		}
		output = append(output, temp)
	}
	return output, nil
}

// MockQueryEscape .
func MockQueryEscape(input []string) []string {
	output := []string{}
	for _, str := range input {
		output = append(output, url.QueryEscape(str))
	}
	return output
}

// MockQueryUnescape .
func MockQueryUnescape(input []string) ([]string, error) {
	output := []string{}
	for _, str := range input {
		temp, err := url.QueryUnescape(str)
		if err != nil {
			return nil, err
		}
		output = append(output, temp)
	}
	return output, nil
}

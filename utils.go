package utils

import (
	"errors"
	"fmt"
	"reflect"
)

func HelloWorld() {
	fmt.Println("Hello World from niquefa/gotools!!!")
}

func SumInt(values ...int) (int, error) {
	if len(values) == 0 {
		return 0, errors.New("no arguments provided")
	}
	sum := 0
	for _, v := range values {
		sum += v
	}
	return sum, nil
}

func SumInt64(values ...int64) (int64, error) {
	if len(values) == 0 {
		return 0, errors.New("no arguments provided")
	}
	sum := int64(0)
	for _, v := range values {
		sum += v
	}
	return sum, nil
}

func SumFloat64(values ...float64) (float64, error) {
	if len(values) == 0 {
		return 0, errors.New("no arguments provided")
	}
	sum := float64(0)
	for _, v := range values {
		sum += v
	}
	return sum, nil
}

func MinInt(values ...int) (int, error) {
	if len(values) == 0 {
		return 0, errors.New("no arguments provided")
	}
	min := values[0]
	for _, v := range values[1:] {
		if v < min {
			min = v
		}
	}
	return min, nil
}

func MinInt64(values ...int64) (int64, error) {
	if len(values) == 0 {
		return 0, errors.New("no arguments provided")
	}
	min := values[0]
	for _, v := range values[1:] {
		if v < min {
			min = v
		}
	}
	return min, nil
}

func MinFloat64(values ...float64) (float64, error) {
	if len(values) == 0 {
		return 0, errors.New("no arguments provided")
	}
	min := values[0]
	for _, v := range values[1:] {
		if v < min {
			min = v
		}
	}
	return min, nil
}

func MaxInt(values ...int) (int, error) {
	if len(values) == 0 {
		return 0, errors.New("no arguments provided")
	}
	max := values[0]
	for _, v := range values[1:] {
		if v > max {
			max = v
		}
	}
	return max, nil
}

func MaxInt64(values ...int64) (int64, error) {
	if len(values) == 0 {
		return 0, errors.New("no arguments provided")
	}
	max := values[0]
	for _, v := range values[1:] {
		if v > max {
			max = v
		}
	}
	return max, nil
}

func MaxFloat64(values ...float64) (float64, error) {
	if len(values) == 0 {
		return 0, errors.New("no arguments provided")
	}
	max := values[0]
	for _, v := range values[1:] {
		if v > max {
			max = v
		}
	}
	return max, nil
}

// CountElements takes a slice of interface{} and returns a map where the keys
// are the unique elements from the slice, and the values are the counts of those
// elements. It checks that all elements in the slice are of the same type and
// returns an error if they are not.
func Frequency(elements []interface{}) (map[interface{}]int, error) {
	if len(elements) == 0 {
		return nil, errors.New("slice is empty")
	}
	firstType := reflect.TypeOf(elements[0])
	for _, elem := range elements[1:] {
		if reflect.TypeOf(elem) != firstType {
			return nil, errors.New("all elements must be of the same type")
		}
	}

	counts := make(map[interface{}]int)
	for _, elem := range elements {
		counts[elem]++
	}

	return counts, nil
}

// RuneFrequency takes a string and returns a map where the keys are the unique runes from the string,
// and the values are the counts of those runes.
func RuneFrequency(s string) map[rune]int {
	counts := make(map[rune]int)
	for _, r := range s {
		counts[r]++
	}
	return counts
}

// ReverseInterface reverses a slice of interface{} in place.
func ReverseInterface(arr []interface{}) {
	n := len(arr)
	for i := 0; i < n/2; i++ {
		arr[i], arr[n-i-1] = arr[n-i-1], arr[i]
	}
}

func NextPermutation(a []int) bool {
	n := len(a)
	i := n - 2
	for i >= 0 && a[i] >= a[i+1] {
		i--
	}
	if i < 0 {
		return false
	}
	j := n - 1
	for j >= 0 && a[j] <= a[i] {
		j--
	}
	a[i], a[j] = a[j], a[i]

	reverse(a[i+1:])

	return true
}

func reverse(a []int) {
	i, j := 0, len(a)-1
	for i < j {
		a[i], a[j] = a[j], a[i]
		i++
		j--
	}
}

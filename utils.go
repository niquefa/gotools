package utils

import (
	"errors"
	"fmt"
	"reflect"
)

func HelloWorld() {
	fmt.Println("Hello World from niquefa/gotools!!!")
}

// Sum takes a variable number of interface{} arguments and returns their sum.
// All arguments must be of the same integer or floating-point type; otherwise, an error is returned.
// If the arguments are floating-point, the return value will also be floating-point.
// If the arguments are integers, the return value will also be an integer.
// If no arguments are provided, an error is returned.
func Sum(a ...interface{}) (interface{}, error) {
	if len(a) == 0 {
		return nil, fmt.Errorf("no arguments provided")
	}

	var sumType reflect.Type
	var sumInt int64
	var sumFloat float64

	for _, v := range a {
		valueType := reflect.TypeOf(v)
		if sumType == nil {
			sumType = valueType
		} else if sumType != valueType {
			return nil, fmt.Errorf("all elements must be of the same type")
		}

		switch value := v.(type) {
		case int, int8, int16, int32, int64:
			intValue := reflect.ValueOf(value).Convert(reflect.TypeOf(int64(0))).Int()
			sumInt += intValue
		case uint, uint8, uint16, uint32, uint64:
			uintValue := reflect.ValueOf(value).Convert(reflect.TypeOf(uint64(0))).Uint()
			sumInt += int64(uintValue)
		case float32, float64:
			floatValue := reflect.ValueOf(value).Convert(reflect.TypeOf(float64(0))).Float()
			sumFloat += floatValue
		default:
			return nil, fmt.Errorf("unsupported type: %T", v)
		}
	}

	if sumType.Kind() == reflect.Float32 || sumType.Kind() == reflect.Float64 {
		return sumFloat, nil
	}

	return sumInt, nil
}

// Min takes a slice of interface{} arguments and returns the minimum value.
// All arguments must be of the same integer or floating-point type; otherwise, an error is returned.
func Min(a ...interface{}) (interface{}, error) {
	if len(a) == 0 {
		return nil, fmt.Errorf("no arguments provided")
	}

	minValue := a[0]
	for _, v := range a[1:] {
		if reflect.TypeOf(v) != reflect.TypeOf(minValue) {
			return nil, fmt.Errorf("all elements must be of the same type")
		}
		switch minValue.(type) {
		case int, int8, int16, int32, int64:
			if v.(int) < minValue.(int) {
				minValue = v
			}
		case float32, float64:
			if v.(float64) < minValue.(float64) {
				minValue = v
			}
		default:
			return nil, fmt.Errorf("unsupported type: %T", v)
		}
	}
	return minValue, nil
}

// Max takes a slice of interface{} arguments and returns the maximum value.
// All arguments must be of the same integer or floating-point type; otherwise, an error is returned.
func Max(a ...interface{}) (interface{}, error) {
	if len(a) == 0 {
		return nil, fmt.Errorf("no arguments provided")
	}

	maxValue := a[0]
	for _, v := range a[1:] {
		if reflect.TypeOf(v) != reflect.TypeOf(maxValue) {
			return nil, fmt.Errorf("all elements must be of the same type")
		}
		switch maxValue.(type) {
		case int, int8, int16, int32, int64:
			if v.(int) > maxValue.(int) {
				maxValue = v
			}
		case float32, float64:
			if v.(float64) > maxValue.(float64) {
				maxValue = v
			}
		default:
			return nil, fmt.Errorf("unsupported type: %T", v)
		}
	}
	return maxValue, nil
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

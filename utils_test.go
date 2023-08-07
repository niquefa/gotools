package utils

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	tests := []struct {
		name    string
		args    []interface{}
		want    interface{}
		wantErr bool
	}{
		{"All integers", []interface{}{int32(1), int32(2), int32(3), int32(4)}, int64(10), false},
		{"All floats", []interface{}{float64(1.1), float64(2.2), float64(3.3), float64(4.4)}, float64(11), false},
		{"Mixed types", []interface{}{int32(1), float64(2.2), int32(3), float64(4.4)}, nil, true},
		{"Empty input", []interface{}{}, nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Sum(tt.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Sum() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMin(t *testing.T) {
	tests := []struct {
		name    string
		args    []interface{}
		want    interface{}
		wantErr bool
	}{
		{"All integers", []interface{}{3, 1, 4, 2}, int(1), false},
		{"All floats", []interface{}{3.3, 1.1, 4.4, 2.2}, float64(1.1), false},
		{"Empty input", []interface{}{}, nil, true},
		{"Mixed types", []interface{}{3, 1.1}, nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Min(tt.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Min() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Min() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMax(t *testing.T) {
	tests := []struct {
		name    string
		args    []interface{}
		want    interface{}
		wantErr bool
	}{
		{"All integers", []interface{}{3, 1, 4, 2}, int(4), false},
		{"All floats", []interface{}{3.3, 1.1, 4.4, 2.2}, float64(4.4), false},
		{"Empty input", []interface{}{}, nil, true},
		{"Mixed types", []interface{}{3, 1.1}, nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Max(tt.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Max() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountElements(t *testing.T) {
	tests := []struct {
		input []interface{}
		want  map[interface{}]int
		err   bool
	}{
		{[]interface{}{1, 2, 3, 1, 2, 3, 3}, map[interface{}]int{1: 2, 2: 2, 3: 3}, false},
		{[]interface{}{"apple", "banana", "apple"}, map[interface{}]int{"apple": 2, "banana": 1}, false},
		{[]interface{}{1, "apple"}, nil, true},
	}

	for _, test := range tests {
		got, err := Frequency(test.input)
		if (err != nil) != test.err || !reflect.DeepEqual(got, test.want) {
			t.Errorf("CountElements(%v) = %v, %v; want %v, %v", test.input, got, err, test.want, test.err)
		}
	}
}

func TestCountRunes(t *testing.T) {
	tests := []struct {
		input string
		want  map[rune]int
	}{
		{"apple", map[rune]int{'a': 1, 'p': 2, 'l': 1, 'e': 1}},
		{"banana", map[rune]int{'b': 1, 'a': 3, 'n': 2}},
		{"", map[rune]int{}},
	}

	for _, test := range tests {
		got := RuneFrequency(test.input)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("CountRunes(%v) = %v; want %v", test.input, got, test.want)
		}
	}
}

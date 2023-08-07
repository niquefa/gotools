package utils

import (
	"reflect"
	"testing"
)

func TestSumInt(t *testing.T) {
	tests := []struct {
		name    string
		args    []int
		want    int
		wantErr bool
	}{
		{"Base case", []int{3, 1, 4, 2}, 10, false},
		{"Empty input", []int{}, 0, true},
		{"One element", []int{3}, 3, false},
		{"Negative Values", []int{-3, 0, -10}, -13, false},
		{"Big Values", []int{1 << 10, 1 << 9, 0}, 1536, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SumInt(tt.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("SumInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SumInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumInt64(t *testing.T) {
	tests := []struct {
		name    string
		args    []int64
		want    int64
		wantErr bool
	}{
		{"Base case", []int64{3, 1, 4, 2}, 10, false},
		{"Empty input", []int64{}, 0, true},
		{"One element", []int64{3}, 3, false},
		{"Negative Values", []int64{-3, 0, -10}, -13, false},
		{"Big Values", []int64{1 << 10, 1 << 9, 0}, 1536, false}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SumInt64(tt.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("SumInt64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SumInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumFloat64(t *testing.T) {
	tests := []struct {
		name    string
		args    []float64
		want    float64
		wantErr bool
	}{
		{"Base case", []float64{3, 1, 4, 2}, 10, false},
		{"Empty input", []float64{}, 0, true},
		{"One element", []float64{3}, 3, false},
		{"Negative Values", []float64{-3, 0, -10}, -13, false},
		{"Big Values", []float64{1 << 10, 1 << 9, 0}, 1536, false}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SumFloat64(tt.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("SumFloat64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SumFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverseInterfaceMore(t *testing.T) {
	tests := []struct {
		input []interface{}
		want  []interface{}
	}{
		{[]interface{}{1}, []interface{}{1}},
		{[]interface{}{1, 2}, []interface{}{2, 1}},
		{[]interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9}, []interface{}{9, 8, 7, 6, 5, 4, 3, 2, 1}},
		{[]interface{}{"a", 1, "b", 2}, []interface{}{2, "b", 1, "a"}},
	}

	for _, test := range tests {
		ReverseInterface(test.input)
		if !reflect.DeepEqual(test.input, test.want) {
			t.Errorf("ReverseInterface(%v) got %v; want %v", test.input, test.input, test.want)
		}
	}
}

func TestMinInt(t *testing.T) {
	tests := []struct {
		name    string
		args    []int
		want    int
		wantErr bool
	}{
		{"Base case", []int{3, 1, 4, 2}, int(1), false},
		{"Empty input", []int{}, 0, true},
		{"One element types", []int{3}, 3, false},
		{"Negative Values", []int{-3, 0, -10}, -10, false},
		{"Big Values", []int{1 << 10, 1 << 9, 0}, 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MinInt(tt.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("MinInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MinInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxInt(t *testing.T) {
	tests := []struct {
		name    string
		args    []int
		want    int
		wantErr bool
	}{
		{"Base case", []int{3, 1, 4, 2}, int(4), false},
		{"Empty input", []int{}, 0, true},
		{"One element types", []int{3}, 3, false},
		{"Negative Values", []int{-3, 0, -10}, 0, false},
		{"Big Values", []int{1 << 10, 1 << 9, 0}, 1 << 10, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MaxInt(tt.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("MaxInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MaxInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinInt64(t *testing.T) {
	tests := []struct {
		name    string
		args    []int64
		want    int64
		wantErr bool
	}{
		{"Base case", []int64{3, 1, 4, 2}, int64(1), false},
		{"Empty input", []int64{}, 0, true},
		{"One element types", []int64{3}, 3, false},
		{"Negative Values", []int64{-3, 0, -10}, -10, false},
		{"Big Values", []int64{1 << 10, 1 << 9, 0}, 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MinInt64(tt.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("MinInt64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MinInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxInt64(t *testing.T) {
	tests := []struct {
		name    string
		args    []int64
		want    int64
		wantErr bool
	}{
		{"Base case", []int64{3, 1, 4, 2}, int64(4), false},
		{"Empty input", []int64{}, 0, true},
		{"One element types", []int64{3}, 3, false},
		{"Negative Values", []int64{-3, 0, -10}, 0, false},
		{"Big Values", []int64{1 << 10, 1 << 9, 0}, 1 << 10, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MaxInt64(tt.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("MaxInt64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MaxInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinFloat64(t *testing.T) {
	tests := []struct {
		name    string
		args    []float64
		want    float64
		wantErr bool
	}{
		{"Base case", []float64{3, 1, 4, 2}, float64(1), false},
		{"Empty input", []float64{}, 0, true},
		{"One element types", []float64{3}, 3, false},
		{"Negative Values", []float64{-3, 0, -10}, -10, false},
		{"Big Values", []float64{1 << 10, 1 << 9, 0}, 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MinFloat64(tt.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("MinFloat64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MinFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxFloat64(t *testing.T) {
	tests := []struct {
		name    string
		args    []float64
		want    float64
		wantErr bool
	}{
		{"Base case", []float64{3, 1, 4, 2}, float64(4), false},
		{"Empty input", []float64{}, 0, true},
		{"One element types", []float64{3}, 3, false},
		{"Negative Values", []float64{-3, 0, -10}, 0, false},
		{"Big Values", []float64{1 << 10, 1 << 9, 0}, 1 << 10, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MaxFloat64(tt.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("MaxFloat64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MaxFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFrequency(t *testing.T) {
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
			t.Errorf("Frequency(%v) = %v, %v; want %v, %v", test.input, got, err, test.want, test.err)
		}
	}
}

func TestRuneFrequency(t *testing.T) {
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
			t.Errorf("RuneFrequency(%v) = %v; want %v", test.input, got, test.want)
		}
	}
}

func TestNextPermutation(t *testing.T) {
	tests := []struct {
		input []int
		want  []int
		ok    bool
	}{
		{[]int{1}, []int{1}, false},
		{[]int{2, 1}, []int{2, 1}, false},
		{[]int{1, 2}, []int{2, 1}, true},
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 5, 4}, true},
		{[]int{5, 4, 3, 2, 1}, []int{5, 4, 3, 2, 1}, false},
		{[]int{1, 2, 3}, []int{1, 3, 2}, true},
		{[]int{3, 2, 1}, []int{3, 2, 1}, false},
		{[]int{1, 1, 5}, []int{1, 5, 1}, true},
		{[]int{3, 3, 3, 3}, []int{3, 3, 3, 3}, false},
		{[]int{4, 3, 2, 1}, []int{4, 3, 2, 1}, false},
	}

	for _, test := range tests {
		originalInput := make([]int, len(test.input))
		copy(originalInput, test.input)
		ok := NextPermutation(test.input)
		if ok != test.ok || !reflect.DeepEqual(test.input, test.want) {
			t.Errorf("nextPermutation(%v) got %v; want %v", originalInput, test.input, test.want)
		}
	}
}

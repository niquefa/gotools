package utils

import (
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
		// Add more test cases as needed
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

package main

import (
	"testing"
)

func TestCalculateTotalModuleMass(t *testing.T) {
	tests := []struct {
		testName string
		mass     int
		want     int
	}{
		{
			testName: "1",
			mass:     14,
			want:     2,
		},
		{
			testName: "2",
			mass:     1969,
			want:     966,
		},
		{
			testName: "3",
			mass:     100756,
			want:     50346,
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			if got := calculateTotalModuleMass(tt.mass); got != tt.want {
				t.Errorf("calculateTotalModuleMass() = [%v], want [%v]", got, tt.want)
			}
		})
	}
}

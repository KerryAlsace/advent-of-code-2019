package main

import (
	"testing"
)

func TestCalculateModuleMass(t *testing.T) {
	tests := []struct {
		testName string
		mass     string
		want     int
	}{
		{
			testName: "3",
			mass:     "12",
			want:     2,
		},
		{
			testName: "3",
			mass:     "14",
			want:     2,
		},
		{
			testName: "3",
			mass:     "1969",
			want:     654,
		},
		{
			testName: "3",
			mass:     "100756",
			want:     33583,
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			if got := calculateModuleMass(tt.mass); got != tt.want {
				t.Errorf("calculateModuleMass() = [%v], want [%v]", got, tt.want)
			}
		})
	}
}

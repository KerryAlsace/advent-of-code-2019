package main

import (
	"testing"
)

func TestGetPositionModesOpcodesAndParameters(t *testing.T) {
	tests := []struct {
		testName   string
		thing      []int
		wantOpcode int
		wantParam  []Param
	}{
		{
			testName:   "1",
			thing:      []int{1002, 4, 3, 4, 33},
			wantOpcode: 02,
			wantParam: []Param{
				Param{
					Parameter:    4,
					PositionMode: 0,
				},
				Param{
					Parameter:    3,
					PositionMode: 1,
				},
				Param{
					Parameter:    4,
					PositionMode: 0,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			gotOpcode, gotParam := getPositionModesOpcodesAndParameters(tt.thing)
			if gotOpcode != tt.wantOpcode {
				t.Errorf("getPositionModesOpcodesAndParameters() = [opcode: %v], want [opcode: %v]", gotOpcode, tt.wantOpcode)
			}

			for i, wp := range tt.wantParam {
				if gotParam[i] != wp {
					t.Errorf("getPositionModesOpcodesAndParameters() = [params: %v], want [params: %v]", gotParam, tt.wantParam)
				}
			}
		})
	}
}

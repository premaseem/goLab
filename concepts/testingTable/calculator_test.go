package main

import "testing"


func Test_multiplication(t *testing.T) {

	type inputParams struct{
		v1 int
		v2 int
	}

	tests := []struct {
		inputParams
		expected int
	}{
		{inputParams{0, 0}, 0},
		{inputParams{1, 1}, 1},
	}

	for _,useCase := range tests{

		got := multiplication(useCase.v1, useCase.v2)
		if got != useCase.expected {
			t.Errorf("got %v, expected %v \n", got, useCase.expected)
		}
	}

}
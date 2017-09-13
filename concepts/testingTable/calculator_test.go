package main
import "testing"

type inputParams struct{
	v1 int
	v2 int
}

type tests []struct {
	inputParams inputParams
	expected int
}

func Test_multiplication(t *testing.T) {

	testCases := tests {
		{inputParams{0, 0}, 0},
		{inputParams{1, 1}, 1},
		{inputParams{-2, -2}, 4},
	}

	for _,useCase := range testCases{
		t.Run("multiplication tests ", func(t *testing.T) {
			if got := multiplication(useCase.inputParams.v1, useCase.inputParams.v2); got != useCase.expected {
				t.Errorf("mul() = %v, want %v", got, useCase.expected)
			}
		})
	}
}


func Test_addition(t *testing.T){

	testCases := tests {
		{inputParams{0, 0}, 0},
		{inputParams{1, 1}, 2},
		{inputParams{v1 : -2, v2: -2}, -4},
	}

	for _,useCase := range testCases{
		t.Run("Addition tests", func(t *testing.T) {
			got := add(useCase.inputParams.v1,useCase.inputParams.v2)
			if got != useCase.expected{
				t.Error("add() = %v , expected %v ", got, useCase.expected)
			}
		})
	}

}
package main
import "testing"

func Test_multiplication(t *testing.T) {
	type inputParams struct{
		v1 int
		v2 int
	}

	tests := []struct {
		inputParams inputParams
		expected int
	}{
		{inputParams{0, 0}, 0},
		{inputParams{1, 1}, 1},
		{inputParams{-2, -2}, 4},
	}

	for _,useCase := range tests{
		t.Run("multiplication tests ", func(t *testing.T) {
			if got := multiplication(useCase.inputParams.v1, useCase.inputParams.v2); got != useCase.expected {
				t.Errorf("mul() = %v, want %v", got, useCase.expected)
			}
		})
	}


}
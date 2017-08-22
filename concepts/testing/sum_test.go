package main

import "testing"

func TestSum(t *testing.T){
	total := Sum(6,4)
	if total != 10{
		t.Errorf("sum was incorrect, got: %d, want : %d ", total,10)
	}

}


package main

import (
	"testing"
	"fmt"
)

func TestSumConventional(t *testing.T) {
	total := Sum(6, 4)
	if total != 10 {
		t.Errorf("sum was incorrect, got: %d, want : %d ", total, 10)
	}

}

func TestSumUsingTestTable(t *testing.T) {

	testtable := []struct {
		scenario string
		value1    int
		value2 	  int
		expectedValue int
	}{
		{"Add positive numbers ",2, 3, 5},
		{"Add negagive numbers ",-4,-8, -12},
		{"Add positive number - demo fail  ",4,8, 10},
		{"Add negative and positive numbers ",10, -3, 7},
	}

	for _, testRow := range testtable {
		processedValue := Sum(testRow.value1, testRow.value2)
		if processedValue != testRow.expectedValue {
			t.Errorf("  Scenario Failed :: %s  expected %d got %d",testRow.scenario,testRow.expectedValue, processedValue)
		}else {
			fmt.Printf(" \n %s :: Scenario Passed, expected %d got %d",testRow.scenario,testRow.expectedValue, processedValue)
		}

	}

}

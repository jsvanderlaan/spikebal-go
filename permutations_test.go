package main

import (
	"testing"
)

func TestPermutationsDedupe(t *testing.T) {
	perms := Permutations{
		&Game{2, 4, 8, 0, 3, 1, 6, 5, 7}, // equal
		&Game{0, 8, 2, 4, 1, 3, 5, 6, 7}, // equal
		&Game{0, 2, 8, 4, 1, 3, 5, 6, 7},
		&Game{0, 2, 8, 4, 1, 3, 5, 7, 6},
	}

	res := perms.Dedupe()
	expectedLen := 3

	if len(res) != expectedLen {
		t.Errorf("Expected %v to have %d entries, but found %d", perms, expectedLen, len(res))
	}
}

func TestPermutationsGetPermutations(t *testing.T) {
	sut := []Player{0, 1, 2}
	result := getPermutations(sut)
	expected := [][]Player{
		{0, 1, 2},
		{0, 2, 1},
		{1, 0, 2},
		{1, 2, 0},
		{2, 0, 1},
		{2, 1, 0},
	}
	for i, exp := range expected {
		for j, item := range exp {
			if result[i][j] != item {
				t.Errorf("Expected result %v to equal %v", result, expected)
			}
		}
	}
}

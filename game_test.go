package main

import (
	"testing"
)

func TestGameEqualsFalse(t *testing.T) {
	game := &Game{2, 4, 8, 0, 3, 1, 6, 5, 7}
	game2 := &Game{0, 8, 2, 4, 1, 3, 5, 6, 7}

	result := game.Equals(game2)

	if result {
		t.Errorf("Expected %v to not equal %v", game, game2)
	}
}

func TestGameEqualsTrue(t *testing.T) {
	game := &Game{0, 8, 2, 4, 1, 3, 5, 6, 7}
	game2 := &Game{0, 8, 2, 4, 1, 3, 5, 6, 7}

	result := game.Equals(game2)

	if !result {
		t.Errorf("Expected %v to equal %v", game, game2)
	}
}

func TestGameSort(t *testing.T) {
	sut := Game{2, 4, 8, 0, 3, 1, 6, 5, 7}

	result := sut.Sort()
	expected := Game{0, 8, 2, 4, 1, 3, 5, 6, 7}

	for i := 0; i < 9; i++ {
		if result[i] != expected[i] {
			t.Errorf("Expected %v to be %v", expected, result)
			break
		}
	}
}

func TestGameOpponents(t *testing.T) {
	sut := Game{2, 4, 8, 0, 3, 1, 6, 5, 7}

	res := sut.Opponents()

	expected := &Matrix{
		{0, 0, 1, 0, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 1, 1, 0, 0},
		{1, 0, 0, 0, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 0, 1, 1, 0, 0},
		{1, 0, 0, 0, 0, 0, 0, 0, 1},
		{0, 1, 0, 1, 0, 0, 0, 0, 0},
		{0, 1, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 1, 0, 1, 0, 0, 0, 0},
	}

	if !res.Equals(expected) {
		t.Errorf("Expected result \n%v to be equal to \n%v", res, expected)
	}
}

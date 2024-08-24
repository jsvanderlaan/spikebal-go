package main

import "math/rand"

type Game [9]Player

func (game *Game) Equals(game2 *Game) bool {
	for i := 0; i < 9; i++ {
		if game[i] != game2[i] {
			return false
		}
	}
	return true
}

func (game *Game) Teammates() *Matrix {
	teammates := NewMatrix()

	for i, pos := range game {
		if i < 8 {
			if i%2 == 0 {
				teammates[pos][game[i+1]]++
			} else {
				teammates[pos][game[i-1]]++
			}
		}
	}

	return teammates
}

func (game *Game) Opponents() *Matrix {
	opponents := NewMatrix()

	for i, pos := range game {
		switch i {
		case 0, 1:
			opponents[pos][game[2]]++
			opponents[pos][game[3]]++
		case 2, 3:
			opponents[pos][game[0]]++
			opponents[pos][game[1]]++
		case 4, 5:
			opponents[pos][game[6]]++
			opponents[pos][game[7]]++
		case 6, 7:
			opponents[pos][game[4]]++
			opponents[pos][game[5]]++
		}
	}

	return opponents
}

func (game *Game) Vrij() *Vrij {
	vrij := NewVrij()
	vrij[game[8]] = true
	return vrij
}

func (game *Game) Sort() *Game {
	for i := 0; i < 8; i += 2 {
		if game[i] > game[i+1] {
			game.swap(i, i+1)
		}
	}

	for i := 0; i < 8; i += 4 {
		if game[i] > game[i+2] {
			game.swap(i, i+2)
		}
	}

	if game[0] > game[4] {
		game.swap(0, 4)
	}
	return game
}

func (game *Game) Shuffle() *Game {
	for i := 0; i < 8; i += 2 {
		if rand.Intn(2) == 1 {
			game.swap(i, i+1)
		}
	}

	for i := 0; i < 8; i += 4 {
		if rand.Intn(2) == 1 {
			game.swap(i, i+2)
		}
	}

	if rand.Intn(2) == 1 {
		game.swap(0, 4)
	}
	return game
}

func (game *Game) swap(i int, j int) *Game {
	len := j - i
	for a := i; a < i+len; a++ {
		game[a], game[a+len] = game[a+len], game[a]
	}
	return game
}

package main

import "fmt"

type Tournament struct {
	Games          []Game
	OpenentMatrix  Matrix
	TeammateMatrix Matrix
	Vrij           Vrij
}

func NewTournament() *Tournament {
	return &Tournament{
		Games:          []Game{},
		OpenentMatrix:  *NewMatrix(),
		TeammateMatrix: *NewMatrix(),
		Vrij:           *NewVrij(),
	}
}

func (t Tournament) String() string {
	res := "Tournament\n\nGames\n"

	for _, game := range t.Games {
		res += fmt.Sprintln(game)
	}
	res += fmt.Sprintln()
	res += fmt.Sprintf("Opponent matrix\n%v\n\n", t.OpenentMatrix)
	res += fmt.Sprintf("Teammate matrix\n%v\n\n", t.TeammateMatrix)
	res += fmt.Sprintf("Vrij\n%v\n\n", t.Vrij)
	return res
}

func (t *Tournament) AddGame(g *Game) *Tournament {
	t.Games = append(t.Games, *g)
	t.OpenentMatrix.Add(g.Opponents())
	t.TeammateMatrix.Add(g.Teammates())
	t.Vrij.Add(g.Vrij())
	return t
}

func (t *Tournament) SubLastGame() *Tournament {
	pop := t.Games[len(t.Games)-1]
	t.Games = t.Games[:len(t.Games)-1]
	t.OpenentMatrix.Sub(pop.Opponents())
	t.TeammateMatrix.Sub(pop.Teammates())
	t.Vrij.Sub(pop.Vrij())
	return t
}

func (t *Tournament) CheckAdd(g *Game) bool {
	vrij := g[8]
	if t.Vrij[vrij] || !t.OpenentMatrix.AddingIsLessThan(g.Opponents(), 2) || !t.TeammateMatrix.AddingIsLessThan(g.Teammates(), 1) {
		return false
	}

	return true
}

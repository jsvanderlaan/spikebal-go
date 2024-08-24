package main

import "math/rand"

type Permutations []*Game

func (game *Game) GetPermutations() Permutations {
	slice := []Player{}
	for _, p := range game {
		slice = append(slice, p)
	}

	res := Permutations{}
	perms := (getPermutations(slice))
	for _, perm := range perms {
		res = append(res, (*Game)(perm))
	}
	return res
}

func getPermutations(arr []Player) [][]Player {
	if len(arr) == 0 {
		return [][]Player{}
	}
	if len(arr) == 1 {
		res := [][]Player{arr[:]}
		return res
	}

	res := [][]Player{}
	for i, el := range arr {
		nextPerms := []Player{}
		for j, x := range arr {
			if j != i {
				nextPerms = append(nextPerms, x)
			}
		}
		perms := getPermutations(nextPerms)
		for _, perm := range perms {
			after := append([]Player{el}, perm...)
			res = append(res, after)
		}
	}
	return res
}

func (p Permutations) Dedupe() Permutations {
	deduped := Permutations{}
	for _, game := range p {
		game.Sort()
		dupe := false
		for _, perm := range deduped {
			if perm.Equals(game) {
				dupe = true
				break
			}
		}
		if !dupe {
			deduped = append(deduped, game)
		}
	}
	return deduped
}

func (p *Permutations) Shuffle() *Permutations {
	arr := *p
	for i := range arr {
		j := rand.Intn(len(arr))
		arr[i], arr[j] = arr[j], arr[i]
	}
	return &arr
}

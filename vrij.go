package main

import "fmt"

type Vrij [9]bool

func NewVrij() *Vrij {
	return &Vrij{false, false, false, false, false, false, false, false, false}
}

func (v Vrij) String() string {
	res := ""

	for i, vrij := range v {
		res += fmt.Sprintf("%d: %t\n", i, vrij)
	}
	res += fmt.Sprintln()
	return res
}

func (v *Vrij) Add(v2 *Vrij) *Vrij {
	for i, vrij := range v {
		v[i] = vrij || v2[i]
	}

	return v
}

func (v *Vrij) Sub(v2 *Vrij) *Vrij {
	for i, vrij := range v {
		if v2[i] {
			if !vrij {
				panic(fmt.Sprintf("Cannot sub %v from %v", v, v2))
			}
			v[i] = false
		}
	}

	return v
}

package main

type Player uint8

func (p Player) String() string {
	switch p {
	case 0:
		return "Rens"
	case 1:
		return "Perry"
	case 2:
		return "Bram"
	case 3:
		return "David"
	case 4:
		return "Erik"
	case 5:
		return "Sjoerd"
	case 6:
		return "Leonie"
	case 7:
		return "Jurre"
	case 8:
		return "Victor"
	default:
		return "Unknown"
	}
}

package model

type Node struct {
	Char rune
	Next map[rune]*Node
	End  bool
}

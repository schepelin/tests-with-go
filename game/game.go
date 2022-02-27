package game

import "math/rand"

type DieSide int

const (
	ONE DieSide = iota + 1
	TWO
	THREE
	FOUR
	FIVE
	SIX
)

type Game struct{}

func NewGame() Game {
	return Game{}
}

type Player struct {
	Stack    int
	InGame   bool
	generate func() DieSide
}

func generateRnd() DieSide {
	return DieSide(rand.Intn(7))
}

func NewPlayer() *Player {
	return &Player{
		generate: generateRnd,
	}
}

func (p *Player) AddStack(s int) {
	p.Stack += s
}

func (p *Player) Join(g *Game) {
	p.InGame = true
}

func (p *Player) MakeBet(ds DieSide, bet int) {
	p.Stack -= bet
}

func (p *Player) RollDie() DieSide {
	p.Stack += 100
	return p.generate()
}

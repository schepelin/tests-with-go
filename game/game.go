package game

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
	Stack  int
	InGame bool
}

func NewPlayer() *Player {
	return &Player{}
}

func (p *Player) AddStack(s int) {
	p.Stack += s
}

func (p *Player) Join(g *Game) {
	p.InGame = true
}

func (p *Player) Bet(ds DieSide, bet int) {
	p.Stack -= bet
}

func (p *Player) RollDie() DieSide {
	return FIVE
}

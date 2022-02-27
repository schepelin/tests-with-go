package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type playerBuilder struct {
	stack     int
	inGame    bool
	betOn     DieSide
	generator func() DieSide
}

func (b *playerBuilder) WithStack(v int) *playerBuilder {
	b.stack = v
	return b
}

func (b *playerBuilder) InGame() *playerBuilder {
	b.inGame = true
	return b
}

func (b *playerBuilder) WithBet(d DieSide, v int) *playerBuilder {
	b.stack -= v
	b.betOn = d
	return b
}

func (b *playerBuilder) Build() *Player {
	return &Player{
		InGame:   b.inGame,
		Stack:    b.stack,
		generate: b.generator,
	}
}

func (b *playerBuilder) StubRndGenerator(f func() DieSide) *playerBuilder {
	b.generator = f
	return b
}

func NewPlayerBuilder() *playerBuilder {
	return &playerBuilder{}
}

func Test_WhenBetWins_PlayerGetsDoubledAmount(t *testing.T) {
	// arrange
	pb := NewPlayerBuilder()
	player := pb.
		StubRndGenerator(func() DieSide {
			return FIVE
		}).
		WithStack(100).
		InGame().
		WithBet(FIVE, 50).
		Build()

	// act
	roll := player.RollDie()

	// assert
	assert.Equal(t, FIVE, roll)
	assert.Equal(t, 150, player.Stack)
}

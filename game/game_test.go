package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGame(t *testing.T) {
	game := NewGame()
	player := NewPlayer()
	player.AddStack(100)
	assert.Equal(t, 100, player.Stack)
	player.Join(&game)
	assert.True(t, player.InGame)

	player.MakeBet(FIVE, 50)
	assert.Equal(t, 50, player.Stack)
	roll := player.RollDie()
	if roll == FIVE {
		assert.Equal(t, 150, player.Stack)
	}
}

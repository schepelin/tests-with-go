package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// A good example of a really bad test
func TestGame(t *testing.T) {
	game := NewGame()
	player := NewPlayer()
	player.AddStack(100)
	assert.Equal(t, 100, player.Stack)
	player.Join(&game)
	assert.True(t, player.InGame)

	player.Bet(FIVE, 50)
	roll := player.RollDie()
	if roll != FIVE {
		assert.Equal(t, 150, player.Stack)
	} else {
		assert.Equal(t, 50, player.Stack)
	}
}

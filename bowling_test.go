package bowling

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRollOne(t *testing.T) {
	game := Game{}
	game.Roll(1)

	for i := 0; i < 19; i++ {
		game.Roll(0)
	}

	expected := 1
	actual := game.GetScores()

	if expected != actual {
		t.Errorf("Score %d, want %d", actual, expected)
	}
}

func TestRollAllOne(t *testing.T) {
	game := Game{}
	for i := 0; i < 20; i++ {
		game.Roll(1)
	}

	expected := 21
	actual := game.GetScores()

	assert.Equal(t, actual, expected, "Score is not expected")
}

func TestOneSpare(t *testing.T) {
	game := Game{}
	game.Roll(2)
	game.Roll(8)
	game.Roll(1)
	game.Roll(2) //spare

	expected := 14
	actual := game.GetScores()

	if expected != actual {
		t.Errorf("Score %d, want %d", actual, expected)
	}
}

func TestOneStrike(t *testing.T) {
	game := Game{}
	game.Roll(10) //strike
	game.Roll(1)
	game.Roll(2)

	expected := 16
	actual := game.GetScores()

	if expected != actual {
		t.Errorf("Score %d, want %d", actual, expected)
	}
}

func TestPerfectGame(t *testing.T) {
	game := Game{}
	for i := 0; i < 12; i++ {
		game.Roll(10)
	}

	expected := 300
	actual := game.GetScores()

	if expected != actual {
		t.Errorf("Score %d, want %d", actual, expected)
	}
}

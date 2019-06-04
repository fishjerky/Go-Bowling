package bowling

import (
	"testing"
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

	expected := 20
	actual := game.GetScores()

	if expected != actual {
		t.Errorf("Score %d, want %d", actual, expected)
	}
}

func TestOneSpare(t *testing.T) {
	game := Game{}
	game.Roll(1)
	game.Roll(2)
	game.Roll(3)
	game.Roll(7) //spare

	expected := 15
	actual := game.GetScores()

	if expected != actual {
		t.Errorf("Score %d, want %d", actual, expected)
	}
}

func TestOneStrike(t *testing.T) {
	game := Game{}
	game.Roll(1)
	game.Roll(2)
	game.Roll(10)
	game.Roll(1) //spare

	expected := 17
	actual := game.GetScores()

	if expected != actual {
		t.Errorf("Score %d, want %d", actual, expected)
	}
}

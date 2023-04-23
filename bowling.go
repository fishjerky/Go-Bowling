package bowling

import _ "fmt"

const (
	maxRolls  = 22
	maxPins   = 10
	maxFrames = 10
)

// Bowling game
type Game struct {
	rolls     [maxRolls]int
	rollIndex int
}

// Roll adds the number of knocked down pins to the rolls array
func (g *Game) Roll(pins int) {
	if g.rollIndex < maxRolls && pins <= maxPins {
		g.rolls[g.rollIndex] = pins
		g.rollIndex++
	}
}

// Score returns the total score for the game
func (g *Game) Score() int {
	var score int
	var frameIndex int

	for i := 0; i < g.rollIndex && frameIndex < maxFrames; i += 1 {
		frameIndex++
		if g.isStrike(i) {
			score += maxPins + g.strikeBonus(i)
		} else if g.isSpare(i) {
			score += maxPins + g.spareBonus(i)
			i++
		} else {
			score += g.sumOfBallsInFrame(i)
			i++
		}
	}
	return score
}

// isSpare returns true if the frame at the given index is a spare
func (g *Game) isSpare(rollIndex int) bool {
	return rollIndex < g.rollIndex-1 && g.rolls[rollIndex]+g.rolls[rollIndex+1] == maxPins
}

// isStrike returns true if the frame at the given index is a strike
func (g *Game) isStrike(rollIndex int) bool {
	return g.rolls[rollIndex] == maxPins
}

// spareBonus returns the bonus points for a spare
func (g *Game) spareBonus(rollIndex int) int {
	return g.rolls[rollIndex+2]
}

// strikeBonus returns the bonus points for a strike
func (g *Game) strikeBonus(rollIndex int) int {
	return g.rolls[rollIndex+1] + g.rolls[rollIndex+2]
}

// sumOfBallsInFrame returns the sum of the two rolls in a frame
func (g *Game) sumOfBallsInFrame(rollIndex int) int {
	return g.rolls[rollIndex] + g.rolls[rollIndex+1]
}

package bowling

//Bowling game
type Game struct {
	inners     [22]int
	innerIndex int
}

//Let roll
func (g *Game) Roll(pin int) {
	g.inners[g.innerIndex] += pin
	g.innerIndex++
}

func (g *Game) GetScores() int {
	scores := 0
	frameIndex := 0
	for i := 0; i < g.innerIndex; i++ {
		scores += g.inners[i]

		//spare
		if g.isSpare(i, frameIndex) {
			//bouns
			scores += g.inners[i-2]
		}

	}

	return scores
}

func (g *Game) isSpare(innerIndex, frameIndex int) bool {
	result := false
	if innerIndex < 2 && frameIndex%2 == 0 {
		return false
	}
	if (g.inners[innerIndex-1] + g.inners[innerIndex]) == 10 {
		result = true
	}
	return result
}

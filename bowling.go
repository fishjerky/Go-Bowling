package bowling

import _ "fmt"

//Bowling game
type Game struct {
	rolls     [22]int
	rollIndex int
}

//Let roll
func (g *Game) Roll(pin int) {
	g.rolls[g.rollIndex] += pin
	g.rollIndex++
}

func (g *Game) GetScores() int {
	scores := 0
	frameIndex := 1
	upperInner := true
	for i := 0; i < g.rollIndex; i++ {
		currentPin := g.rolls[i]
		scores += currentPin
		
		//spare
		if g.isSpare(i) {
			//bouns
			scores += currentPin
			frameIndex++
			upperInner = true
			continue
		}

		//spare
		if g.isStrike(i) {
			
			//is Last frame
			if frameIndex == 10 {
				continue
			}

			//bouns
			scores +=  g.rolls[i+1]
			scores +=  g.rolls[i+2]

			upperInner = true
			frameIndex++
		}
		if upperInner == false {
			frameIndex++
		}
		upperInner = false
		//fmt.Println(frameIndex,i , scores)
	}

	return scores
}

func (g *Game) isSpare(rollIndex int) bool {
	result := false
	if rollIndex < 2 {
		return false
	}
	if (g.rolls[rollIndex-1] + g.rolls[rollIndex-2]) == 10 {
		result = true
	}
	return result
}

func (g *Game) isStrike(rollIndex int) bool {
	result := false
	if g.rolls[rollIndex] == 10 {
		result = true
	}
	return result
}

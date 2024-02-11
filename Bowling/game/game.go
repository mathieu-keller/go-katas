package game

type Frame struct {
	PinsRolled []int
	Score      int
}

type Game struct {
	frames []Frame
}

func NewGame() *Game {
	return &Game{frames: make([]Frame, 10)}
}

func (game *Game) AddRoll(pins int) {
	for i := 0; i < 10; i++ {
		if len(game.frames[i].PinsRolled) < 2 && game.frames[i].Score < 10 {
			game.roll(pins, i, 10)
			break
		}
		if i == 9 {
			game.handleLastFrame(pins, i)
			break
		}
	}
}

func (game *Game) handleLastFrame(pins int, i int) {
	if len(game.frames[i].PinsRolled) == 0 {
		game.roll(pins, i, 10)
	} else if len(game.frames[i].PinsRolled) == 1 {
		game.handleSingleRollLastFrame(pins, i)
	} else {
		game.handleMultipleRollsLastFrame(pins, i)
	}
}

func (game *Game) handleSingleRollLastFrame(pins int, i int) {
	if game.frames[i].PinsRolled[0] == 10 {
		game.roll(pins, i, 20)
	} else {
		game.roll(pins, i, 10)
	}
}

func (game *Game) handleMultipleRollsLastFrame(pins int, i int) {
	if game.frames[i].PinsRolled[0]+game.frames[i].PinsRolled[1] == 10 {
		game.roll(pins, i, 20)
	} else if game.frames[i].PinsRolled[0]+game.frames[i].PinsRolled[1] == 20 {
		game.roll(pins, i, 30)
	}
}

func (game *Game) roll(pins int, frameIndex int, pinsAvailable int) {
	pinsAvailable = game.calculatePinsAvailable(pins, frameIndex, pinsAvailable)
	if pinsAvailable < pins {
		panic("Too many pins")
	}
	game.frames[frameIndex].PinsRolled = append(game.frames[frameIndex].PinsRolled, pins)
	game.frames[frameIndex].Score += pins
	game.updateScores(pins, frameIndex)
}

func (game *Game) calculatePinsAvailable(pins int, frameIndex int, pinsAvailable int) int {
	for _, pin := range game.frames[frameIndex].PinsRolled {
		pinsAvailable -= pin
	}
	return pinsAvailable
}

func (game *Game) updateScores(pins int, frameIndex int) {
	if game.isSpare(frameIndex) {
		game.frames[frameIndex-1].Score += pins
	}
	if game.isStrike(frameIndex) {
		game.frames[frameIndex-1].Score += pins
		if frameIndex > 1 && len(game.frames[frameIndex-2].PinsRolled) == 1 && len(game.frames[frameIndex].PinsRolled) == 1 {
			game.frames[frameIndex-2].Score += pins
		}
	}
}

func (game *Game) isSpare(frameIndex int) bool {
	return frameIndex > 0 && len(game.frames[frameIndex].PinsRolled) == 1 && game.frames[frameIndex-1].Score == 10 && game.frames[frameIndex-1].PinsRolled[0] != 10
}

func (game *Game) isStrike(frameIndex int) bool {
	return frameIndex > 0 && len(game.frames[frameIndex-1].PinsRolled) == 1 && len(game.frames[frameIndex].PinsRolled) < 3
}

func (game *Game) Frames() []Frame {
	return game.frames
}

func (game *Game) TotalScore() int {
	totalScore := 0
	for _, frame := range game.frames {
		totalScore += frame.Score
	}
	return totalScore
}

func (game *Game) Over() bool {
	for i, frame := range game.frames {
		if len(frame.PinsRolled) == 0 || (len(frame.PinsRolled) < 2 && frame.PinsRolled[0] != 10) {
			return false
		}
		if i == 9 && !game.isLastFrameOver(frame) {
			return false
		}
	}
	return true
}

func (game *Game) isLastFrameOver(frame Frame) bool {
	if len(frame.PinsRolled) < 3 {
		for _, pin := range frame.PinsRolled {
			if pin == 10 {
				return false
			}
		}
		if len(frame.PinsRolled) == 2 && frame.PinsRolled[0]+frame.PinsRolled[1] == 10 {
			return false
		}
	}
	return true
}

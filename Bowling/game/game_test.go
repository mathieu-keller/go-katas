package game

import (
	"testing"
)

func TestSpare(t *testing.T) {
	game := NewGame()
	game.AddRoll(6)
	game.AddRoll(4)
	game.AddRoll(5)
	if game.Frames()[0].Score != 15 {
		t.Error("Expected 15, got ", game.Frames()[0].Score)
	}
	assertScore(t, 20, game.TotalScore())
}

func TestStrike(t *testing.T) {
	game := NewGame()
	game.AddRoll(10)
	game.AddRoll(4)
	game.AddRoll(5)
	if len(game.Frames()[0].PinsRolled) != 1 {
		t.Error("Expected 1, got ", len(game.Frames()[0].PinsRolled))
	}
	if len(game.Frames()[1].PinsRolled) != 2 {
		t.Error("Expected 2, got ", len(game.Frames()[1].PinsRolled))
	}
	if game.Frames()[0].Score != 19 {
		t.Error("Expected 19, got ", game.Frames()[0].Score)
	}
	if game.Frames()[1].Score != 9 {
		t.Error("Expected 9, got ", game.Frames()[1].Score)
	}
	assertScore(t, 28, game.TotalScore())
}

func TestDoubleStrike(t *testing.T) {
	game := NewGame()
	game.AddRoll(10)

	game.AddRoll(10)

	game.AddRoll(5)
	game.AddRoll(4)
	if len(game.Frames()[0].PinsRolled) != 1 {
		t.Error("Expected 1, got ", len(game.Frames()[0].PinsRolled))
	}
	if len(game.Frames()[1].PinsRolled) != 1 {
		t.Error("Expected 1, got ", len(game.Frames()[1].PinsRolled))
	}
	if len(game.Frames()[2].PinsRolled) != 2 {
		t.Error("Expected 2, got ", len(game.Frames()[2].PinsRolled))
	}
	if game.Frames()[0].Score != 25 {
		t.Error("Expected 25, got ", game.Frames()[0].Score)
	}
	if game.Frames()[1].Score != 19 {
		t.Error("Expected 19, got ", game.Frames()[0].Score)
	}
	assertScore(t, 53, game.TotalScore())
}

func TestSpareStrike(t *testing.T) {
	game := NewGame()
	game.AddRoll(6)
	game.AddRoll(4)
	game.AddRoll(10)
	if len(game.Frames()[0].PinsRolled) != 2 {
		t.Error("Expected 2, got ", len(game.Frames()[0].PinsRolled))
	}
	if len(game.Frames()[1].PinsRolled) != 1 {
		t.Error("Expected 1, got ", len(game.Frames()[1].PinsRolled))
	}
	if game.Frames()[0].Score != 20 {
		t.Error("Expected 20, got ", game.Frames()[0].Score)
	}
	if game.Frames()[1].Score != 10 {
		t.Error("Expected 10, got ", game.Frames()[0].Score)
	}
	assertScore(t, 30, game.TotalScore())
}

func TestStrikeSpare(t *testing.T) {
	game := NewGame()
	game.AddRoll(10)
	game.AddRoll(4)
	game.AddRoll(6)
	if len(game.Frames()[0].PinsRolled) != 1 {
		t.Error("Expected 1, got ", len(game.Frames()[0].PinsRolled))
	}
	if len(game.Frames()[1].PinsRolled) != 2 {
		t.Error("Expected 2, got ", len(game.Frames()[1].PinsRolled))
	}
	if game.Frames()[0].Score != 20 {
		t.Error("Expected 20, got ", game.Frames()[0].Score)
	}
	if game.Frames()[1].Score != 10 {
		t.Error("Expected 10, got ", game.Frames()[0].Score)
	}
	assertScore(t, 30, game.TotalScore())
}

func TestNormalRoll(t *testing.T) {
	game := NewGame()
	game.AddRoll(5)
	game.AddRoll(4)
	if game.Frames()[0].Score != 9 {
		t.Error("Expected 9, got ", game.Frames()[0].Score)
	}
}

func TestHoleGame(t *testing.T) {
	game := NewGame()
	game.AddRoll(1)
	assertScore(t, 1, game.TotalScore())
	game.AddRoll(4)
	assertScore(t, 5, game.TotalScore())

	game.AddRoll(4)
	assertScore(t, 9, game.TotalScore())
	game.AddRoll(5)
	assertScore(t, 14, game.TotalScore())

	game.AddRoll(6)
	assertScore(t, 20, game.TotalScore())
	game.AddRoll(4)
	assertScore(t, 24, game.TotalScore())

	game.AddRoll(5)
	assertScore(t, 34, game.TotalScore())
	game.AddRoll(5)
	assertScore(t, 39, game.TotalScore())

	game.AddRoll(10)
	assertScore(t, 59, game.TotalScore())

	game.AddRoll(0)
	assertScore(t, 59, game.TotalScore())
	game.AddRoll(1)
	assertScore(t, 61, game.TotalScore())

	game.AddRoll(7)
	assertScore(t, 68, game.TotalScore())
	game.AddRoll(3)
	assertScore(t, 71, game.TotalScore())

	game.AddRoll(6)
	assertScore(t, 83, game.TotalScore())
	game.AddRoll(4)
	assertScore(t, 87, game.TotalScore())

	game.AddRoll(10)
	assertScore(t, 107, game.TotalScore())

	game.AddRoll(2)
	assertScore(t, 111, game.TotalScore())
	game.AddRoll(8)
	assertScore(t, 127, game.TotalScore())
	game.AddRoll(6)
	assertScore(t, 133, game.TotalScore())
	if game.Over() != true {
		t.Error("Expected true, got ", game.Over())
	}
}

func TestStrikeHoleGame(t *testing.T) {
	game := NewGame()
	game.AddRoll(10)
	assertScore(t, 10, game.TotalScore())

	game.AddRoll(10)
	assertScore(t, 30, game.TotalScore())

	game.AddRoll(10)
	assertScore(t, 60, game.TotalScore())

	game.AddRoll(10)
	assertScore(t, 90, game.TotalScore())

	game.AddRoll(10)
	assertScore(t, 120, game.TotalScore())

	game.AddRoll(10)
	assertScore(t, 150, game.TotalScore())

	game.AddRoll(10)
	assertScore(t, 180, game.TotalScore())

	game.AddRoll(10)
	assertScore(t, 210, game.TotalScore())

	game.AddRoll(10)
	assertScore(t, 240, game.TotalScore())

	game.AddRoll(10)
	assertScore(t, 270, game.TotalScore())
	game.AddRoll(10)
	assertScore(t, 290, game.TotalScore())
	game.AddRoll(10)
	assertScore(t, 300, game.TotalScore())
	if game.Over() != true {
		t.Error("Expected true, got ", game.Over())
	}
}

func TestSpareHoleGame(t *testing.T) {
	game := NewGame()
	game.AddRoll(5)
	assertScore(t, 5, game.TotalScore())
	game.AddRoll(5)
	assertScore(t, 10, game.TotalScore())

	game.AddRoll(5)
	assertScore(t, 20, game.TotalScore())
	game.AddRoll(5)
	assertScore(t, 25, game.TotalScore())

	game.AddRoll(5)
	assertScore(t, 35, game.TotalScore())
	game.AddRoll(5)
	assertScore(t, 40, game.TotalScore())

	game.AddRoll(5)
	assertScore(t, 50, game.TotalScore())
	game.AddRoll(5)
	assertScore(t, 55, game.TotalScore())

	game.AddRoll(5)
	assertScore(t, 65, game.TotalScore())
	game.AddRoll(5)
	assertScore(t, 70, game.TotalScore())

	game.AddRoll(5)
	assertScore(t, 80, game.TotalScore())
	game.AddRoll(5)
	assertScore(t, 85, game.TotalScore())

	game.AddRoll(5)
	assertScore(t, 95, game.TotalScore())
	game.AddRoll(5)
	assertScore(t, 100, game.TotalScore())

	game.AddRoll(5)
	assertScore(t, 110, game.TotalScore())
	game.AddRoll(5)
	assertScore(t, 115, game.TotalScore())

	game.AddRoll(5)
	assertScore(t, 125, game.TotalScore())
	game.AddRoll(5)
	assertScore(t, 130, game.TotalScore())

	game.AddRoll(5)
	assertScore(t, 140, game.TotalScore())
	game.AddRoll(5)
	assertScore(t, 145, game.TotalScore())
	game.AddRoll(5)
	assertScore(t, 150, game.TotalScore())
	if game.Over() != true {
		t.Error("Expected true, got ", game.Over())
	}
}

func assertScore(t *testing.T, expected int, actual int) {
	if expected != actual {
		t.Error("Expected ", expected, ", got ", actual)
	}
}

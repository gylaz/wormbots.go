package worm_test

import (
	"testing"

	"github.com/gylaz/wormbots.go/worm"
)

var scenarios = []struct {
	worm        worm.Worm
	expectedDir string
	expectedX   int
	expectedY   int
}{
	{worm.Worm{"up", worm.MaxX, worm.MinY}, "left", worm.MaxX - 1, worm.MinY},
	{worm.Worm{"up", worm.MinX, worm.MinY}, "right", worm.MinX + 1, worm.MinY},
	{worm.Worm{"down", worm.MaxX, worm.MaxY}, "left", worm.MaxX - 1, worm.MaxY},
	{worm.Worm{"down", worm.MinX, worm.MaxY}, "right", worm.MinY + 1, worm.MaxY},
	{worm.Worm{"right", worm.MaxX, worm.MinY}, "down", worm.MaxX, worm.MinY + 1},
	{worm.Worm{"right", worm.MaxX, worm.MaxY}, "up", worm.MaxX, worm.MaxY - 1},
	{worm.Worm{"left", worm.MinX, worm.MinY}, "down", worm.MinX, worm.MinY + 1},
	{worm.Worm{"left", worm.MinX, worm.MaxY}, "up", worm.MinX, worm.MaxY - 1},
}

func TestMove(t *testing.T) {
	for _, scenario := range scenarios {
		var worm = scenario.worm

		worm.Move()

		if worm.Dir != scenario.expectedDir {
			t.Error("Worm did not turn")
		}

		if worm.X != scenario.expectedX || worm.Y != scenario.expectedY {
			t.Error("Worm did not move")
		}
	}
}

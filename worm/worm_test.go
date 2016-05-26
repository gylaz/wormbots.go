package worm_test

import (
	"testing"

	"github.com/gylaz/wormbots.go/worm"
)

func TestMovingUpIntoTopRightCorner(t *testing.T) {
	w := worm.Worm{Dir: "up", X: worm.MaxX, Y: worm.MinY}

	w.Move()

	if w.Dir != "left" {
		t.Error("Worm did not turn")
	}

	if w.X != worm.MaxX-1 && w.Y != worm.MinY {
		t.Error("Worm did not move")
	}
}

func TestWormMovingRightIntoTopRightCorner(t *testing.T) {
	w := worm.Worm{Dir: "right", X: worm.MaxX, Y: worm.MinY}

	w.Move()

	if w.Dir != "down" {
		t.Error("Worm did not turn")
	}

	if w.X != worm.MaxX && w.Y != worm.MinY+1 {
		t.Error("Worm did not move")
	}
}

func TestMovingRightIntoBottomRightCorner(t *testing.T) {
	w := worm.Worm{Dir: "right", X: worm.MaxX, Y: worm.MaxY}

	w.Move()

	if w.Dir != "up" {
		t.Error("Worm did not turn")
	}

	if w.X != worm.MaxX && w.Y != worm.MaxY-1 {
		t.Error("Worm did not move")
	}
}

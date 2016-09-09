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

	if w.X != worm.MaxX-1 || w.Y != worm.MinY {
		t.Error("Worm did not move")
	}
}

func TestMovingUpIntoTopLeftCorner(t *testing.T) {
	w := worm.Worm{Dir: "up", X: worm.MinX, Y: worm.MinY}

	w.Move()

	if w.Dir != "right" {
		t.Error("Worm did not turn")
	}

	if w.X != worm.MinX+1 || w.Y != worm.MinY {
		t.Error("Worm did not move")
	}
}

func TestMovingDownIntoBottomRightCorner(t *testing.T) {
	w := worm.Worm{Dir: "down", X: worm.MaxX, Y: worm.MaxY}

	w.Move()

	if w.Dir != "left" {
		t.Error("Worm did not turn")
	}

	if w.X != worm.MaxX-1 || w.Y != worm.MaxY {
		t.Error("Worm did not move")
	}
}

func TestMovingDownIntoBottomLeftCorner(t *testing.T) {
	w := worm.Worm{Dir: "down", X: worm.MinX, Y: worm.MaxY}

	w.Move()

	if w.Dir != "right" {
		t.Error("Worm did not turn")
	}

	if w.X != worm.MinX+1 || w.Y != worm.MaxY {
		t.Error("Worm did not move")
	}
}

func TestWormMovingRightIntoTopRightCorner(t *testing.T) {
	w := worm.Worm{Dir: "right", X: worm.MaxX, Y: worm.MinY}

	w.Move()

	if w.Dir != "down" {
		t.Error("Worm did not turn")
	}

	if w.X != worm.MaxX || w.Y != worm.MinY+1 {
		t.Error("Worm did not move")
	}
}

func TestMovingRightIntoBottomRightCorner(t *testing.T) {
	w := worm.Worm{Dir: "right", X: worm.MaxX, Y: worm.MaxY}

	w.Move()

	if w.Dir != "up" {
		t.Error("Worm did not turn")
	}

	if w.X != worm.MaxX || w.Y != worm.MaxY-1 {
		t.Error("Worm did not move")
	}
}

func TestWormMovingLeftIntoTopLeftCorner(t *testing.T) {
	w := worm.Worm{Dir: "left", X: worm.MinX, Y: worm.MinY}

	w.Move()

	if w.Dir != "down" {
		t.Error("Worm did not turn")
	}

	if w.X != worm.MinX || w.Y != worm.MinY+1 {
		t.Error("Worm did not move")
	}
}

func TestMovingLeftIntoBottomLeftCorner(t *testing.T) {
	w := worm.Worm{Dir: "left", X: worm.MinX, Y: worm.MaxY}

	w.Move()

	if w.Dir != "up" {
		t.Error("Worm did not turn")
	}

	if w.X != worm.MinX || w.Y != worm.MaxY-1 {
		t.Error("Worm did not move")
	}
}

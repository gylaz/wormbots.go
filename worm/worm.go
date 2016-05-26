package worm

import (
	"math/rand"
)

const MaxX = 320
const MaxY = 240
const MinX = 0
const MinY = 0

type Worm struct {
	Dir string `json:"dir"`
	X   int    `json:"x"`
	Y   int    `json:"y"`
}

func New() Worm {
	return Worm{Dir: "up", X: rand.Intn(MaxX), Y: rand.Intn(MaxY)}
}

func (w *Worm) Move() {
	w.setDirection()
	w.advancePosition()
}

func (w *Worm) setDirection() {
	dirs := w.availableDirs()
	w.Dir = dirs[rand.Intn(len(dirs))]
}

func (w *Worm) advancePosition() {
	switch w.Dir {
	case "up":
		w.Y--
	case "down":
		w.Y++
	case "right":
		w.X++
	case "left":
		w.X--
	}
}

func (w *Worm) availableDirs() []string {
	dirs := []string{}

	switch w.Dir {
	case "up":
		dirs = append(dirs, w.availableHorizontalDirs()...)

		if w.Y > MinY {
			dirs = append(dirs, "up")
		}
	case "down":
		dirs = append(dirs, w.availableHorizontalDirs()...)

		if w.Y < MaxY {
			dirs = append(dirs, "down")
		}
	case "left":
		dirs = append(dirs, w.availableVerticalDirs()...)

		if w.X > MaxX {
			dirs = append(dirs, "left")
		}
	case "right":
		dirs = append(dirs, w.availableVerticalDirs()...)

		if w.X < MaxX {
			dirs = append(dirs, "right")
		}
	}

	return dirs
}

func (w *Worm) availableHorizontalDirs() []string {
	if w.X == MinX {
		return []string{"right"}
	} else if w.X == MaxX {
		return []string{"left"}
	} else {
		return []string{"right", "left"}
	}
}

func (w *Worm) availableVerticalDirs() []string {
	if w.Y == MinY {
		return []string{"down"}
	} else if w.Y == MaxY {
		return []string{"up"}
	} else {
		return []string{"up", "down"}
	}
}

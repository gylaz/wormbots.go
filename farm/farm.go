package farm

import "github.com/gylaz/wormbots.go/worm"

type Farm struct {
	Worms []worm.Worm
}

func New() Farm {
	farm := Farm{make([]worm.Worm, 10)}

	for i := 0; i < len(farm.Worms); i++ {
		farm.Worms[i] = worm.New()
	}

	return farm
}

func (f *Farm) Tick() {
	for i := 0; i < len(f.Worms); i++ {
		f.Worms[i].Move()
	}
}

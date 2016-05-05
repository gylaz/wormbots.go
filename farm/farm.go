package farm

import "math/rand"

type Farm struct {
	Worms []worm
}

func New() Farm {
	farm := Farm{make([]worm, 10)}

	for i := 0; i < len(farm.Worms); i++ {
		w := worm{Dir: "up", X: rand.Intn(640), Y: rand.Intn(360)}
		farm.Worms[i] = w
	}

	return farm
}

func (f *Farm) Tick() {
	for i := 0; i < len(f.Worms); i++ {
		f.Worms[i].Move()
	}
}

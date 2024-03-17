package reptiles

type Dragon struct {
	hydra
}

func NewDragon(name string) *Dragon {
	return &Dragon{
		hydra: *NewHydra(name),
	}
}

package reptiles

type Dragon struct {
	hydra
}

func NewDragon(name string) (*Dragon, error) {
	hydra, err := NewHydra(name)
	if err != nil {
		return &Dragon{}, err
	}
	return &Dragon{
		hydra: *hydra,
	}, nil
}

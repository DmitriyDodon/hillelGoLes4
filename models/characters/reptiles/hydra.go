package reptiles

import (
	"errors"
	"fmt"
	custommath "les4/customMath"
	"math/rand"
)

const hydraMinHeadsQuantity = 50
const hydraMaxHeadsQuantity = 150
const hydraHeadsQuantityToWin = 200

type hydra struct {
	headsQuantity int
	name          string
}

func NewHydra(name string) *hydra {
	return &hydra{
		headsQuantity: int(custommath.GenerateRandomNumberByRange(hydraMinHeadsQuantity, hydraMaxHeadsQuantity)),
		name:          name,
	}
}

func (h *hydra) HeadsQuantity() int {
	return h.headsQuantity
}

func (h *hydra) TakeDamage(damage int) {
	h.headsQuantity -= damage
	fmt.Printf("%s губить %d голiв.\n", h.name, damage)
	if damage > 0 {
		for i := 1; i <= damage; i++ {
			h.HeadsGrowth(i)
		}
	}
}

// Multiple Heads: The Hydra is known for having multiple heads,
// often portrayed as regenerating two new heads for every one that is cut off.
// todo tests
func (h *hydra) HeadsGrowth(numberHeadLost int) {
	newHeadsQuantity := randomHeadsNumberWithProbability()
	h.headsQuantity += newHeadsQuantity
	fmt.Printf("%s вiдрощує %d знову за вiдрублену голову #%d.\n", h.name, newHeadsQuantity, numberHeadLost)
}

func (h *hydra) Name() string {
	return h.name
}

func (h *hydra) HasHeadsLeft() bool {
	return h.headsQuantity > 0
}

func (h *hydra) IsWon() (bool, error) {
	wonFlag := h.headsQuantity >= hydraHeadsQuantityToWin

	if h.headsQuantity > 190 {
		return wonFlag, errors.New("перемога близько")
	}

	return wonFlag, nil
}

func randomHeadsNumberWithProbability() int {
	randomNumber := rand.Intn(100)
	switch {
	case randomNumber < 35:
		return 0
	case randomNumber < 70:
		return 1
	case randomNumber < 90:
		return 2
	default:
		return 3
	}
}

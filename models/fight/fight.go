package fight

import (
	"fmt"
	"les4/models/characters/humans"
	"les4/models/characters/reptiles"

	log "github.com/sirupsen/logrus"
)

type Fight struct {
	hero           *humans.Hero
	dragon         *reptiles.Dragon
	figthIteration int
}

func NewFight(hero *humans.Hero, dragon *reptiles.Dragon) *Fight {
	return &Fight{
		hero:           hero,
		dragon:         dragon,
		figthIteration: 0,
	}
}

func (f *Fight) PrintFightersInfo() {
	fmt.Printf("Dragon info: %#v\n", f.dragon)
	fmt.Printf("Hero info info: %#v\n", f.hero)
}

func (f *Fight) Start() {
	for {
		f.figthIteration++
		fmt.Printf("\n################## Нова %d iтерацiя бою ################################\n", f.figthIteration)

		f.hero.AttackDragon(f.dragon)

		if !f.dragon.HasHeadsLeft() {
			f.hero.WinVoice(f.dragon)
			break
		}

		isDragonWon, err := f.dragon.IsWon()

		if err != nil {
			log.Info(err.Error())
			fmt.Println(err.Error())
		}

		if isDragonWon {
			f.hero.LooseVoice()
		}
		fmt.Printf("У дракона залишилось %d голiв\n", f.dragon.HeadsQuantity())
	}
}

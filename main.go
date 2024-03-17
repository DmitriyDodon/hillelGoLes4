package main

import (
	"fmt"
	"les4/models/characters/humans"
	"les4/models/characters/reptiles"
	"les4/models/fight"

	facker "github.com/go-faker/faker/v4"
)

func main() {
	hero := humans.NewHero()
	dragon := reptiles.NewDragon(fmt.Sprintf("%s Hydrovich", facker.Name()))

	fight := fight.NewFight(hero, dragon)

	fight.PrintFightersInfo()

	fight.Start()

}

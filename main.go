package main

import (
	"fmt"
	"les4/models/characters/humans"
	"les4/models/characters/reptiles"
	"les4/models/fight"

	facker "github.com/go-faker/faker/v4"
	log "github.com/sirupsen/logrus"
)

func main() {
	hero, err := humans.NewHero()
	if err != nil {
		log.Error(err.Error())
		return
	}

	dragon, err := reptiles.NewDragon(fmt.Sprintf("%s Hydrovich", facker.Name()))
	if err != nil {
		log.Error(err.Error())
		return
	}

	fight := fight.NewFight(hero, dragon)

	fight.PrintFightersInfo()

	err = fight.Start()
	if err != nil {
		log.Error(err.Error())
		return
	}

}

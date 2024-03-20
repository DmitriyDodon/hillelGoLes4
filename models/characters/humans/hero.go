package humans

import (
	"fmt"
	"les4/models/characters/reptiles"

	custommath "les4/customMath"
)

const heroMinStrength = 1
const heroMaxStrength = 5

type Hero struct {
	name    string
	stength int
}

func NewHero() (*Hero, error) {
	heroName, err := generateHeroName()
	if err != nil {
		return &Hero{}, err
	}
	return &Hero{
		name:    heroName,
		stength: 1,
	}, nil
}

func generateHeroName() (string, error) {
	firstNames := [10]string{"Вищий", "Відмінний", "Відбірний", "Найвидатніший", "Найдосконаліший", "Найжвящіший", "Найжвучніший", "Найжвантаженіший", "Найповніший", "Найенергійніший"}
	middleNames := [10]string{"Старий", "Стриманий", "Сивоволосий", "Строений", "Стійкий", "Червоний", "Здоровий", "Засмаглий", "Жвавий", "Сп'янілий"}
	patronymics := [10]string{"Камінь", "Камінчик", "Скеля", "Галька", "Коник", "Жеребець", "Кобила", "Конячка", "Кінниця", "Стовбур"}

	firstNameIndex, err := custommath.GenerateRandomNumber(9)
	if err != nil {
		return "", err
	}
	middleNameIndex, err := custommath.GenerateRandomNumber(9)
	if err != nil {
		return "", err
	}
	patronymicsIndex, err := custommath.GenerateRandomNumber(9)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(
		"%s %s %s",
		firstNames[firstNameIndex],
		middleNames[middleNameIndex],
		patronymics[patronymicsIndex],
	), nil
}

func (h Hero) AttackDragon(dragon *reptiles.Dragon) error {
	err := h.updateStrengthWithRandomNumber()
	if err != nil {
		return err
	}
	fmt.Printf("%s наносить удар по противнику та вiдрубає в нього %d голови.\n", h.name, h.stength)
	dragon.TakeDamage(h.stength)
	return nil
}

func (h *Hero) updateStrengthWithRandomNumber() error {
	heroStrength, err := custommath.GenerateRandomNumberByRange(heroMinStrength, heroMaxStrength)
	if err != nil {
		return err
	}
	h.stength = int(heroStrength)
	return nil
}

func (h Hero) LooseVoice() {
	panic("О нi, я програв i через це дракон буде й далi буде вбивати мирних людей.")
}

func (h Hero) WinVoice(dragon reptiles.Dragon) {
	fmt.Printf("Ураа! Я перемiг %s. З цього моменту свiту бiльше нiчого не загрожує!\n", dragon.Name())
}

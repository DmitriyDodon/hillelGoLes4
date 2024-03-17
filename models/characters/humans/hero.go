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

func NewHero() *Hero {
	return &Hero{
		name:    generateHeroName(),
		stength: 1,
	}
}

func generateHeroName() string {
	firstNames := [10]string{"Вищий", "Відмінний", "Відбірний", "Найвидатніший", "Найдосконаліший", "Найжвящіший", "Найжвучніший", "Найжвантаженіший", "Найповніший", "Найенергійніший"}
	middleNames := [10]string{"Старий", "Стриманий", "Сивоволосий", "Строений", "Стійкий", "Червоний", "Здоровий", "Засмаглий", "Жвавий", "Сп'янілий"}
	patronymics := [10]string{"Камінь", "Камінчик", "Скеля", "Галька", "Коник", "Жеребець", "Кобила", "Конячка", "Кінниця", "Стовбур"}

	return fmt.Sprintf(
		"%s %s %s",
		firstNames[custommath.GenerateRandomNumber(9)],
		middleNames[custommath.GenerateRandomNumber(9)],
		patronymics[custommath.GenerateRandomNumber(9)],
	)
}

func (h *Hero) AttackDragon(dragon *reptiles.Dragon) {
	h.updateStrengthWithRandomNumber()
	fmt.Printf("%s наносить удар по противнику та вiдрубає в нього %d голови.\n", h.name, h.stength)
	dragon.TakeDamage(h.stength)
}

func (h *Hero) updateStrengthWithRandomNumber() {
	h.stength = int(custommath.GenerateRandomNumberByRange(heroMinStrength, heroMaxStrength))
}

func (h *Hero) LooseVoice() {
	panic("О нi, я програв i через це дракон буде й далi буде вбивати мирних людей.")
}

func (h *Hero) WinVoice(dragon *reptiles.Dragon) {
	fmt.Printf("Ураа! Я перемiг %s. З цього моменту свiту бiльше нiчого не загрожує!\n", dragon.Name())
}

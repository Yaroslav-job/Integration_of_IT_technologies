package villages

import (
	"fmt"       // импорт пакета для форматированного ввода/вывода
	"math/rand" // импорт пакета для работы с случайными числами

	"example.com/City/city" // импорт модуля city
)

var villageList = []*city.Village{
	{
		Name:         "Крестьянская деревня", // название деревни
		RequiredArmy: 5,                      // необходимое количество воинов для захвата
		Bonus:        "+5 жителей за ход",    // бонус за захват
		BonusEffect: func(c *city.City) { // эффект бонуса
			c.Residents += 5 // увеличение количества жителей на 5
		},
		PlunderEffect: func(c *city.City) { // эффект разграбления
			c.Treasury += 50 // увеличение казны на 50 золота
		},
		Contested: false, // флаг, указывающий, что деревня не захвачена
	},
	{
		Name:         "Горная деревня",
		RequiredArmy: 10,
		Bonus:        "+5 золота за ход",
		BonusEffect: func(c *city.City) {
			c.Treasury += 5
		},
		PlunderEffect: func(c *city.City) {
			c.Treasury += 50
		},
		Contested: false,
	},
	{
		Name:         "Рыбацкая деревня",
		RequiredArmy: 8,
		Bonus:        "+3 счастья",
		BonusEffect: func(c *city.City) {
			c.Happiness += 3
			if c.Happiness > 10 {
				c.Happiness = 10
			}
		},
		PlunderEffect: func(c *city.City) {
			c.Treasury += 50
		},
		Contested: false,
	},
	{
		Name:         "Форт разбойников",
		RequiredArmy: 15,
		Bonus:        "+7 к армии, -2 счастья",
		BonusEffect: func(c *city.City) {
			c.Army += 7
			c.Happiness -= 2
		},
		PlunderEffect: func(c *city.City) {
			c.Treasury += 50
		},
		Contested: false,
	},
	{
		Name:         "Священная деревня",
		RequiredArmy: 12,
		Bonus:        "+2 счастья, -10% шанс эпидемии",
		BonusEffect: func(c *city.City) {
			c.Happiness += 2
			if c.Happiness > 10 {
				c.Happiness = 10
			}
			c.EpidemicRiskReduction *= 0.9
		},
		PlunderEffect: func(c *city.City) {
			c.Treasury += 50
		},
		Contested: false,
	},
	{
		Name:         "Торговый пост",
		RequiredArmy: 20,
		Bonus:        "+15 золота за ход",
		BonusEffect: func(c *city.City) {
			c.Treasury += 15
		},
		PlunderEffect: func(c *city.City) {
			c.Treasury += 50
		},
		Contested: false,
	},
	{
		Name:         "Военный лагерь",
		RequiredArmy: 25,
		Bonus:        "+10 к армии, -5 золота",
		BonusEffect: func(c *city.City) {
			c.Army += 10
			c.Treasury -= 5
		},
		PlunderEffect: func(c *city.City) {
			c.Treasury += 50
		},
		Contested: false,
	},
	{
		Name:         "Заброшенная крепость",
		RequiredArmy: 18,
		Bonus:        "+5 к армии, +3 счастья",
		BonusEffect: func(c *city.City) {
			c.Army += 5
			c.Happiness += 3
			if c.Happiness > 10 {
				c.Happiness = 10
			}
		},
		PlunderEffect: func(c *city.City) {
			c.Treasury += 50
		},
		Contested: false,
	},
	{
		Name:         "Деревня охотников",
		RequiredArmy: 10,
		Bonus:        "+3 золота, +2 жителей за ход",
		BonusEffect: func(c *city.City) {
			c.Treasury += 3
			c.Residents += 2
		},
		PlunderEffect: func(c *city.City) {
			c.Treasury += 50
		},
		Contested: false,
	},
	{
		Name:         "Работорговцы",
		RequiredArmy: 22,
		Bonus:        "+20 жителей, -5 счастья",
		BonusEffect: func(c *city.City) {
			c.Residents += 20
			c.Happiness -= 5
		},
		PlunderEffect: func(c *city.City) {
			c.Treasury += 50
		},
		Contested: false,
	},
}

func AttackVillageMenu(c *city.City) {
	fmt.Println("========================================")
	fmt.Println("Выберите деревню для атаки:")
	for i, village := range villageList {
		if !village.Contested {
			fmt.Printf("%d. %s (%d+ воинов) → %s\n", i+1, village.Name, village.RequiredArmy, village.Bonus)
		}
	}
	fmt.Println("11. Отмена")
	fmt.Print("Введите номер цели: ")
	var choice int
	fmt.Scan(&choice) // считывание выбора пользователя
	if choice < 1 || choice > len(villageList) {
		fmt.Println("❌ Неверный выбор, попробуйте снова. ❌")
		return
	}
	village := villageList[choice-1]
	if c.Army < village.RequiredArmy {
		fmt.Println("Недостаточно воинов для атаки.")
		return
	}
	attackVillage(c, village) // вызов функции атаки деревни
}

func attackVillage(c *city.City, village *city.Village) {
	fmt.Printf("Вы атакуете %s...\n", village.Name)
	var successChance int
	if c.Army >= village.RequiredArmy*2 {
		successChance = 80 // шанс успеха 80%, если армия в два раза больше необходимой
	} else if c.Army >= village.RequiredArmy {
		successChance = 50 // шанс успеха 50%, если армия равна необходимой
	} else {
		successChance = 20 // шанс успеха 20%, если армия меньше необходимой
	}
	if rand.Intn(100) < successChance {
		fmt.Println("🎉 Победа! Деревня теперь ваша!")
		fmt.Println("Выберите, что делать с деревней:")
		fmt.Println("1. Интегрировать (получать", village.Bonus, ")")
		fmt.Println("2. Разграбить (+50 золота, деревня сгорит)")
		fmt.Print("Введите номер действия: ")
		var action int
		fmt.Scan(&action) // считывание выбора пользователя
		if action == 1 {
			c.Villages[village.Name] = village // интеграция деревни
			village.Contested = true           // установка флага захвата
		} else if action == 2 {
			village.PlunderEffect(c) // эффект разграбления
		} else {
			fmt.Println("❌ Неверный выбор, попробуйте снова. ❌")
		}
		// Уменьшите количество воинов после успешной битвы
		c.Army -= village.RequiredArmy
		if c.Army < 0 {
			c.Army = 0
		}
	} else {
		fmt.Println("❌ Атака провалилась. Вы потеряли 30% армии.")
		c.Army = int(float64(c.Army) * 0.7) // уменьшение армии на 30%
		if c.Army < 0 {
			c.Army = 0
		}
	}
}

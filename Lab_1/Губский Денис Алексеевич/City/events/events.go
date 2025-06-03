package events

import (
	"fmt"       // импорт пакета для форматированного ввода/вывода
	"math/rand" // импорт пакета для работы с случайными числами
	"time"      // импорт пакета для работы со временем

	"example.com/City/city" // импорт модуля city
)

var lastFloodTurn int     // переменная для отслеживания последнего наводнения
var lastBadEventTurn int  // переменная для отслеживания последнего плохого события
var lastGoodEventTurn int // переменная для отслеживания последнего хорошего события

func init() {
	rand.Seed(time.Now().UnixNano()) // инициализация генератора случайных чисел
}

func TriggerEvent(c *city.City, currentTurn int) {
	if c.Happiness <= 0 || c.Treasury <= 0 {
		fmt.Println("========================================")
		fmt.Println("💀 Событие: Бунт! Горожане восстали из-за низкого уровня счастья или пустой казны.")
		// Логика бунта
		c.Residents -= 10              // уменьшение количества жителей
		c.Happiness = 0                // установка уровня счастья в 0
		c.Treasury = 0                 // установка казны в 0
		lastBadEventTurn = currentTurn // обновление последнего хода плохого события
		fmt.Println("========================================")
		fmt.Println("💀 Игра окончена! 💀")
		c.Residents = 0
		c.Treasury = 0
		return
	}

	event := rand.Intn(8) // увеличение диапазона для новых событий
	fmt.Println("========================================")

	// Проверка на частоту плохих и хороших событий
	if event <= 2 && currentTurn-lastBadEventTurn < 4 {
		fmt.Println("========================================")
		return
	}
	if event > 2 && currentTurn-lastGoodEventTurn < 2 {
		fmt.Println("========================================")
		return
	}
	fmt.Println("📅 Событие произошло!")
	switch event {
	case 0:
		fmt.Println("🦠 Событие: Эпидемия! Много жителей умерло.")
		deathRate := 0.6
		if c.Buildings["Лазарет"] > 0 {
			deathRate -= 0.3 // уменьшение смертности на 30% за первый лазарет
			if c.Buildings["Лазарет"] > 1 {
				deathRate -= 0.1 // уменьшение смертности на 10% за второй лазарет
			}
		}
		if deathRate < 0 {
			deathRate = 0
		}
		deaths := int(c.Residents * deathRate)
		c.Residents *= (1 - deathRate)
		fmt.Printf("Умерло жителей: %d\n", deaths)
		c.CheckCityLevel()             // проверка уровня города после изменения количества жителей
		lastBadEventTurn = currentTurn // обновление последнего хода плохого события
	case 1:
		if rand.Float64() > c.DefenseChance || c.Army == 0 {
			fmt.Println("⚔️ Событие: Варвары напали! -1 здание, -10% жителей.")
			deaths := int(c.Residents * 0.1)
			warriorDeaths := 0
			if c.Army > 0 {
				warriorDeaths = rand.Intn(c.Army) + 1
				if warriorDeaths > deaths {
					warriorDeaths = deaths
				}
				c.Army -= warriorDeaths
				deaths -= warriorDeaths
			}
			c.Residents *= (1 - float64(deaths)/c.Residents)
			fmt.Printf("Умерло жителей: %d, воинов: %d\n", deaths, warriorDeaths)
			c.CheckCityLevel() // проверка уровня города после изменения количества жителей
			for building := range c.Buildings {
				if building != "Дом Воеводы" {
					fmt.Printf("Уничтожено здание: %s\n", building)
					c.DestroyBuilding(building) // использование метода DestroyBuilding
					break
				}
			}
			lastBadEventTurn = currentTurn // обновление последнего хода плохого события
		} else {
			fmt.Println("⚔️ Событие: Варвары напали, но были отбиты!")
			// Рассчитайте количество погибших воинов в зависимости от сложности варваров
			warriorDeaths := rand.Intn(10) + 1 // случайное количество погибших воинов от 1 до 10
			if warriorDeaths > c.Army {
				warriorDeaths = c.Army
			}
			c.Army -= warriorDeaths
			fmt.Printf("Погибло воинов: %d\n", warriorDeaths)
		}
		c.UpdateIncome() // перерасчет дохода после гибели воинов
	case 2:
		if currentTurn-lastFloodTurn >= 8 {
			fmt.Println("🌊 Событие: Наводнение! -1 здание.")
			for building := range c.Buildings {
				if building != "Дом Воеводы" {
					fmt.Printf("Уничтожено здание: %s\n", building)
					c.DestroyBuilding(building) // использование метода DestroyBuilding
					break
				}
			}
			lastFloodTurn = currentTurn    // обновление последнего хода наводнения
			lastBadEventTurn = currentTurn // обновление последнего хода плохого события
		}
	case 3:
		fmt.Println("🌾 Событие: Хороший урожай! +15-25% жителей.")
		increaseRate := 0.15 + rand.Float64()*0.1
		c.Residents *= (1 + increaseRate)
		c.CheckCityLevel()              // проверка уровня города после изменения количества жителей
		lastGoodEventTurn = currentTurn // обновление последнего хода хорошего события
	case 4:
		fmt.Println("🛍️ Событие: Торговцы привезли товары! +10-25% золота.")
		increaseRate := 0.1 + rand.Float64()*0.15
		c.Treasury += int(float64(c.Treasury) * increaseRate)
		lastGoodEventTurn = currentTurn // обновление последнего хода хорошего события
	case 5:
		fmt.Println("🏗️ Событие: Странник построил здание бесплатно.")
		buildings := []string{"Ферма", "Рынок", "Лазарет", "Трактир", "Стены", "Башня", "Шахта", "Пристань", "Набережная", "Церковь", "Дом лекаря"}
		randomBuilding := buildings[rand.Intn(len(buildings))]
		c.Buildings[randomBuilding]++
		if randomBuilding == "Дом лекаря" {
			c.PopulationGrowth += 1 // увеличение рождаемости при постройке дома лекаря
		}
		fmt.Printf("Странник построил: %s\n", randomBuilding)
		lastGoodEventTurn = currentTurn // обновление последнего хода хорошего события
	case 6:
		fmt.Println("🔥 Событие: Пожар! -1 здание, -5% жителей.")
		deaths := int(c.Residents * 0.05)
		c.Residents *= (1 - 0.05)
		fmt.Printf("Умерло жителей: %d\n", deaths)
		for building := range c.Buildings {
			if building != "Дом Воеводы" {
				fmt.Printf("Уничтожено здание: %s\n", building)
				c.DestroyBuilding(building) // использование метода DestroyBuilding
				break
			}
		}
		lastBadEventTurn = currentTurn // обновление последнего хода плохого события
	case 7:
		fmt.Println("🎉 Событие: Праздник! +5 счастья, -10 золота.")
		c.Happiness += 5
		if c.Happiness > 10 {
			c.Happiness = 10
		}
		c.Treasury -= 10
		lastGoodEventTurn = currentTurn // обновление последнего хода хорошего события
	}

	// Обновление показателей после события
	c.UpdateIncome()
	c.UpdateLevel()
	c.CheckCityLevel()
	c.UpdateArmy()
	c.UpdatePopulationGrowth()
	c.UpdateDefenseChance() // обновление шанса нападения варваров
	c.UpdateEpidemicRisk()  // обновление риска эпидемий
	c.UpdateFloodRisk()     // обновление риска наводнения

	// Проверка на окончание игры
	if c.Residents <= 0 {
		fmt.Println("💀 Игра окончена! Все жители погибли. 💀")
		c.Residents = 0
		c.Treasury = 0
	} else if c.Treasury < 0 {
		fmt.Println("💀 Игра окончена! Казна пуста. 💀")
		c.Residents = 0
		c.Treasury = 0
	}
	fmt.Println("========================================")
}

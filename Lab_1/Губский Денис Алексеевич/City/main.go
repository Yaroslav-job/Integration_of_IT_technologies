package main

import (
	"fmt"       // импорт пакета для форматированного ввода/вывода
	"math/rand" // импорт пакета для работы с случайными числами
	"os"        // импорт пакета для работы с операционной системой
	"time"      // импорт пакета для работы со временем

	"example.com/City/city"     // импорт модуля city
	"example.com/City/council"  // импорт модуля council
	"example.com/City/events"   // импорт модуля events
	"example.com/City/villages" // импорт модуля villages
)

func main() {
	rand.Seed(time.Now().UnixNano()) // инициализация генератора случайных чисел
	c := city.NewCity()              // создание нового города
	turns := 0                       // инициализация счетчика ходов
	for {
		c.DisplayState() // отображение текущего состояния города
		fmt.Println("🏗️ 1. Построить здание")
		fmt.Println("⚔️ 2. Нанять воинов")
		fmt.Println("🏘️ 3. Захватить деревню") // добавьте новый пункт меню
		fmt.Println("⏩ 4. Завершить ход")
		fmt.Println("🚪 5. Выход")
		fmt.Print("Выберите действие: ")
		var choice int
		fmt.Scan(&choice) // считывание выбора пользователя
		switch choice {
		case 1:
			buildingMenu(c) // вызов меню строительства
		case 2:
			hireArmy(c) // вызов функции найма армии
		case 3:
			villages.AttackVillageMenu(c) // вызов меню захвата деревень
		case 4:
			turns++     // увеличение счетчика ходов
			c.EndTurn() // завершение текущего хода
			if turns%5 == 0 {
				council.TriggerCouncil(c)  // вызов совета каждые 5 ходов
				c.UpdateIncome()           // обновление дохода
				c.UpdateLevel()            // обновление уровня города
				c.CheckCityLevel()         // проверка уровня города
				c.UpdateArmy()             // обновление армии
				c.UpdatePopulationGrowth() // обновление роста населения
			}
			if turns >= 6 && rand.Intn(6)+6 <= turns {
				events.TriggerEvent(c, turns) // вызов случайного события
				c.UpdateIncome()              // обновление дохода
				c.UpdateLevel()               // обновление уровня города
				c.CheckCityLevel()            // проверка уровня города
				c.UpdateArmy()                // обновление армии
				c.UpdatePopulationGrowth()    // обновление роста населения
				if c.Residents <= 0 || c.Treasury < 0 {
					fmt.Println("💀 Игра окончена! 💀")
					os.Exit(0) // завершение игры, если жители или казна в минусе
				}
			}
			if c.Treasury < 0 {
				events.TriggerEvent(c, turns) // вызов события бунта при отрицательной казне
				c.UpdateIncome()              // обновление дохода
				c.UpdateLevel()               // обновление уровня города
				c.CheckCityLevel()            // проверка уровня города
				c.UpdateArmy()                // обновление армии
				c.UpdatePopulationGrowth()    // обновление роста населения
				if c.Residents <= 0 || c.Treasury < 0 {
					fmt.Println("💀 Игра окончена! 💀")
					os.Exit(0) // завершение игры, если жители или казна в минусе
				}
			}
			if c.Happiness <= 0 {
				events.TriggerEvent(c, turns) // вызов события бунта при нулевом или отрицательном уровне счастья
				c.UpdateIncome()              // обновление дохода
				c.UpdateLevel()               // обновление уровня города
				c.CheckCityLevel()            // проверка уровня города
				c.UpdateArmy()                // обновление армии
				c.UpdatePopulationGrowth()    // обновление роста населения
				if c.Residents <= 0 || c.Treasury < 0 {
					fmt.Println("💀 Игра окончена! 💀")
					os.Exit(0) // завершение игры, если жители или казна в минусе
				}
			}
			c.UpdateLevel()  // обновление уровня города
			c.UpdateIncome() // обновление дохода
			if c.IsGameOver() {
				if c.Residents <= 0 {
					fmt.Println("💀 Игра окончена! Все жители погибли. 💀")
				} else if c.Treasury < 0 {
					fmt.Println("💀 Игра окончена! Казна пуста. 💀")
				}
				os.Exit(0) // завершение игры, если игра окончена
			}
		case 5:
			return // выход из игры
		default:
			fmt.Println("❌ Неверный выбор, попробуйте снова. ❌")
		}
		fmt.Println("========================================")
		fmt.Println(" ")
	}
}

func buildingMenu(c *city.City) {
	fmt.Println("🏢 Доступные здания:")
	fmt.Println("1. 🌾 Ферма (30 золота) [+5 доход, +5% к риску эпидемии]")
	fmt.Println("2. 🏪 Рынок (50 золота) [+10 доход, повышает шанс торговцев]")
	fmt.Println("3. 🏥 Лазарет (70 золота) [-4 доход, снижает смертность от эпидемии]")
	fmt.Println("4. 🍻 Трактир (40 золота) [+3 счастья, -3 доход]")
	fmt.Println("5. 🛡️ Стены (60 золота) [-10% шанс нападения варваров]")
	fmt.Println("6. 🏰 Башня (80 золота) [-20% шанс нападения варваров]")
	fmt.Println("7. 🏋️ Казарма (100 золота) [-5 доход, 1 воин за каждых 3 жителей]")
	fmt.Println("8. ⛏️ Шахта (90 золота) [+8 доход, -3 счастья, +10% риск эпидемии]")
	fmt.Println("9. 🚢 Пристань (120 золота) [+12 доход, +10% шанс торговцев/странников]")
	fmt.Println("10. 🌊 Набережная (50 золота) [-2 доход, +4 счастья]")
	fmt.Println("11. ⛪ Церковь (70 золота) [+15% к доходу, +5 счастья]")
	fmt.Println("12. 🏡 Дом лекаря (110 золота) [-9 доход, +3 счастья, +1 к рождаемости]")
	fmt.Println("13. 🏛️ Академия (150 золота) [+20 доход, +5 счастья, +2 к рождаемости]")
	fmt.Println("14. 🏭 Фабрика (200 золота) [+30 доход, -5 счастья, +10% риск эпидемии]")
	fmt.Println("15. 🔙 Назад")
	fmt.Print("Выберите здание для постройки: ")
	var choice int
	fmt.Scan(&choice) // считывание выбора пользователя
	switch choice {
	case 1:
		if c.Treasury >= 30 {
			c.Treasury -= 30       // уменьшение казны на 30 золота
			c.Income += 5.0        // увеличение дохода на 5
			c.Buildings["Ферма"]++ // увеличение количества ферм
		} else {
			fmt.Println("Недостаточно золота.")
		}
	case 2:
		if c.Treasury >= 50 {
			c.Treasury -= 50       // уменьшение казны на 50 золота
			c.Income += 10.0       // увеличение дохода на 10
			c.Buildings["Рынок"]++ // увеличение количества рынков
		} else {
			fmt.Println("Недостаточно золота.")
		}
	case 3:
		if c.Treasury >= 70 {
			if c.Buildings["Лазарет"] < 2 {
				c.Treasury -= 70         // уменьшение казны на 70 золота
				c.Income -= 4.0          // уменьшение дохода на 4
				c.Buildings["Лазарет"]++ // увеличение количества лазаретов
			} else {
				fmt.Println("Можно построить не более двух лазаретов.")
			}
		} else {
			fmt.Println("Недостаточно золота.")
		}
	case 4:
		if c.Treasury >= 40 {
			c.Treasury -= 40 // уменьшение казны на 40 золота
			c.Happiness += 3 // увеличение счастья на 3
			if c.Happiness > 10 {
				c.Happiness = 10 // ограничение счастья до 10
			}
			c.Income -= 3.0          // уменьшение дохода на 3
			c.Buildings["Трактир"]++ // увеличение количества трактиров
		} else {
			fmt.Println("Недостаточно золота.")
		}
	case 5:
		if c.Treasury >= 60 {
			c.Treasury -= 60       // уменьшение казны на 60 золота
			c.Defense += 10        // увеличение защиты на 10
			c.Buildings["Стены"]++ // увеличение количества стен
		} else {
			fmt.Println("Недостаточно золота.")
		}
	case 6:
		if c.Treasury >= 80 {
			c.Treasury -= 80       // уменьшение казны на 80 золота
			c.Defense += 20        // увеличение защиты на 20
			c.Buildings["Башня"]++ // увеличение количества башен
		} else {
			fmt.Println("Недостаточно золота.")
		}
	case 7:
		if c.Treasury >= 100 {
			c.Treasury -= 100           // уменьшение казны на 100 золота
			c.Income -= 5.0             // уменьшение дохода на 5
			c.Buildings["Казарма"]++    // увеличение количества казарм
			warriors := c.Residents / 5 // 20% жителей становятся воинами
			c.Army += int(warriors)     // увеличение армии
			c.Residents -= warriors     // уменьшение количества жителей
		} else {
			fmt.Println("Недостаточно золота.")
		}
	case 8:
		if c.Treasury >= 90 {
			c.Treasury -= 90       // уменьшение казны на 90 золота
			c.Income += 8.0        // увеличение дохода на 8
			c.Happiness -= 3       // уменьшение счастья на 3
			c.Buildings["Шахта"]++ // увеличение количества шахт
		} else {
			fmt.Println("Недостаточно золота.")
		}
	case 9:
		if c.Treasury >= 120 {
			c.Treasury -= 120         // уменьшение казны на 120 золота
			c.Income += 12.0          // увеличение дохода на 12
			c.Buildings["Пристань"]++ // увеличение количества пристаней
		} else {
			fmt.Println("Недостаточно золота.")
		}
	case 10:
		if c.Treasury >= 50 {
			c.Treasury -= 50 // уменьшение казны на 50 золота
			c.Income -= 2.0  // уменьшение дохода на 2
			c.Happiness += 4 // увеличение счастья на 4
			if c.Happiness > 10 {
				c.Happiness = 10 // ограничение счастья до 10
			}
			c.Buildings["Набережная"]++ // увеличение количества набережных
		} else {
			fmt.Println("Недостаточно золота.")
		}
	case 11:
		if c.Treasury >= 70 {
			c.Treasury -= 70           // уменьшение казны на 70 золота
			c.Income = c.Income * 1.15 // увеличение дохода на 15%
			c.Happiness += 5           // увеличение счастья на 5
			if c.Happiness > 10 {
				c.Happiness = 10 // ограничение счастья до 10
			}
			c.Buildings["Церковь"]++ // увеличение количества церквей
		} else {
			fmt.Println("Недостаточно золота.")
		}
	case 12:
		if c.Treasury >= 110 {
			c.Treasury -= 110 // уменьшение казны на 110 золота
			c.Income -= 9.0   // уменьшение дохода на 9
			c.Happiness += 3  // увеличение счастья на 3
			if c.Happiness > 10 {
				c.Happiness = 10 // ограничение счастья до 10
			}
			c.PopulationGrowth += 1     // увеличение рождаемости на 1
			c.Buildings["Дом лекаря"]++ // увеличение количества домов лекарей
		} else {
			fmt.Println("Недостаточно золота.")
		}
	case 13:
		if c.Treasury >= 150 {
			c.Treasury -= 150 // уменьшение казны на 150 золота
			c.Income += 20.0  // увеличение дохода на 20
			c.Happiness += 5  // увеличение счастья на 5
			if c.Happiness > 10 {
				c.Happiness = 10 // ограничение счастья до 10
			}
			c.PopulationGrowth += 2   // увеличение рождаемости на 2
			c.Buildings["Академия"]++ // увеличение количества академий
		} else {
			fmt.Println("Недостаточно золота.")
		}
	case 14:
		if c.Treasury >= 200 {
			c.Treasury -= 200        // уменьшение казны на 200 золота
			c.Income += 30.0         // увеличение дохода на 30
			c.Happiness -= 5         // уменьшение счастья на 5
			c.Buildings["Фабрика"]++ // увеличение количества фабрик
		} else {
			fmt.Println("Недостаточно золота.")
		}
	case 15:
		return // выход из меню строительства
	default:
		fmt.Println("❌ Неверный выбор, попробуйте снова. ❌")
	}
}

func hireArmy(c *city.City) {
	fmt.Printf("Сейчас в городе: %d воинов\n", c.Army)
	fmt.Println("Стоимость одного воина: 10 золота")
	fmt.Printf("Максимально доступно: %d (жителей: %d)\n", int(c.Residents/3), int(c.Residents))
	fmt.Print("Введите количество воинов для найма (или 0 для отмены): ")
	var count int
	fmt.Scan(&count) // считывание количества воинов для найма
	if count > 0 && c.Treasury >= count*10 {
		c.Treasury -= count * 10      // уменьшение казны на стоимость воинов
		c.Army += count               // увеличение армии на количество воинов
		c.Residents -= float64(count) // уменьшение количества жителей
		c.UpdateIncome()              // обновление дохода
		c.CheckCityLevel()            // проверка уровня города
	} else if count > 0 {
		fmt.Println("Недостаточно золота.")
	}
}

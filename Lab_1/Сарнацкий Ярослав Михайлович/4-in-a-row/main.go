package main

import (
	"fmt"
)

func main() {
	var choise int

	clear()

	fmt.Println("Выберите тип игры: \n1 - Player vs Player \n2 - Player vs Bot")
	for {
		fmt.Scan(&choise)
		if choise == 1 || choise == 2 {
			break
		} else {
			fmt.Println("Вы ввели недопустимое значение")
		}
	}

	board_start()

	if choise == 1 {
		game("player")
	} else if choise == 2 {
		game("bot")
	}
}

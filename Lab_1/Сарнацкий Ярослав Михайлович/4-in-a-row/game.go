package main

import (
	"fmt"
	"os"
)

func game(type_game string) {
	for {
		input("[🟢]")
		board_print()
		check_win("player1", "[🟢]")

		if type_game == "player" {
			input("[🔴]")
			board_print()
			check_win("player_2", "[🔴]")
		} else if type_game == "bot" {
			bot("[🔴]")
			board_print()
			check_win("bot", "[🔴]")
		}
	}
}

func game_over() {
	var expectation int
	fmt.Println("ВВедите: \n1 - чтобы начать новую игру \n2 - чтобы выйти из игры")
	for {
		fmt.Scan(&expectation)
		if expectation == 1 {
			main()
		} else if expectation == 2 {
			os.Exit(0)
		} else {
			fmt.Printf("Вы ввели недопустимое значение. Попробуйте ещё раз:")
		}
	}
}

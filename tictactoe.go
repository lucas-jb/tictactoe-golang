package main

import "fmt"

func main() {
	posiciones := [9]string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	jugador := "X"
	state := true
	for state {
		menu(posiciones)
		read(&posiciones, jugador)
		cambio(&jugador)
		comprobar(posiciones, &state, &jugador)
	}
	fmt.Print("El jugador " + jugador + " ha ganado la partida!")
}

func read(posiciones *[9]string, jugador string) {
	fmt.Print(jugador + " introduce la posicion:")
	var first int
	fmt.Scanln(&first)
	first = first - 1
	if first >= 0 && first <= 8 {
		if posiciones[first] != "X" && posiciones[first] != "O" {
			posiciones[first] = jugador
		} else {
			read(posiciones, jugador)
		}
	}
}

func menu(posiciones [9]string) {
	fmt.Println("[" + posiciones[0] + "] [" + posiciones[1] + "] [" + posiciones[2] + "]")
	fmt.Println("[" + posiciones[3] + "] [" + posiciones[4] + "] [" + posiciones[5] + "]")
	fmt.Println("[" + posiciones[6] + "] [" + posiciones[7] + "] [" + posiciones[8] + "]")
}

func cambio(jugador *string) {
	if *jugador == "X" {
		*jugador = "O"
	} else {
		*jugador = "X"
	}
}

func comprobar(posiciones [9]string, state *bool, jugador *string) {
	//1 2 3
	if posiciones[0] == posiciones[1] && posiciones[1] == posiciones[2] {
		*state = false
		*jugador = posiciones[0]
	}
	// 4 5 6
	if posiciones[3] == posiciones[4] && posiciones[4] == posiciones[5] {
		*state = false
		*jugador = posiciones[3]
	}
	// 7 8 9
	if posiciones[6] == posiciones[7] && posiciones[7] == posiciones[8] {
		*state = false
		*jugador = posiciones[6]
	}
	// 1 4 7
	if posiciones[0] == posiciones[3] && posiciones[3] == posiciones[6] {
		*state = false
		*jugador = posiciones[0]
	}
	// 2 5 8
	if posiciones[1] == posiciones[4] && posiciones[4] == posiciones[7] {
		*state = false
		*jugador = posiciones[1]
	}
	// 3 6 9
	if posiciones[2] == posiciones[5] && posiciones[5] == posiciones[8] {
		*state = false
		*jugador = posiciones[2]
	}
	// 1 5 9
	if posiciones[0] == posiciones[4] && posiciones[4] == posiciones[8] {
		*state = false
		*jugador = posiciones[0]
	}
	// 3 5 7
	if posiciones[2] == posiciones[4] && posiciones[4] == posiciones[6] {
		*state = false
		*jugador = posiciones[2]
	}
}

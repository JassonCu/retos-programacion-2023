package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type Room struct {
	enigma   string
	hasGhost bool
	hasDoor  bool
	hasCandy bool
}

type Mansion struct {
	rooms     [4][4]Room
	playerRow int
	playerCol int
}

func main() {
	rand.Seed(time.Now().UnixNano())

	mansion := initializeMansion()
	fmt.Println("Bienvenido a la mansión abandonada. Tu misión es encontrar la habitación de los dulces.")

	for {
		currentRoom := &mansion.rooms[mansion.playerRow][mansion.playerCol]

		if currentRoom.hasDoor {
			fmt.Println("¡Has encontrado una puerta! Deberás responder el siguiente enigma para avanzar.")
			askEnigma(currentRoom.enigma)

			if currentRoom.hasGhost && rand.Float32() < 0.1 {
				fmt.Println("¡Oh no! Un fantasma apareció. Deberás responder dos preguntas para salir de esta habitación.")
				askEnigma(generateRandomEnigma())
				askEnigma(generateRandomEnigma())
			}

			moveDirection := askDirection()
			if moveDirection == "salir" {
				if currentRoom.hasCandy {
					fmt.Println("¡Felicidades! ¡Has encontrado la habitación de los dulces! Has ganado el juego.")
					break
				} else {
					fmt.Println("Esta no es la habitación de los dulces. Sigue buscando.")
				}
			} else {
				movePlayer(&mansion, moveDirection)
			}
		} else {
			fmt.Println("¡Has encontrado la habitación de los dulces! ¡Felicidades! Has ganado el juego.")
			break
		}
	}
}

func initializeMansion() Mansion {
	mansion := Mansion{}

	// Inicializar las habitaciones con enigmas y puertas
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			mansion.rooms[i][j] = Room{
				enigma:   generateRandomEnigma(),
				hasDoor:  rand.Float32() < 0.8, // 80% de probabilidad de tener puerta
				hasCandy: false,
				hasGhost: rand.Float32() < 0.1, // 10% de probabilidad de tener fantasma
			}
		}
	}

	// Colocar la puerta de entrada y la habitación de los dulces
	mansion.rooms[0][0].hasDoor = true
	mansion.rooms[3][2].hasCandy = true

	return mansion
}

func askEnigma(enigma string) {
	var userAnswer string
	fmt.Println("Enigma:", enigma)
	fmt.Print("Tu respuesta: ")
	fmt.Scanln(&userAnswer)

	if !checkAnswer(enigma, userAnswer) {
		fmt.Println("Respuesta incorrecta. No puedes avanzar hasta resolver el enigma.")
		askEnigma(enigma)
	}
}

func askDirection() string {
	var direction string
	fmt.Print("¿Hacia dónde quieres ir? (norte/sur/este/oeste/salir): ")
	fmt.Scanln(&direction)
	return direction
}

func movePlayer(mansion *Mansion, direction string) {
	switch direction {
	case "norte":
		if mansion.playerRow > 0 {
			mansion.playerRow--
		} else {
			fmt.Println("No puedes ir al norte desde aquí.")
		}
	case "sur":
		if mansion.playerRow < 3 {
			mansion.playerRow++
		} else {
			fmt.Println("No puedes ir al sur desde aquí.")
		}
	case "este":
		if mansion.playerCol < 3 {
			mansion.playerCol++
		} else {
			fmt.Println("No puedes ir al este desde aquí.")
		}
	case "oeste":
		if mansion.playerCol > 0 {
			mansion.playerCol--
		} else {
			fmt.Println("No puedes ir al oeste desde aquí.")
		}
	default:
		fmt.Println("Dirección no válida. Intenta de nuevo.")
	}
}

func generateRandomEnigma() string {
	enigmas := []string{
		"¿Qué tiene ojos y no puede ver?",
		"Cuanto más lo quitan, más grande se vuelve. ¿Qué es?",
		"Es más poderoso que Dios. Es más malvado que el diablo. Los ricos lo necesitan, los pobres lo tienen. Y si lo comes, morirás. ¿Qué es?",
		"¿Qué tiene llaves pero no abre cerraduras?",
	}

	return enigmas[rand.Intn(len(enigmas))]
}

func checkAnswer(correctAnswer, userAnswer string) bool {
	return normalizeString(correctAnswer) == normalizeString(userAnswer)
}

func normalizeString(s string) string {
	// Convertir a minúsculas y eliminar espacios al inicio y al final
	return strings.ToLower(strings.TrimSpace(s))
}
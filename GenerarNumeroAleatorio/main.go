package main

import (
	"fmt"
	"math/rand"
	//"time"
)

func main() {
	//rand.Seed(time.Now().UnixNano())
	//fmt.Println(rand.Intn(100))
	jugar()
}

func jugar() {
	numAleatorio := rand.Intn(100)
	var numIngresado int
	var intentos int
	const maxIntentos = 10

	for intentos < maxIntentos {
		intentos++
		fmt.Printf("Ingrese un número (intentos restantes: %d): ", maxIntentos-intentos+1)
		fmt.Scanln(&numIngresado)

		if numIngresado == numAleatorio {
			fmt.Println("Felicitaciones, adivinastes el número!")
			jugarNuevamente()
			return
		} else if numIngresado < numAleatorio {
			fmt.Println("El número a adivinar es mayor.")
		} else if numIngresado > numAleatorio {
			fmt.Println("El número a adivinar es menor.")
		}
	}

	fmt.Println("Se acabaron los intentos. El número era: ", numAleatorio)
	jugarNuevamente()
}

func jugarNuevamente() {
	var eleccion string
	fmt.Println("¿Quieres jugar nuevamente? (s/n): ")
	fmt.Scanln(&eleccion)
	switch eleccion {
	case "s":
		jugar()

	case "n":
		fmt.Println("Gracias por jugar!")

	default:
		fmt.Println("Elección inválida, Inténtelo nuevamente")
		jugarNuevamente()
	}
}

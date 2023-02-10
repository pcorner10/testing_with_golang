package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// print a welcome message
	intro()

	// Crear un canal para indicar cuando el usuario quiere salir
	doneChan := make(chan bool)

	// Iniciar una goroutine para leer el input del usuario y correr el programa
	go readUserInput(doneChan)

	// bloquear hasta que doneChan reciba un valor
	<-doneChan

	// cerramos el canal
	close(doneChan)

	// despedir el programa
	fmt.Println("Goodbay!.")

}

func readUserInput(doneChan chan bool) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		res, done := checkNumbers(scanner)

		if done {
			doneChan <- true
			return
		}

		fmt.Println(res)
		prompt()

	}
}

func checkNumbers(scanner *bufio.Scanner) (string, bool) {
	// leer el input del usuario
	scanner.Scan()

	// checar si el usuario quiere salir
	if strings.EqualFold(scanner.Text(), "q") {
		return "", true
	}

	// intentamos convertir el input a entero
	numToCheck, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return "Por favor ingresa un numero entero", false
	}

	_, msg := isPrime(numToCheck)
	return msg, false

}

func intro()  {
	fmt.Println("Is it Prime?")
	fmt.Println("------------")
	fmt.Println("Introduce un número y te diremos si es un primo o no. Usa 'q' para salir.")
	prompt()
}

func prompt() {
	fmt.Print("-> ")
}

func isPrime(n int) (bool, string) {
	// 0 an 1 are not prime by definition
	if n == 0 || n == 1 {
		return false, fmt.Sprintf("%d is not prime, by definition!", n)
	}

	// negative numbers are not prime
	if n < 0 {
		return false, "Negative numbers are not prime, by definition!"
	}

	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			// not a prime number
			return false, fmt.Sprintf("%d is not a primer number because it is divisible by %d", n, i)
		}
	}

	return true, fmt.Sprintf("%d is a prime number!", n)
}
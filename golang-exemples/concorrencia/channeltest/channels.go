package main

import "fmt"

func TurnOn(status string, c chan string) {
	c <- status
}

func MoveOn(status string, c chan string) {
	c <- status
}

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	// PROCESSAMENTO CONCORRENTE ... PROCESSANDO OS DUAS FUNÇÕES AO MESMO TEMPO
	go TurnOn("ligar", c1)
	go MoveOn("Andar", c2)

	fmt.Println(<-c1)
	fmt.Println(<-c2)
}

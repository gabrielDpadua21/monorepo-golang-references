package main

import (
	"runtime/debug"
)

// Chama a funcao stack
func f3() {
	debug.PrintStack()
}

// Chama a funcao 3
func f2() {
	f3()
}

// Chama a funcao 2
func f1() {
	f2()
}

// Chama a funcao 1
func main() {
	f1()
}

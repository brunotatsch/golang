package main

import "fmt"

func main() {
	fmt.Println("Hello, World!!!")
}

// Execute ontime
// # go run main.go

// Create executable
// # go build main.go

// Create executable with name
// # go build -o "hello" main.go

// Create executable with crosscompilation (to another s.o)
// # go build -o "hello" main.go

// A Primeira letra do nome indica se o objeto é publico ou privado, maiusculas publicas e minusculas privadas
// Func foo() {} função privada
// Func Foo() {} função Publica

package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Jogo da Adivinhação")
	fmt.Println("Um número aleatório será sorteado, sua missão é tentar acertar ele!")

	scanner := bufio.NewScanner(os.Stdin)

	println(scanner)
	chutes := [10]int64{}

	x := rand.Int64N(101)

	for i := range chutes {
		fmt.Println("Qual o seu palpite?")
		scanner.Scan()

		chute := scanner.Text()
		chute = strings.TrimSpace(chute)
		chuteInt, err := strconv.ParseInt(chute, 10, 64)

		if err != nil {
			fmt.Println("informe apenas números inteiros!")
			return
		}

		switch {
		case chuteInt < x:
			fmt.Println("Você errou! O número é maior que", chuteInt)
		case chuteInt > x:
			fmt.Println("Você errou! O número é menor que", chuteInt)
		case chuteInt == x:
			fmt.Printf("Parabéns! Você acertou !!!! O número que era: %d\n"+
				"Voce teve %d tentativa(s).\n"+
				"Estas foram as suas tentativas: %v\n", x, i+1, chutes[:i])
			return
		}

		chutes[i] = chuteInt

	}

	fmt.Printf("Infelizmente suas chances acabaram e você não acertou o número que era: %d\n"+
		"Voce teve 10 tentativas.\n"+
		"Estas foram as suas tentativas: %v\n", x, chutes)
}

package main

import (
	"fmt"

	"github.com/curso-fullcycle-3-0/internal"
)

func main() {
	soma := internal.Calc(10, 10)
	fmt.Println("Calculo realizado com sucesso")
	fmt.Printf("A soma é %d", soma)
}

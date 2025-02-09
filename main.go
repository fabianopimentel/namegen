package main

import (
	"fmt"
	"log"

	"github.com/fabianopimentel/namegen/internal/generator"
)

func main() {
	name, err := generator.GenerateName()
	if err != nil {
		log.Fatalf("Erro ao gerar nome: %v", err)
	}

	fmt.Println(name)
}

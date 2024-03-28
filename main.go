package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fabianopimentel/namegen/dictionaries"
)

func main() {
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)

	adjectivesIndex := rng.Intn(len(dictionaries.Adjectives))
	nameIndex := rng.Intn(len(dictionaries.Names))

	fmt.Printf("%s-%s\n", dictionaries.Adjectives[adjectivesIndex], dictionaries.Names[nameIndex])
}

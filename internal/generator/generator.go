package generator

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fabianopimentel/namegen/internal/dictionary"
)

func GenerateName() (string, error) {
	adjectives, err := dictionary.LoadAllFiles(dictionary.Adjectives, "resources/adjectives")
	if err != nil {
		return "", fmt.Errorf("erro ao carregar adjetivos: %w", err)
	}

	names, err := dictionary.LoadAllFiles(dictionary.Names, "resources/names")
	if err != nil {
		return "", fmt.Errorf("erro ao carregar nomes: %w", err)
	}

	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)

	adjectivesIndex := rng.Intn(len(adjectives))
	nameIndex := rng.Intn(len(names))

	return fmt.Sprintf("%s-%s", adjectives[adjectivesIndex], names[nameIndex]), nil
}

package dictionary

import (
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
)

//go:embed resources/adjectives/*.json
var Adjectives embed.FS

//go:embed resources/names/*.json
var Names embed.FS

func LoadAllFiles(files embed.FS, dirPath string) ([]string, error) {
	var combined []string

	dirEntries, err := fs.ReadDir(files, dirPath)
	if err != nil {
		return nil, fmt.Errorf("erro ao ler o diretório embutido: %w", err)
	}

	for _, entry := range dirEntries {
		data, err := files.ReadFile(fmt.Sprintf("%s/%s", dirPath, entry.Name()))
		if err != nil {
			return nil, fmt.Errorf("erro ao ler o arquivo %s: %w", entry.Name(), err)
		}

		var words []string
		if err := json.Unmarshal(data, &words); err != nil {
			return nil, fmt.Errorf("erro ao fazer unmarshal do arquivo %s: %w", entry.Name(), err)
		}

		combined = append(combined, words...)
	}

	return combined, nil
}

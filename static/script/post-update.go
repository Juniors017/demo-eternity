package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

type FrontMatter struct {
	Title string `yaml:"title"`
	// Ajoutez d'autres champs du frontmatter au besoin
}

func main() {
	// Spécifiez le chemin du répertoire où se trouvent vos fichiers .md
	dirPath := "content/work/bunny/"

	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		fmt.Println("Erreur lors de la lecture du répertoire:", err)
		return
	}

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".md") {
			filePath := dirPath + "/" + file.Name()

			// Lire le contenu du fichier
			content, err := ioutil.ReadFile(filePath)
			if err != nil {
				fmt.Println("Erreur lors de la lecture du fichier", file.Name(), ":", err)
				continue
			}

			// Analyser le frontmatter YAML
			frontMatter := FrontMatter{}
			if err := yaml.Unmarshal(content, &frontMatter); err != nil {
				fmt.Println("Erreur lors de l'analyse du frontmatter dans le fichier", file.Name(), ":", err)
				continue
			}

			// Mettre à jour le nom du fichier
			newFileName := strings.ReplaceAll(frontMatter.Title, " ", "") + ".md"
			newFilePath := dirPath + "/" + newFileName

			// Renommer le fichier
			err = os.Rename(filePath, newFilePath)
			if err != nil {
				fmt.Println("Erreur lors du changement de nom du fichier", file.Name(), ":", err)
				continue
			}

			fmt.Println("Le fichier", file.Name(), "a été renommé en", newFileName)
		}
	}
}

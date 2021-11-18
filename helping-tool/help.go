package helping_tool

import "fmt"

func Help() {
	fmt.Printf("Bienvenue dans l'aide\n")
	fmt.Printf("Pour utiliser le script il suffit de mettre en premier paramètre\nle fichier d'entrée au format \".csv\" et en second paramètre le fichier de sortie au format \".txt\"\n")
	fmt.Printf("Exemple: ./converter emails.csv list_email.txt\n")
}

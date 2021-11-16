package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

type Client struct {
	code  string
	nom   string
	email string
	adress string
}

func main() {
	var Arguments []string
	var InputFile string
	var OutputFile string

	Arguments = os.Args[1:]

	if len(Arguments) == 1 {
		if Arguments[0] == "help" {
			Help()
			os.Exit(1)
		}
	} else if len(Arguments) == 2 {
		InputFile = Arguments[0]
		OutputFile = Arguments[1]
	}


	//TODO : MUST BE DELETED, IT'S FOR DEBUG
	//inputfile = "emails.csv"
	//outputfile = "list.txt"

	records, errRecords := readData(InputFile)

	if errRecords != nil {
		log.Fatal(errRecords)
	}

	for _, record := range records {

		matched, _ := regexp.MatchString(`\b[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Z|a-z]{2,}\b`, record[2])

		if matched == true {
			client := Client{
				code:  record[0],
				nom:   record[1],
				email: record[2],
				adress: record[3],
			}

			// DEBUG
			// fmt.Printf("Client Code: %s\nNom: %s\nEmail: %s\nAdress: %s\n--- --- ---\n", client.code, client.nom, client.email, client.adress)
			fmt.Printf("Email: %s\n", client.email)

			f, errFile := os.OpenFile(OutputFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

			if errFile != nil {
				log.Fatal(errFile)
			}

			defer func(f *os.File) {
				errDeferClose := f.Close()
				if errDeferClose != nil {
					log.Fatal(errDeferClose)
				}
			}(f)

			FinalEmail := strings.TrimSpace(client.email)
			_, errWriteEmail := f.WriteString(FinalEmail + "\n")

			if errWriteEmail != nil {
				log.Fatal(errWriteEmail)
			}

			fmt.Println("Email was successfully wroten\n")
		}
	}
}

func readData(fileName string) ([][]string, error) {

	f, errOpenFile := os.Open(fileName)

	if errOpenFile != nil {
		return [][]string{}, errOpenFile
	}

	defer f.Close()

	r := csv.NewReader(f)
	r.Comma = ';'
	r.LazyQuotes = true

	// Skip first line -> Header
	if _, errReadCSV := r.Read(); errReadCSV != nil {
		return [][]string{}, errReadCSV
	}

	records, errReadAll := r.ReadAll()
	if errReadAll != nil {
		return [][]string{}, errReadAll
	}

	return records, nil
}

func Help(){
	fmt.Printf("Bienvenue dans l'aide\n")
	fmt.Printf("Pour utiliser le script il suffit de mettre en premier paramètre\nle fichier d'entrée au format \".csv\" et en second paramètre le fichier de sortie au format \".txt\"\n")
	fmt.Printf("Exemple: ./converter emails.csv list_email.txt")
}

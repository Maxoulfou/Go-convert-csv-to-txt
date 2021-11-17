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
	code   string
	nom    string
	email  string
	adress string
}

func main() {
	var Arguments []string
	var InputFile string
	var OutputFile string

	Arguments = os.Args[1:]

	//Setup log file
	f, errLogs := os.OpenFile("log.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if errLogs != nil {
		log.Fatalf("Error opening file: %v", errLogs)
	}

	defer func(f *os.File) {
		errDeferCloseLogFile := f.Close()
		if errDeferCloseLogFile != nil {
			log.Printf(errDeferCloseLogFile.Error())
		}
	}(f)

	log.SetOutput(f)
	log.Println("Start")
	fmt.Println("Start ...")

	// ---

	if len(Arguments) == 0 {
		fmt.Println("There is no argument, please execute './converter help'\n")
	} else if len(Arguments) == 1 {
		if Arguments[0] == "help" {
			Help()
			os.Exit(1)
		} else if Arguments[0] != "help" {
			fmt.Printf("The only argument " + Arguments[0] + " is unknown\n")
		}
	} else if len(Arguments) == 2 {
		RegexInput, _ := regexp.MatchString(`^[\w,\s-]+\.(csv)$`, Arguments[0])
		RegexOutput, _ := regexp.MatchString(`^[\w,\s-]+\.(txt)$`, Arguments[1])
		if RegexInput {
			InputFile = Arguments[0]
			fmt.Printf("Your input file: " + InputFile + " is valid\n")
		} else {
			fmt.Printf("Your input filename must has [dot]csv extension and no special char !\n")
		}
		if RegexOutput {
			OutputFile = Arguments[1]
			fmt.Printf("Your output file: " + OutputFile + " is valid\n")
		} else {
			fmt.Printf("Your output filename must has [dot]txt extension and no special char !\n")
		}

	} else if len(Arguments) > 2 {
		fmt.Printf("There is too much arguments\n")
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
		fmt.Print(record)
		if record[2] != "mail@nomail.com" {
			if matched == true {
				client := Client{
					code:   record[0],
					nom:    record[1],
					email:  record[2],
					adress: record[3],
				}

				// DEBUG
				// fmt.Printf("Client Code: %s\nNom: %s\nEmail: %s\nAdress: %s\n--- --- ---\n", client.code, client.nom, client.email, client.adress)
				// log.Printf("Email: %+v\n", client.email)

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

				log.Printf("Email %+v was successfully wroten\n", client.email)
			}
		}
	}
	log.Println("End")
	fmt.Println("End !")
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

func Help() {
	fmt.Printf("Bienvenue dans l'aide\n")
	fmt.Printf("Pour utiliser le script il suffit de mettre en premier paramètre\nle fichier d'entrée au format \".csv\" et en second paramètre le fichier de sortie au format \".txt\"\n")
	fmt.Printf("Exemple: ./converter emails.csv list_email.txt\n")
}

package main

import (
	helping_tool "brochier.xyz/converter/helping-tool"
	reading_tool "brochier.xyz/converter/reading-tool"
	"brochier.xyz/converter/sorting-tool"
	"fmt"
	"log"
	"math/big"
	"os"
	"regexp"
	"strings"
	"time"
)

type Client struct {
	code   string
	nom    string
	email  string
	adress string
}

func main() {
	start := time.Now()
	r := new(big.Int)
	fmt.Println(r.Binomial(1000, 10))

	var Arguments []string
	var InputFile string
	var OutputFile string

	Arguments = os.Args[1:]

	// --- Arguments --- //
	if len(Arguments) == 0 {
		fmt.Println("There is no argument, please execute './converter help'\n")
	} else if len(Arguments) == 1 {
		if Arguments[0] == "help" {
			helping_tool.Help()
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
	// ---End Arguments--- //

	//Setup log file
	LogFile, errLogs := os.OpenFile("log.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if errLogs != nil {
		log.Fatalf("Error opening file: %v", errLogs.Error())
	}

	defer func(LogFile *os.File) {
		errDeferCloseLogFile := LogFile.Close()
		if errDeferCloseLogFile != nil {
			log.Printf(errDeferCloseLogFile.Error())
		}
	}(LogFile)

	// Setup temp file
	TempFile, errTmpFile := os.OpenFile("tmp.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)

	if errTmpFile != nil {
		log.Fatalf("Error opening file: %v", errTmpFile.Error())
	}

	defer func(TempFile *os.File) {
		errDeferClose := TempFile.Close()
		if errDeferClose != nil {
			log.Fatal(errDeferClose.Error())
		}
	}(TempFile)

	// Setup final file
	FinalFileList, errFinalFile := os.OpenFile(OutputFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)

	if errFinalFile != nil {
		log.Fatalf("Error opening file: %v", errFinalFile.Error())
	}

	defer func(FinalFileList *os.File) {
		errDeferClose := FinalFileList.Close()
		if errDeferClose != nil {
			log.Fatal(errDeferClose.Error())
		}
	}(FinalFileList)

	log.SetOutput(LogFile)
	log.Println("Start")
	fmt.Println("Start ...")

	// Read Data's, stock in records var
	records, errRecords := reading_tool.ReadData(InputFile)
	if errRecords != nil {
		log.Fatal(errRecords.Error())
	}

	// Read records line by line as record var
	for _, record := range records {

		matched, _ := regexp.MatchString(`\b[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Z|a-z]{2,}\b`, record[2])

		if record[2] != "mail@nomail.com" {
			if matched == true {
				client := Client{
					code:   record[0],
					nom:    record[1],
					email:  record[2],
					adress: record[3],
				}

				FinalEmail := strings.TrimSpace(client.email)
				_, errWriteEmail := TempFile.WriteString(FinalEmail + "\n")

				if errWriteEmail != nil {
					log.Fatal(errWriteEmail.Error())
				}

				log.Printf("Email %+v was successfully wroten\n", client.email)
			}
		}
	}

	log.Println("Start sorting email")

	// TODO : fix the sort -> All uppercase email are not sort with precedent lowercase email

	MailList := sorting_tool.SortEmail("tmp.txt")
	for _, MailItem := range MailList {
		MailItem = strings.TrimSpace(MailItem)
		MailItem = strings.ToLower(MailItem)
		_, ErrorWriteFinalEmail := FinalFileList.WriteString(MailItem + "\n")
		if ErrorWriteFinalEmail != nil {
			log.Fatal(ErrorWriteFinalEmail.Error())
		}
	}

	TempFile.Close() // If TempFile not close, he's supposed to be used by another process
	errRemoveTmpFile := os.Remove(TempFile.Name())
	if errRemoveTmpFile != nil {
		log.Fatal(errRemoveTmpFile.Error())
	}

	log.Println("End")
	fmt.Println("End !")

	// Calculate time elapsed
	elapsed := time.Since(start)
	log.Printf("Elapsed time : %s\n", elapsed)
	fmt.Printf("Elapsed time : %s\n", elapsed)
}

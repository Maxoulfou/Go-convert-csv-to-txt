package sorting_tool

import (
	filetool "brochier.xyz/converter/file-tool"
	"fmt"
	"log"
	"os"
	"sort"
)

func SortEmail(SortingFile string) []string {
	// Open email list file
	MailFile, errOpenMailFile := os.OpenFile(SortingFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if errOpenMailFile != nil {
		log.Fatalf("Error opening file: %v", errOpenMailFile)
	}

	defer func(MailFile *os.File) {
		errDeferCloseLogFile := MailFile.Close()
		if errDeferCloseLogFile != nil {
			log.Fatalf("Error opening file: %v", errDeferCloseLogFile.Error())
		}
	}(MailFile)

	lineNB, ErrLineNB := filetool.LineCounter(MailFile)
	if ErrLineNB != nil {
		log.Printf(ErrLineNB.Error())
	}
	fmt.Printf("Line in %+v: %+v\n", SortingFile, lineNB)

	MailList := filetool.ReadLine(SortingFile)

	sort.Strings(MailList)
	log.Printf("\n----sorted----\n")
	for _, item := range MailList {
		log.Printf("item: %+v\n", item)
	}
	log.Printf("\n----end sorted----\n")

	return MailList
}

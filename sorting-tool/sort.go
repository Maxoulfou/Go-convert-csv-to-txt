package sorting_tool

import (
	filetool "brochier.xyz/converter/file-tool"
	"fmt"
	"log"
	"os"
	"sort"
)

func SortEmail(SortingFile string) []string {
	//Setup log file
	LogFile, errLogs := os.OpenFile("sorting-tool.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if errLogs != nil {
		log.Fatalf("Error opening file: %v", errLogs)
	}

	defer func(f *os.File) {
		errDeferCloseLogFile := f.Close()
		if errDeferCloseLogFile != nil {
			log.Printf(errDeferCloseLogFile.Error())
		}
	}(LogFile)

	// Open email list file
	MailFile, errOpenMailFile := os.OpenFile(SortingFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if errOpenMailFile != nil {
		log.Fatalf("Error opening file: %v", errOpenMailFile)
	}

	defer func(f *os.File) {
		errDeferCloseLogFile := f.Close()
		if errDeferCloseLogFile != nil {
			log.Printf(errDeferCloseLogFile.Error())
		}
	}(MailFile)

	lineNB, ErrLineNB := filetool.LineCounter(MailFile)
	if ErrLineNB != nil {
		log.Printf(ErrLineNB.Error())
	}
	fmt.Printf("Line in %+v: %+v\n", SortingFile, lineNB)

	MailList := filetool.ReadLine(SortingFile)
	for _, item := range MailList {
		fmt.Printf("item: %+v\n", item)
	}

	sort.Strings(MailList)
	fmt.Printf("\n----sorted:----")
	for _, item := range MailList {
		fmt.Printf("item: %+v\n", item)
	}

	return MailList
}

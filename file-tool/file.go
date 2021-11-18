package file_tool

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"os"
)

func LineCounter(r io.Reader) (int, error) {
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}
	}
}

/*
	BenchmarkBuffioScan   500      6408963 ns/op     4208 B/op    2 allocs/op
	BenchmarkBytesCount   500      4323397 ns/op     8200 B/op    1 allocs/op
	BenchmarkBytes32k     500      3650818 ns/op     65545 B/op   1 allocs/op
*/

func ReadLine(file string) []string {
	OpenedFile, err := os.Open(file)

	if err != nil {
		log.Fatalf("failed to open")
	}

	scanner := bufio.NewScanner(OpenedFile)

	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	OpenedFile.Close()

	return text

	/*
		DEBUG
		for _, eachLn := range text {
			fmt.Print("New line: ")
			fmt.Println(eachLn)
		}
		fmt.Println("---------")
	*/
}

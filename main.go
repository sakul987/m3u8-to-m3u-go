package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	sourceFile := flag.String("source", "", "File path of source .m3u8 file")
	//destFile := flag.String("dest","","File path of destination .m3u file")

	//sourceLocation := flag.String("sloc","","File path of directory containing audio files used in source .m3u8")
	//destLocation := flag.String("dloc","","File path of directory containing audio files used in destination .m3u")
	flag.Parse()

	//fmt.Println(*sourceFile)
	fmt.Println(loadSource(*sourceFile))
}

func loadSource(sourceFile string) []string {
	var lines []string
	file, err := os.Open(sourceFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		if !strings.HasPrefix(text, "/") {
			continue
		}
		lines = append(lines, text)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	file.Close()

	return lines
}

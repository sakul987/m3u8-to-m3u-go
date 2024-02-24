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
	destFile := flag.String("dest", "", "File path of destination .m3u file")

	sourceLocation := flag.String("sloc", "", "File path of directory containing audio files used in source .m3u8")
	destLocation := flag.String("dloc", "", "File path of directory containing audio files used in destination .m3u")
	flag.Parse()

	oldLines := loadSource(*sourceFile)
	fmt.Println("m3u8 file loaded")

	newLines := convert(oldLines, *sourceLocation, *destLocation)
	fmt.Println("old lines converted")

	writeDest(newLines, *destFile)
	fmt.Printf("m3u file has been saved at %v\n", *destFile)
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

func convert(lines []string, sourceLoc string, destLoc string) []string {
	newLines := []string{"#EXTM3U"}

	for i, line := range lines {
		//add empty line
		newLines = append(newLines, "")
		newLines = append(newLines, extinf(line, sourceLoc, i))
		newLines = append(newLines, strings.ReplaceAll(line, sourceLoc, destLoc))
	}

	return newLines
}

func extinf(line string, sourceLoc string, i int) string {
	name := strings.ReplaceAll(line, sourceLoc, "")

	return fmt.Sprintf("#EXTINF:%v,%v", i, name)
}

func writeDest(lines []string, destFile string) {
	file, err := os.OpenFile(destFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			fmt.Println(err)
			os.Exit(4)
		}
	}

	err = writer.Flush()
	if err != nil {
		fmt.Println(err)
		os.Exit(5)
	}
}

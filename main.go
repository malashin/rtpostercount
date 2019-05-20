package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

var inputPath = "input.txt"
var outputPath = "output.txt"
var resolution string
var re = regexp.MustCompile(`^(?:sd|hd)_\d{4}(?:_3d)?(?:_\w+)+__(?:\w+_)*poster(190x230|350x500|525x300|780x100|810x498|270x390|1620x996|503x726|1140x726|3510x1089|100x100|140x140|1170x363|570x363)(?:#[0-9a-fA-F]{8})?\.(?:jpg|png)$`)
var poster = []string{"350x500", "525x300", "780x100", "810x498", "1620x996", "1140x726", "3510x1089", "100x100", "140x140", "1170x363", "570x363"}
var rescale = []string{"190x230", "270x390", "503x726"}

func main() {
	files, err := readLines(inputPath)
	if err != nil {
		panic(err)
	}
	var output []string
	for _, file := range files {
		if !re.MatchString(file) {
			fmt.Println("WRONG FILENAME")
		}

		if re.MatchString(file) {
			resolution = re.ReplaceAllString(file, "${1}")
		}

		var s = "ERROR"
		switch {
		case contains(poster, resolution):
			s = "\t1"
		case contains(rescale, resolution):
			s = "\t\t1"
		}

		output = append(output, file+s+"\n")
	}
	writeStringArrayToFile(outputPath, output, 0775)
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// contains reports whether string is in string slice.
func contains(slice []string, s string) bool {
	for _, v := range slice {
		if v == s {
			return true
		}
	}
	return false
}

func writeStringArrayToFile(filename string, strArray []string, perm os.FileMode) {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, perm)
	if err != nil {
		log.Panic(err)
	}
	defer f.Close()
	for _, v := range strArray {
		if _, err = f.WriteString(v); err != nil {
			log.Panic(err)
		}
	}
}

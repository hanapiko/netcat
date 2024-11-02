package readart

import (
	"os"
	"log"
	"bufio"
)

func ReadArt(filename string) (art string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		art += scanner.Text() + "\n"
	}
	if err := scanner.Err(); err != nil {
		log.Println("Error reading from file:", err)
	}
	return
}
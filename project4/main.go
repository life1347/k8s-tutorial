package main

import (
	"os"
	"log"
	"bufio"
	"fmt"
)

func main() {
	logPath := os.Getenv("LOG_PATH")
	f, err := os.OpenFile(logPath, os.O_RDONLY, 0644)
	if err != nil {
		log.Printf("Error opening log file: %v", err)
	}
	defer f.Close()

	reader := bufio.NewReader(f)

	for {
		l, _, err := reader.ReadLine()
		if err != nil {
			log.Fatalf("error: %v", err)
		}
		fmt.Println(string(l))
	}
}

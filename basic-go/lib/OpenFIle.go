package lib

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func OpenFile() {

	wd, _ := os.Getwd()
	fmt.Println("Current working directory:", wd)

	if file, err := os.Open("/Users/users/Documents/go-learn/basic-go/data.txt"); err != nil {
		log.Fatal("Gagal membuka file:", err)
	} else {
		defer file.Close()
		fmt.Println("File berhasil dibuka")
		

		fmt.Printf("Isi file %s:\n\n", wd)

		
		scanner := bufio.NewScanner(file)

		lineNumber := 1
		for scanner.Scan() {
			line := scanner.Text()
			fmt.Printf("Baris %d: %s\n", lineNumber, line)
			lineNumber++
		}

		
		if err := scanner.Err(); err != nil {
			log.Fatalf("Error saat membaca file: %v", err)
		}
	}
}

// Package storage provides functions to read and manipulate files.
package storage

import (
	"fmt"
	"log"
	"os"
)

func ReadWrite(f string) {
	file, err := os.Open(f)
	if err != nil {
		log.Println("File failed to open:", err)
		return
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		log.Println("Stat fialed:", err)
		return
	}

	data := make([]byte, stat.Size())
	numBytes, err := file.Read(data)
	if err != nil {
		log.Println("Read failed:", err)
		return
	}

	fmt.Printf("read %d bytes\n", numBytes)
	fmt.Println(string(data[:numBytes]))
}

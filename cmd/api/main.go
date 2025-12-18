package main

import (
	"fmt"
	"log"

	"github.com/momoney3/scrubs/internal/storage"
)

func main() {
	foo, err := storage.ReadMem("./data.txt")
	if err != nil {
		log.Fatalln("did not work")
	}
	fmt.Println(foo)

	bar, err := storage.Reads("./data.txt")
	if err != nil {
		log.Fatalln("short read did not work", err)
	}
	fmt.Println(bar)
}

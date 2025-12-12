package main

import (
	"bufio"
	"fmt"
	"os"
)

func FRead(file string) (content string) {
	f, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	r := bufio.NewReader(f)

	for {
		line, err := r.ReadString('\n')
		if err != nil {
			break
		}
		content := line
		return content
	}
	defer f.Close()
	return fmt.Fprintln("Nothing Found")
}

func main() {
	FRead("test.txt")
}

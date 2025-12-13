package main

import (
	"fmt"
	"log"
	"strconv"
)

// func getFile(name string) {
// 	file, err := os.Open(name)
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()
// 	data, err := make([]byte, 100)
// 	for {
// 		count, err := file.Read(data)
// 		process(data[:count])
// 		if err != nil {
// 			if errors.Is(err, io.EOF) {
// 				return nil
// 			}
// 			return err
// 		}
// 	}
// }

func main() {
	str := "abcdefghai"
	num, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		log.Fatalln("Cannot Pars string", err)
	}
	fmt.Println("numberis", num)
}

func Conv(str string) (num int64, err error) {
	num, err = strconv.ParseInt(str, 10, 64)
	return
}

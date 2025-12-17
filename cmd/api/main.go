package main

import (
	"strconv"
)

func conv(str string) (numb int64, err error) {
	numb, err = strconv.ParseInt(str, 2, 64)
	return
}

func main() {
}

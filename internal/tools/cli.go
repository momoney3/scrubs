package tools

import (
	"fmt"
	"os"
)

func CommCLI() {
	if len(os.Args) < 2 {
		fmt.Println("expect hello")
		return
	}
	comment := os.Args[1]

	switch comment {
	case "hello":
		name := "world"
		if len(os.Args) > 2 {
			name = os.Args[2]
		}
		fmt.Println("hello ", name)

	default:
		fmt.Println("unknown comman", comment)
	}
}

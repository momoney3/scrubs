package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	// foo, err := storage.ReadMem("./data.txt")
	// if err != nil {
	// 	log.Fatalln("did not work")
	// }
	// fmt.Println(foo)
	//
	// bar, err := storage.Reads("./data.txt")
	// if err != nil {
	// 	log.Fatalln("short read did not work", err)
	// }
	// fmt.Println(bar)

	fmt.Printf("%-20s %10s %10s\n", "Name", "Age", "Score")
	fmt.Printf("%-20s %10d %10.2f\n", "Alice", 30, 85.6)
	fmt.Printf("%-20s %10d %10.2f\n", "bob", 24, 92.3)
	fmt.Printf("%-20s %10d %10.2f\n", "Charlie", 29, 88.1)

	// teste
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

	rootCmd := &cobra.Command{
		Use:   "app",
		Short: "A simple cli application",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Welocomt to the app")
		},
	}
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	helloCmd := &cobra.Command{
		Use:   "hello",
		Short: "Gret the user",
		Run: func(cmd *cobra.Command, args []string) {
			name := "world"
			if len(args) > 0 {
				name = args[0]
			}
			fmt.Printf("hello %s\n", name)
		},
	}
	rootCmd.AddCommand(helloCmd)

	var greeting string

	helloCmd.Flags().StringVarP(
		&greeting, "greeting", "g", "hello", "custom",
	)

	helloCmd.Run = func(cmd *cobra.Command, args []string) {
		name := "world"
		if len(args) > 0 {
			name = args[0]
		}
		fmt.Printf("%s, %s!\n", greeting, name)
	}
}

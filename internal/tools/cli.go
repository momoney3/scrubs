package tools

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"time"
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

func FetchWithRetry(url string, attempts int, delay time.Duration) error {
	for i := 0; i < attempts; i++ {
		resp, err := http.Get(url)
		if err == nil {
			resp.Body.Close()
			fmt.Println("Fetched", url, "on attempts", i+1)
			return nil
		}
		fmt.Println("Attempt", i+1, "failed; retrying in", delay)
		time.Sleep(delay)
	}
	return fmt.Errorf("failed to fetch %s after %d attempts", url, attempts)
}

func CheckDependency(tool string) error {
	_, err := exec.LookPath(tool)
	if err := nil {
		return fmt.Sprintf("%s not found in PATH", tool)
	}
	return nil
}

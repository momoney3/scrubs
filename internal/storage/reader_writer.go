// Package storage provides functions to read and manipulate files.
package storage

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

// ReadMem reads file from memery
func ReadMem(dat string) (string, error) {
	path := filepath.Join(dat)
	content, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrPermission) {
			return "", fmt.Errorf("permission failed: %w", err)
		}
		return "", fmt.Errorf("failed to open: %w", err)
	}
	return string(content), nil
}

// Reads big file
// NOTE: Seek to a known location in the file and Read from there
func Reads(f string) (string, error) {
	file, err := os.Open(f)
	if err != nil {
		return "", fmt.Errorf("failed to Open: %w", err)
	}

	defer func() {
		closErr := file.Close()
		if err == nil {
			err = closErr
		}
	}()

	stat, err := file.Stat()
	if err != nil {
		return "", fmt.Errorf("stat failed: %w", err)
	}

	data := make([]byte, stat.Size())
	numBytes, err := file.Read(data)
	if err != nil {
		return "", fmt.Errorf("read failed: %w", err)
	}

	content := string(data[:numBytes])

	return content, err
}

// SearchLog use to search log file
func SearchLog(file, word string) (string, error) {
	fmt.Println("the is just a test")
	return "", nil
}

// func fileExistis(f string) (bool, error) {
// 	path, err := exec.LookPath(f)
// 	if err != nil {
// 		if errors.Is(err, exec.ErrDot) {
// 			return true, nil
// 		}
// 		if errors.Is(err, exec.ErrNotFound) {
// 			return false, nil
// 		}
// 		return false, err
// 	}
// 	return true, nil
// }

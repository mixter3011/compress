package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	// Prompt the user to enter the file name
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the file name: ")
	fileName, _ := reader.ReadString('\n')
	fileName = strings.TrimSpace(fileName)

	// Specify the starting directory for the file search
	startDir := "C:\\Users\\Fobic"

	// Find the file
	filePath, err := findFile(fileName, startDir)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Compress the file
	compressedFilePath, err := compressFile(filePath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("File compressed successfully. Compressed file saved at:", compressedFilePath)
}

func findFile(fileName string, startDir string) (string, error) {
	// Walk through the file system starting from the specified directory
	var foundFilePath string
	err := filepath.Walk(startDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.Name() == fileName {
			foundFilePath = path
			return io.EOF // stop walking
		}
		return nil
	})
	if err == io.EOF {
		return foundFilePath, nil
	}
	if err != nil {
		return "", err
	}
	return "", fmt.Errorf("file '%s' not found", fileName)
}

func compressFile(filePath string) (string, error) {
	// Determine the extension of the file
	fileExt := filepath.Ext(filePath)

	// Destination path for compressed file
	compressedFilePath := filepath.Join(os.Getenv("USERPROFILE"), "Desktop", "compressed"+fileExt+".zip")

	// Command to compress file using zip
	cmd := exec.Command("zip", compressedFilePath, filePath)

	// Run the command
	err := cmd.Run()
	if err != nil {
		return "", err
	}

	return compressedFilePath, nil
}

package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

type Suite struct {
	directory string
}

func NewSuite(directory string) (instance *Suite) {
	return &Suite{
		directory: directory,
	}
}

func (instance *Suite) Run() (err error) {
	directories, err := os.ReadDir(instance.directory)
	if err != nil {
		return err
	}

	// Iterate the directory, finding files viable for execution
	for _, directory := range directories {
		if !directory.IsDir() {
			continue
		}

		// Check if the directory contains a main.go file
		mainFilePath := filepath.Join(instance.directory, directory.Name(), "main.go")
		_, err = os.Stat(mainFilePath)
		if err != nil {
			if os.IsNotExist(err) {
				continue
			}

			return fmt.Errorf("failed to stat %s: %w", mainFilePath, err)
		}

		// Compile the main.go file
		executablePath := directory.Name() + ".out"
		compileCmd := exec.Command("go", "build", "-o", executablePath, mainFilePath)
		compileCmd.Stdout = os.Stdout
		compileCmd.Stderr = os.Stderr

		err = compileCmd.Run()
		if err != nil {
			return fmt.Errorf("failed to compile %s: %w", mainFilePath, err)
		}
		defer os.Remove(executablePath)

		// Run the compiled executable
		runCmd := exec.Command("./" + executablePath)
		runCmd.Stdout = os.Stdout
		runCmd.Stderr = os.Stderr

		err = runCmd.Run()
		if err != nil {
			return fmt.Errorf("failed to run %s: %w", executablePath, err)
		}
	}

	return nil
}

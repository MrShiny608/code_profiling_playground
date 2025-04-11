package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"gopkg.in/yaml.v3"
)

func ReadConfig() (data map[string]any, err error) {
	_, callerFile, _, ok := runtime.Caller(1)
	if !ok {
		return nil, fmt.Errorf("could not get caller information")
	}

	currentDirectory := filepath.Dir(callerFile)
	parentDirectory, err := filepath.Abs(filepath.Join(currentDirectory, ".."))
	if err != nil {
		return nil, fmt.Errorf("could not get parent directory: %w", err)
	}

	configFile := filepath.Join(parentDirectory, "config.yaml")

	fileContent, err := os.ReadFile(configFile)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(fileContent, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func WriteConfig(data map[string]any) (err error) {
	_, callerFile, _, ok := runtime.Caller(1)
	if !ok {
		return fmt.Errorf("could not get caller information")
	}

	currentDirectory := filepath.Dir(callerFile)
	configFile := filepath.Join(currentDirectory, "config.yaml")

	fileContent, err := yaml.Marshal(data)
	if err != nil {
		return err
	}

	err = os.WriteFile(configFile, fileContent, 0644)
	if err != nil {
		return err
	}

	return nil
}

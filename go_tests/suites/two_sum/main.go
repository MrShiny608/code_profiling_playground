package main

import (
	"math/rand"
	"path/filepath"
	"runtime"

	"go_tests/go_tests/utils"
)

func createSuite() (suite *utils.Suite) {
	_, filePath, _, _ := runtime.Caller(0)
	currentDirectory := filepath.Dir(filePath)

	return utils.NewSuite(currentDirectory)
}

func main() {
	// Prepare the config files
	iterations := 100000000
	dataRange := 1000
	dataSize := 5

	// Generate a slice of integers from 1 to 1000
	numbers := make([]int, dataRange)
	for i := range dataRange {
		numbers[i] = i + 1
	}

	// Select dataSize unique integers from the slice
	for i := range dataSize {
		// Minus i as we shift data a long but don't reduce the size of the slice
		index := rand.Intn(dataRange - i)

		// Keep the selected value to be injected at the end
		selected := numbers[index]

		// Remove the selected integer by copying later data over it
		copy(numbers[index:], numbers[index+1:])

		// Put the newly selected entry at the end of the slice
		numbers[dataRange-1] = selected
	}

	data := numbers[dataRange-dataSize:]
	target := data[rand.Intn(dataSize)] + data[rand.Intn(dataSize)]

	utils.WriteConfig(map[string]any{
		"iterations": iterations,
		"target":     target,
		"data":       data,
	})

	// Run the test suite
	suite := createSuite()
	err := suite.Run()
	if err != nil {
		panic(err)
	}
}

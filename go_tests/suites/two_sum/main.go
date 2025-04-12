package main

import (
	"math/rand"
	"path/filepath"
	"runtime"
	"time"

	"go_tests/go_tests/utils"
)

func createSuite() (suite *utils.Suite) {
	_, filePath, _, _ := runtime.Caller(0)
	currentDirectory := filepath.Dir(filePath)

	return utils.NewSuite(currentDirectory)
}

func main() {
	// Prepare the config files
	duration := time.Minute * 5
	dataSizes := make([]int64, 100)
	for i := range dataSizes {
		dataSizes[i] = int64(i+1) * 1000
	}

	dataRange := int64(1)
	for _, size := range dataSizes {
		if dataRange < size {
			dataRange = size
		}
	}

	// Generate a slice of integers from 1 to 1000
	numbers := make([]int64, dataRange)
	for i := range dataRange {
		numbers[i] = i + 1
	}

	testConfigs := make([]map[string]any, len(dataSizes))

	for i, dataSize := range dataSizes {
		// Select dataSize unique integers from the slice
		for j := range dataSize {
			// Minus i as we shift data a long but don't reduce the size of the slice
			index := rand.Int63n(dataRange - j)

			// Keep the selected value to be injected at the end
			selected := numbers[index]

			// Remove the selected integer by copying later data over it
			copy(numbers[index:], numbers[index+1:])

			// Put the newly selected entry at the end of the slice
			numbers[dataRange-1] = selected
		}

		data := make([]int64, dataSize)
		copy(data, numbers[dataRange-dataSize:])

		// Set the target to an unachievable level so we can test
		// the worse case scenario
		target := dataRange + 1

		testConfigs[i] = map[string]any{
			"target": target,
			"data":   data,
		}
	}

	utils.WriteConfig(map[string]any{
		"duration":     duration.Seconds(),
		"test_configs": testConfigs,
	})

	// Run the test suite
	suite := createSuite()
	err := suite.Run()
	if err != nil {
		panic(err)
	}
}

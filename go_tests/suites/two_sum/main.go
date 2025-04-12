package main

import (
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

	testConfigs := make([]map[string]any, len(dataSizes))
	for i, dataSize := range dataSizes {
		// We don't actually need any real data, as we're intentionally hitting
		// worst case scenario so don't generate a consistent dataset for use across
		// tests, just let them generate an array of zeros of the correct size

		// Set the target to an unachievable level so we can test the worse case scenario
		target := 1

		testConfigs[i] = map[string]any{
			"target":    target,
			"data_size": dataSize,
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

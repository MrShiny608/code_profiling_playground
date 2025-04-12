package main

import (
	"go_tests/go_tests/utils"
	"time"
)

func createTest(data []int64, target int64) (work utils.Callable) {
	work = func() (indices any) {
		length := int64(len(data))

		for i := int64(0); i < length; i++ {
			compliment := target - data[i]

			for j := i + 1; j < length; j++ {
				if data[j] == compliment {
					return []int64{i, j}
				}
			}
		}

		return nil
	}

	return work
}

func main() {
	config, err := utils.ReadConfig()
	if err != nil {
		panic(err)
	}

	duration := time.Second * time.Duration(config["duration"].(int))
	testConfigs := config["test_configs"].([]interface{})

	tests := make([]utils.Test, len(testConfigs))
	for i := range testConfigs {
		testConfig := testConfigs[i].(map[string]any)

		target := int64(testConfig["target"].(int))
		dataSize := int64(testConfig["data_size"].(int))
		data := make([]int64, dataSize)

		tests[i] = utils.Test{
			Work: createTest(data, target),
			N:    dataSize,
		}
	}

	profile := utils.NewProfile("Brute Force", duration, tests)
	profile.Run()
}

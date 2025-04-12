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
		data_interface := testConfig["data"].([]interface{})
		n := int64(len(data_interface))
		data := make([]int64, n)

		for j := range data_interface {
			data[j] = int64(data_interface[j].(int))
		}

		target := int64(testConfig["target"].(int))
		tests[i] = utils.Test{
			Work: createTest(data, target),
			N:    n,
		}
	}

	profile := utils.NewProfile("Brute Force", duration, tests)
	profile.Run()
}

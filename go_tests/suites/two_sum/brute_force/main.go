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
	target := int64(config["target"].(int))
	dataSize := int64(config["data_size"].(int))

	data := make([]int64, dataSize)
	for i := range data {
		data[i] = int64(i) + 1
	}

	test := utils.Test{
		Work: createTest(data, target),
		N:    dataSize,
	}

	profile := utils.NewProfile("Brute Force", duration, test)
	profile.Run()
}

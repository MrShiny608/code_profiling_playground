package main

import (
	"go_tests/go_tests/utils"
	"time"
)

func createTest() (work utils.Callable2_1[[]int64, int64, []int64]) {
	work = func(data []int64, target int64) (indices []int64) {
		length := int64(len(data))

		for i := int64(0); i < length; i++ {
			complement := target - data[i]

			for j := i + 1; j < length; j++ {
				if data[j] == complement {
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

	test := utils.Test2_1[[]int64, int64, []int64]{
		Work:   createTest(),
		N:      dataSize,
		Input1: data,
		Input2: target,
	}

	profile := utils.NewProfile2_1("Brute Force", duration, test)
	profile.Run()
}

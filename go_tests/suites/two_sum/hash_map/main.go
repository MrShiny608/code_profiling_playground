package main

import (
	"go_tests/go_tests/utils"
	"time"
)

func createTest() (work utils.Callable2_1[[]int64, int64, []int64]) {
	work = func(data []int64, target int64) (indices []int64) {
		hashmap := make(map[int64]int64)

		for i, a := range data {
			complement := target - a

			index, ok := hashmap[complement]
			if ok {
				return []int64{index, int64(i)}
			}

			hashmap[a] = int64(i)
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

	profile := utils.NewProfile2_1("Hashmap", duration, test)
	profile.Run()
}

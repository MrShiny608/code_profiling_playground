package main

import (
	"go_tests/go_tests/utils"
	"time"
)

func createTest(data []int64, target int64) (work utils.Callable) {
	work = func() (indices any) {
		hashmap := make(map[int64]int64)
		for i, a := range data {
			compliment := target - a

			index, ok := hashmap[compliment]
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

	test := utils.Test{
		Work: createTest(data, target),
		N:    dataSize,
	}

	profile := utils.NewProfile("Hashmap", duration, test)
	profile.Run()
}

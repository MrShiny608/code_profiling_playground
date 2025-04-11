package main

import (
	"go_tests/go_tests/utils"
)

func createTest(data []int, target int) (work utils.Callable) {
	work = func() (result any) {
		hashmap := make(map[int]int)
		for i, a := range data {
			compliment := target - a

			index, ok := hashmap[compliment]
			if ok {
				return []int{i, index}
			}

			hashmap[a] = i
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

	data_interface := config["data"].([]interface{})
	data := make([]int, len(data_interface))
	for i := range data_interface {
		data[i] = data_interface[i].(int)
	}

	target := config["target"].(int)
	iterations := config["iterations"].(int)

	profile := utils.NewProfile("Hashmap (optimised)", iterations, createTest(data, target))
	profile.Run()
}

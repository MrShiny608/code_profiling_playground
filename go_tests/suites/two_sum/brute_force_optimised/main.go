package main

import (
	"go_tests/go_tests/utils"
)

func createTest(data []int, target int) (work utils.Callable) {
	work = func() (result any) {
		length := len(data)

		for i := 0; i < length; i++ {
			compliment := target - data[i]

			for j := i + 1; j < length; j++ {
				if data[j] == compliment {
					return []int{i, j}
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

	data_interface := config["data"].([]interface{})
	data := make([]int, len(data_interface))
	for i := range data_interface {
		data[i] = data_interface[i].(int)
	}

	target := config["target"].(int)
	iterations := config["iterations"].(int)

	profile := utils.NewProfile("Brute Force (optimised)", iterations, createTest(data, target))
	profile.Run()
}

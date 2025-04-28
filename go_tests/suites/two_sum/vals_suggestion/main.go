package main

import (
	"go_tests/go_tests/utils"
	"sort"
	"time"
)

type Tuple struct {
	value int64
	index int64
}

func createTest() (work utils.Callable2_1[[]int64, int64, []int64]) {
	work = func(data []int64, target int64) (indices []int64) {
		length := int64(len(data))
		dataAndIndices := make([]Tuple, length)
		for i := int64(0); i < length; i++ {
			dataAndIndices[i] = Tuple{
				value: data[i],
				index: i,
			}
		}

		sort.Slice(dataAndIndices, func(i, j int) bool {
			return dataAndIndices[i].value < dataAndIndices[j].value
		})

		i := int64(0)
		j := length - 1
		for i <= j {
			total := dataAndIndices[i].value + dataAndIndices[j].value
			if total == target {
				return []int64{dataAndIndices[i].index, dataAndIndices[j].index}
			} else if total < target {
				i += 1
			} else if total > target {
				j -= 1
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

	profile := utils.NewProfile2_1("Val's Suggestion", duration, test)
	profile.Run()
}

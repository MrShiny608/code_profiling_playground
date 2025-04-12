package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTest(t *testing.T) {
	t.Parallel()

	type args struct {
		data   []int64
		target int64
	}
	type result struct {
		indices []int64
	}
	type testConfig struct {
		name   string
		args   *args
		result *result
	}
	configs := []testConfig{
		{
			name: "returns the correct indices",
			args: &args{
				data:   []int64{1, 2, 3},
				target: 5,
			},
			result: &result{
				indices: []int64{1, 2},
			},
		},
		{
			name: "returns nil when target isn't achievable",
			args: &args{
				data:   []int64{1, 2, 3},
				target: 6,
			},
			result: &result{
				indices: nil,
			},
		},
	}

	for _, config := range configs {
		t.Run(config.name, func(t *testing.T) {
			// Arrange
			args := config.args
			result := config.result
			work := createTest(args.data, args.target)

			// Act
			_indices := work()

			// Cast as utils.Callable interface returns any
			var indices []int64 = nil
			if _indices != nil {
				indices = _indices.([]int64)
			}

			// Assert
			assert.Equal(t, result.indices, indices)
		})
	}
}

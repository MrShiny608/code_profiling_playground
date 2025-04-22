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
		i int64
		j int64
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
				i: 1,
				j: 2,
			},
		},
		{
			name: "returns nil when target isn't achievable",
			args: &args{
				data:   []int64{1, 2, 3},
				target: -1,
			},
			result: &result{
				i: -1,
				j: -1,
			},
		},
	}

	for _, config := range configs {
		t.Run(config.name, func(t *testing.T) {
			// Arrange
			args := config.args
			result := config.result
			work := createTest()

			// Act
			i, j := work(args.data, args.target)

			// Assert
			assert.Equal(t, result.i, i)
			assert.Equal(t, result.j, j)
		})
	}
}

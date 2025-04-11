package utils

import (
	"fmt"
	"time"
)

type Callable func() (result any)

type Test struct {
	Work Callable
	N    int64
}

type Profile struct {
	name     string
	tests    []Test
	duration time.Duration
}

func NewProfile(name string, duration time.Duration, test []Test) (instance *Profile) {
	return &Profile{
		name:     name,
		tests:    test,
		duration: duration,
	}
}

func (instance *Profile) Run() {
	fmt.Printf("Running test in Go\n - name: %s\n", instance.name)

	// Run the tests
	for _, test := range instance.tests {
		fmt.Printf(" - N=%d", test.N)

		iterations := int64(0)
		start := time.Now()

		for time.Since(start) < instance.duration {
			test.Work()
			iterations++
		}

		duration := time.Since(start)
		fmt.Printf(" %dns\n", duration.Nanoseconds()/iterations)
	}
}

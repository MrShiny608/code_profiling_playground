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
	test     Test
	duration time.Duration
}

func NewProfile(name string, duration time.Duration, test Test) (instance *Profile) {
	return &Profile{
		name:     name,
		test:     test,
		duration: duration,
	}
}

func (instance *Profile) Run() {
	fmt.Printf("[Go] %s - N=%d", instance.name, instance.test.N)

	// Run the test
	test := instance.test
	iterations := int64(0)
	start := time.Now()

	for time.Since(start) < instance.duration {
		test.Work()
		iterations++
	}

	duration := time.Since(start)
	fmt.Printf(" %dns\n", duration.Nanoseconds()/iterations)
}

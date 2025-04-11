package utils

import (
	"fmt"
	"time"
)

type Callable func() (result any)

type Profile struct {
	name       string
	test       Callable
	iterations int
	duration   time.Duration
}

func NewProfile(name string, iterations int, test Callable) (instance *Profile) {
	return &Profile{
		name:       name,
		test:       test,
		iterations: iterations,
		duration:   0,
	}
}

func (instance *Profile) Run() {
	fmt.Printf("Running test in Go\n - name: %s\n", instance.name)

	// Run the test
	start := time.Now()

	for range instance.iterations {
		instance.test()
	}

	instance.duration += time.Since(start)

	fmt.Printf(" - duration: %ds\n - iterations: %d\n - average: %dns\n", int(instance.duration.Seconds()), instance.iterations, instance.duration.Nanoseconds()/int64(instance.iterations))
}

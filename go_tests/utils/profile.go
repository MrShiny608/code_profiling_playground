package utils

import (
	"fmt"
	"time"
)

// Go really doesn't make this easy, we'll need to define a type for each function signature we need

type Callable2_1[T1 any, T2 any, R1 any] func(T1, T2) R1
type Callable2_2[T1 any, T2 any, R1 any, R2 any] func(T1, T2) (R1, R2)

type Test2_1[T1 any, T2 any, R1 any] struct {
	N      int64
	Work   Callable2_1[T1, T2, R1]
	Input1 T1
	Input2 T2
}

type Test2_2[T1 any, T2 any, R1 any, R2 any] struct {
	N      int64
	Work   Callable2_2[T1, T2, R1, R2]
	Input1 T1
	Input2 T2
}

type Profile2_1[T1 any, T2 any, R1 any] struct {
	name     string
	test     Test2_1[T1, T2, R1]
	duration time.Duration
}

func NewProfile2_1[T1 any, T2 any, R1 any](name string, duration time.Duration, test Test2_1[T1, T2, R1]) (instance *Profile2_1[T1, T2, R1]) {
	return &Profile2_1[T1, T2, R1]{
		name:     name,
		test:     test,
		duration: duration,
	}
}

func (instance *Profile2_1[T1, T2, R1]) Run() {
	fmt.Printf("[Go] %s - N=%d", instance.name, instance.test.N)

	// Run the test
	test := instance.test
	iterations := int64(0)
	start := time.Now()

	for time.Since(start) < instance.duration {
		test.Work(test.Input1, test.Input2)
		iterations++
	}

	duration := time.Since(start)
	fmt.Printf(" %dns\n", duration.Nanoseconds()/iterations)
}

type Profile2_2[T1 any, T2 any, R1 any, R2 any] struct {
	name     string
	test     Test2_2[T1, T2, R1, R2]
	duration time.Duration
}

func NewProfile2_2[T1 any, T2 any, R1 any, R2 any](name string, duration time.Duration, test Test2_2[T1, T2, R1, R2]) (instance *Profile2_2[T1, T2, R1, R2]) {
	return &Profile2_2[T1, T2, R1, R2]{
		name:     name,
		test:     test,
		duration: duration,
	}
}

func (instance *Profile2_2[T1, T2, R1, R2]) Run() {
	fmt.Printf("[Go] %s - N=%d", instance.name, instance.test.N)

	// Run the test
	test := instance.test
	iterations := int64(0)
	start := time.Now()

	for time.Since(start) < instance.duration {
		test.Work(test.Input1, test.Input2)
		iterations++
	}

	duration := time.Since(start)
	fmt.Printf(" %dns\n", duration.Nanoseconds()/iterations)
}

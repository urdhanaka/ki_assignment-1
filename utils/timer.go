package utils

import (
	"fmt"
	"time"
)

func timer(name string) func() time.Duration {
	start := time.Now()
	return func() time.Duration {
		actualTime := time.Since(start) - 1*time.Second
		fmt.Printf("[%s] time = %v\n", name, actualTime)
		return time.Since(start)
	}
}

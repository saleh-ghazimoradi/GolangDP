package singleton

import (
	"fmt"
	"sync"
	"testing"
)

func TestSingletonConcurrency(t *testing.T) {
	var wg sync.WaitGroup
	n := 100
	instances := make([]*Singleton, n)

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			instances[idx] = GetInstance()
		}(i)
	}
	wg.Wait()

	for i := 1; i < n; i++ {
		if instances[i] != instances[0] {
			t.Errorf("Singleton instances are not the same: %p vs %p", instances[i], instances[0])
		}
	}

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			s := GetInstance()
			s.SetData(fmt.Sprintf("data-%d", i))
		}(i)
	}
	wg.Wait()
}

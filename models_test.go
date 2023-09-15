package robinson

import (
	"fmt"
	"sync"
	"testing"
)

func TestCrusoe_Call(t *testing.T) {
	crusoe := NewCrusoe[int32]()
	var wg sync.WaitGroup
	f := func(i int32) int32 {
		return i + 1
	}
	for j := 0; j < 1000; j++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			crusoe.Call(f)
		}()
	}

	wg.Wait()
	if crusoe.Get() != 1000 {
		t.Errorf("strange value returned: %v", crusoe.Get())
	}
}

func TestCrusoe_Get_Int(t *testing.T) {
	tests := []struct {
		value int
		want  *Crusoe[int]
	}{
		{123, &Crusoe[int]{}},
		{456, &Crusoe[int]{}},
		{789, &Crusoe[int]{}},
		{1, &Crusoe[int]{}},
	}
	crusoe := NewCrusoe[int]()
	if fmt.Sprintf("%[1]T", crusoe) != "*robinson.Crusoe[int]" {
		t.Errorf("strange cache type returned: %T", crusoe)
	}
	for _, scenario := range tests {
		t.Run(fmt.Sprintf("%T %[1]v", scenario.value), func(t *testing.T) {
			crusoe.Set(scenario.value)
			cacheValue := crusoe.Get()
			if fmt.Sprintf("%[1]T", cacheValue) != "int" {
				t.Errorf("strange value type returned: %T", cacheValue)
			}
		})
	}

	{
		cacheValue := crusoe.Get()
		if cacheValue != tests[len(tests)-1].value {
			t.Errorf("strange value returned: %v", cacheValue)
		}
	}

	for _, scenario := range tests {
		t.Run(fmt.Sprintf("%T %[1]v", scenario.value), func(t *testing.T) {
			crusoe.Set(scenario.value)
			cacheValue := crusoe.Get()
			if fmt.Sprintf("%[1]T", cacheValue) != "int" {
				t.Errorf("strange value type returned: %T", cacheValue)
			}
			if cacheValue != scenario.value {
				t.Errorf("strange value returned: %v", cacheValue)
			}
		})
	}
}

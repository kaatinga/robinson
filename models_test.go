package robinson_test

import (
	"fmt"
	"sync"
	"testing"

	"github.com/kaatinga/robinson"
)

func TestCrusoe_Call(t *testing.T) {
	crusoe := robinson.NewCrusoe[int32]()
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
		want  *robinson.Crusoe[int]
	}{
		{123, &robinson.Crusoe[int]{}},
		{456, &robinson.Crusoe[int]{}},
		{789, &robinson.Crusoe[int]{}},
		{1, &robinson.Crusoe[int]{}},
	}
	crusoe := robinson.NewCrusoe[int]()
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

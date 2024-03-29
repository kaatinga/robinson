package robinson_test

import (
	"fmt"
	"sync"
	"testing"

	"github.com/kaatinga/robinson"
)

func TestCrusoe_Call(t *testing.T) {
	crusoe := robinson.NewCrusoePointer[int32]()
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
	var tests []struct {
		value int
	}

	// fuzzing
	for i := 0; i < 1000; i++ {
		tests = append(tests, struct{ value int }{value: i})
	}

	crusoe := robinson.NewCrusoePointer[int]()
	if fmt.Sprintf("%[1]T", crusoe) != "*robinson.Crusoe[int]" {
		t.Errorf("strange cache type returned: %T", crusoe)
	}

	wg := sync.WaitGroup{}
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func(i int) {
			for _, scenario := range tests {
				t.Run(fmt.Sprintf("%T %[1]v, loop %d", scenario.value, i), func(t *testing.T) {
					crusoe.Set(scenario.value)
					_ = crusoe.Get()
				})
			}

			wg.Done()
		}(i)
	}
	wg.Wait()

	t.Run("get the last value", func(t *testing.T) {
		cacheValue := crusoe.Get()
		if cacheValue != tests[len(tests)-1].value {
			t.Errorf("strange value returned: %v", cacheValue)
		}
	})
}

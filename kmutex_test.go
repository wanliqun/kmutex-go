package kmutex_test

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wanliqun/kmutex-go"
)

func TestKMutex(t *testing.T) {
	km := kmutex.NewKMutex()
	lockKey := kmutex.MutexKey("kmutex")

	testNum := int64(0)
	testTimes := int64(2000)

	wg := &sync.WaitGroup{}

	for i := int64(0); i < testTimes; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			km.Lock(lockKey)
			defer km.Unlock(lockKey)

			testNum = testNum + 1
		}()
	}

	wg.Wait()

	assert.EqualValues(t, testTimes, testNum)
}

// Additional test for multiple keys
func TestKMutexMultipleKeys(t *testing.T) {
	km := kmutex.NewKMutex()
	keys := []kmutex.MutexKey{"key1", "key2", "key3", "key4", "key5"}

	testTimes := 1000
	counter := make(map[kmutex.MutexKey]int64)

	mu := &sync.Mutex{}
	wg := &sync.WaitGroup{}

	for i := 0; i < testTimes; i++ {
		for _, key := range keys {
			wg.Add(1)
			go func(k kmutex.MutexKey) {
				defer wg.Done()

				km.Lock(k)
				defer km.Unlock(k)

				mu.Lock()
				counter[k]++
				mu.Unlock()
			}(key)
		}
	}

	wg.Wait()

	for _, key := range keys {
		assert.EqualValues(t, testTimes, counter[key])
	}
}

func BenchmarkKMutex(b *testing.B) {
	km := kmutex.NewKMutex()
	lockKey := kmutex.MutexKey("benchKey")

	for i := 0; i < b.N; i++ {
		km.Lock(lockKey)
		km.Unlock(lockKey)
	}
}

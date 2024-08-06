package kmutex

import (
	"math/rand"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestKMutex(t *testing.T) {
	kmutex := NewKMutex()
	lockKey := MutexKey("kmutex")

	testNum := int64(0)
	testTimes := int64(2000)

	wg := &sync.WaitGroup{}

	for i := int64(0); i < testTimes; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			kmutex.Lock(lockKey)
			defer kmutex.Unlock(lockKey)

			amt := time.Duration(rand.Intn(100))
			time.Sleep(time.Millisecond * amt)

			testNum = testNum + 1
		}()
	}

	wg.Wait()

	assert.EqualValues(t, testTimes, testNum)
}

// Additional test for multiple keys
func TestKMutexMultipleKeys(t *testing.T) {
	kmutex := NewKMutex()
	keys := []MutexKey{"key1", "key2", "key3", "key4", "key5"}

	testTimes := 1000
	counter := make(map[MutexKey]int64)
	mu := &sync.Mutex{}
	wg := &sync.WaitGroup{}

	for i := 0; i < testTimes; i++ {
		for _, key := range keys {
			wg.Add(1)
			go func(k MutexKey) {
				defer wg.Done()

				kmutex.Lock(k)
				defer kmutex.Unlock(k)

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
	kmutex := NewKMutex()
	lockKey := MutexKey("benchKey")

	for i := 0; i < b.N; i++ {
		kmutex.Lock(lockKey)
		kmutex.Unlock(lockKey)
	}
}

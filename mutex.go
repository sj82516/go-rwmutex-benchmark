package mutex_benchmark

import (
	"math/rand"
	"sync"
)

func LockTest(lockType string, times int, writeRate float64){
	wg := sync.WaitGroup{}
	wg.Add(times)

	rwlock := sync.RWMutex{}
	lock := sync.Mutex{}

	for i:=0; i < times; i++ {
		isWrite := rand.Float64() < writeRate
		if lockType == "rw" {
			go FakeRWExec(&rwlock, &wg, isWrite)
		}else {
			go FakeExec(&lock, &wg, isWrite)
		}
	}
	wg.Wait()
}

func FakeExec(lock *sync.Mutex, wg *sync.WaitGroup, isWrite bool) {
	defer wg.Done()
	lock.Lock()
	if isWrite {
		FakeWrite()
	} else {
		FakeRead()
	}
	lock.Unlock()
}

func FakeRWExec(lock *sync.RWMutex, wg *sync.WaitGroup, isWrite bool) {
	defer wg.Done()
	if isWrite {
		lock.Lock()
		FakeWrite()
		lock.Unlock()
	} else {
		lock.RLock()
		FakeRead()
		lock.RUnlock()
	}
}
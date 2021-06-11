package mutex_benchmark

import (
	"testing"
)

const (
	lockTimes = 1e3
	writeRate = 0.2
)

func BenchmarkMutexTest(b *testing.B) {
	LockTest("", lockTimes, writeRate)
}

func BenchmarkRWMutexTest(b *testing.B) {
	LockTest("rw", lockTimes, writeRate)
}

func BenchmarkFakeWrite(b *testing.B) {
	FakeWrite()
}

func BenchmarkFakeRead(b *testing.B) {
	FakeRead()
}
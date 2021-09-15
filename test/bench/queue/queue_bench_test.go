package test

import (
	"testing"

	"github.com/Juliusan/go.infinitechannel/src/queue"
)

// General warning: Go's benchmark utility (go test -bench .) increases the number of
// iterations until the benchmarks take a reasonable amount of time to run; memory usage
// is *NOT* considered. On my machine, these benchmarks hit around ~1GB before they've had
// enough, but if you have less than that available and start swapping, then all bets are off.
var q *queue.PriorityQueue

func BenchmarkQueueAdd10k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		q = queue.NewDefaultPriorityQueue()
		b.StartTimer()

		for i := 0; i < 10000; i++ {
			q.Add(i)
		}
	}
}

func BenchmarkQueueAdd10kPrioritising(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		q = queue.NewPriorityQueue(func(i interface{}) bool {
			return i.(int)%2 == 0
		})
		b.StartTimer()

		for i := 0; i < 10000; i++ {
			q.Add(i)
		}
	}
}

func BenchmarkQueueRemove10k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		q = queue.NewDefaultPriorityQueue()
		for i := 0; i < 10000; i++ {
			q.Add(i)
		}
		b.StartTimer()
		for i := 0; i < 10000; i++ {
			q.Remove()
		}
	}
}

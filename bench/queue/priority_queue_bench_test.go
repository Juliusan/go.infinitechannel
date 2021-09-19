package queue_bench

import (
	"testing"

	"github.com/Juliusan/go.infinitechannel/src/queue"
)

func BenchmarkDefaultPriorityQueueAdd10k(b *testing.B) {
	benchmarkQueueAdd10k(func() queue.Queue { return queue.NewDefaultPriorityQueue() }, b)
}

func BenchmarkPriorityQueueAdd10k(b *testing.B) {
	makeQueueFun := func() queue.Queue {
		return queue.NewPriorityQueue(func(i interface{}) bool {
			return i.(int)%2 == 0
		})
	}
	benchmarkQueueAdd10k(makeQueueFun, b)
}

//--

func BenchmarkDefaultPriorityQueueRemove10k(b *testing.B) {
	benchmarkDefaultQueueRemove10k(func() queue.Queue { return queue.NewDefaultPriorityQueue() }, b)
}

func BenchmarkPriorityQueueRemove10k(b *testing.B) {
	makeQueueFun := func() queue.Queue {
		return queue.NewPriorityQueue(func(i interface{}) bool {
			return i.(int)%2 == 0
		})
	}
	benchmarkDefaultQueueRemove10k(makeQueueFun, b)
}

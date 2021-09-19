package queue_bench

import (
	"testing"

	"github.com/Juliusan/go.infinitechannel/src/queue"
)

var q queue.Queue

func benchmarkQueueAdd10k(makeQueueFun func() queue.Queue, b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		q = makeQueueFun()
		b.StartTimer()

		for i := 0; i < 10000; i++ {
			q.Add(i)
		}
	}
}

//--

func benchmarkDefaultQueueRemove10k(makeQueueFun func() queue.Queue, b *testing.B) {
	benchmarkQueueRemove10k(makeQueueFun, 10000, b)
}

func benchmarkQueueRemove10k(makeQueueFun func() queue.Queue, elementsToAdd int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		q = makeQueueFun()
		for i := 0; i < elementsToAdd; i++ {
			q.Add(i)
		}
		b.StartTimer()
		for i := 0; i < 10000; i++ {
			q.Remove()
		}
	}
}

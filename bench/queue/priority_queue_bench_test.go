package queue_bench

import (
	"testing"

	"github.com/Juliusan/go.infinitechannel/src/queue"
	"github.com/Juliusan/go.infinitechannel/src/util"
)

func BenchmarkDefaultPriorityQueueAdd10k(b *testing.B) {
	benchmarkQueueAdd10k(func() queue.Queue { return queue.NewDefaultPriorityQueue() }, b)
}

func BenchmarkPriorityQueueAdd10k(b *testing.B) {
	benchmarkPriorityQueueAdd10k(func(fun func(i interface{}) bool) queue.Queue { return queue.NewPriorityQueue(fun) }, b)
}

func benchmarkPriorityQueueAdd10k(makePriorityQueueFun func(func(i interface{}) bool) queue.Queue, b *testing.B) {
	makeQueueFun := func() queue.Queue { return makePriorityQueueFun(util.PriorityFunMod10) }
	benchmarkQueueAdd10k(makeQueueFun, b)
}

//--

func BenchmarkDefaultPriorityQueueRemove10k(b *testing.B) {
	benchmarkDefaultQueueRemove10k(func() queue.Queue { return queue.NewDefaultPriorityQueue() }, b)
}

func BenchmarkPriorityQueueRemove10k(b *testing.B) {
	benchmarkPriorityQueueRemove10k(func(fun func(i interface{}) bool) queue.Queue { return queue.NewPriorityQueue(fun) }, b)
}

func benchmarkPriorityQueueRemove10k(makePriorityQueueFun func(func(i interface{}) bool) queue.Queue, b *testing.B) {
	makeQueueFun := func() queue.Queue { return makePriorityQueueFun(util.PriorityFunMod10) }
	benchmarkDefaultQueueRemove10k(makeQueueFun, b)
}

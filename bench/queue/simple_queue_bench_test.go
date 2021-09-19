package queue_bench

import (
	"testing"

	"github.com/Juliusan/go.infinitechannel/src/queue"
)

func BenchmarkSimpleQueueAdd10k(b *testing.B) {
	benchmarkQueueAdd10k(func() queue.Queue { return queue.NewSimpleQueue() }, b)
}

func BenchmarkSimpleQueueRemove10k(b *testing.B) {
	benchmarkDefaultQueueRemove10k(func() queue.Queue { return queue.NewSimpleQueue() }, b)
}

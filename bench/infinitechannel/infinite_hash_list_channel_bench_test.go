package infinitechannel_bench

import (
	"testing"

	"github.com/Juliusan/go.infinitechannel/src/infinitechannel"
)

func BenchmarkInfiniteDefaultLimitedPriorityHashListChannelTransfer10kSerial(b *testing.B) {
	benchmarkDefaultInfiniteChannelTransfer10kSerial(func() *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteDefaultLimitedPriorityHashListChannel()
	}, b)
}

func BenchmarkInfinitePriorityLimitedPriorityHashListChannelTransfer10kSerial(b *testing.B) {
	benchmarkInfinitePriorityChannelTransfer10kSerial(func(fun func(i interface{}) bool) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfinitePriorityLimitedPriorityHashListChannel(fun)
	}, b)
}

func BenchmarkInfiniteLimitLimitedPriorityHashListChannelNoLimitTransfer10kSerial(b *testing.B) {
	benchmarkInfiniteLimitLimitedPriorityChannelNoLimitTransfer10kSerial(func(limit int) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteLimitLimitedPriorityHashListChannel(limit)
	}, b)
}

func BenchmarkInfiniteLimitLimitedPriorityHashListChannelTransfer10kSerial(b *testing.B) {
	benchmarkInfiniteLimitLimitedPriorityChannelTransfer10kSerial(func(limit int) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteLimitLimitedPriorityHashListChannel(limit)
	}, b)
}

func BenchmarkInfiniteLimitPriorityLimitedPriorityHashListChannelNoLimitTransfer10kSerial(b *testing.B) {
	benchmarkInfiniteLimitedPriorityChannelNoLimitTransfer10kSerial(func(fun func(interface{}) bool, limit int) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteLimitPriorityLimitedPriorityHashListChannel(fun, limit)
	}, b)
}

func BenchmarkInfiniteLimitPriorityLimitedPriorityHashListChannelTransfer10kSerial(b *testing.B) {
	benchmarkInfiniteLimitedPriorityChannelTransfer10kSerial(func(fun func(interface{}) bool, limit int) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteLimitPriorityLimitedPriorityHashListChannel(fun, limit)
	}, b)
}

func BenchmarkInfiniteHashLimitedPriorityHashListChannelTransfer10kSerial(b *testing.B) {
	benchmarkDefaultInfiniteChannelTransfer10kSerial(func() *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteHashLimitedPriorityHashListChannel()
	}, b)
}

func BenchmarkInfinitePriorityHashLimitedPriorityHashListChannelTransfer10kSerial(b *testing.B) {
	benchmarkInfinitePriorityChannelTransfer10kSerial(func(fun func(i interface{}) bool) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfinitePriorityHashLimitedPriorityHashListChannel(fun)
	}, b)
}

func BenchmarkInfiniteLimitHashLimitedPriorityHashListChannelNoLimitTransfer10kSerial(b *testing.B) {
	benchmarkInfiniteLimitLimitedPriorityChannelNoLimitTransfer10kSerial(func(limit int) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteLimitHashLimitedPriorityHashListChannel(limit)
	}, b)
}

func BenchmarkInfiniteLimitHashLimitedPriorityHashListChannelTransfer10kSerial(b *testing.B) {
	benchmarkInfiniteLimitLimitedPriorityChannelTransfer10kSerial(func(limit int) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteLimitHashLimitedPriorityHashListChannel(limit)
	}, b)
}

func BenchmarkInfiniteLimitedPriorityHashListChannelNoLimitTransfer10kSerial(b *testing.B) {
	benchmarkInfiniteLimitedPriorityChannelNoLimitTransfer10kSerial(func(fun func(interface{}) bool, limit int) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteLimitedPriorityHashListChannel(fun, limit)
	}, b)
}

func BenchmarkInfiniteLimitedPriorityHashListChannelTransfer10kSerial(b *testing.B) {
	benchmarkInfiniteLimitedPriorityChannelTransfer10kSerial(func(fun func(interface{}) bool, limit int) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteLimitedPriorityHashListChannel(fun, limit)
	}, b)
}

//--

func BenchmarkInfiniteDefaultLimitedPriorityHashListChannelTransfer10kConcurrent(b *testing.B) {
	benchmarkDefaultInfiniteChannelTransfer10kConcurrent(func() *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteDefaultLimitedPriorityHashListChannel()
	}, b)
}

func BenchmarkInfinitePriorityLimitedPriorityHashListChannelTransfer10kConcurrent(b *testing.B) {
	benchmarkInfinitePriorityChannelTransfer10kConcurrent(func(fun func(i interface{}) bool) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfinitePriorityLimitedPriorityHashListChannel(fun)
	}, b)
}

func BenchmarkInfiniteLimitLimitedPriorityHashListChannelNoLimitTransfer10kConcurrent(b *testing.B) {
	benchmarkInfiniteLimitLimitedPriorityChannelNoLimitTransfer10kConcurrent(func(limit int) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteLimitLimitedPriorityHashListChannel(limit)
	}, b)
}

func BenchmarkInfiniteLimitLimitedPriorityHashListChannelTransfer10kConcurrent(b *testing.B) {
	benchmarkInfiniteLimitLimitedPriorityChannelTransfer10kConcurrent(func(limit int) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteLimitLimitedPriorityHashListChannel(limit)
	}, b)
}

func BenchmarkInfiniteLimitPriorityLimitedPriorityHashListChannelNoLimitTransfer10kConcurrent(b *testing.B) {
	benchmarkInfiniteLimitedPriorityChannelNoLimitTransfer10kConcurrent(func(fun func(interface{}) bool, limit int) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteLimitPriorityLimitedPriorityHashListChannel(fun, limit)
	}, b)
}

func BenchmarkInfiniteLimitPriorityLimitedPriorityHashListChannelTransfer10kConcurrent(b *testing.B) {
	benchmarkInfiniteLimitedPriorityChannelTransfer10kConcurrent(func(fun func(interface{}) bool, limit int) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteLimitPriorityLimitedPriorityHashListChannel(fun, limit)
	}, b)
}

func BenchmarkInfiniteHashLimitedPriorityHashListChannelTransfer10kConcurrent(b *testing.B) {
	benchmarkDefaultInfiniteChannelTransfer10kConcurrent(func() *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteHashLimitedPriorityHashListChannel()
	}, b)
}

func BenchmarkInfinitePriorityHashLimitedPriorityHashListChannelTransfer10kConcurrent(b *testing.B) {
	benchmarkInfinitePriorityChannelTransfer10kConcurrent(func(fun func(i interface{}) bool) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfinitePriorityHashLimitedPriorityHashListChannel(fun)
	}, b)
}

func BenchmarkInfiniteLimitHashLimitedPriorityHashListChannelNoLimitTransfer10kConcurrent(b *testing.B) {
	benchmarkInfiniteLimitLimitedPriorityChannelNoLimitTransfer10kConcurrent(func(limit int) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteLimitHashLimitedPriorityHashListChannel(limit)
	}, b)
}

func BenchmarkInfiniteLimitHashLimitedPriorityHashListChannelTransfer10kConcurrent(b *testing.B) {
	benchmarkInfiniteLimitLimitedPriorityChannelTransfer10kConcurrent(func(limit int) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteLimitHashLimitedPriorityHashListChannel(limit)
	}, b)
}

func BenchmarkInfiniteLimitedPriorityHashListChannelNoLimitTransfer10kConcurrent(b *testing.B) {
	benchmarkInfiniteLimitedPriorityChannelNoLimitTransfer10kConcurrent(func(fun func(interface{}) bool, limit int) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteLimitedPriorityHashListChannel(fun, limit)
	}, b)
}

func BenchmarkInfiniteLimitedPriorityHashListChannelTransfer10kConcurrent(b *testing.B) {
	benchmarkInfiniteLimitedPriorityChannelTransfer10kConcurrent(func(fun func(interface{}) bool, limit int) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteLimitedPriorityHashListChannel(fun, limit)
	}, b)
}

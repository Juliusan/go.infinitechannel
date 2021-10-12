package infinitechannel_bench

import (
	"testing"

	"github.com/Juliusan/go.infinitechannel/src/infinitechannel"
)

func BenchmarkInfiniteDefaultLimitedPriorityHashChannelTransfer10kSerial(b *testing.B) {
	benchmarkDefaultInfiniteChannelTransfer10kSerial(func() *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteDefaultLimitedPriorityHashChannel()
	}, b)
}

func BenchmarkInfinitePriorityLimitedPriorityHashChannelTransfer10kSerial(b *testing.B) {
	benchmarkInfinitePriorityChannelTransfer10kSerial(func(fun func(i interface{}) bool) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfinitePriorityLimitedPriorityHashChannel(fun)
	}, b)
}

func BenchmarkInfiniteLimitLimitedPriorityHashChannelNoLimitTransfer10kSerial(b *testing.B) {
	benchmarkInfiniteLimitLimitedPriorityChannelNoLimitTransfer10kSerial(func(limit int) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteLimitLimitedPriorityHashChannel(limit)
	}, b)
}

func BenchmarkInfiniteLimitLimitedPriorityHashChannelTransfer10kSerial(b *testing.B) {
	benchmarkInfiniteLimitLimitedPriorityChannelTransfer10kSerial(func(limit int) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteLimitLimitedPriorityHashChannel(limit)
	}, b)
}

func BenchmarkInfiniteLimitPriorityLimitedPriorityHashChannelNoLimitTransfer10kSerial(b *testing.B) {
	benchmarkInfiniteLimitedPriorityChannelNoLimitTransfer10kSerial(func(fun func(interface{}) bool, limit int) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteLimitPriorityLimitedPriorityHashChannel(fun, limit)
	}, b)
}

func BenchmarkInfiniteLimitPriorityLimitedPriorityHashChannelTransfer10kSerial(b *testing.B) {
	benchmarkInfiniteLimitedPriorityChannelTransfer10kSerial(func(fun func(interface{}) bool, limit int) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteLimitPriorityLimitedPriorityHashChannel(fun, limit)
	}, b)
}

func BenchmarkInfiniteHashLimitedPriorityHashChannelTransfer10kSerial(b *testing.B) {
	benchmarkDefaultInfiniteChannelTransfer10kSerial(func() *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteHashLimitedPriorityHashChannel(func(elem interface{}) interface{} { return elem })
	}, b)
}

func BenchmarkInfinitePriorityHashLimitedPriorityHashChannelTransfer10kSerial(b *testing.B) {
	benchmarkInfinitePriorityChannelTransfer10kSerial(func(fun func(i interface{}) bool) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfinitePriorityHashLimitedPriorityHashChannel(fun, func(elem interface{}) interface{} { return elem })
	}, b)
}

func BenchmarkInfiniteLimitHashLimitedPriorityHashChannelNoLimitTransfer10kSerial(b *testing.B) {
	benchmarkInfiniteLimitLimitedPriorityChannelNoLimitTransfer10kSerial(func(limit int) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteLimitHashLimitedPriorityHashChannel(limit, func(elem interface{}) interface{} { return elem })
	}, b)
}

func BenchmarkInfiniteLimitHashLimitedPriorityHashChannelTransfer10kSerial(b *testing.B) {
	benchmarkInfiniteLimitLimitedPriorityChannelTransfer10kSerial(func(limit int) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteLimitHashLimitedPriorityHashChannel(limit, func(elem interface{}) interface{} { return elem })
	}, b)
}

func BenchmarkInfiniteLimitedPriorityHashChannelNoLimitTransfer10kSerial(b *testing.B) {
	benchmarkInfiniteLimitedPriorityChannelNoLimitTransfer10kSerial(func(fun func(interface{}) bool, limit int) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteLimitedPriorityHashChannel(fun, limit, func(elem interface{}) interface{} { return elem })
	}, b)
}

func BenchmarkInfiniteLimitedPriorityHashChannelTransfer10kSerial(b *testing.B) {
	benchmarkInfiniteLimitedPriorityChannelTransfer10kSerial(func(fun func(interface{}) bool, limit int) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteLimitedPriorityHashChannel(fun, limit, func(elem interface{}) interface{} { return elem })
	}, b)
}

//--

func BenchmarkInfiniteDefaultLimitedPriorityHashChannelTransfer10kConcurrent(b *testing.B) {
	benchmarkDefaultInfiniteChannelTransfer10kConcurrent(func() *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteDefaultLimitedPriorityHashChannel()
	}, b)
}

func BenchmarkInfinitePriorityLimitedPriorityHashChannelTransfer10kConcurrent(b *testing.B) {
	benchmarkInfinitePriorityChannelTransfer10kConcurrent(func(fun func(i interface{}) bool) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfinitePriorityLimitedPriorityHashChannel(fun)
	}, b)
}

func BenchmarkInfiniteLimitLimitedPriorityHashChannelNoLimitTransfer10kConcurrent(b *testing.B) {
	benchmarkInfiniteLimitLimitedPriorityChannelNoLimitTransfer10kConcurrent(func(limit int) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteLimitLimitedPriorityHashChannel(limit)
	}, b)
}

func BenchmarkInfiniteLimitLimitedPriorityHashChannelTransfer10kConcurrent(b *testing.B) {
	benchmarkInfiniteLimitLimitedPriorityChannelTransfer10kConcurrent(func(limit int) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteLimitLimitedPriorityHashChannel(limit)
	}, b)
}

func BenchmarkInfiniteLimitPriorityLimitedPriorityHashChannelNoLimitTransfer10kConcurrent(b *testing.B) {
	benchmarkInfiniteLimitedPriorityChannelNoLimitTransfer10kConcurrent(func(fun func(interface{}) bool, limit int) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteLimitPriorityLimitedPriorityHashChannel(fun, limit)
	}, b)
}

func BenchmarkInfiniteLimitPriorityLimitedPriorityHashChannelTransfer10kConcurrent(b *testing.B) {
	benchmarkInfiniteLimitedPriorityChannelTransfer10kConcurrent(func(fun func(interface{}) bool, limit int) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteLimitPriorityLimitedPriorityHashChannel(fun, limit)
	}, b)
}

func BenchmarkInfiniteHashLimitedPriorityHashChannelTransfer10kConcurrent(b *testing.B) {
	benchmarkDefaultInfiniteChannelTransfer10kConcurrent(func() *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteHashLimitedPriorityHashChannel(func(elem interface{}) interface{} { return elem })
	}, b)
}

func BenchmarkInfinitePriorityHashLimitedPriorityHashChannelTransfer10kConcurrent(b *testing.B) {
	benchmarkInfinitePriorityChannelTransfer10kConcurrent(func(fun func(i interface{}) bool) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfinitePriorityHashLimitedPriorityHashChannel(fun, func(elem interface{}) interface{} { return elem })
	}, b)
}

func BenchmarkInfiniteLimitHashLimitedPriorityHashChannelNoLimitTransfer10kConcurrent(b *testing.B) {
	benchmarkInfiniteLimitLimitedPriorityChannelNoLimitTransfer10kConcurrent(func(limit int) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteLimitHashLimitedPriorityHashChannel(limit, func(elem interface{}) interface{} { return elem })
	}, b)
}

func BenchmarkInfiniteLimitHashLimitedPriorityHashChannelTransfer10kConcurrent(b *testing.B) {
	benchmarkInfiniteLimitLimitedPriorityChannelTransfer10kConcurrent(func(limit int) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteLimitHashLimitedPriorityHashChannel(limit, func(elem interface{}) interface{} { return elem })
	}, b)
}

func BenchmarkInfiniteLimitedPriorityHashChannelNoLimitTransfer10kConcurrent(b *testing.B) {
	benchmarkInfiniteLimitedPriorityChannelNoLimitTransfer10kConcurrent(func(fun func(interface{}) bool, limit int) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteLimitedPriorityHashChannel(fun, limit, func(elem interface{}) interface{} { return elem })
	}, b)
}

func BenchmarkInfiniteLimitedPriorityHashChannelTransfer10kConcurrent(b *testing.B) {
	benchmarkInfiniteLimitedPriorityChannelTransfer10kConcurrent(func(fun func(interface{}) bool, limit int) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteLimitedPriorityHashChannel(fun, limit, func(elem interface{}) interface{} { return elem })
	}, b)
}

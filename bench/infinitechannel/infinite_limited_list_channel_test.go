package infinitechannel_bench

import (
	"testing"

	"github.com/Juliusan/go.infinitechannel/src/infinitechannel"
)

func BenchmarkInfiniteDefaultLimitedPriorityListChannelTransfer10kSerial(b *testing.B) {
	benchmarkDefaultInfiniteChannelTransfer10kSerial(func() *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteDefaultLimitedPriorityListChannel()
	}, b)
}

func BenchmarkInfinitePriorityLimitedPriorityListChannelTransfer10kSerial(b *testing.B) {
	benchmarkInfinitePriorityChannelTransfer10kSerial(func(fun func(i interface{}) bool) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfinitePriorityLimitedPriorityListChannel(fun)
	}, b)
}

func BenchmarkInfiniteLimitLimitedPriorityListChannelNoLimitTransfer10kSerial(b *testing.B) {
	benchmarkInfiniteLimitLimitedPriorityChannelNoLimitTransfer10kSerial(func(limit int) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteLimitLimitedPriorityListChannel(limit)
	}, b)
}

func benchmarkInfiniteLimitLimitedPriorityChannelNoLimitTransfer10kSerial(makeLimitedChannelFun func(limit int) *infinitechannel.InfiniteChannel, b *testing.B) {
	benchmarkDefaultInfiniteChannelTransfer10kSerial(func() *infinitechannel.InfiniteChannel {
		return makeLimitedChannelFun(12000)
	}, b)
}

func BenchmarkInfiniteLimitLimitedPriorityListChannelTransfer10kSerial(b *testing.B) {
	benchmarkInfiniteLimitLimitedPriorityChannelTransfer10kSerial(func(limit int) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteLimitLimitedPriorityListChannel(limit)
	}, b)
}

func benchmarkInfiniteLimitLimitedPriorityChannelTransfer10kSerial(makeLimitedChannelFun func(limit int) *infinitechannel.InfiniteChannel, b *testing.B) {
	benchmarkInfiniteChannelTransfer10kSerial(func() *infinitechannel.InfiniteChannel {
		return makeLimitedChannelFun(10000)
	}, 12000, b)
}

func BenchmarkInfiniteLimitedPriorityListChannelNoLimitTransfer10kSerial(b *testing.B) {
	benchmarkInfiniteLimitedPriorityChannelNoLimitTransfer10kSerial(func(fun func(interface{}) bool, limit int) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteLimitedPriorityListChannel(fun, limit)
	}, b)
}

func benchmarkInfiniteLimitedPriorityChannelNoLimitTransfer10kSerial(makeLimitedPriorityChannelFun func(fun func(interface{}) bool, limit int) *infinitechannel.InfiniteChannel, b *testing.B) {
	benchmarkInfinitePriorityChannelTransfer10kSerial(func(fun func(i interface{}) bool) *infinitechannel.InfiniteChannel {
		return makeLimitedPriorityChannelFun(fun, 12000)
	}, b)
}

func BenchmarkInfiniteLimitedPriorityListChannelTransfer10kSerial(b *testing.B) {
	benchmarkInfiniteLimitedPriorityChannelTransfer10kSerial(func(fun func(interface{}) bool, limit int) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteLimitedPriorityListChannel(fun, limit)
	}, b)
}

func benchmarkInfiniteLimitedPriorityChannelTransfer10kSerial(makeLimitedPriorityChannelFun func(fun func(interface{}) bool, limit int) *infinitechannel.InfiniteChannel, b *testing.B) {
	makeInfiniteChannelFun := func() *infinitechannel.InfiniteChannel {
		return makeLimitedPriorityChannelFun(func(i interface{}) bool {
			return i.(int)%2 == 0
		}, 10000)
	}
	benchmarkInfiniteChannelTransfer10kSerial(makeInfiniteChannelFun, 12000, b)
}

//--

func BenchmarkInfiniteDefaultLimitedPriorityListChannelTransfer10kConcurrent(b *testing.B) {
	benchmarkDefaultInfiniteChannelTransfer10kConcurrent(func() *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteDefaultLimitedPriorityListChannel()
	}, b)
}

func BenchmarkInfinitePriorityLimitedPriorityListChannelTransfer10kConcurrent(b *testing.B) {
	benchmarkInfinitePriorityChannelTransfer10kConcurrent(func(fun func(i interface{}) bool) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfinitePriorityLimitedPriorityListChannel(fun)
	}, b)
}

func BenchmarkInfiniteLimitLimitedPriorityListChannelNoLimitTransfer10kConcurrent(b *testing.B) {
	benchmarkInfiniteLimitLimitedPriorityChannelNoLimitTransfer10kConcurrent(func(limit int) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteLimitLimitedPriorityListChannel(limit)
	}, b)
}

func benchmarkInfiniteLimitLimitedPriorityChannelNoLimitTransfer10kConcurrent(makeLimitedChannelFun func(limit int) *infinitechannel.InfiniteChannel, b *testing.B) {
	benchmarkDefaultInfiniteChannelTransfer10kConcurrent(func() *infinitechannel.InfiniteChannel {
		return makeLimitedChannelFun(12000)
	}, b)
}

func BenchmarkInfiniteLimitLimitedPriorityListChannelTransfer10kConcurrent(b *testing.B) {
	benchmarkInfiniteLimitLimitedPriorityChannelTransfer10kConcurrent(func(limit int) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteLimitLimitedPriorityListChannel(limit)
	}, b)
}

func benchmarkInfiniteLimitLimitedPriorityChannelTransfer10kConcurrent(makeLimitedChannelFun func(limit int) *infinitechannel.InfiniteChannel, b *testing.B) {
	benchmarkInfiniteChannelTransfer10kConcurrent(func() *infinitechannel.InfiniteChannel {
		return makeLimitedChannelFun(10000)
	}, 12000, b)
}

func BenchmarkInfiniteLimitedPriorityListChannelNoLimitTransfer10kConcurrent(b *testing.B) {
	benchmarkInfiniteLimitedPriorityChannelNoLimitTransfer10kConcurrent(func(fun func(interface{}) bool, limit int) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteLimitedPriorityListChannel(fun, limit)
	}, b)
}

func benchmarkInfiniteLimitedPriorityChannelNoLimitTransfer10kConcurrent(makeLimitedPriorityChannelFun func(fun func(interface{}) bool, limit int) *infinitechannel.InfiniteChannel, b *testing.B) {
	benchmarkInfinitePriorityChannelTransfer10kConcurrent(func(fun func(i interface{}) bool) *infinitechannel.InfiniteChannel {
		return makeLimitedPriorityChannelFun(fun, 12000)
	}, b)
}

func BenchmarkInfiniteLimitedPriorityListChannelTransfer10kConcurrent(b *testing.B) {
	benchmarkInfiniteLimitedPriorityChannelTransfer10kConcurrent(func(fun func(interface{}) bool, limit int) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteLimitedPriorityListChannel(fun, limit)
	}, b)
}

func benchmarkInfiniteLimitedPriorityChannelTransfer10kConcurrent(makeLimitedPriorityChannelFun func(fun func(interface{}) bool, limit int) *infinitechannel.InfiniteChannel, b *testing.B) {
	makeInfiniteChannelFun := func() *infinitechannel.InfiniteChannel {
		return makeLimitedPriorityChannelFun(func(i interface{}) bool {
			return i.(int)%2 == 0
		}, 10000)
	}
	benchmarkInfiniteChannelTransfer10kConcurrent(makeInfiniteChannelFun, 12000, b)
}

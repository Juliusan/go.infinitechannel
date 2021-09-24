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

func BenchmarkInfiniteLimitLimitedPriorityListChannelTransfer10kSerial(b *testing.B) {
	benchmarkInfiniteLimitLimitedPriorityChannelTransfer10kSerial(func(limit int) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteLimitLimitedPriorityListChannel(limit)
	}, b)
}

func BenchmarkInfiniteLimitedPriorityListChannelNoLimitTransfer10kSerial(b *testing.B) {
	benchmarkInfiniteLimitedPriorityChannelNoLimitTransfer10kSerial(func(fun func(interface{}) bool, limit int) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteLimitedPriorityListChannel(fun, limit)
	}, b)
}

func BenchmarkInfiniteLimitedPriorityListChannelTransfer10kSerial(b *testing.B) {
	benchmarkInfiniteLimitedPriorityChannelTransfer10kSerial(func(fun func(interface{}) bool, limit int) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteLimitedPriorityListChannel(fun, limit)
	}, b)
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

func BenchmarkInfiniteLimitLimitedPriorityListChannelTransfer10kConcurrent(b *testing.B) {
	benchmarkInfiniteLimitLimitedPriorityChannelTransfer10kConcurrent(func(limit int) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteLimitLimitedPriorityListChannel(limit)
	}, b)
}

func BenchmarkInfiniteLimitedPriorityListChannelNoLimitTransfer10kConcurrent(b *testing.B) {
	benchmarkInfiniteLimitedPriorityChannelNoLimitTransfer10kConcurrent(func(fun func(interface{}) bool, limit int) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteLimitedPriorityListChannel(fun, limit)
	}, b)
}

func BenchmarkInfiniteLimitedPriorityListChannelTransfer10kConcurrent(b *testing.B) {
	benchmarkInfiniteLimitedPriorityChannelTransfer10kConcurrent(func(fun func(interface{}) bool, limit int) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteLimitedPriorityListChannel(fun, limit)
	}, b)
}

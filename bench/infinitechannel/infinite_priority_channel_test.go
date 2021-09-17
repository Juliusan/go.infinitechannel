package infinitechannel_bench

import (
	"testing"

	"github.com/Juliusan/go.infinitechannel/src/infinitechannel"
)

func BenchmarkInfiniteDefaultPriorityChannelTransfer10kSerial(b *testing.B) {
	benchmarkInfiniteChannelTransfer10kSerial(func() *infinitechannel.InfiniteChannel { return infinitechannel.NewInfiniteDefaultPriorityChannel() }, b)
}

func BenchmarkInfinitePriorityChannelTransfer10kSerial(b *testing.B) {
	makeInfiniteChannelFun := func() *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfinitePriorityChannel(func(i interface{}) bool {
			return i.(int)%2 == 0
		})
	}
	benchmarkInfiniteChannelTransfer10kSerial(makeInfiniteChannelFun, b)
}

func BenchmarkInfiniteDefaultPriorityChannelTransfer10kConcurrent(b *testing.B) {
	benchmarkInfiniteChannelTransfer10kConcurrent(func() *infinitechannel.InfiniteChannel { return infinitechannel.NewInfiniteDefaultPriorityChannel() }, b)
}

func BenchmarkInfinitePriorityChannelTransfer10kConcurrent(b *testing.B) {
	makeInfiniteChannelFun := func() *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfinitePriorityChannel(func(i interface{}) bool {
			return i.(int)%2 == 0
		})
	}
	benchmarkInfiniteChannelTransfer10kConcurrent(makeInfiniteChannelFun, b)
}

package infinitechannel_bench

import (
	"testing"

	"github.com/Juliusan/go.infinitechannel/src/infinitechannel"
	"github.com/Juliusan/go.infinitechannel/src/util"
)

func BenchmarkInfiniteDefaultPriorityChannelTransfer10kSerial(b *testing.B) {
	benchmarkDefaultInfiniteChannelTransfer10kSerial(func() *infinitechannel.InfiniteChannel { return infinitechannel.NewInfiniteDefaultPriorityChannel() }, b)
}

func BenchmarkInfinitePriorityChannelTransfer10kSerial(b *testing.B) {
	benchmarkInfinitePriorityChannelTransfer10kSerial(func(fun func(i interface{}) bool) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfinitePriorityChannel(fun)
	}, b)
}

func benchmarkInfinitePriorityChannelTransfer10kSerial(makeInfinitePriorityChannelFun func(func(i interface{}) bool) *infinitechannel.InfiniteChannel, b *testing.B) {
	makeInfiniteChannelFun := func() *infinitechannel.InfiniteChannel { return makeInfinitePriorityChannelFun(util.PriorityFunMod10) }
	benchmarkDefaultInfiniteChannelTransfer10kSerial(makeInfiniteChannelFun, b)
}

//--

func BenchmarkInfiniteDefaultPriorityChannelTransfer10kConcurrent(b *testing.B) {
	benchmarkDefaultInfiniteChannelTransfer10kConcurrent(func() *infinitechannel.InfiniteChannel { return infinitechannel.NewInfiniteDefaultPriorityChannel() }, b)
}

func BenchmarkInfinitePriorityChannelTransfer10kConcurrent(b *testing.B) {
	benchmarkInfinitePriorityChannelTransfer10kConcurrent(func(fun func(i interface{}) bool) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfinitePriorityChannel(fun)
	}, b)
}

func benchmarkInfinitePriorityChannelTransfer10kConcurrent(makeInfinitePriorityChannelFun func(func(i interface{}) bool) *infinitechannel.InfiniteChannel, b *testing.B) {
	makeInfiniteChannelFun := func() *infinitechannel.InfiniteChannel { return makeInfinitePriorityChannelFun(util.PriorityFunMod10) }
	benchmarkDefaultInfiniteChannelTransfer10kConcurrent(makeInfiniteChannelFun, b)
}

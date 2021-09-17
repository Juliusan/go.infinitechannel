package infinitechannel_bench

import (
	"testing"

	"github.com/Juliusan/go.infinitechannel/src/infinitechannel"
)

func BenchmarkInfiniteSimpleChannelTransfer10Serial(b *testing.B) {
	benchmarkInfiniteChannelTransfer10kSerial(func() *infinitechannel.InfiniteChannel { return infinitechannel.NewInfiniteChannel() }, b)
}

func BenchmarkInfiniteSimpleChannelTransfer10Concurrent(b *testing.B) {
	benchmarkInfiniteChannelTransfer10kConcurrent(func() *infinitechannel.InfiniteChannel { return infinitechannel.NewInfiniteChannel() }, b)
}

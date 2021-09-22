package infinitechannel_bench

import (
	"testing"

	"github.com/Juliusan/go.infinitechannel/src/infinitechannel"
)

func BenchmarkInfiniteSimpleChannelTransfer10kSerial(b *testing.B) {
	benchmarkDefaultInfiniteChannelTransfer10kSerial(func() *infinitechannel.InfiniteChannel { return infinitechannel.NewInfiniteChannel() }, b)
}

func BenchmarkInfiniteSimpleChannelTransfer10kConcurrent(b *testing.B) {
	benchmarkDefaultInfiniteChannelTransfer10kConcurrent(func() *infinitechannel.InfiniteChannel { return infinitechannel.NewInfiniteChannel() }, b)
}

package infinitechannel_bench

import (
	"sync"
	"testing"

	"github.com/Juliusan/go.infinitechannel/src/infinitechannel"
)

var ch *infinitechannel.InfiniteChannel

func benchmarkInfiniteChannelTransfer10kSerial(makeInfiniteChannelFun func() *infinitechannel.InfiniteChannel, b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		ch = makeInfiniteChannelFun()
		b.StartTimer()

		for i := 0; i < 10000; i++ {
			ch.In() <- i
		}
		for i := 0; i < 10000; i++ {
			<-ch.Out()
		}
	}
}

func benchmarkInfiniteChannelTransfer10kConcurrent(makeInfiniteChannelFun func() *infinitechannel.InfiniteChannel, b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		var wg sync.WaitGroup
		wg.Add(2)
		ch = makeInfiniteChannelFun()
		b.StartTimer()

		go func() {
			for i := 0; i < 10000; i++ {
				ch.In() <- i
			}
			wg.Done()
		}()

		go func() {
			for i := 0; i < 10000; i++ {
				<-ch.Out()
			}
			wg.Done()
		}()

		wg.Wait()
	}
}

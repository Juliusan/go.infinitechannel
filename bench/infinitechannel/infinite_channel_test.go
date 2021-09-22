package infinitechannel_bench

import (
	"sync"
	"testing"

	"github.com/Juliusan/go.infinitechannel/src/infinitechannel"
)

var ch *infinitechannel.InfiniteChannel

func benchmarkDefaultInfiniteChannelTransfer10kSerial(makeInfiniteChannelFun func() *infinitechannel.InfiniteChannel, b *testing.B) {
	benchmarkInfiniteChannelTransfer10kSerial(makeInfiniteChannelFun, 10000, b)
}

func benchmarkInfiniteChannelTransfer10kSerial(makeInfiniteChannelFun func() *infinitechannel.InfiniteChannel, elementsToWrite int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		ch = makeInfiniteChannelFun()
		moreToWrite := elementsToWrite - 10000
		for i := 0; i < moreToWrite; i++ {
			ch.In() <- i
		}
		b.StartTimer()

		for i := moreToWrite; i < elementsToWrite; i++ {
			ch.In() <- i
		}
		for i := 0; i < 10000; i++ {
			<-ch.Out()
		}
	}
}

//--

func benchmarkDefaultInfiniteChannelTransfer10kConcurrent(makeInfiniteChannelFun func() *infinitechannel.InfiniteChannel, b *testing.B) {
	benchmarkInfiniteChannelTransfer10kConcurrent(makeInfiniteChannelFun, 10000, b)
}

func benchmarkInfiniteChannelTransfer10kConcurrent(makeInfiniteChannelFun func() *infinitechannel.InfiniteChannel, elementsToWrite int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		var wg sync.WaitGroup
		wg.Add(2)
		ch = makeInfiniteChannelFun()
		moreToWrite := elementsToWrite - 10000
		for i := 0; i < moreToWrite; i++ {
			ch.In() <- i
		}
		b.StartTimer()

		go func() {
			for i := moreToWrite; i < elementsToWrite; i++ {
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

package infinitechannel_test

import (
	"github.com/Juliusan/go.infinitechannel/src/infinitechannel"
	"testing"
)

func TestInfiniteChannel(t *testing.T) {
	//var ch Channel

	ch := infinitechannel.NewInfiniteChannel()
	testChannel(t, "infinite channel", ch)

	ch = infinitechannel.NewInfiniteChannel()
	testChannelPair(t, "infinite channel", ch, ch)

	ch = infinitechannel.NewInfiniteChannel()
	testChannelConcurrentAccessors(t, "infinite channel", ch)
}

func BenchmarkInfiniteChannelSerial(b *testing.B) {
	ch := infinitechannel.NewInfiniteChannel()
	for i := 0; i < b.N; i++ {
		ch.In() <- nil
	}
	for i := 0; i < b.N; i++ {
		<-ch.Out()
	}
}

func BenchmarkInfiniteChannelParallel(b *testing.B) {
	ch := infinitechannel.NewInfiniteChannel()
	go func() {
		for i := 0; i < b.N; i++ {
			<-ch.Out()
		}
		ch.Close()
	}()
	for i := 0; i < b.N; i++ {
		ch.In() <- nil
	}
	<-ch.Out()
}

func BenchmarkInfiniteChannelTickTock(b *testing.B) {
	ch := infinitechannel.NewInfiniteChannel()
	for i := 0; i < b.N; i++ {
		ch.In() <- nil
		<-ch.Out()
	}
}

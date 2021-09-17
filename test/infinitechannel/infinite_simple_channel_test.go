package infinitechannel_test

import (
	"testing"

	"github.com/Juliusan/go.infinitechannel/src/infinitechannel"
)

func TestInfiniteSimpleChannelWriteReadLen(t *testing.T) {
	testInfiniteChannelWriteReadLen(infinitechannel.NewInfiniteChannel(), 1000, func(index int) int { return index }, t)
}

func TestInfiniteSimpleChannelConcurrentWriteReadLen(t *testing.T) {
	result := func(index int) int { return index }
	testInfiniteChannelConcurrentWriteReadLen(infinitechannel.NewInfiniteChannel(), 1000, &result, t)
}

package infinitechannel_test

import (
	"testing"

	"github.com/Juliusan/go.infinitechannel/src/infinitechannel"
	"github.com/Juliusan/go.infinitechannel/src/util"
)

func TestInfiniteDefaultLimitedPriorityChannelWriteReadLen(t *testing.T) {
	testDefaultInfiniteChannelWriteReadLen(infinitechannel.NewInfiniteDefaultLimitedPriorityChannel(), 1000, func(index int) int { return index }, t)
}

func TestInfinitePriorityLimitedPriorityChannelWriteReadLen(t *testing.T) {
	testInfinitePriorityChannelWriteReadLen(func(fun func(i interface{}) bool) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfinitePriorityLimitedPriorityChannel(fun)
	}, t)
}

func TestInfiniteLimitLimitedPriorityChannelNoLimitWriteReadLen(t *testing.T) {
	testInfiniteLimitLimitedPriorityChannelNoLimitWriteReadLen(func(limit int) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteLimitLimitedPriorityChannel(limit)
	}, t)
}

func testInfiniteLimitLimitedPriorityChannelNoLimitWriteReadLen(makeLimitedChannelFun func(limit int) *infinitechannel.InfiniteChannel, t *testing.T) {
	testDefaultInfiniteChannelWriteReadLen(makeLimitedChannelFun(1200), 1000, func(index int) int { return index }, t)
}

func TestInfiniteLimitLimitedPriorityChannelWriteReadLen(t *testing.T) {
	testInfiniteLimitLimitedPriorityChannelWriteReadLen(func(limit int) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteLimitLimitedPriorityChannel(limit)
	}, t)
}

func testInfiniteLimitLimitedPriorityChannelWriteReadLen(makeLimitedChannelFun func(limit int) *infinitechannel.InfiniteChannel, t *testing.T) {
	limit := 800
	elementsToAdd := 1000
	indexDiff := elementsToAdd - limit
	result := func(index int) int {
		return index + indexDiff
	}
	testInfiniteChannelWriteReadLen(makeLimitedChannelFun(limit), elementsToAdd, limit, result, t)
}

func TestInfiniteLimitedPriorityChannelNoLimitWriteReadLen(t *testing.T) {
	testInfiniteLimitedPriorityChannelNoLimitWriteReadLen(func(fun func(i interface{}) bool, limit int) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteLimitedPriorityChannel(fun, limit)
	}, t)
}

func testInfiniteLimitedPriorityChannelNoLimitWriteReadLen(makeLimitedPriorityChannelFun func(fun func(i interface{}) bool, limit int) *infinitechannel.InfiniteChannel, t *testing.T) {
	testInfinitePriorityChannelWriteReadLen(func(fun func(i interface{}) bool) *infinitechannel.InfiniteChannel {
		return makeLimitedPriorityChannelFun(fun, 1200)
	}, t)
}

func TestInfiniteLimitedPriorityChannelWriteReadLen(t *testing.T) {
	testInfiniteLimitedPriorityChannelWriteReadLen(func(fun func(i interface{}) bool, limit int) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteLimitedPriorityChannel(fun, limit)
	}, t)
}

func testInfiniteLimitedPriorityChannelWriteReadLen(makeLimitedPriorityChannelFun func(fun func(i interface{}) bool, limit int) *infinitechannel.InfiniteChannel, t *testing.T) {
	limit := 800
	ch := makeLimitedPriorityChannelFun(util.PriorityFunMod3, limit)
	result := func(index int) int {
		if index <= 333 {
			return -3*index + 999
		} else {
			if index%2 == 0 {
				return 3*index/2 - 200
			} else {
				return (3*index - 401) / 2
			}
		}
	}
	testInfiniteChannelWriteReadLen(ch, 1000, limit, result, t)
}

//--

func TestInfiniteDefaultLimitedPriorityChannelConcurrentWriteReadLen(t *testing.T) {
	result := func(index int) int { return index }
	testDefaultInfiniteChannelConcurrentWriteReadLen(infinitechannel.NewInfiniteDefaultLimitedPriorityChannel(), 1000, &result, t)
}

func TestInfinitePriorityLimitedPriorityChannelConcurrentWriteReadLen(t *testing.T) {
	testInfinitePriorityChannelConcurrentWriteReadLen(func(fun func(i interface{}) bool) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfinitePriorityLimitedPriorityChannel(fun)
	}, t)
}

func TestInfiniteLimitLimitedPriorityChannelNoLimitConcurrentWriteReadLen(t *testing.T) {
	testInfiniteLimitLimitedPriorityChannelNoLimitConcurrentWriteReadLen(func(limit int) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteLimitLimitedPriorityChannel(limit)
	}, t)
}

func testInfiniteLimitLimitedPriorityChannelNoLimitConcurrentWriteReadLen(makeLimitedChannelFun func(limit int) *infinitechannel.InfiniteChannel, t *testing.T) {
	result := func(index int) int { return index }
	testDefaultInfiniteChannelConcurrentWriteReadLen(makeLimitedChannelFun(1200), 1000, &result, t)
}

func TestInfiniteLimitLimitedPriorityChannelConcurrentWriteReadLen(t *testing.T) {
	testInfiniteLimitLimitedPriorityChannelConcurrentWriteReadLen(func(limit int) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteLimitLimitedPriorityChannel(limit)
	}, t)
}

func testInfiniteLimitLimitedPriorityChannelConcurrentWriteReadLen(makeLimitedChannelFun func(limit int) *infinitechannel.InfiniteChannel, t *testing.T) {
	testInfiniteChannelConcurrentWriteReadLen(makeLimitedChannelFun(800), 1000, 800, nil, t)
}

func TestInfiniteLimitedPriorityChannelNoLimitConcurrentWriteReadLen(t *testing.T) {
	testInfiniteLimitedPriorityChannelNoLimitConcurrentWriteReadLen(func(fun func(i interface{}) bool, limit int) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteLimitedPriorityChannel(fun, limit)
	}, t)
}

func testInfiniteLimitedPriorityChannelNoLimitConcurrentWriteReadLen(makeLimitedPriorityChannelFun func(fun func(i interface{}) bool, limit int) *infinitechannel.InfiniteChannel, t *testing.T) {
	testInfinitePriorityChannelConcurrentWriteReadLen(func(fun func(i interface{}) bool) *infinitechannel.InfiniteChannel {
		return makeLimitedPriorityChannelFun(fun, 1200)
	}, t)
}

func TestInfiniteLimitedPriorityChannelConcurrentWriteReadLen(t *testing.T) {
	testInfiniteLimitedPriorityChannelConcurrentWriteReadLen(func(fun func(i interface{}) bool, limit int) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteLimitedPriorityChannel(fun, limit)
	}, t)
}

func testInfiniteLimitedPriorityChannelConcurrentWriteReadLen(makeLimitedPriorityChannelFun func(fun func(i interface{}) bool, limit int) *infinitechannel.InfiniteChannel, t *testing.T) {
	limit := 800
	ch := makeLimitedPriorityChannelFun(util.PriorityFunMod3, limit)
	testInfiniteChannelConcurrentWriteReadLen(ch, 1000, limit, nil, t)
}

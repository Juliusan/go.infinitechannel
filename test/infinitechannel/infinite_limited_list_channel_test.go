package infinitechannel_test

import (
	"testing"

	"github.com/Juliusan/go.infinitechannel/src/infinitechannel"
)

func TestInfiniteDefaultLimitedPriorityListChannelWriteReadLen(t *testing.T) {
	testDefaultInfiniteChannelWriteReadLen(infinitechannel.NewInfiniteDefaultLimitedPriorityListChannel(), 1000, func(index int) int { return index }, t)
}

func TestInfinitePriorityLimitedPriorityListChannelWriteReadLen(t *testing.T) {
	testInfinitePriorityChannelWriteReadLen(func(fun func(i interface{}) bool) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfinitePriorityLimitedPriorityListChannel(fun)
	}, t)
}

func TestInfiniteLimitLimitedPriorityListChannelNoLimitWriteReadLen(t *testing.T) {
	testInfiniteLimitLimitedPriorityChannelNoLimitWriteReadLen(func(limit int) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteLimitLimitedPriorityListChannel(limit)
	}, t)
}

func testInfiniteLimitLimitedPriorityChannelNoLimitWriteReadLen(makeLimitedChannelFun func(limit int) *infinitechannel.InfiniteChannel, t *testing.T) {
	testDefaultInfiniteChannelWriteReadLen(makeLimitedChannelFun(1200), 1000, func(index int) int { return index }, t)
}

func TestInfiniteLimitLimitedPriorityListChannelWriteReadLen(t *testing.T) {
	testInfiniteLimitLimitedPriorityChannelWriteReadLen(func(limit int) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteLimitLimitedPriorityListChannel(limit)
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

func TestInfiniteLimitedPriorityListChannelNoLimitWriteReadLen(t *testing.T) {
	testInfiniteLimitedPriorityChannelNoLimitWriteReadLen(func(fun func(i interface{}) bool, limit int) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteLimitedPriorityListChannel(fun, limit)
	}, t)
}

func testInfiniteLimitedPriorityChannelNoLimitWriteReadLen(makeLimitedPriorityChannelFun func(fun func(i interface{}) bool, limit int) *infinitechannel.InfiniteChannel, t *testing.T) {
	testInfinitePriorityChannelWriteReadLen(func(fun func(i interface{}) bool) *infinitechannel.InfiniteChannel {
		return makeLimitedPriorityChannelFun(fun, 1200)
	}, t)
}

func TestInfiniteLimitedPriorityListChannelWriteReadLen(t *testing.T) {
	testInfiniteLimitedPriorityChannelWriteReadLen(func(fun func(i interface{}) bool, limit int) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteLimitedPriorityListChannel(fun, limit)
	}, t)
}

func testInfiniteLimitedPriorityChannelWriteReadLen(makeLimitedPriorityChannelFun func(fun func(i interface{}) bool, limit int) *infinitechannel.InfiniteChannel, t *testing.T) {
	limit := 800
	ch := makeLimitedPriorityChannelFun(func(i interface{}) bool {
		return i.(int)%3 == 0
	}, limit)
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

func TestInfiniteDefaultLimitedPriorityListChannelConcurrentWriteReadLen(t *testing.T) {
	result := func(index int) int { return index }
	testDefaultInfiniteChannelConcurrentWriteReadLen(infinitechannel.NewInfiniteDefaultLimitedPriorityListChannel(), 1000, &result, t)
}

func TestInfinitePriorityLimitedPriorityListChannelConcurrentWriteReadLen(t *testing.T) {
	testInfinitePriorityChannelConcurrentWriteReadLen(func(fun func(i interface{}) bool) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfinitePriorityLimitedPriorityListChannel(fun)
	}, t)
}

func TestInfiniteLimitLimitedPriorityListChannelNoLimitConcurrentWriteReadLen(t *testing.T) {
	testInfiniteLimitLimitedPriorityChannelNoLimitConcurrentWriteReadLen(func(limit int) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteLimitLimitedPriorityListChannel(limit)
	}, t)
}

func testInfiniteLimitLimitedPriorityChannelNoLimitConcurrentWriteReadLen(makeLimitedChannelFun func(limit int) *infinitechannel.InfiniteChannel, t *testing.T) {
	result := func(index int) int { return index }
	testDefaultInfiniteChannelConcurrentWriteReadLen(makeLimitedChannelFun(1200), 1000, &result, t)
}

func TestInfiniteLimitLimitedPriorityListChannelConcurrentWriteReadLen(t *testing.T) {
	testInfiniteLimitLimitedPriorityChannelConcurrentWriteReadLen(func(limit int) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteLimitLimitedPriorityListChannel(limit)
	}, t)
}

func testInfiniteLimitLimitedPriorityChannelConcurrentWriteReadLen(makeLimitedChannelFun func(limit int) *infinitechannel.InfiniteChannel, t *testing.T) {
	testInfiniteChannelConcurrentWriteReadLen(makeLimitedChannelFun(800), 1000, 800, nil, t)
}

func TestInfiniteLimitedPriorityListChannelNoLimitConcurrentWriteReadLen(t *testing.T) {
	testInfiniteLimitedPriorityChannelNoLimitConcurrentWriteReadLen(func(fun func(i interface{}) bool, limit int) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteLimitedPriorityListChannel(fun, limit)
	}, t)
}

func testInfiniteLimitedPriorityChannelNoLimitConcurrentWriteReadLen(makeLimitedPriorityChannelFun func(fun func(i interface{}) bool, limit int) *infinitechannel.InfiniteChannel, t *testing.T) {
	testInfinitePriorityChannelConcurrentWriteReadLen(func(fun func(i interface{}) bool) *infinitechannel.InfiniteChannel {
		return makeLimitedPriorityChannelFun(fun, 1200)
	}, t)
}

func TestInfiniteLimitedPriorityListChannelConcurrentWriteReadLen(t *testing.T) {
	testInfiniteLimitedPriorityChannelConcurrentWriteReadLen(func(fun func(i interface{}) bool, limit int) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteLimitedPriorityListChannel(fun, limit)
	}, t)
}

func testInfiniteLimitedPriorityChannelConcurrentWriteReadLen(makeLimitedPriorityChannelFun func(fun func(i interface{}) bool, limit int) *infinitechannel.InfiniteChannel, t *testing.T) {
	limit := 800
	ch := makeLimitedPriorityChannelFun(func(i interface{}) bool {
		return i.(int)%3 == 0
	}, limit)
	testInfiniteChannelConcurrentWriteReadLen(ch, 1000, limit, nil, t)
}

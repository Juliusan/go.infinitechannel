package infinitechannel_test

import (
	"sync"
	"testing"
	"time"

	"github.com/Juliusan/go.infinitechannel/src/infinitechannel"
)

func testDefaultInfiniteChannelWriteReadLen(ch *infinitechannel.InfiniteChannel, elementsToWrite int, result func(index int) int, t *testing.T) {
	testInfiniteChannelWriteReadLen(ch, elementsToWrite, elementsToWrite, result, t)
}

func testInfiniteChannelWriteReadLen(ch *infinitechannel.InfiniteChannel, elementsToWrite int, elementsToRead int, result func(index int) int, t *testing.T) {
	for i := 0; i < elementsToWrite; i++ {
		ch.In() <- i
	}
	obtained := ch.Len()
	if obtained != elementsToRead {
		t.Errorf("expected full channel length %d, obtained %d", elementsToWrite, obtained)
	}
	ch.Close()
	obtained = ch.Len()
	if obtained != elementsToRead {
		t.Errorf("expected closed channel length %d, obtained %d", elementsToWrite, obtained)
	}
	for i := 0; i < elementsToRead; i++ {
		val := <-ch.Out()
		expected := result(i)
		obtained := val.(int)
		if obtained != expected {
			t.Errorf("read %d obtained %d instead of %d", i, obtained, expected)
		}
	}
}

//--

func testDefaultInfiniteChannelConcurrentWriteReadLen(ch *infinitechannel.InfiniteChannel, elementsToWrite int, result *func(index int) int, t *testing.T) {
	testInfiniteChannelConcurrentWriteReadLen(ch, elementsToWrite, elementsToWrite, result, t)
}

func testInfiniteChannelConcurrentWriteReadLen(ch *infinitechannel.InfiniteChannel, elementsToWrite int, elementsToRead int, result *func(index int) int, t *testing.T) {
	var wg sync.WaitGroup
	written := 0
	read := 0
	wg.Add(2)

	go func() {
		for i := 0; i < elementsToWrite; i++ {
			ch.In() <- i
			written++
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < elementsToRead; i++ {
			val := <-ch.Out()
			if result != nil {
				expected := (*result)(i)
				obtained := val.(int)
				if obtained != expected {
					t.Errorf("concurent read %d obtained %d instead of %d", i, obtained, expected)
				}
			}
			read++
		}
		wg.Done()
	}()

	go func() {
		for {
			length := ch.Len()
			t.Logf("current channel length is %d", length)
			// no asserts here - the read/write process is asynchronious
			time.Sleep(10 * time.Millisecond)
		}
	}()

	wg.Wait()
	if written != elementsToWrite {
		t.Errorf("concurent write written %d should have %d", written, elementsToWrite)
	}
	if read != elementsToRead {
		t.Errorf("concurent read read %d should have %d", read, elementsToRead)
	}
}

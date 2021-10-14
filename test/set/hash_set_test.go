package queue_test

import (
	"testing"

	"github.com/Juliusan/go.infinitechannel/src/set"
)

type TestElem struct {
	hashField      int
	importantField int
	uselessField   int
}

var _ set.Hashable = &TestElem{}

func (te *TestElem) GetHash() interface{} {
	return te.hashField
}

func (te *TestElem) Equals(elem interface{}) bool {
	other, ok := elem.(*TestElem)
	if !ok {
		return false
	}
	return te.hashField == other.hashField && te.importantField == other.importantField
}

func TestHashSetSimple(t *testing.T) {
	testElems := []*TestElem{
		&TestElem{1, 2, 3},
		&TestElem{4, 5, 6},
		&TestElem{7, 8, 9},
		&TestElem{6, 5, 4},
		&TestElem{3, 2, 1},
	}
	addResults := make([]bool, len(testElems))
	for i := 0; i < len(addResults); i++ {
		addResults[i] = true
	}
	testHashSetCreateSizeAddContainsRemove(t, testElems, addResults, func(i int) int { return i })
}

func TestHashSetCollisions(t *testing.T) {
	testElems := []*TestElem{
		&TestElem{1, 2, 3},
		&TestElem{1, 4, 5},
		&TestElem{6, 7, 8},
		&TestElem{1, 9, 0},
		&TestElem{1, 7, 6},
		&TestElem{1, 3, 2},
		&TestElem{1, 1, 2},
		&TestElem{0, 1, 2},
	}
	addResults := make([]bool, len(testElems))
	for i := 0; i < len(addResults); i++ {
		addResults[i] = true
	}
	testHashSetCreateSizeAddContainsRemove(t, testElems, addResults, func(i int) int { return i })
}

func TestHashSetDuplicates(t *testing.T) {
	testElems := []*TestElem{
		&TestElem{1, 2, 3},
		&TestElem{4, 5, 6},
		&TestElem{1, 7, 8},
		&TestElem{1, 2, 3},
		&TestElem{1, 2, 4},
		&TestElem{1, 5, 6},
		&TestElem{7, 8, 9},
	}
	addResults := []bool{
		true,
		true,
		true,
		false,
		false,
		true,
		true,
	}
	testHashSetCreateSizeAddContainsRemove(t, testElems, addResults, func(i int) int { return i })
}

func TestHashSetReverse(t *testing.T) {
	testElems := []*TestElem{
		&TestElem{1, 2, 3},
		&TestElem{4, 5, 6},
		&TestElem{1, 7, 8},
		&TestElem{1, 2, 3},
		&TestElem{1, 2, 4},
		&TestElem{1, 5, 6},
		&TestElem{7, 8, 9},
	}
	addResults := []bool{
		true,
		true,
		true,
		false,
		false,
		true,
		true,
	}
	testHashSetCreateSizeAddContainsRemove(t, testElems, addResults, func(i int) int { return len(testElems) - i - 1 })
}

func TestHashSetRandom(t *testing.T) {
	testElems := []*TestElem{
		&TestElem{1, 2, 3},
		&TestElem{4, 5, 6},
		&TestElem{1, 7, 8},
		&TestElem{1, 2, 3},
		&TestElem{1, 2, 4},
		&TestElem{1, 5, 6},
		&TestElem{7, 8, 9},
	}
	addResults := []bool{
		true,
		true,
		true,
		false,
		false,
		true,
		true,
	}
	removeOrder := []int{5, 2, 6, 3, 4, 0, 1}
	testHashSetCreateSizeAddContainsRemove(t, testElems, addResults, func(i int) int { return removeOrder[i] })
}

func TestHashSetDouble(t *testing.T) {
	testElems := []*TestElem{
		&TestElem{1, 2, 3},
		&TestElem{4, 5, 6},
		&TestElem{7, 8, 9},
		&TestElem{1, 2, 3},
		&TestElem{1, 2, 4},
		&TestElem{1, 3, 5},
		&TestElem{4, 5, 5},
		&TestElem{6, 7, 8},
		&TestElem{1, 4, 7},
		&TestElem{1, 2, 0},
	}
	addResults := []bool{
		true,
		true,
		true,
		false,
		false,
		true,
		false,
		true,
		true,
		false,
	}
	hs := set.NewHashSet()
	randomRemoveOrder := []int{7, 4, 2, 9, 5, 6, 0, 3, 1, 8}
	testHashSetSizeAddContainsRemove(t, hs, testElems, addResults, func(i int) int { return i })
	testHashSetSizeAddContainsRemove(t, hs, testElems, addResults, func(i int) int { return len(testElems) - i - 1 })
	testHashSetSizeAddContainsRemove(t, hs, testElems, addResults, func(i int) int { return randomRemoveOrder[i] })
	testHashSetSizeAddContainsRemove(t, hs, testElems, addResults, func(i int) int { return i })
}

func testHashSetCreateSizeAddContainsRemove(t *testing.T, testElems []*TestElem, addResults []bool, removeOrderFun func(int) int) {
	hs := set.NewHashSet()
	testHashSetSizeAddContainsRemove(t, hs, testElems, addResults, removeOrderFun)
}

func testHashSetSizeAddContainsRemove(t *testing.T, hs *set.HashSet, testElems []*TestElem, addResults []bool, removeOrderFun func(int) int) {
	expectedSize := 0
	size := hs.Size()
	if size != expectedSize {
		t.Errorf("Newly initialised set size is %v, expected %v", size, expectedSize)
	}
	for i := 0; i < len(testElems); i++ {
		expectedAddResult := addResults[i]
		addResult := hs.Add(testElems[i])
		if addResult != expectedAddResult {
			t.Errorf("Error adding testElem %v: %v (obtained %v, expected %v)", i, testElems[i], addResult, expectedAddResult)
		}
		if expectedAddResult {
			expectedSize++
		}
		size = hs.Size()
		if size != expectedSize {
			t.Errorf("Size after %v add iteration is %v, expected %v", i, size, expectedSize)
		}
	}

	for i := 0; i < len(testElems); i++ {
		containsResult := hs.Contains(testElems[i])
		if !containsResult {
			t.Errorf("Set does not contain testElem %v: %v", i, testElems[i])
		}
	}

	for i := 0; i < len(testElems); i++ {
		removeIndex := removeOrderFun(i)
		if addResults[removeIndex] {
			removeResult := hs.Remove(testElems[removeIndex])
			if !removeResult {
				t.Errorf("Error removing iteration %v testElem %v: %v", i, removeIndex, testElems[removeIndex])
			}
			expectedSize--
			size = hs.Size()
			if size != expectedSize {
				t.Errorf("Size after %v remove iteration is %v, expected %v", i, size, expectedSize)
			}
		}
	}

	size = hs.Size()
	if size != expectedSize {
		t.Errorf("Empty set size is %v, expected %v", size, expectedSize)
	}
}

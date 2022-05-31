package ecstatic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSparseSet(t *testing.T) {
	capacity, highest := uint32(11), uint32(100)
	tooBig := uint32(666)
	var testSet *SparseSet[int]

	t.Run("create sparse set", func(t *testing.T) {
		testSet = NewSparseSet[int](capacity, highest)
		assert.Equal(t, capacity, testSet.capacity)
		assert.Equal(t, uint32(0), testSet.n)
		assert.Equal(t, capacity, uint32(len(testSet.dense)))
		assert.Equal(t, capacity, uint32(len(testSet.payload)))
		assert.Equal(t, highest, testSet.highest)
		assert.Equal(t, highest+1, uint32(len(testSet.sparse)))
	})

	t.Run("insert returns true for capacity-2 elements", func(t *testing.T) {
		assert.Equal(t, uint32(0), testSet.n)
		for i := 0; i <= int(capacity)-2; i++ {
			assert.True(t, testSet.Insert(uint32(i), i))
			assert.Equal(t, i+1, int(testSet.n))
		}
		assert.Equal(t, []uint32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0}, testSet.dense)
		assert.Equal(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0}, testSet.payload)
		assert.Equal(t, uint32(10), testSet.n)
	})

	t.Run("insert returns false if element is too big", func(t *testing.T) {
		assert.Equal(t, uint32(10), testSet.n)
		assert.False(t, testSet.Insert(tooBig, 0))
		assert.Equal(t, uint32(10), testSet.n)
	})

	t.Run("insert returns false if set is full", func(t *testing.T) {
		assert.Equal(t, uint32(10), testSet.n)
		// Fill set with another item.
		assert.True(t, testSet.Insert(capacity-1, int(capacity-1)))
		// Now insert will fail no matter what el is.
		assert.False(t, testSet.Insert(capacity, int(capacity)))
		assert.Equal(t, uint32(11), testSet.n)
	})

	t.Run("insert updates existing element with new payload", func(t *testing.T) {
		assert.Equal(t, uint32(11), testSet.n)
		assert.Equal(t, testSet.payload[0], 0)
		assert.True(t, testSet.Insert(0, 1))
		assert.Equal(t, testSet.payload[0], 1)
		assert.Equal(t, uint32(11), testSet.n)
	})

	t.Run("delete element that does not exist returns false", func(t *testing.T) {
		assert.Equal(t, uint32(11), testSet.n)
		assert.False(t, testSet.Delete(99))
		assert.Equal(t, uint32(11), testSet.n)
	})

	t.Run("deleting first element moves last to first", func(t *testing.T) {
		assert.Equal(t, uint32(11), testSet.n)
		assert.Equal(t, []uint32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, testSet.dense)
		assert.Equal(t, []int{1, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, testSet.payload)
		assert.True(t, testSet.Delete(0))
		// One less. Last is first now. "Dead" element remains unchanged.
		assert.Equal(t, uint32(10), testSet.n)
		assert.Equal(t, []uint32{10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, testSet.dense)
		assert.Equal(t, []int{10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, testSet.payload)
	})

	t.Run("deleting all remaining elements succeeds", func(t *testing.T) {
		assert.Equal(t, uint32(10), testSet.n)
		assert.Equal(t, []uint32{10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, testSet.dense)
		assert.Equal(t, []int{10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, testSet.payload)
		for i := uint32(1); i <= 10; i++ {
			assert.True(t, testSet.Delete(i))
		}
		// The remaining values are all dead and have no meaning.
		assert.Equal(t, uint32(0), testSet.n)
		assert.Equal(t, []uint32{10, 9, 8, 7, 6, 5, 6, 7, 8, 9, 10}, testSet.dense)
		assert.Equal(t, []int{10, 9, 8, 7, 6, 5, 6, 7, 8, 9, 10}, testSet.payload)
	})

	t.Run("clear sets n to 0", func(t *testing.T) {
		// Fill test set.
		for i := uint32(0); i < capacity; i++ {
			assert.True(t, testSet.Insert(i, int(i)))
		}
		assert.Equal(t, uint32(11), testSet.n)
		assert.Equal(t, []uint32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, testSet.dense)
		assert.Equal(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, testSet.payload)
		// Now it's full. Adding more fails.
		assert.False(t, testSet.Insert(capacity, int(capacity)))
		// Now we clear.
		testSet.Clear()
		// Only n has changed.
		assert.Equal(t, uint32(0), testSet.n)
		assert.Equal(t, []uint32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, testSet.dense)
		assert.Equal(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, testSet.payload)
		// Now we can fill the set again with (other) values.
		for i := uint32(0); i < capacity; i++ {
			assert.True(t, testSet.Insert(i, int(i)+10))
		}
		assert.Equal(t, uint32(11), testSet.n)
		assert.Equal(t, []uint32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, testSet.dense)
		assert.Equal(t, []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}, testSet.payload)
	})

	// TODO assert sparse in test cases.
}

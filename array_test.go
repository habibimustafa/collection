package arr

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCollectingArray(t *testing.T) {
	arrString := []string{"Hello", "World"}
	strCollection := List(arrString)
	assert.Equal(t, len(arrString), len(strCollection))
	assert.Equal(t, len(arrString), strCollection.Size())
	assert.Equal(t, "Hello World", strCollection.Implode(" "))

	arrInt := []int{172, 20, 100, 255}
	intCollection := List(arrInt)
	assert.Equal(t, len(arrInt), len(intCollection))
	assert.Equal(t, len(arrInt), intCollection.Size())
	assert.Equal(t, "172.20.100.255", intCollection.Implode("."))

	arrRune := []rune{'h', 'e', 'l', 'l', 'o'}
	runeCollection := List(arrRune)
	assert.Equal(t, len(arrRune), len(runeCollection))
	assert.Equal(t, len(arrRune), runeCollection.Size())
	assert.Equal(t, "104.101.108.108.111", runeCollection.Implode("."))

	arrMap := map[string]string{"First Name": "John", "Last Name": "Doe"}
	mapCollection := List(arrMap)
	assert.Equal(t, len(arrMap), len(mapCollection))
	assert.Equal(t, len(arrMap), mapCollection.Size())
	assert.Equal(t, "John Doe", mapCollection.Implode(" "))
}

func TestGetAllItems(t *testing.T) {
	collection := Array{"Hello", "World"}
	assert.Equal(t, []interface{}{"Hello", "World"}, collection.All())
	assert.Equal(t, "Hello", collection.All()[0])
	assert.Equal(t, []interface{}{"Hello"}, collection.All()[:1])
	assert.Equal(t, []interface{}{"World"}, collection.All()[1:])
}

func TestGetItem(t *testing.T) {
	collection := Array{"Hello", "World"}
	assert.Equal(t, "Hello", collection.Get(0))
	assert.Equal(t, "World", collection.Get(1))
}

func TestGetFirstAndLastItem(t *testing.T) {
	collection := Array{"Hello", "Middle", "World"}
	assert.Equal(t, "Hello", collection.First())
	assert.Equal(t, "World", collection.Last())
}

func TestCollectionIsNotEmpty(t *testing.T) {
	collection := Array{"Hello", "World"}
	assert.True(t, collection.IsNotEmpty())
}

func TestCollectionAppend(t *testing.T) {
	collection := Array{"Hello", "World"}
	collection = collection.Append("Hi")

	assert.Equal(t, 3, len(collection))
	assert.Equal(t, "Hi", collection[2])
}

func TestCollectionPrepend(t *testing.T) {
	collection := Array{"Hello", "World"}
	collection = collection.Prepend("Hi")

	assert.Equal(t, 3, len(collection))
	assert.Equal(t, "Hi", collection[0])
}

func TestCollectionImplode(t *testing.T) {
	collection := Array{"Hello", "World"}
	assert.Equal(t, "Hello World", collection.Implode(" "))
}

func TestCollectionKeys(t *testing.T) {
	collection := List(map[string]string{"first": "John", "last": "Doe"})
	assert.Equal(t, []interface{}{0, 1}, collection.Keys())
}

func TestCollectionIndex(t *testing.T) {
	collection := Array{"Hello", "World"}
	assert.Equal(t, 0, collection.Index("Hello"))
	assert.Equal(t, 1, collection.Index("World"))
	assert.Equal(t, nil, collection.Index("Random"))

	collection = List(map[string]string{"first": "John", "last": "Doe"})
	assert.Equal(t, 0, collection.Index("John"))
	assert.Equal(t, 1, collection.Index("Doe"))
	assert.Equal(t, nil, collection.Index("Random"))
}

func TestCollectionHas(t *testing.T) {
	collection := Array{"Hello", "World"}
	assert.True(t, collection.Has("Hello"))
	assert.False(t, collection.Has("Random"))
}

func TestCollectionEach(t *testing.T) {
	collectionBefore := Array{"Hello", "World"}
	collectionAfter := collectionBefore.Each(func(value interface{}, index int) {
		value = fmt.Sprintf("[%v] %v", index, value)
	})

	assert.Equal(t, collectionAfter, collectionBefore)
}

func TestCollectionMap(t *testing.T) {
	collection := Array{"Hello", "World"}
	collection = collection.Map(func(item interface{}) interface{} {
		return fmt.Sprintf("- %v\n", item)
	})

	assert.Equal(t, Array{"- Hello\n", "- World\n"}, collection)
}

func TestCollectionFilter(t *testing.T) {
	collection := Array{"Hello", "World"}
	collection = collection.Filter(func(item interface{}) bool {
		return item != "Hello"
	})

	assert.Equal(t, Array{"World"}, collection)
}

func TestCollectionWhenNotEmptyOnEmpty(t *testing.T) {
	collection := Array{}
	collection = collection.WhenNotEmpty(func(collection Array) interface{} {
		return collection.Prepend("Hi")
	})

	assert.Equal(t, Array{}, collection)
}

func TestCollectionWhenNotEmptyOnNotEmpty(t *testing.T) {
	collection := Array{"Hello", "World"}
	collection = collection.WhenNotEmpty(func(collection Array) interface{} {
		return collection.Prepend("Hi")
	})

	assert.Equal(t, Array{"Hi", "Hello", "World"}, collection)
}

func TestCollectionChunk(t *testing.T) {
	collection := Array{"2607", "f0d0", "1002", "0051", "0000", "0000", "0000", "0004"}

	expected := Array{
		Array{"2607", "f0d0", "1002", "0051", "0000", "0000", "0000", "0004"},
	}

	assert.Equal(t, expected, collection.Chunk(8))

	expected = Array{
		Array{"2607", "f0d0", "1002", "0051", "0000", "0000"},
		Array{"0000", "0004"},
	}

	assert.Equal(t, expected, collection.Chunk(6))

	expected = Array{
		Array{"2607", "f0d0", "1002", "0051"},
		Array{"0000", "0000", "0000", "0004"},
	}

	assert.Equal(t, expected, collection.Chunk(4))

	expected = Array{
		Array{"2607", "f0d0"},
		Array{"1002", "0051"},
		Array{"0000", "0000"},
		Array{"0000", "0004"},
	}

	assert.Equal(t, expected, collection.Chunk(2))

	assert.Equal(t, collection, collection.Chunk(0))
	assert.Equal(t, collection, collection.Chunk(-1))
}

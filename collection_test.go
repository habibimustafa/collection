package array

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCollectingArray(t *testing.T) {
	arrString := []string{"Hello", "World"}
	strCollection := Collect(arrString)
	assert.Equal(t, len(arrString), len(strCollection))
	assert.Equal(t, len(arrString), strCollection.Size())
	assert.Equal(t, "Hello World", strCollection.Implode(" "))

	arrInt := []int{172, 20, 100, 255}
	intCollection := Collect(arrInt)
	assert.Equal(t, len(arrInt), len(intCollection))
	assert.Equal(t, len(arrInt), intCollection.Size())
	assert.Equal(t, "172.20.100.255", intCollection.Implode("."))

	arrRune := []rune{'h', 'e', 'l', 'l', 'o'}
	runeCollection := Collect(arrRune)
	assert.Equal(t, len(arrRune), len(runeCollection))
	assert.Equal(t, len(arrRune), runeCollection.Size())
	assert.Equal(t, "104.101.108.108.111", runeCollection.Implode("."))

	arrMap := map[string]string{"First Name": "John", "Last Name": "Doe"}
	mapCollection := Collect(arrMap)
	assert.Equal(t, len(arrMap), len(mapCollection))
	assert.Equal(t, len(arrMap), mapCollection.Size())
	assert.Equal(t, "John Doe", mapCollection.Implode(" "))
}

func TestCollectionIsNotEmpty(t *testing.T) {
	collection := Collection{"Hello", "World"}
	assert.True(t, collection.IsNotEmpty())
}

func TestCollectionAppend(t *testing.T) {
	collection := Collection{"Hello", "World"}
	collection = collection.Append("Hi")

	assert.Equal(t, 3, len(collection))
	assert.Equal(t, "Hi", collection[2])
}

func TestCollectionPrepend(t *testing.T) {
	collection := Collection{"Hello", "World"}
	collection = collection.Prepend("Hi")

	assert.Equal(t, 3, len(collection))
	assert.Equal(t, "Hi", collection[0])
}

func TestCollectionImplode(t *testing.T) {
	collection := Collection{"Hello", "World"}
	assert.Equal(t, "Hello World", collection.Implode(" "))
}

func TestCollectionHas(t *testing.T) {
	collection := Collection{"Hello", "World"}
	assert.True(t, collection.Has("Hello"))
	assert.False(t, collection.Has("Random"))
}

func TestCollectionEach(t *testing.T) {
	collectionBefore := Collection{"Hello", "World"}
	collectionAfter := collectionBefore.Each(func(value interface{}, index int) {
		value = fmt.Sprintf("[%v] %v", index, value)
	})

	assert.Equal(t, collectionAfter, collectionBefore)
}

func TestCollectionMap(t *testing.T) {
	collection := Collection{"Hello", "World"}
	collection = collection.Map(func(item interface{}) interface{} {
		return fmt.Sprintf("- %v\n", item)
	})

	assert.Equal(t, Collection{"- Hello\n", "- World\n"}, collection)
}

func TestCollectionFilter(t *testing.T) {
	collection := Collection{"Hello", "World"}
	collection = collection.Filter(func(item interface{}) bool {
		return item != "Hello"
	})

	assert.Equal(t, Collection{"World"}, collection)
}

func TestCollectionWhenNotEmptyOnEmpty(t *testing.T) {
	collection := Collection{}
	collection = collection.WhenNotEmpty(func(collection Collection) interface{} {
		return collection.Prepend("Hi")
	})

	assert.Equal(t, Collection{}, collection)
}

func TestCollectionWhenNotEmptyOnNotEmpty(t *testing.T) {
	collection := Collection{"Hello", "World"}
	collection = collection.WhenNotEmpty(func(collection Collection) interface{} {
		return collection.Prepend("Hi")
	})

	assert.Equal(t, Collection{"Hi", "Hello", "World"}, collection)
}

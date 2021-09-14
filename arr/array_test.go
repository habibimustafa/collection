package arr

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateArray(t *testing.T) {
	arrString := []string{"Hello", "World"}
	strArray := List(arrString)
	assert.Equal(t, len(arrString), len(strArray))
	assert.Equal(t, len(arrString), strArray.Size())
	assert.Equal(t, "Hello World", strArray.Implode(" "))

	arrInt := []int{172, 20, 100, 255}
	intArray := List(arrInt)
	assert.Equal(t, len(arrInt), len(intArray))
	assert.Equal(t, len(arrInt), intArray.Size())
	assert.Equal(t, "172.20.100.255", intArray.Implode("."))

	arrRune := []rune{'h', 'e', 'l', 'l', 'o'}
	runeArray := List(arrRune)
	assert.Equal(t, len(arrRune), len(runeArray))
	assert.Equal(t, len(arrRune), runeArray.Size())
	assert.Equal(t, "104.101.108.108.111", runeArray.Implode("."))

	arrMap := map[string]string{"First Name": "John", "Last Name": "Doe"}
	mapArray := List(arrMap)
	assert.Equal(t, len(arrMap), len(mapArray))
	assert.Equal(t, len(arrMap), mapArray.Size())
	assert.Equal(t, "John Doe", mapArray.Implode(" "))
}

func TestArrayGetAllItems(t *testing.T) {
	array := Array{"Hello", "World"}
	assert.Equal(t, []interface{}{"Hello", "World"}, array.All())
	assert.Equal(t, "Hello", array.All()[0])
	assert.Equal(t, []interface{}{"Hello"}, array.All()[:1])
	assert.Equal(t, []interface{}{"World"}, array.All()[1:])
}

func TestArrayGetItem(t *testing.T) {
	array := Array{"Hello", "World"}
	assert.Equal(t, "Hello", array.Get(0))
	assert.Equal(t, "World", array.Get(1))
}

func TestArrayGetFirstAndLastItem(t *testing.T) {
	array := Array{"Hello", "Middle", "World"}
	assert.Equal(t, "Hello", array.First())
	assert.Equal(t, "World", array.Last())
	assert.PanicsWithValue(t, "cannot get first element from empty array", func() { Array{}.First() })
	assert.PanicsWithValue(t, "cannot get last element from empty array", func() { Array{}.Last() })
}

func TestArrayIsNotEmpty(t *testing.T) {
	array := Array{"Hello", "World"}
	assert.True(t, array.IsNotEmpty())
}

func TestArrayAppend(t *testing.T) {
	array := Array{"Hello", "World"}
	array = array.Append("Hi")

	assert.Equal(t, 3, len(array))
	assert.Equal(t, "Hi", array[2])
}

func TestArrayPrepend(t *testing.T) {
	array := Array{"Hello", "World"}
	array = array.Prepend("Hi")

	assert.Equal(t, 3, len(array))
	assert.Equal(t, "Hi", array[0])
}

func TestArrayImplode(t *testing.T) {
	array := Array{"Hello", "World"}
	assert.Equal(t, "Hello World", array.Implode(" "))
}

func TestArrayKeys(t *testing.T) {
	array := List(map[string]string{"first": "John", "last": "Doe"})
	assert.Equal(t, []interface{}{0, 1}, array.Keys())
}

func TestArrayIndex(t *testing.T) {
	array := Array{"Hello", "World"}
	assert.Equal(t, 0, array.Index("Hello"))
	assert.Equal(t, 1, array.Index("World"))
	assert.Equal(t, -1, array.Index("Random"))

	array = List(map[string]string{"first": "John", "last": "Doe"})
	assert.Equal(t, 0, array.Index("John"))
	assert.Equal(t, 1, array.Index("Doe"))
	assert.Equal(t, -1, array.Index("Random"))
}

func TestArrayHas(t *testing.T) {
	array := Array{"Hello", "World"}
	assert.True(t, array.Has("Hello"))
	assert.False(t, array.Has("Random"))
}

func TestArrayEach(t *testing.T) {
	arrayBefore := Array{"Hello", "World"}
	arrayAfter := arrayBefore.Each(func(value interface{}, index int) {
		value = fmt.Sprintf("[%v] %v", index, value)
	})

	assert.Equal(t, arrayAfter, arrayBefore)
}

func TestArrayMap(t *testing.T) {
	array := Array{"Hello", "World"}
	array = array.Map(func(item interface{}) interface{} {
		return fmt.Sprintf("- %v\n", item)
	})

	assert.Equal(t, Array{"- Hello\n", "- World\n"}, array)
}

func TestArrayFilter(t *testing.T) {
	array := Array{"Hello", "World"}
	array = array.Filter(func(item interface{}) bool {
		return item != "Hello"
	})

	assert.Equal(t, Array{"World"}, array)
}

func TestArrayWhenNotEmptyOnEmpty(t *testing.T) {
	array := Array{}
	array = array.WhenNotEmpty(func(array Array) interface{} {
		return array.Prepend("Hi")
	})

	assert.Equal(t, Array{}, array)
}

func TestArrayWhenNotEmptyOnNotEmpty(t *testing.T) {
	array := Array{"Hello", "World"}
	array = array.WhenNotEmpty(func(array Array) interface{} {
		return array.Prepend("Hi")
	})

	assert.Equal(t, Array{"Hi", "Hello", "World"}, array)
}

func TestArrayChunk(t *testing.T) {
	array := Array{"2607", "f0d0", "1002", "0051", "0000", "0000", "0000", "0004"}

	expected := Array{
		Array{"2607", "f0d0", "1002", "0051", "0000", "0000", "0000", "0004"},
	}

	assert.Equal(t, expected, array.Chunk(8))

	expected = Array{
		Array{"2607", "f0d0", "1002", "0051", "0000", "0000"},
		Array{"0000", "0004"},
	}

	assert.Equal(t, expected, array.Chunk(6))

	expected = Array{
		Array{"2607", "f0d0", "1002", "0051"},
		Array{"0000", "0000", "0000", "0004"},
	}

	assert.Equal(t, expected, array.Chunk(4))

	expected = Array{
		Array{"2607", "f0d0"},
		Array{"1002", "0051"},
		Array{"0000", "0000"},
		Array{"0000", "0004"},
	}

	assert.Equal(t, expected, array.Chunk(2))

	assert.Equal(t, array, array.Chunk(0))
	assert.Equal(t, array, array.Chunk(-1))
}

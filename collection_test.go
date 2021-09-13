package collection

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var arrString = []string{"Hello", "World", "Are", "You", "Ready"}
var arrMap = map[string]interface{}{"First Name": "John", "Last Name": "Doe"}

func TestCreateCollection(t *testing.T) {
	strCollection := Collect(arrString)
	assert.Equal(t, len(arrString), strCollection.Size())

	mapCollection := Collect(arrMap)
	assert.Equal(t, len(arrMap), mapCollection.Size())
}

func TestCollectionGetAllItems(t *testing.T) {
	strCollection := Collect(arrString)
	assert.Equal(t, map[interface{}]interface{}{0: "Hello", 1: "World", 2: "Are", 3: "You", 4: "Ready"}, strCollection.All())
	assert.Equal(t, []interface{}{0, 1, 2, 3, 4}, strCollection.Keys().All())
	assert.Equal(t, []interface{}{"Hello", "World", "Are", "You", "Ready"}, strCollection.Values().All())
	assert.Equal(t, "Hello World Are You Ready", strCollection.Values().Implode(" "))

	mapCollection := Collect(arrMap)
	assert.Equal(t, map[interface{}]interface{}{"First Name": "John", "Last Name": "Doe"}, mapCollection.All())
	assert.Equal(t, []interface{}{"First Name", "Last Name"}, mapCollection.Keys().All())
	assert.Equal(t, []interface{}{"John", "Doe"}, mapCollection.Values().All())
	assert.Equal(t, "John Doe", mapCollection.Values().Implode(" "))
}

func TestCollectionGetFirstAndLastItems(t *testing.T) {
	strCollection := Collect(arrString)
	assert.Equal(t, map[interface{}]interface{}{0: "Hello"}, strCollection.First())
	assert.Equal(t, map[interface{}]interface{}{4: "Ready"}, strCollection.Last())
	assert.Equal(t, map[interface{}]interface{}{3: "You"}, strCollection.Get(3))

	mapCollection := Collect(arrMap)
	assert.Equal(t, map[interface{}]interface{}{"First Name": "John"}, mapCollection.First())
	assert.Equal(t, map[interface{}]interface{}{"Last Name": "Doe"}, mapCollection.Last())
	assert.Equal(t, map[interface{}]interface{}{"Last Name": "Doe"}, mapCollection.Get(1))
}

func TestCollectionSlicing(t *testing.T) {
	strCollection := Collect(arrString)
	assert.Equal(t, map[interface{}]interface{}{2: "Are", 3: "You", 4: "Ready"}, strCollection.Slice(2))
	assert.Equal(t, map[interface{}]interface{}{2: "Are", 3: "You", 4: "Ready"}, strCollection.Slice(2, 5))
	assert.Equal(t, map[interface{}]interface{}{2: "Are", 3: "You"}, strCollection.Slice(2, 3))
}

func TestCollectionAppend(t *testing.T) {
	strCollection := Collect(arrString)
	appended := strCollection.Append(20, "Haha")
	assert.Equal(t, 20, appended.Keys().Last())
	assert.Equal(t, "Haha", appended.Values().Last())
	assert.Equal(t, map[interface{}]interface{}{20: "Haha"}, appended.Last())
	assert.PanicsWithValue(t, "the new key is already exists", func() { strCollection.Append(2, "Haha") })
	assert.PanicsWithValue(t, "the new key type is different", func() { strCollection.Append('a', "Haha") })
	assert.NotPanics(t, func() { strCollection.Append(20, 2021) })
}

func TestCollectionPrepend(t *testing.T) {
	strCollection := Collect(arrString)
	prepended := strCollection.Prepend(-5, "Haha")
	assert.Equal(t, -5, prepended.Keys().First())
	assert.Equal(t, "Haha", prepended.Values().First())
	assert.Equal(t, map[interface{}]interface{}{-5: "Haha"}, prepended.First())
	assert.PanicsWithValue(t, "the new key is already exists", func() { strCollection.Prepend(2, "Haha") })
	assert.PanicsWithValue(t, "the new key type is different", func() { strCollection.Prepend('a', "Haha") })
	assert.NotPanics(t, func() { strCollection.Prepend(-5, 2021) })
}

func TestCollectionContains(t *testing.T) {
	strCollection := Collect(arrString)
	assert.True(t, strCollection.Contains(4, "Ready"))
	assert.False(t, strCollection.Contains(2, "Ready"))

	mapCollection := Collect(arrMap)
	assert.True(t, mapCollection.Contains("First Name", "John"))
	assert.False(t, mapCollection.Contains("Last Name", "John"))
}

func TestCollectionEach(t *testing.T) {
	idx := 0
	strCollection := Collect(arrString)
	strCollection.Each(func(value interface{}, key interface{}, index int) {
		assert.Equal(t, idx, index)
		assert.Equal(t, strCollection.Values().Get(idx), value)
		assert.Equal(t, strCollection.Keys().Get(idx), key)
		idx++
	})

	idx = 0
	mapCollection := Collect(arrMap)
	mapCollection.Each(func(value interface{}, key interface{}, index int) {
		assert.Equal(t, idx, index)
		assert.Equal(t, mapCollection.Values().Get(idx), value)
		assert.Equal(t, mapCollection.Keys().Get(idx), key)
		idx++
	})
}

func TestCollectionMap(t *testing.T) {
	strCollection := Collect(arrString)
	colmap := strCollection.Map(func(value interface{}, key interface{}, index int) (newValue interface{}, newKey interface{}) {
		assert.Equal(t, strCollection.Values().Get(index), value)
		assert.Equal(t, strCollection.Keys().Get(index), key)
		return "- " + value.(string), rune(key.(int) + 97)
	})

	assert.Equal(t, []interface{}{"- Hello", "- World", "- Are", "- You", "- Ready"}, colmap.Values().All())
	assert.Equal(t, []interface{}{'a', 'b', 'c', 'd', 'e'}, colmap.Keys().All())

	mapCollection := Collect(arrMap)
	colmap = mapCollection.Map(func(value interface{}, key interface{}, index int) (newValue interface{}, newKey interface{}) {
		assert.Equal(t, mapCollection.Values().Get(index), value)
		assert.Equal(t, mapCollection.Keys().Get(index), key)
		return "> " + value.(string), rune(index + 65)
	})

	assert.Equal(t, []interface{}{"> John", "> Doe"}, colmap.Values().All())
	assert.Equal(t, []interface{}{'A', 'B'}, colmap.Keys().All())
}

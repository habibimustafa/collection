package collection

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateCollection(t *testing.T) {
	arrString := []string{"Hello", "World", "Are", "You", "Ready"}
	strCollection := Collect(arrString)
	assert.Equal(t, len(arrString), strCollection.Size())
	assert.Equal(t, map[interface{}]interface{}{0: "Hello", 1: "World", 2: "Are", 3: "You", 4: "Ready"}, strCollection.All())
	assert.Equal(t, []interface{}{0, 1, 2, 3, 4}, strCollection.Keys().All())
	assert.Equal(t, []interface{}{"Hello", "World", "Are", "You", "Ready"}, strCollection.Values().All())
	assert.Equal(t, "Hello World Are You Ready", strCollection.Values().Implode(" "))
	assert.Equal(t, map[interface{}]interface{}{0: "Hello"}, strCollection.First())
	assert.Equal(t, map[interface{}]interface{}{4: "Ready"}, strCollection.Last())
	assert.Equal(t, map[interface{}]interface{}{3: "You"}, strCollection.Get(3))
	assert.Equal(t, map[interface{}]interface{}{2: "Are", 3: "You", 4: "Ready"}, strCollection.Slice(2))
	assert.Equal(t, map[interface{}]interface{}{2: "Are", 3: "You", 4: "Ready"}, strCollection.Slice(2, 5))
	assert.Equal(t, map[interface{}]interface{}{2: "Are", 3: "You"}, strCollection.Slice(2, 3))
	assert.True(t, strCollection.Contains(4, "Ready"))
	assert.False(t, strCollection.Contains(2, "Ready"))

	idx := 0
	strCollection.Each(func(value interface{}, key interface{}, index int) {
		assert.Equal(t, idx, index)
		assert.Equal(t, strCollection.Values().Get(idx), value)
		assert.Equal(t, strCollection.Keys().Get(idx), key)
		idx++
	})

	colmap := strCollection.Map(func(value interface{}, key interface{}, index int) (newValue interface{}, newKey interface{}) {
		assert.Equal(t, strCollection.Values().Get(index), value)
		assert.Equal(t, strCollection.Keys().Get(index), key)
		return "- " + value.(string), rune(key.(int) + 97)
	})

	assert.Equal(t, []interface{}{"- Hello", "- World", "- Are", "- You", "- Ready"}, colmap.Values().All())
	assert.Equal(t, []interface{}{'a', 'b', 'c', 'd', 'e'}, colmap.Keys().All())

	arrMap := map[string]string{"First Name": "John", "Last Name": "Doe"}
	mapCollection := Collect(arrMap)
	assert.Equal(t, len(arrMap), mapCollection.Size())
	assert.Equal(t, map[interface{}]interface{}{"First Name": "John", "Last Name": "Doe"}, mapCollection.All())
	assert.Equal(t, []interface{}{"First Name", "Last Name"}, mapCollection.Keys().All())
	assert.Equal(t, []interface{}{"John", "Doe"}, mapCollection.Values().All())
	assert.Equal(t, "John Doe", mapCollection.Values().Implode(" "))
	assert.True(t, mapCollection.Contains("First Name", "John"))
	assert.False(t, mapCollection.Contains("Last Name", "John"))

	idx = 0
	mapCollection.Each(func(value interface{}, key interface{}, index int) {
		assert.Equal(t, idx, index)
		assert.Equal(t, mapCollection.Values().Get(idx), value)
		assert.Equal(t, mapCollection.Keys().Get(idx), key)
		idx++
	})

	colmap = mapCollection.Map(func(value interface{}, key interface{}, index int) (newValue interface{}, newKey interface{}) {
		assert.Equal(t, mapCollection.Values().Get(index), value)
		assert.Equal(t, mapCollection.Keys().Get(index), key)
		return "> " + value.(string), rune(index + 65)
	})

	assert.Equal(t, []interface{}{"> John", "> Doe"}, colmap.Values().All())
	assert.Equal(t, []interface{}{'A', 'B'}, colmap.Keys().All())
}

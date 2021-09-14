package collection

import (
	"github.com/habibimustafa/collection/arr"
	"github.com/habibimustafa/collection/sort"
	"reflect"
)

type Collection interface {
	// Size count the collection items
	Size() int

	// All get all the items
	All() map[interface{}]interface{}

	// Keys get array of the keys
	Keys() arr.Array

	// Values get array of the values
	Values() arr.Array

	// Get gets item by index
	Get(index int) map[interface{}]interface{}

	// GetValue gets value by key
	GetValue(key interface{}) interface{}

	// First gets the first item
	First() map[interface{}]interface{}

	// Last gets the last item
	Last() map[interface{}]interface{}

	// Slice gets slice of items
	Slice(slice ...int) map[interface{}]interface{}

	// Contains is collection contains key with value
	Contains(key interface{}, val interface{}) bool

	// Has is collection has provided keys
	Has(keys ...interface{}) bool

	// Append add new item to last position
	Append(key interface{}, val interface{}) Collection

	// Prepend add new item to first position
	Prepend(key interface{}, val interface{}) Collection

	// Set update the existing item when its exist
	// when not exist, it will add new item to last position
	Set(key interface{}, val interface{}) Collection

	// Unset remove item by key
	Unset(key interface{}) Collection

	// Remove alias of Unset method
	Remove(key interface{}) Collection

	// Except gets all items except provided keys
	Except(keys ...interface{}) Collection

	// Only gets all items that match with provided keys
	Only(keys ...interface{}) Collection

	// Each looping each item
	Each(callback func(value interface{}, key interface{}, index int)) Collection

	// Map converts each item into new format
	Map(callback func(value interface{}, key interface{}, index int) (newValue interface{}, newKey interface{})) Collection

	// Filter remove unmatched items from the collection
	Filter(callback func(value interface{}, key interface{}, index int) bool) Collection

	// Where alias of Filter method
	Where(callback func(value interface{}, key interface{}, index int) bool) Collection
}

type collect struct {
	keys   []interface{}
	values []interface{}
}

func Collect(collection interface{}) Collection {
	if collection == nil {
		return collect{}
	}

	val := reflect.ValueOf(collection)
	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		var keys []interface{}
		var values []interface{}
		for i := 0; i < val.Len(); i++ {
			keys = append(keys, i)
			values = append(values, val.Index(i).Interface())
		}
		return collect{keys: keys, values: values}
	case reflect.Map:
		sorted := sort.Sort(val)

		var keys []interface{}
		for _, k := range sorted.Key {
			keys = append(keys, k.Interface())
		}

		var values []interface{}
		for _, v := range sorted.Value {
			values = append(values, v.Interface())
		}

		return collect{keys: keys, values: values}
	default:
		panic("collection: collection type must be a slice, array, map, or nil")
		return nil
	}
}

func (c collect) Size() int {
	return len(c.keys)
}

func (c collect) All() map[interface{}]interface{} {
	m := map[interface{}]interface{}{}
	for i, key := range c.keys {
		m[key] = c.values[i]
	}
	return m
}

func (c collect) Keys() arr.Array {
	return arr.List(c.keys)
}

func (c collect) Values() arr.Array {
	return arr.List(c.values)
}

func (c collect) Get(index int) map[interface{}]interface{} {
	m := map[interface{}]interface{}{}
	m[c.keys[index]] = c.values[index]
	return m
}

func (c collect) GetValue(key interface{}) interface{} {
	index := c.Keys().Index(key)
	if index > -1 {
		return c.Values().Get(index)
	}

	return nil
}

func (c collect) First() map[interface{}]interface{} {
	return c.Get(0)
}

func (c collect) Last() map[interface{}]interface{} {
	return c.Get(c.Size() - 1)
}

func (c collect) Slice(slice ...int) map[interface{}]interface{} {
	m := map[interface{}]interface{}{}
	if len(slice) < 1 {
		return m
	}

	start := slice[0]
	end := c.Size() - 1

	if len(slice) >= 2 && slice[1] <= end {
		end = slice[1]
	}

	for i := start; i <= end; i++ {
		m[c.keys[i]] = c.values[i]
	}

	return m
}

func (c collect) Contains(key interface{}, value interface{}) bool {
	return c.All()[key] == value
}

func (c collect) Has(keys ...interface{}) bool {
	if len(keys) < 1 {
		return false
	}

	for _, k := range keys {
		if !c.Keys().Has(k) {
			return false
		}
	}

	return true
}

func (c collect) Append(key interface{}, value interface{}) Collection {
	c.validateKey(key)
	return collect{
		keys:   c.Keys().Append(key).All(),
		values: c.Values().Append(value).All(),
	}
}

func (c collect) Prepend(key interface{}, value interface{}) Collection {
	c.validateKey(key)
	return collect{
		keys:   c.Keys().Prepend(key).All(),
		values: c.Values().Prepend(value).All(),
	}
}

func (c collect) Set(key interface{}, value interface{}) Collection {
	if !c.Keys().Has(key) {
		return c.Append(key, value)
	}

	index := c.Keys().Index(key)
	var values []interface{}
	for i, v := range c.Values().All() {
		if i == index {
			values = append(values, value)
			continue
		}
		values = append(values, v)
	}

	return collect{
		keys:   c.Keys().All(),
		values: values,
	}
}

func (c collect) Unset(key interface{}) Collection {
	if !c.Keys().Has(key) {
		panic("the inputted key is not exist in this collection")
	}

	removedKey := key
	return c.Filter(func(value interface{}, key interface{}, index int) bool {
		return key != removedKey
	})
}

func (c collect) Remove(key interface{}) Collection {
	return c.Unset(key)
}

func (c collect) Except(keys ...interface{}) Collection {
	return c.Filter(func(value interface{}, key interface{}, index int) bool {
		return !arr.List(keys).Has(key)
	})
}

func (c collect) Only(keys ...interface{}) Collection {
	return c.Filter(func(value interface{}, key interface{}, index int) bool {
		return arr.List(keys).Has(key)
	})
}

func (c collect) Each(callback func(value interface{}, key interface{}, index int)) Collection {
	for i := 0; i < c.Size(); i++ {
		callback(c.values[i], c.keys[i], i)
	}
	return c
}

func (c collect) Map(callback func(value interface{}, key interface{}, index int) (newValue interface{}, newKey interface{})) Collection {
	var keys []interface{}
	var values []interface{}
	for i := 0; i < c.Size(); i++ {
		newValue, newKey := callback(c.values[i], c.keys[i], i)
		values = append(values, newValue)
		keys = append(keys, newKey)
	}
	return collect{keys: keys, values: values}
}

func (c collect) Filter(callback func(value interface{}, key interface{}, index int) bool) Collection {
	var keys []interface{}
	var values []interface{}
	for i := 0; i < c.Size(); i++ {
		if !callback(c.values[i], c.keys[i], i) {
			continue
		}

		values = append(values, c.Values().Get(i))
		keys = append(keys, c.Keys().Get(i))
	}

	return collect{keys: keys, values: values}
}

func (c collect) Where(callback func(value interface{}, key interface{}, index int) bool) Collection {
	return c.Filter(callback)
}

func (c collect) validateKey(key interface{}) {
	if c.Keys().Has(key) {
		panic("the new key is already exists")
	}

	if c.Keys().Size() > 0 {
		keyType := reflect.TypeOf(c.Keys().First()).Kind()
		newKeyType := reflect.TypeOf(key).Kind()
		if keyType != newKeyType {
			panic("the new key type is different")
		}
	}
}

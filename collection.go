package collection

import (
	"github.com/habibimustafa/collection/arr"
	"github.com/habibimustafa/collection/sort"
	"reflect"
)

type Collection interface {
	Size() int
	All() map[interface{}]interface{}
	Get(index int) map[interface{}]interface{}
	GetValue(key interface{}) interface{}
	First() map[interface{}]interface{}
	Last() map[interface{}]interface{}
	Slice(slice ...int) map[interface{}]interface{}
	Contains(key interface{}, val interface{}) bool
	Append(key interface{}, val interface{}) Collection
	Prepend(key interface{}, val interface{}) Collection
	Set(key interface{}, val interface{}) Collection
	Unset(key interface{}) Collection
	Keys() arr.Array
	Values() arr.Array
	Each(callback func(value interface{}, key interface{}, index int)) Collection
	Map(callback func(value interface{}, key interface{}, index int) (newValue interface{}, newKey interface{})) Collection
	Filter(callback func(value interface{}, key interface{}, index int) bool) Collection
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

func (c collect) Get(index int) map[interface{}]interface{} {
	m := map[interface{}]interface{}{}
	m[c.keys[index]] = c.values[index]
	return m
}

func (c collect) GetValue(key interface{}) interface{} {
	index := c.Keys().Index(key)
	if index != nil {
		return c.Values().Get(index.(int))
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

func (c collect) Keys() arr.Array {
	return arr.List(c.keys)
}

func (c collect) Values() arr.Array {
	return arr.List(c.values)
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

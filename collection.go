package collection

import (
	"github.com/habibimustafa/collection/arr"
	"reflect"
)

type Collection interface {
	Size() int
	All() map[interface{}]interface{}
	Get(index int) map[interface{}]interface{}
	First() map[interface{}]interface{}
	Last() map[interface{}]interface{}
	Slice(slice ...int) map[interface{}]interface{}
	Contains(key interface{}, val interface{}) bool
	Append(key interface{}, val interface{}) Collection
	Prepend(key interface{}, val interface{}) Collection
	Keys() arr.Array
	Values() arr.Array
	Each(callback func(value interface{}, key interface{}, index int)) Collection
	Map(callback func(value interface{}, key interface{}, index int) (newValue interface{}, newKey interface{})) Collection
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
		var keys []interface{}
		var values []interface{}
		for _, k := range val.MapKeys() {
			keys = append(keys, k.Interface())
			values = append(values, val.MapIndex(k).Interface())
		}
		return collect{keys: keys, values: values}
	default:
		panic("collection: collection type must be a slice, array or map")
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
	c.validateNewItem(key, value)
	return collect{
		keys:   c.Keys().Append(key).All(),
		values: c.Values().Append(value).All(),
	}
}

func (c collect) Prepend(key interface{}, value interface{}) Collection {
	c.validateNewItem(key, value)
	return collect{
		keys:   c.Keys().Prepend(key).All(),
		values: c.Values().Prepend(value).All(),
	}
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

func (c collect) validateNewItem(key interface{}, value interface{}) {
	if c.Keys().Has(key) {
		panic("the new key is already exists")
	}

	if c.Keys().Size() > 0 {
		keyType := reflect.TypeOf(c.Keys().First()).Kind()
		newKeyType := reflect.TypeOf(key).Kind()
		if keyType != newKeyType {
			panic("the new key type is different")
		}

		valType := reflect.TypeOf(c.Values().First()).Kind()
		newValType := reflect.TypeOf(value).Kind()
		if valType != newValType {
			panic("the new value type is different")
		}
	}
}

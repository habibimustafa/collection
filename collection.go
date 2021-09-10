package collection

import (
	"github.com/habibimustafa/collection/arr"
	"log"
	"reflect"
)

type Collection struct {
	keys   []interface{}
	values []interface{}
}

func Collect(collection interface{}) *Collection {
	val := reflect.ValueOf(collection)
	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		var keys []interface{}
		var values []interface{}
		for i := 0; i < val.Len(); i++ {
			keys = append(keys, i)
			values = append(values, val.Index(i).Interface())
		}
		return &Collection{keys, values}
	case reflect.Map:
		var keys []interface{}
		var values []interface{}
		for _, k := range val.MapKeys() {
			keys = append(keys, k.Interface())
			values = append(values, val.MapIndex(k).Interface())
		}
		return &Collection{keys, values}
	default:
		log.Fatalln("collection: collection type must be a slice, Array or map")
	}

	return &Collection{}
}

func (c Collection) Size() int {
	return len(c.keys)
}

func (c Collection) All() map[interface{}]interface{} {
	m := map[interface{}]interface{}{}
	for i, key := range c.keys {
		m[key] = c.values[i]
	}
	return m
}

func (c Collection) Get(index int) map[interface{}]interface{} {
	m := map[interface{}]interface{}{}
	m[c.keys[index]] = c.values[index]
	return m
}

func (c Collection) First() map[interface{}]interface{} {
	return c.Get(0)
}

func (c Collection) Last() map[interface{}]interface{} {
	return c.Get(c.Size() - 1)
}

func (c Collection) Slice(slice ...int) map[interface{}]interface{} {
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

func (c Collection) Keys() arr.Array {
	return arr.List(c.keys)
}

func (c Collection) Values() arr.Array {
	return arr.List(c.values)
}

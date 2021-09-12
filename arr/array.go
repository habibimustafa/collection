package arr

import (
	"bytes"
	"fmt"
	"math"
	"reflect"
)

type Array []interface{}

func List(list interface{}) Array {
	val := reflect.ValueOf(list)
	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		c := Array{}
		for i := 0; i < val.Len(); i++ {
			c = append(c, val.Index(i).Interface())
		}
		return c
	case reflect.Map:
		c := Array{}
		for _, k := range val.MapKeys() {
			c = append(c, val.MapIndex(k).Interface())
		}
		return c
	default:
		panic("list: list type must be a slice, array or map")
		return nil
	}
}

func (a Array) All() []interface{} {
	return a
}

func (a Array) Get(index int) interface{} {
	return a[index]
}

func (a Array) Size() int {
	return len(a)
}

func (a Array) First() interface{} {
	if a.Size() < 1 {
		panic("cannot get first element from empty array")
	}
	return a[0]
}

func (a Array) Last() interface{} {
	if a.Size() < 1 {
		panic("cannot get last element from empty array")
	}
	return a[a.Size()-1]
}

func (a Array) IsNotEmpty() bool {
	return a.Size() > 0
}

func (a Array) Append(item interface{}) Array {
	return append(a, item)
}

func (a Array) Prepend(item interface{}) Array {
	return append(Array{item}, a...)
}

func (a Array) Implode(glue string) string {
	var buf bytes.Buffer
	for i, str := range a {
		if i > 0 {
			buf.WriteString(glue)
		}

		buf.WriteString(fmt.Sprintf("%v", str))
	}
	return buf.String()
}

func (a Array) Keys() []interface{} {
	var keys []interface{}
	for k, _ := range a {
		keys = append(keys, k)
	}
	return keys
}

func (a Array) Index(value interface{}) interface{} {
	for index, item := range a {
		if item == value {
			return index
		}
	}

	return nil
}

func (a Array) Has(value interface{}) bool {
	for _, item := range a {
		if value == item {
			return true
		}
	}
	return false
}

func (a Array) Each(callback func(item interface{}, index int)) Array {
	itemsCopy := a
	for i, item := range itemsCopy {
		callback(item, i)
	}
	return a
}

func (a Array) Map(callback func(item interface{}) interface{}) Array {
	var newCollection Array
	for _, item := range a {
		newCollection = append(newCollection, callback(item))
	}
	return newCollection
}

func (a Array) Filter(callback func(item interface{}) bool) Array {
	var newCollection Array
	for _, item := range a {
		if callback(item) {
			newCollection = append(newCollection, item)
		}
	}
	return newCollection
}

func (a Array) WhenNotEmpty(callback func(collection Array) interface{}) Array {
	if a.IsNotEmpty() {
		result := callback(a)
		if newCollection, ok := result.(Array); ok {
			return newCollection
		}
	}

	return a
}

func (a Array) Chunk(size int) interface{} {
	if size <= 0 {
		return a
	}

	length := len(a)
	chunks := int(math.Ceil(float64(length) / float64(size)))

	var newCollection Array
	for i, end := 0, 0; chunks > 0; chunks-- {
		end = (i + 1) * size
		if end > length {
			end = length
		}
		newCollection = append(newCollection, a[i*size:end])
		i++
	}

	return newCollection
}

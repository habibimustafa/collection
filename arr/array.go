package arr

import (
	"bytes"
	"fmt"
	"math"
	"reflect"
)

type Array []interface{}

func List(list interface{}) Array {
	if list == nil {
		return Array{}
	}

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
		panic("list: list type must be a slice, array, map, or nil")
		return nil
	}
}

// All gets all items
func (a Array) All() []interface{} {
	return a
}

// Get gets item by index
func (a Array) Get(index int) interface{} {
	return a[index]
}

// Size count items
func (a Array) Size() int {
	return len(a)
}

// First get the first item
func (a Array) First() interface{} {
	if a.IsEmpty() {
		panic("cannot get first element from empty array")
	}
	return a[0]
}

// Last get the last item
func (a Array) Last() interface{} {
	if a.IsEmpty() {
		panic("cannot get last element from empty array")
	}
	return a[a.Size()-1]
}

// IsEmpty is array has no items
func (a Array) IsEmpty() bool {
	return a.Size() < 1
}

// IsNotEmpty is array has items
func (a Array) IsNotEmpty() bool {
	return !a.IsEmpty()
}

// Append add new item to last position
func (a Array) Append(item interface{}) Array {
	return append(a, item)
}

// Prepend add new item to first position
func (a Array) Prepend(item interface{}) Array {
	return append(Array{item}, a...)
}

// Implode merge items with glue into string
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

// Keys gets the array keys
func (a Array) Keys() []interface{} {
	var keys []interface{}
	for k, _ := range a {
		keys = append(keys, k)
	}
	return keys
}

// Index get the index of value
func (a Array) Index(value interface{}) interface{} {
	for index, item := range a {
		if item == value {
			return index
		}
	}

	return nil
}

// Has is array has provided value
func (a Array) Has(value interface{}) bool {
	for _, item := range a {
		if value == item {
			return true
		}
	}
	return false
}

// Each looping each item
func (a Array) Each(callback func(item interface{}, index int)) Array {
	itemsCopy := a
	for i, item := range itemsCopy {
		callback(item, i)
	}
	return a
}

// Map converts each item into new format
func (a Array) Map(callback func(item interface{}) interface{}) Array {
	var newCollection Array
	for _, item := range a {
		newCollection = append(newCollection, callback(item))
	}
	return newCollection
}

// Filter remove unmatched items from the array
func (a Array) Filter(callback func(item interface{}) bool) Array {
	var newCollection Array
	for _, item := range a {
		if callback(item) {
			newCollection = append(newCollection, item)
		}
	}
	return newCollection
}

// WhenNotEmpty executes callback when array is not empty
func (a Array) WhenNotEmpty(callback func(collection Array) interface{}) Array {
	if a.IsNotEmpty() {
		result := callback(a)
		if newCollection, ok := result.(Array); ok {
			return newCollection
		}
	}

	return a
}

// Chunk splits items into separated array
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

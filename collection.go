package arr

import (
	"bytes"
	"fmt"
	"log"
	"math"
	"reflect"
)

type Collection []interface{}

func Collect(list interface{}) Collection {
	val := reflect.ValueOf(list)
	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		c := Collection{}
		for i := 0; i < val.Len(); i++ {
			c = append(c, val.Index(i).Interface())
		}
		return c
	case reflect.Map:
		c := Collection{}
		for _, k := range val.MapKeys() {
			c = append(c, val.MapIndex(k).Interface())
		}
		return c
	default:
		log.Fatalln("list: list type must be a slice, array or map")
	}

	return nil
}

func (c Collection) Size() int {
	return len(c)
}

func (c Collection) IsNotEmpty() bool {
	return c.Size() > 0
}

func (c Collection) Append(item interface{}) Collection {
	return append(c, item)
}

func (c Collection) Prepend(item interface{}) Collection {
	newItems := Collection{item}
	return append(newItems, c...)
}

func (c Collection) Implode(glue string) string {
	var buf bytes.Buffer
	for i, str := range c {
		if i > 0 {
			buf.WriteString(glue)
		}

		buf.WriteString(fmt.Sprintf("%v", str))
	}
	return buf.String()
}

func (c Collection) Has(value interface{}) bool {
	for _, item := range c {
		if value == item {
			return true
		}
	}

	return false
}

func (c Collection) Each(callback func(item interface{}, index int)) Collection {
	itemsCopy := c
	for i, item := range itemsCopy {
		callback(item, i)
	}
	return c
}

func (c Collection) Map(callback func(item interface{}) interface{}) Collection {
	var newItems Collection
	for _, item := range c {
		newItems = append(newItems, callback(item))
	}
	return newItems
}

func (c Collection) Filter(callback func(item interface{}) bool) Collection {
	var newItems Collection
	for _, item := range c {
		if callback(item) {
			newItems = append(newItems, item)
		}
	}
	return newItems
}

func (c Collection) WhenNotEmpty(callback func(collection Collection) interface{}) Collection {
	if c.IsNotEmpty() {
		result := callback(c)
		if newCollection, ok := result.(Collection); ok {
			return newCollection
		}
	}

	return c
}

func (c Collection) Chunk(size int) interface{} {
	if size <= 0 {
		return c
	}

	length := len(c)
	chunks := int(math.Ceil(float64(length) / float64(size)))

	var newItems Collection
	for i, end := 0, 0; chunks > 0; chunks-- {
		end = (i + 1) * size
		if end > length {
			end = length
		}
		newItems = append(newItems, c[i*size:end])
		i++
	}

	return newItems
}

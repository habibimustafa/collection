package array

import (
	"bytes"
	"fmt"
	"log"
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

func (items Collection) Size() int {
	return len(items)
}

func (items Collection) IsNotEmpty() bool {
	return items.Size() > 0
}

func (items Collection) Append(item interface{}) Collection {
	return append(items, item)
}

func (items Collection) Prepend(item interface{}) Collection {
	newItems := Collection{item}
	return append(newItems, items...)
}

func (items Collection) Implode(glue string) string {
	var buf bytes.Buffer
	for i, str := range items {
		if i > 0 {
			buf.WriteString(glue)
		}

		buf.WriteString(fmt.Sprintf("%v", str))
	}
	return buf.String()
}

func (items Collection) Has(value interface{}) bool {
	for _, item := range items {
		if value == item {
			return true
		}
	}

	return false
}

func (items Collection) Each(callback func(item interface{}, index int)) Collection {
	itemsCopy := items
	for i, item := range itemsCopy {
		callback(item, i)
	}
	return items
}

func (items Collection) Map(callback func(item interface{}) interface{}) Collection {
	var newItems Collection
	for _, item := range items {
		newItems = append(newItems, callback(item))
	}
	return newItems
}

func (items Collection) Filter(callback func(item interface{}) bool) Collection {
	var newItems Collection
	for _, item := range items {
		if callback(item) {
			newItems = append(newItems, item)
		}
	}
	return newItems
}

func (items Collection) WhenNotEmpty(callback func(collection Collection) interface{}) Collection {
	if items.IsNotEmpty() {
		result := callback(items)
		if newCollection, ok := result.(Collection); ok {
			return newCollection
		}
	}

	return items
}

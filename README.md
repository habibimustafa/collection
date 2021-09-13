# collection
Golang Collection Module is a helper for managing array, slice and map.

## Usage
```bash
go get github.com/habibimustafa/collection
```

## Example
```go
arrString := []string{"Hello", "World", "Are", "You", "Ready"}
strCollection := Collect(arrString)

strCollection.Size()
strCollection.Contains(0, "Hello")
strCollection.Keys().Size()
strCollection.Values().Has("Ready")
```

```go
arrMap = map[string]interface{}{"First Name": "John", "Last Name": "Doe", "Age": 28}
mapCollection := Collect(arrMap)

mapCollection.Size()
mapCollection.Contains("First Name", "John")
mapCollection.Keys().All().Map(...)
mapCollection.Values().Has("Ready")
```

For more usage examples please see the test files.
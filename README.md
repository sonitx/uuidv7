# uuidv7
UUID V7 library for Golang

## How to use

### Install uuid v7 library

* From your project:

```go
go get github.com/sonitx/uuidv7
```

### Add UUID V7 in your code

#### Get UUID V7 by []byte, use function `uuidv7.GetUUIDv7()`

```go
uuidVal, err := uuidv7.GetUUIDv7()
	if err != nil {
		fmt.Errorf("Cannot generate UUID V7, error %v", err)
		return
	}
	fmt.Printf("%x", uuidVal)
```

#### Get UUID V7 by string, use function `uuidv7.GetUUIDv7String()`

```go
uuidVal, err := uuidv7.GetUUIDv7String()
if err != nil {
	fmt.Errorf("Cannot generate UUID V7, error %v", err)
	return
}
fmt.Printf(uuidVal)
```

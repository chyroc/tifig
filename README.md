# tifig
wrap tifig with golang

convert heic to png

depend on : https://github.com/monostream/tifig

## usage

### go

```go
import "github.com/Chyroc/tifig"

output, err := tifig.Convert("/path/gile.heic", "")
if err != nil {
	panic(err)
}

fmt.Println(output)
```

### binary

see: https://github.com/monostream/tifig

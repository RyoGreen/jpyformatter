# jpyformatter

## Instalation:

```
go get github.com/RyoGreen/jpyformatter
```

## Example

```
package main

import (
	"fmt"

	formatter "github.com/RyoGreen/jpyformatter"
)

func main() {
	var price = 1000
	p, _ := formatter.Format(price, true, true)
	pWithPrefix, _ := formatter.Format(price, true, false)
	pWithSuffix, _ := formatter.Format(price, false, true)

	fmt.Println(p)           // => ¥1,000-
	fmt.Println(pWithPrefix) // => ¥1,000
	fmt.Println(pWithSuffix) // => 1,000-

	p = formatter.FormatX(price, true, true)
	pWithPrefix = formatter.FormatX(price, true, false)
	pWithSuffix = formatter.FormatX(price, false, true)

	fmt.Println(p)           // => ¥1,000-
	fmt.Println(pWithPrefix) // => ¥1,000
	fmt.Println(pWithSuffix) // => 1,000-
}

```

## Available Data Types

- int int8 int16 int32 int64
- uint uint8 uint16 uint32 uint64 uintptr
- float64

## Note

Price must be between 99,999,999,999,999 yen and -99,999,999,999,999 yen

## License
MIT

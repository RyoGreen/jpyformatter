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
	"log"

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
}

```

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
	p, err := formatter.Format(price, true, true)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(p)  // => ¥1,000-
}

```

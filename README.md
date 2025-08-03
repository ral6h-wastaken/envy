# envy
A simple GoLang env file parsing and access library 

## usage
In your Go project directory, run 
```
go get github.com/ral6h-wastaken/envy@latest
```

Then in your code you can use something like:

```go
package main

import (
	"fmt"
	"github.com/ral6h-wastaken/envy"
)

//cat $(pwd)/.env -> ciao=como estas

func main() {
	vars := envy.GetInstance("")
	fmt.Println(vars.Get("ciao")) //como estas
}
```

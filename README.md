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

## multiple .env files 
Envy provides support for multiple env files at the same time: 
```go
//cat $(pwd)/.env -> ciao=como estas
//cat $(pwd)/secodary.env -> ciao=come stai

func main() {
	vars := envy.GetInstance("")
	fmt.Println(vars.Get("ciao")) //como estas

	vars2 := envy.GetInstance("./secondary.env")
	fmt.Println(vars2.Get("ciao")) //come stai
}
```

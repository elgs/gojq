# gojq
JSON query in Golang.

```go
package main

import (
	"fmt"
	"github.com/elgs/gojq"
)

var jsonArray = `
[
  {
    "name": "elgs",
    "gender": "m",
    "skills": [
      "Golang",
      "Java",
      "C"
    ]
  },
  {
    "name": "enny",
    "gender": "f",
    "skills": [
      "IC",
      "Electric design",
      "Verification"
    ]
  },
  {
    "name": "sam",
    "gender": "m",
	"pet": null,
    "skills": [
      "Eating",
      "Sleeping",
      "Crawling"
    ]
  }
]
`

func main() {
	parser, err := gojq.NewStringQuery(jsonArray)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(parser.Parse("[0].name"))       // elgs <nil>
	fmt.Println(parser.Parse("[1].gender"))     // f <nil>
	fmt.Println(parser.Parse("[2].skills.[1]")) // Sleeping <nil>
	fmt.Println(parser.Parse("[2].hello"))      // <nil> hello does not exist.
	fmt.Println(parser.Parse("[2].pet"))        // <nil> <nil>
}
```

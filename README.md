# gojq
JSON query in Golang.

## Install
`go get -u github.com/elgs/gojq`

## Query from JSON Object

```go
package main

import (
	"fmt"
	"github.com/elgs/gojq"
)

var jsonObj = `
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
`

func main() {
	parser, err := gojq.NewStringQuery(jsonObj)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(parser.Parse("name"))       // sam <nil>
	fmt.Println(parser.Parse("gender"))     // m <nil>
	fmt.Println(parser.Parse("skills.[1]")) // Sleeping <nil>
	fmt.Println(parser.Parse("hello"))      // <nil> hello does not exist.
	fmt.Println(parser.Parse("pet"))        // <nil> <nil>
}
```

## Query from JSON Array

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

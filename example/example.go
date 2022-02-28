package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/fxmwl/jsonpath"
)

const s = `
{
	"a": {
		"b":123,
		"c":12
	}
}
`

func main() {
	js := jsonpath.New("default")

	if err := js.Parse("{.a.b}"); err != nil {
		panic(err)
	}

	data := make(map[string]interface{})

	if err := json.Unmarshal([]byte(s), &data); err != nil {
		panic(err)
	}

	result := bytes.NewBuffer([]byte(""))

	if err := js.Execute(result, data); err != nil {
		panic(err)
	}

	fmt.Println(result.String())

}

package main

import (
	"context"
	"fmt"

	"github.com/risor-io/risor"
)

func main() {
	ctx := context.Background()
	code := `
resp := fetch(url)
print(resp.text())
`
	_, err := risor.Eval(ctx, code, risor.WithGlobal("url", "https://httpbin.org/anything"))
	if err != nil {
		fmt.Println(err)
	}
}

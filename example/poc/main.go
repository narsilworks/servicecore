package main

import (
	"fmt"

	"github.com/narsilworks/servicecore"
	impl "github.com/narsilworks/servicecore/impl"
)

func main() {
	sc, err := servicecore.Create(map[string]any{
		"id": "SUPERSERVICE",
	})
	if err != nil {
		return
	}

	rc := impl.NewRedisCache("", "", 0)

	sc.Set().Cache(rc)
	c, err := sc.Get().Cache()
	if err != nil {
		fmt.Println(err)
		return
	}

	c.Del("")
}

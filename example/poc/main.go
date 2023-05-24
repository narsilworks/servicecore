package main

import (
	"fmt"

	"github.com/narsilworks/servicecore"
	impl "github.com/narsilworks/servicecore/impl"
)

func main() {
	sc, err := servicecore.Create(
		map[string]any{
			"id":        "SUPERSERVICE",
			"name":      "Super Service Application",
			"copyright": "Copyright 1976, ACME, Inc.",
		})
	if err != nil {
		return
	}

	rc := impl.NewRedisCache("", "", 0)

	// Set cache
	sc.Set().Cache(rc)
	c, err := sc.Get().Cache()
	if err != nil {
		fmt.Println(err)
		return
	}

	c.Del("")

	// Serve
	sc.Serve()
}

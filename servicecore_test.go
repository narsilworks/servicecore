package servicecore

import (
	"testing"

	"github.com/narsilworks/servicecore/impl"
)

func TestSetter(t *testing.T) {
	sc, err := Create(map[string]any{})
	if err != nil {
		return
	}

	sc.Set().Cache(impl.NewRedisCache("", "", 0))
	c, err := sc.Get().Cache()
	if err != nil {
		return
	}

	c.Del("")
}

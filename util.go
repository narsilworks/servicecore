package servicecore

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/narsilworks/servicecore/ifcs"
	ssd "github.com/shopspring/decimal"
	"golang.org/x/exp/constraints"
)

type FieldTypeConstraint interface {
	constraints.Ordered | time.Time | ssd.Decimal | bool | byte
}

func mapValue[T any](identity *map[string]any, key string, out *T) {
	if identity == nil || len(*identity) == 0 {
		return
	}

	val, ok := (*identity)[key]
	if !ok {
		return
	}

	if rval, ok := val.(T); ok {
		*out = rval
	}
}

func getZero[T FieldTypeConstraint]() T {
	var result T
	return result
}

func isNullOrEmpty[T FieldTypeConstraint](value *T) bool {
	return value == nil || *value == getZero[T]()
}

func val[T FieldTypeConstraint](value *T) T {
	if value == nil {
		return getZero[T]()
	}

	return *value
}

func loadConfig(src string, cfg *ifcs.IConfiguration) error {

	var (
		err       error
		b         []byte
		isLocFile bool
	)

	if !(strings.HasPrefix(src, `http://`) || strings.HasPrefix(src, `https://`)) {
		isLocFile = true
	}

	if isLocFile {
		if b, err = os.ReadFile(src); err != nil {
			return err
		}
	} else {
		if b, err = func() ([]byte, error) {
			var ob []byte
			nr, err := http.Get(src)
			if err != nil {
				return ob, err
			}
			defer nr.Body.Close()
			ob, err = io.ReadAll(nr.Body)
			if err != nil {
				return ob, err
			}
			return ob, nil
		}(); err != nil {
			return err
		}
	}

	if len(b) == 0 {
		return errors.New(`no data from source`)
	}

	err = json.Unmarshal(b, cfg)
	if err != nil {
		return err
	}

	(*cfg).SetFileName(src)

	return nil
}

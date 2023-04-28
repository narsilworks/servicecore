package dumm

import "errors"

type Cache struct {
}

func (obj Cache) Set(key string, value []byte) error {
	return errors.New("dummy cache")
}

func (obj Cache) Get(dst []byte, key string) []byte {
	return []byte{}
}

func (obj Cache) GetWithErr(key string) ([]byte, error) {
	return []byte{}, errors.New("dummy cache")
}

func (obj Cache) Del(keyPattern string) error {
	return errors.New("dummy cache")
}

func (obj Cache) Has(key string) bool {
	return false
}

func (obj Cache) Reset() {

}

func (obj Cache) ListKeys() []string {
	return []string{}
}

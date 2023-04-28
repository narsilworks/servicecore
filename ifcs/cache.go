package ifcs

// ICache interface for cache
type ICache interface {
	Set(key string, value []byte) error
	Get(dst []byte, key string) []byte
	GetWithErr(key string) ([]byte, error)
	Del(keyPattern string) error
	Has(key string) bool
	Reset()
	ListKeys() []string
}

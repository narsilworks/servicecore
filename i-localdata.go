package servicecore

type ILocalData interface {
	Fetch(bucket string, key string) (data []byte, err error)
	Pluck(bucket string, key string) ([]byte, error)
	Delete(bucket string, key string) error
	Store(bucket string, key string, data []byte) error
	Close() error
}

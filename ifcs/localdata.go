package ifcs

type ILocalData interface {
	Fetch(bucket, key string) (data []byte, err error)
	Pluck(bucket, key string) ([]byte, error)
	Delete(bucket, key string) error
	Store(bucket, key string, data []byte) error
	Close() error
}

package ifcs

type IQueue interface {
	Publish([]byte) error
}

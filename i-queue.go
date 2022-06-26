package servicecore

type IQueue interface {
	Publish(evtPlusPayload []byte) error
}
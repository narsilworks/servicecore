package ifcs

type IMessage interface {
	SetBody(string) error
	Body() string
}

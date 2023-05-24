package ifcs

type INotification interface {
	Send(m IMessage) error
}

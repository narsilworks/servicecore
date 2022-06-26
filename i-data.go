package servicecore

type IData interface {
	Open() error
	Handler() dhl.DataHelperLite
	Close() error
}

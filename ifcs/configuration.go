package ifcs

type IConfiguration interface {
	Save() error
	Reload() error
}

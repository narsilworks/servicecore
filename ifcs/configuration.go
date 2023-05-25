package ifcs

type IConfiguration interface {
	Save(fileName string) error
	Reload() error
}

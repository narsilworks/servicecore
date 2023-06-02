package ifcs

type IConfiguration interface {
	Save(fn string) error
	Reload() error
	GetFileName() string
	SetFileName(fn string)
}

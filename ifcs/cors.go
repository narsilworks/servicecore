package ifcs

type ICORS interface {
	Headers(headers ...string) []string
	Methods(methods ...string) []string
	Origins(origins ...string) []string
	Valid() bool
}

package dumm

type CORS struct {
}

func (obj *CORS) Headers(headers ...string) []string {
	return []string{}
}

func (obj *CORS) Methods(methods ...string) []string {
	return []string{}
}

func (obj *CORS) Origins(origins ...string) []string {
	return []string{}
}

func (obj *CORS) Valid() bool {
	return false
}

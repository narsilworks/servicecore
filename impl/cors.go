package imp

import "strings"

type CORSInfo struct {
	headers []string
	methods []string
	origins []string
}

// Headers set or get CORS allowed headers
func (obj *CORSInfo) Headers(headers ...string) []string {
	if len(headers) == 0 {
		return obj.headers
	}

	for _, h := range headers {
		for _, ah := range obj.headers {
			if strings.EqualFold(ah, h) {
				return obj.headers
			}
		}
		obj.headers = append(obj.headers, h)
	}

	return obj.headers
}

// Methods set or get CORS allowed methods
func (obj *CORSInfo) Methods(methods ...string) []string {
	if len(methods) == 0 {
		return obj.methods
	}

	for _, h := range methods {
		for _, ah := range obj.methods {
			if strings.EqualFold(ah, h) {
				return obj.methods
			}
		}
		obj.methods = append(obj.methods, h)
	}

	return obj.methods
}

// Origins set or get CORS allowed origins
func (obj *CORSInfo) Origins(origins ...string) []string {
	if len(origins) == 0 {
		return obj.origins
	}

	for _, h := range origins {
		for _, ah := range obj.origins {
			if strings.EqualFold(ah, h) {
				return obj.origins
			}
		}
		obj.origins = append(obj.origins, h)
	}

	return obj.origins
}

package cwctl

// EP represents a CW endpoint.
type EP int

const (
	EPToken = iota
	EPTokenRefresh
)

var epStrings = [...]string{
	`/cwa/api/v1/apitoken`,
	`/cwa/api/v1/apitoken/refresh`,
}

func (ep EP) String() string {
	return epStrings[ep]
}

func getMethod(ep EP) (method string) {
	switch ep {
	case EPToken, EPTokenRefresh:
		return `POST`
	default:
		return `GET`
	}
}

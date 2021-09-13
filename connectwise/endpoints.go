package connectwise

import "sort"

// EP represents a CW endpoint.
type EP int

const (
	EPToken = iota
	EPTokenRefresh
	EPClients
)

var epStrings = [...]string{
	`/cwa/api/v1/apitoken`,
	`/cwa/api/v1/apitoken/refresh`,
	`/cwa/api/v1/clients`,
}

func (ep EP) String() string {
	return epStrings[ep]
}

var setEmpty = struct{}{}
var availEPs = map[EP]struct{}{
	EPClients: setEmpty,
}

func epAvailable(ep EP) bool {
	if _, ok := availEPs[ep]; ok {
		return true
	}
	return false
}

// Endpoint contains the ID and Path for a CW API endpoint.
type EndPoint struct {
	ID   EP
	Path string
}

// GetEndPoints retrieves all available EndPoints.
func GetEndPoints() []EndPoint {
	EPs := make([]EndPoint, len(availEPs))
	i := 0
	for k := range availEPs {
		EPs[i] = EndPoint{
			ID:   k,
			Path: k.String(),
		}
		i++
	}
	sort.SliceStable(EPs, func(i, j int) bool {
		return EPs[i].ID < EPs[j].ID
	})
	return EPs
}

func getMethod(ep EP) (method string) {
	switch ep {
	case EPToken, EPTokenRefresh:
		return `POST`
	default:
		return `GET`
	}
}

package connectwise

import (
	"fmt"
	"sort"
	"strings"
)

// EP represents a CW endpoint.
type EP int

const (
	EPToken = iota
	EPTokenRefresh
	EPClients
	EPComputers
	EPComputer
	EPComputerCmdExec
	EPComputerCmdHistory
	EPCommands
)

var epStrings = [...]string{
	`/cwa/api/v1/apitoken`,
	`/cwa/api/v1/apitoken/refresh`,
	`/cwa/api/v1/clients`,
	`/cwa/api/v1/computers`,
	`/cwa/api/v1/computers/%v`,
	`/cwa/api/v1/computers/%v/commandexecute`,
	`/cwa/api/v1/computers/%v/commandhistory`,
	`/cwa/api/v1/commands`,
}

func (ep EP) String(args ...interface{}) string {
	n := strings.Count(epStrings[ep], `%v`)
	a := len(args)
	switch {
	case n == a:
	case n == 0:
		args = []interface{}{}
	case n > a:
		delta := n - len(args)
		for i := 0; i < delta; i++ {
			args = append(args, `%v`)
		}
	case a > n:
		args = args[:n]
	}
	return fmt.Sprintf(epStrings[ep], args...)
}

var setEmpty = struct{}{}
var availEPs = map[EP]struct{}{
	EPClients:            setEmpty,
	EPComputers:          setEmpty,
	EPComputer:           setEmpty,
	EPComputerCmdExec:    setEmpty,
	EPComputerCmdHistory: setEmpty,
	EPCommands:           setEmpty,
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
func GetEndPoints(args ...interface{}) []EndPoint {
	EPs := make([]EndPoint, len(availEPs))
	i := 0
	for k := range availEPs {
		EPs[i] = EndPoint{
			ID:   k,
			Path: k.String(args...),
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
	case EPComputerCmdExec:
		return `POST`
	default:
		return `GET`
	}
}

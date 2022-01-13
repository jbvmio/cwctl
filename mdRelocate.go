package cwctl

import (
	"encoding/json"
	"fmt"

	"github.com/jbvmio/cwctl/connectwise"
)

type RelocatePayload struct {
	EntityType int
	TargetType int
	TargetId   string
	EntityIds  []string
}

type RelocateResults struct {
	SendToResults               []SendToResults
	ContainsUnsuccessfulResults bool
}

type SendToResults struct {
	EntityId      int
	ResultDetails ResultDetails
}

type ResultDetails struct {
	ResultStatus int
	ReasonCode   int
	Message      string
}

func RelocateComputers(C *connectwise.Client, params *connectwise.Parameters, body RelocatePayload) (RelocateResults, error) {
	var (
		resource RelocateResults
		desc     string         = `relocating computers`
		ep       connectwise.EP = connectwise.EPSendTo
	)
	if len(body.EntityIds) < 1 {
		return resource, fmt.Errorf("error %s: %s", desc, `missing computerIDs`)
	}
	if body.TargetId == "" {
		return resource, fmt.Errorf("error %s: %s", desc, `missing destination location id`)
	}
	body.EntityType = 1
	body.TargetType = 2

	b, err := C.Post(ep, params, body)
	if err != nil {
		return resource, fmt.Errorf("error %s: %w", desc, err)
	}
	err = json.Unmarshal(b, &resource)
	return resource, err
}

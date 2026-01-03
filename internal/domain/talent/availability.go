package talent

import (
	"encoding/json"
	"errors"
)

var ErrInvalidAvailability = errors.New("invalid availability")

type AvailabilityStatus string

const (
	AvailabilityUnknown     AvailabilityStatus = "unknown"
	AvailabilityAvailable   AvailabilityStatus = "available"
	AvailabilityBusy        AvailabilityStatus = "busy"
	AvailabilityUnavailable AvailabilityStatus = "unavailable"
)

// Availability is a value object representing current availability.
type Availability struct {
	status AvailabilityStatus
}

func NewAvailability(status AvailabilityStatus) (Availability, error) {
	if !status.valid() {
		return Availability{}, ErrInvalidAvailability
	}

	return Availability{status: status}, nil
}

func Available() Availability {
	return Availability{status: AvailabilityAvailable}
}

func (a Availability) Status() AvailabilityStatus {
	if a.status == "" {
		return AvailabilityUnknown
	}

	return a.status
}

func (a Availability) MarshalJSON() ([]byte, error) {
	return json.Marshal(a.Status())
}

func (a *Availability) UnmarshalJSON(data []byte) error {
	var value AvailabilityStatus
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	availability, err := NewAvailability(value)
	if err != nil {
		return err
	}

	*a = availability
	return nil
}

func (s AvailabilityStatus) valid() bool {
	switch s {
	case AvailabilityUnknown, AvailabilityAvailable, AvailabilityBusy, AvailabilityUnavailable:
		return true
	default:
		return false
	}
}

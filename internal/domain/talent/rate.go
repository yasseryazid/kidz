package talent

import (
	"encoding/json"
	"errors"
	"strings"
)

var ErrInvalidRate = errors.New("invalid rate")

type RateUnit string

const (
	RateUnitHour  RateUnit = "hour"
	RateUnitDay   RateUnit = "day"
	RateUnitMonth RateUnit = "month"
)

// Rate is a value object describing a compensation rate.
type Rate struct {
	amount   int64
	currency string
	unit     RateUnit
}

func NewRate(amount int64, currency string, unit RateUnit) (*Rate, error) {
	normalizedCurrency := strings.TrimSpace(currency)
	if amount <= 0 || normalizedCurrency == "" || !unit.valid() {
		return nil, ErrInvalidRate
	}

	return &Rate{
		amount:   amount,
		currency: strings.ToUpper(normalizedCurrency),
		unit:     unit,
	}, nil
}

func (r Rate) Amount() int64 {
	return r.amount
}

func (r Rate) Currency() string {
	return r.currency
}

func (r Rate) Unit() RateUnit {
	return r.unit
}

func (r Rate) MarshalJSON() ([]byte, error) {
	payload := struct {
		Amount   int64    `json:"amount"`
		Currency string   `json:"currency"`
		Unit     RateUnit `json:"unit"`
	}{
		Amount:   r.amount,
		Currency: r.currency,
		Unit:     r.unit,
	}

	return json.Marshal(payload)
}

func (r *Rate) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		*r = Rate{}
		return nil
	}

	var payload struct {
		Amount   int64    `json:"amount"`
		Currency string   `json:"currency"`
		Unit     RateUnit `json:"unit"`
	}
	if err := json.Unmarshal(data, &payload); err != nil {
		return err
	}

	rate, err := NewRate(payload.Amount, payload.Currency, payload.Unit)
	if err != nil {
		return err
	}

	*r = *rate
	return nil
}

func (u RateUnit) valid() bool {
	switch u {
	case RateUnitHour, RateUnitDay, RateUnitMonth:
		return true
	default:
		return false
	}
}

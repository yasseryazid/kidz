package talent

import (
	"encoding/json"
	"errors"
	"strings"
)

var ErrInvalidSkill = errors.New("invalid skill")

// Skill is a value object representing a normalized skill label.
type Skill struct {
	name string
}

func NewSkill(name string) (Skill, error) {
	normalized := strings.TrimSpace(name)
	if normalized == "" {
		return Skill{}, ErrInvalidSkill
	}

	return Skill{name: normalized}, nil
}

func (s Skill) Name() string {
	return s.name
}

func (s Skill) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.name)
}

func (s *Skill) UnmarshalJSON(data []byte) error {
	var value string
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	skill, err := NewSkill(value)
	if err != nil {
		return err
	}

	*s = skill
	return nil
}

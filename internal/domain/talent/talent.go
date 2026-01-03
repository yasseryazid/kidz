package talent

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

var (
	ErrInvalidTalent = errors.New("invalid talent")
	ErrNotFound      = errors.New("talent not found")
)

// Talent represents a person profile in the talent list.
type Talent struct {
	ID           string       `json:"id"`
	Name         string       `json:"name"`
	Title        string       `json:"title"`
	Skills       []Skill      `json:"skills"`
	Rate         *Rate        `json:"rate,omitempty"`
	Availability Availability `json:"availability"`
	CreatedAt    time.Time    `json:"created_at"`
}

// New creates a new Talent after basic validation.
func New(name, title string, skills []string) (Talent, error) {
	normalizedName := strings.TrimSpace(name)
	if normalizedName == "" {
		return Talent{}, ErrInvalidTalent
	}

	normalizedSkills, err := buildSkills(skills)
	if err != nil {
		return Talent{}, fmt.Errorf("%w: %v", ErrInvalidTalent, err)
	}

	return Talent{
		Name:         normalizedName,
		Title:        strings.TrimSpace(title),
		Skills:       normalizedSkills,
		Availability: Available(),
		CreatedAt:    time.Now().UTC(),
	}, nil
}

func buildSkills(values []string) ([]Skill, error) {
	if len(values) == 0 {
		return []Skill{}, nil
	}

	seen := make(map[string]struct{}, len(values))
	result := make([]Skill, 0, len(values))
	for _, value := range values {
		skill, err := NewSkill(value)
		if err != nil {
			return nil, err
		}
		if _, exists := seen[skill.name]; exists {
			continue
		}
		seen[skill.name] = struct{}{}
		result = append(result, skill)
	}

	return result, nil
}

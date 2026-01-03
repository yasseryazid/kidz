package talent

import "context"

// Repository defines storage operations for Talent.
type Repository interface {
	List(ctx context.Context) ([]Talent, error)
	Create(ctx context.Context, talent Talent) (Talent, error)
	GetByID(ctx context.Context, id string) (Talent, error)
}

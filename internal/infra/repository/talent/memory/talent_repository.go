package memory

import (
	"context"
	"strconv"
	"sync"

	domain "github.com/yasseryazid/boilerpart/internal/domain/talent"
)

// TalentRepository is an in-memory implementation of the talent repository.
type TalentRepository struct {
	mu     sync.RWMutex
	items  map[string]domain.Talent
	nextID int
}

func NewTalentRepository() *TalentRepository {
	return &TalentRepository{
		items:  make(map[string]domain.Talent),
		nextID: 1,
	}
}

func (r *TalentRepository) List(ctx context.Context) ([]domain.Talent, error) {
	_ = ctx

	r.mu.RLock()
	defer r.mu.RUnlock()

	results := make([]domain.Talent, 0, len(r.items))
	for _, talent := range r.items {
		results = append(results, talent)
	}

	return results, nil
}

func (r *TalentRepository) Create(ctx context.Context, talent domain.Talent) (domain.Talent, error) {
	_ = ctx

	r.mu.Lock()
	defer r.mu.Unlock()

	id := strconv.Itoa(r.nextID)
	r.nextID++

	talent.ID = id
	r.items[id] = talent

	return talent, nil
}

func (r *TalentRepository) GetByID(ctx context.Context, id string) (domain.Talent, error) {
	_ = ctx

	r.mu.RLock()
	defer r.mu.RUnlock()

	talent, ok := r.items[id]
	if !ok {
		return domain.Talent{}, domain.ErrNotFound
	}

	return talent, nil
}

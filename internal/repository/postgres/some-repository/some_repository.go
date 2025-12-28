package some_repository

type SomeRepository struct{}

func New() *SomeRepository {
	return &SomeRepository{}
}

func (r *SomeRepository) Create() {}

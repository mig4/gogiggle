package swapi

type inMemoryRepository struct {
	data []Character
}

// NewInMemoryRepository factory for in-memory repository
func NewInMemoryRepository() *inMemoryRepository {
	return &inMemoryRepository{data: TestCharacters}
}

func (r *inMemoryRepository) GetAll() ([]Character, error) {
	return r.data, nil
}

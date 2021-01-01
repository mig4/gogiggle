package swapi

// Port of the SW-API data backend
type Port interface {
	GetAll() ([]Character, error)
}

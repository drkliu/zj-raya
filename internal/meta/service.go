package meta
type MetaService interface {
	Repository
}
type service struct {
	Repository
}
func NewService(repository *Repository) MetaService {
	return &service{
		Repository: *repository,
	}
}
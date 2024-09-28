package auth

type Place struct {
	ID    int
	Name  string
	Image string
}

type PlaceUsecase interface {
	GetPlace() ([]Place, error)
}

type RepoUsecase struct {
	repos Repository
}

func NewRepoUsecase(repos *Repository) *RepoUsecase {
	return &RepoUsecase{repos: *repos}
}

func (i *RepoUsecase) GetPlace() ([]Place, error) {
	places, err := i.repos.getPlaces()
	if err != nil {
		return nil, err
	}
	return places, nil
}

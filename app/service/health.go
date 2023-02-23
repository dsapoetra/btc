package service

type HealthService struct {
	//db repository.IArticleRepository
}

type IHealthService interface {
	//FindOneArticle(id uuid.UUID) (*repo.Article, error)
	//CreateArticle(article *http.CreateArticleRequest) error

	HealthCheck() error
}

func NewHealthService() IHealthService {
	return &HealthService{
		//db: repo,
	}
}

func (h *HealthService) HealthCheck() error {
	//res, err := a.db.FindOneArticle(id)

	//if err != nil {
	//	return nil, err
	//}

	return nil
}

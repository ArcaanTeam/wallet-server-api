package service

type ITestRepo interface {
	Ping() string
}

type TestService struct {
	repo ITestRepo
}

func NewTestService(repo ITestRepo) *TestService {
	return &TestService{
		repo: repo,
	}
}

func (service *TestService) Ping() string {
	return service.repo.Ping()
}
